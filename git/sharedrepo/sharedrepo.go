// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sharedrepo

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/harness/gitness/errors"
	"github.com/harness/gitness/git/api"
	"github.com/harness/gitness/git/command"
	"github.com/harness/gitness/git/sha"
	"github.com/harness/gitness/git/tempdir"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/slices"
)

type SharedRepo struct {
	repoPath       string
	sourceRepoPath string
}

// NewSharedRepo creates a new temporary bare repository.
func NewSharedRepo(
	baseTmpDir string,
	sourceRepoPath string,
) (*SharedRepo, error) {
	if sourceRepoPath == "" {
		return nil, errors.New("repository path can't be empty")
	}

	var buf [5]byte
	_, _ = rand.Read(buf[:])
	id := base32.StdEncoding.EncodeToString(buf[:])

	repoPath, err := tempdir.CreateTemporaryPath(baseTmpDir, id)
	if err != nil {
		return nil, fmt.Errorf("failed to create shared repository directory: %w", err)
	}

	t := &SharedRepo{
		repoPath:       repoPath,
		sourceRepoPath: sourceRepoPath,
	}

	return t, nil
}

func (r *SharedRepo) Close(ctx context.Context) {
	if err := tempdir.RemoveTemporaryPath(r.repoPath); err != nil {
		log.Ctx(ctx).Err(err).
			Str("path", r.repoPath).
			Msg("Failed to remove temporary shared directory")
	}
}

func (r *SharedRepo) Init(ctx context.Context, alternates ...string) error {
	cmd := command.New("init", command.WithFlag("--bare"))

	if err := cmd.Run(ctx, command.WithDir(r.repoPath)); err != nil {
		return fmt.Errorf("failed to initialize bare git repository directory: %w", err)
	}

	if err := func() error {
		alternatesFilePath := filepath.Join(r.repoPath, "objects", "info", "alternates")
		f, err := os.OpenFile(alternatesFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
		if err != nil {
			return fmt.Errorf("failed to create alternates file: %w", err)
		}

		defer func() { _ = f.Close() }()

		data := filepath.Join(r.sourceRepoPath, "objects")
		if _, err = fmt.Fprintln(f, data); err != nil {
			return fmt.Errorf("failed to write alternates file: %w", err)
		}

		for _, alternate := range alternates {
			if _, err = fmt.Fprintln(f, alternate); err != nil {
				return fmt.Errorf("failed to write alternates file: %w", err)
			}
		}

		return nil
	}(); err != nil {
		return fmt.Errorf("failed to make the alternates file in shared repository: %w", err)
	}

	return nil
}

func (r *SharedRepo) Directory() string {
	return r.repoPath
}

// SetDefaultIndex sets the git index to our HEAD.
func (r *SharedRepo) SetDefaultIndex(ctx context.Context) error {
	cmd := command.New("read-tree", command.WithArg("HEAD"))

	if err := cmd.Run(ctx, command.WithDir(r.repoPath)); err != nil {
		return fmt.Errorf("failed to initialize shared repository index to HEAD: %w", err)
	}

	return nil
}

// SetIndex sets the git index to the provided treeish.
func (r *SharedRepo) SetIndex(ctx context.Context, treeish sha.SHA) error {
	cmd := command.New("read-tree", command.WithArg(treeish.String()))

	if err := cmd.Run(ctx, command.WithDir(r.repoPath)); err != nil {
		return fmt.Errorf("failed to initialize shared repository index to %q: %w", treeish, err)
	}

	return nil
}

// ClearIndex clears the git index.
func (r *SharedRepo) ClearIndex(ctx context.Context) error {
	cmd := command.New("read-tree", command.WithFlag("--empty"))

	if err := cmd.Run(ctx, command.WithDir(r.repoPath)); err != nil {
		return fmt.Errorf("failed to clear shared repository index: %w", err)
	}

	return nil
}

// LsFiles checks if the given filename arguments are in the index.
func (r *SharedRepo) LsFiles(
	ctx context.Context,
	filenames ...string,
) ([]string, error) {
	cmd := command.New("ls-files",
		command.WithFlag("-z"),
		command.WithPostSepArg(filenames...),
	)

	stdout := bytes.NewBuffer(nil)

	err := cmd.Run(ctx, command.WithDir(r.repoPath), command.WithStdout(stdout))
	if err != nil {
		return nil, fmt.Errorf("failed to list files in shared repository's git index: %w", err)
	}

	files := make([]string, 0)
	for _, line := range bytes.Split(stdout.Bytes(), []byte{'\000'}) {
		files = append(files, string(line))
	}

	return files, nil
}

// RemoveFilesFromIndex removes the given files from the index.
func (r *SharedRepo) RemoveFilesFromIndex(
	ctx context.Context,
	filenames ...string,
) error {
	cmd := command.New("update-index",
		command.WithFlag("--remove"),
		command.WithFlag("-z"),
		command.WithFlag("--index-info"))

	stdin := bytes.NewBuffer(nil)
	for _, file := range filenames {
		if file != "" {
			stdin.WriteString("0 0000000000000000000000000000000000000000\t")
			stdin.WriteString(file)
			stdin.WriteByte('\000')
		}
	}

	if err := cmd.Run(ctx, command.WithDir(r.repoPath), command.WithStdin(stdin)); err != nil {
		return fmt.Errorf("failed to update-index in shared repo: %w", err)
	}

	return nil
}

// WriteGitObject writes the provided content to the object db and returns its hash.
func (r *SharedRepo) WriteGitObject(
	ctx context.Context,
	content io.Reader,
) (sha.SHA, error) {
	cmd := command.New("hash-object",
		command.WithFlag("-w"),
		command.WithFlag("--stdin"))

	stdout := bytes.NewBuffer(nil)

	err := cmd.Run(ctx,
		command.WithDir(r.repoPath),
		command.WithStdin(content),
		command.WithStdout(stdout))
	if err != nil {
		return sha.None, fmt.Errorf("failed to hash-object in shared repo: %w", err)
	}

	return sha.New(stdout.String())
}

// GetTreeSHA returns the tree SHA of the rev.
func (r *SharedRepo) GetTreeSHA(
	ctx context.Context,
	rev string,
) (sha.SHA, error) {
	cmd := command.New("show",
		command.WithFlag("--no-patch"),
		command.WithFlag("--format=%T"),
		command.WithArg(rev),
	)
	stdout := &bytes.Buffer{}
	err := cmd.Run(ctx,
		command.WithDir(r.repoPath),
		command.WithStdout(stdout),
	)
	if err != nil {
		if strings.Contains(err.Error(), "ambiguous argument") {
			return sha.None, errors.NotFound("could not resolve git revision %q", rev)
		}
		return sha.None, fmt.Errorf("failed to get tree sha: %w", err)
	}

	return sha.New(stdout.String())
}

// ShowFile dumps show file and write to io.Writer.
func (r *SharedRepo) ShowFile(
	ctx context.Context,
	filePath string,
	rev string,
	writer io.Writer,
) error {
	file := strings.TrimSpace(rev) + ":" + strings.TrimSpace(filePath)

	cmd := command.New("show", command.WithArg(file))

	if err := cmd.Run(ctx, command.WithDir(r.repoPath), command.WithStdout(writer)); err != nil {
		return fmt.Errorf("failed to show file in shared repo: %w", err)
	}

	return nil
}

// AddObjectToIndex adds the provided object hash to the index with the provided mode and path.
func (r *SharedRepo) AddObjectToIndex(
	ctx context.Context,
	mode string,
	objectHash sha.SHA,
	objectPath string,
) error {
	cmd := command.New("update-index",
		command.WithFlag("--add"),
		command.WithFlag("--replace"),
		command.WithFlag("--cacheinfo", mode, objectHash.String(), objectPath))

	if err := cmd.Run(ctx, command.WithDir(r.repoPath)); err != nil {
		if matched, _ := regexp.MatchString(".*Invalid path '.*", err.Error()); matched {
			return errors.InvalidArgument("invalid path '%s'", objectPath)
		}
		return fmt.Errorf("failed to add object to index in shared repo (path=%s): %w", objectPath, err)
	}

	return nil
}

// WriteTree writes the current index as a tree to the object db and returns its hash.
func (r *SharedRepo) WriteTree(ctx context.Context) (sha.SHA, error) {
	cmd := command.New("write-tree")

	stdout := bytes.NewBuffer(nil)

	if err := cmd.Run(ctx, command.WithDir(r.repoPath), command.WithStdout(stdout)); err != nil {
		return sha.None, fmt.Errorf("failed to write-tree in shared repo: %w", err)
	}

	return sha.New(stdout.String())
}

// MergeTree merges commits in git index.
func (r *SharedRepo) MergeTree(
	ctx context.Context,
	commitMergeBase, commitTarget, commitSource sha.SHA,
) (sha.SHA, []string, error) {
	cmd := command.New("merge-tree",
		command.WithFlag("--write-tree"),
		command.WithFlag("--name-only"),
		command.WithFlag("--no-messages"),
		command.WithArg(commitTarget.String()),
		command.WithArg(commitSource.String()))

	if !commitMergeBase.IsEmpty() {
		cmd.Add(command.WithFlag("--merge-base=" + commitMergeBase.String()))
	}

	stdout := bytes.NewBuffer(nil)

	err := cmd.Run(ctx,
		command.WithDir(r.repoPath),
		command.WithStdout(stdout))

	// no error: the output is just the tree object SHA
	if err == nil {
		return sha.Must(stdout.String()), nil, nil
	}

	// exit code=1: the output is the tree object SHA, and list of files in conflict.
	if cErr := command.AsError(err); cErr != nil && cErr.ExitCode() == 1 {
		output := strings.TrimSpace(stdout.String())
		lines := strings.Split(output, "\n")
		if len(lines) < 2 {
			log.Ctx(ctx).Err(err).Str("output", output).Msg("unexpected output of merge-tree in shared repo")
			return sha.None, nil, fmt.Errorf("unexpected output of merge-tree in shared repo: %w", err)
		}
		return sha.Must(lines[0]), lines[1:], nil
	}

	return sha.None, nil, fmt.Errorf("failed to merge-tree in shared repo: %w", err)
}

// CommitTree creates a commit from a given tree for the user with provided message.
func (r *SharedRepo) CommitTree(
	ctx context.Context,
	author, committer *api.Signature,
	treeHash sha.SHA,
	message string,
	signoff bool,
	parentCommits ...sha.SHA,
) (sha.SHA, error) {
	cmd := command.New("commit-tree",
		command.WithArg(treeHash.String()),
		command.WithAuthorAndDate(
			author.Identity.Name,
			author.Identity.Email,
			author.When,
		),
		command.WithCommitterAndDate(
			committer.Identity.Name,
			committer.Identity.Email,
			committer.When,
		),
	)

	for _, parentCommit := range parentCommits {
		cmd.Add(command.WithFlag("-p", parentCommit.String()))
	}

	// temporary no signing
	cmd.Add(command.WithFlag("--no-gpg-sign"))

	messageBytes := new(bytes.Buffer)
	_, _ = messageBytes.WriteString(message)
	_, _ = messageBytes.WriteString("\n")

	if signoff {
		// Signed-off-by
		_, _ = messageBytes.WriteString("\n")
		_, _ = messageBytes.WriteString("Signed-off-by: ")
		_, _ = messageBytes.WriteString(fmt.Sprintf("%s <%s>", committer.Identity.Name, committer.Identity.Email))
	}

	stdout := bytes.NewBuffer(nil)

	err := cmd.Run(ctx,
		command.WithDir(r.repoPath),
		command.WithStdout(stdout),
		command.WithStdin(messageBytes))
	if err != nil {
		return sha.None, fmt.Errorf("failed to commit-tree in shared repo: %w", err)
	}

	return sha.New(stdout.String())
}

// CommitSHAsForRebase returns list of SHAs of the commits between the two git revisions
// for a rebase operation - in the order they should be rebased in.
func (r *SharedRepo) CommitSHAsForRebase(
	ctx context.Context,
	target, source sha.SHA,
) ([]sha.SHA, error) {
	// the command line arguments are mostly matching default `git rebase` behavior.
	// Only difference is we use `--date-order` (matches github behavior) whilst `git rebase` uses `--topo-order`.
	// Git Rebase's rev-list: https://github.com/git/git/blob/v2.41.0/sequencer.c#L5703-L5714
	cmd := command.New("rev-list",
		command.WithFlag("--max-parents=1"), // exclude merge commits
		command.WithFlag("--cherry-pick"),   // drop commits that already exist on target
		command.WithFlag("--reverse"),
		command.WithFlag("--right-only"), // only return commits from source
		command.WithFlag("--date-order"), // childs always before parents, otherwise by commit time stamp
		command.WithArg(target.String()+"..."+source.String()))

	stdout := bytes.NewBuffer(nil)

	if err := cmd.Run(ctx, command.WithDir(r.repoPath), command.WithStdout(stdout)); err != nil {
		return nil, fmt.Errorf("failed to rev-list in shared repo: %w", err)
	}

	var commitSHAs []sha.SHA

	scan := bufio.NewScanner(stdout)
	for scan.Scan() {
		commitSHA := sha.Must(scan.Text())
		commitSHAs = append(commitSHAs, commitSHA)
	}
	if err := scan.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rev-list output in shared repo: %w", err)
	}

	return commitSHAs, nil
}

// MergeBase returns number of commits between the two git revisions.
func (r *SharedRepo) MergeBase(
	ctx context.Context,
	rev1, rev2 string,
) (string, error) {
	cmd := command.New("merge-base", command.WithArg(rev1), command.WithArg(rev2))

	stdout := bytes.NewBuffer(nil)

	if err := cmd.Run(ctx, command.WithDir(r.repoPath), command.WithStdout(stdout)); err != nil {
		return "", fmt.Errorf("failed to merge-base in shared repo: %w", err)
	}

	return strings.TrimSpace(stdout.String()), nil
}

func (r *SharedRepo) CreateFile(
	ctx context.Context,
	treeishSHA sha.SHA,
	filePath, mode string,
	payload []byte,
) error {
	// only check path availability if a source commit is available (empty repo won't have such a commit)
	if !treeishSHA.IsEmpty() {
		if err := r.checkPathAvailability(ctx, treeishSHA, filePath, true); err != nil {
			return err
		}
	}

	objectSHA, err := r.WriteGitObject(ctx, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("createFile: error hashing object: %w", err)
	}

	// Add the object to the index
	if err = r.AddObjectToIndex(ctx, mode, objectSHA, filePath); err != nil {
		return fmt.Errorf("createFile: error creating object: %w", err)
	}

	return nil
}

func (r *SharedRepo) UpdateFile(
	ctx context.Context,
	treeishSHA sha.SHA,
	filePath string,
	objectSHA sha.SHA,
	mode string,
	payload []byte,
) error {
	// get file mode from existing file (default unless executable)
	entry, err := r.getFileEntry(ctx, treeishSHA, objectSHA, filePath)
	if err != nil {
		return err
	}

	if entry.IsExecutable() {
		mode = "100755"
	}

	objectSHA, err = r.WriteGitObject(ctx, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("updateFile: error hashing object: %w", err)
	}

	if err = r.AddObjectToIndex(ctx, mode, objectSHA, filePath); err != nil {
		return fmt.Errorf("updateFile: error updating object: %w", err)
	}

	return nil
}

func (r *SharedRepo) MoveFile(
	ctx context.Context,
	treeishSHA sha.SHA,
	filePath string,
	objectSHA sha.SHA,
	mode string,
	payload []byte,
) error {
	newPath, newContent, err := parseMovePayload(payload)
	if err != nil {
		return err
	}

	// ensure file exists and matches SHA
	entry, err := r.getFileEntry(ctx, treeishSHA, objectSHA, filePath)
	if err != nil {
		return err
	}

	// ensure new path is available
	if err = r.checkPathAvailability(ctx, treeishSHA, newPath, false); err != nil {
		return err
	}

	var fileHash sha.SHA
	var fileMode string
	if newContent != nil {
		hash, err := r.WriteGitObject(ctx, bytes.NewReader(newContent))
		if err != nil {
			return fmt.Errorf("moveFile: error hashing object: %w", err)
		}

		fileHash = hash
		fileMode = mode
		if entry.IsExecutable() {
			const filePermissionExec = "100755"
			fileMode = filePermissionExec
		}
	} else {
		fileHash = entry.SHA
		fileMode = entry.Mode.String()
	}

	if err = r.AddObjectToIndex(ctx, fileMode, fileHash, newPath); err != nil {
		return fmt.Errorf("moveFile: add object error: %w", err)
	}

	if err = r.RemoveFilesFromIndex(ctx, filePath); err != nil {
		return fmt.Errorf("moveFile: remove object error: %w", err)
	}

	return nil
}

func (r *SharedRepo) DeleteFile(ctx context.Context, filePath string) error {
	filesInIndex, err := r.LsFiles(ctx, filePath)
	if err != nil {
		return fmt.Errorf("deleteFile: listing files error: %w", err)
	}
	if !slices.Contains(filesInIndex, filePath) {
		return errors.NotFound("file path %s not found", filePath)
	}

	if err = r.RemoveFilesFromIndex(ctx, filePath); err != nil {
		return fmt.Errorf("deleteFile: remove object error: %w", err)
	}
	return nil
}

func (r *SharedRepo) getFileEntry(
	ctx context.Context,
	treeishSHA sha.SHA,
	objectSHA sha.SHA,
	path string,
) (*api.TreeNode, error) {
	entry, err := api.GetTreeNode(ctx, r.repoPath, treeishSHA.String(), path, false)
	if errors.IsNotFound(err) {
		return nil, errors.NotFound("path %s not found", path)
	}
	if err != nil {
		return nil, fmt.Errorf("getFileEntry: failed to get tree for path %s: %w", path, err)
	}

	// If a SHA was given and the SHA given doesn't match the SHA of the fromTreePath, throw error
	if !objectSHA.IsEmpty() && !objectSHA.Equal(entry.SHA) {
		return nil, errors.InvalidArgument("sha does not match for path %s [given: %s, expected: %s]",
			path, objectSHA, entry.SHA)
	}

	return entry, nil
}

// checkPathAvailability ensures that the path is available for the requested operation.
// For the path where this file will be created/updated, we need to make
// sure no parts of the path are existing files or links except for the last
// item in the path which is the file name, and that shouldn't exist IF it is
// a new file OR is being moved to a new path.
func (r *SharedRepo) checkPathAvailability(
	ctx context.Context,
	treeishSHA sha.SHA,
	filePath string,
	isNewFile bool,
) error {
	parts := strings.Split(filePath, "/")
	subTreePath := ""

	for index, part := range parts {
		subTreePath = path.Join(subTreePath, part)

		entry, err := api.GetTreeNode(ctx, r.repoPath, treeishSHA.String(), subTreePath, false)
		if err != nil {
			if errors.IsNotFound(err) {
				// Means there is no item with that name, so we're good
				break
			}
			return fmt.Errorf("checkPathAvailability: failed to get tree entry for path %s: %w", subTreePath, err)
		}

		switch {
		case index < len(parts)-1:
			if !entry.IsDir() {
				return errors.Conflict("a file already exists where you're trying to create a subdirectory [path: %s]",
					subTreePath)
			}
		case entry.IsLink():
			return errors.Conflict("a symbolic link already exist where you're trying to create a subdirectory [path: %s]",
				subTreePath)
		case entry.IsDir():
			return errors.Conflict("a directory already exists where you're trying to create a subdirectory [path: %s]",
				subTreePath)
		case filePath != "" || isNewFile:
			return errors.Conflict("file path %s already exists", filePath)
		}
	}
	return nil
}

// MoveObjects moves git object from the shared repository to the original repository.
func (r *SharedRepo) MoveObjects(ctx context.Context) error {
	if r.sourceRepoPath == "" {
		return errors.New("shared repo not initialized with a repository")
	}

	srcDir := path.Join(r.repoPath, "objects")
	dstDir := path.Join(r.sourceRepoPath, "objects")

	var files []fileEntry

	err := filepath.WalkDir(srcDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(srcDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		// avoid coping anything in the info/
		if strings.HasPrefix(relPath, "info/") {
			return nil
		}

		fileName := filepath.Base(relPath)

		files = append(files, fileEntry{
			fileName: fileName,
			fullPath: path,
			relPath:  relPath,
			priority: filePriority(fileName),
		})

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to list files of shared repository directory: %w", err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].priority < files[j].priority // 0 is top priority, 5 is lowest priority
	})

	for _, f := range files {
		dstPath := filepath.Join(dstDir, f.relPath)

		err = os.MkdirAll(filepath.Dir(dstPath), os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory for git object: %w", err)
		}

		// Try to move the file

		errRename := os.Rename(f.fullPath, dstPath)
		if errRename == nil {
			log.Ctx(ctx).Debug().
				Str("object", f.relPath).
				Msg("moved git object")
			continue
		}

		// Try to copy the file

		copyError := func() error {
			srcFile, err := os.Open(f.fullPath)
			if err != nil {
				return fmt.Errorf("failed to open source file: %w", err)
			}
			defer func() { _ = srcFile.Close() }()

			dstFile, err := os.Create(dstPath)
			if err != nil {
				return fmt.Errorf("failed to create target file: %w", err)
			}
			defer func() { _ = dstFile.Close() }()

			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return fmt.Errorf("failed to copy file content: %w", err)
			}

			return nil
		}()
		if copyError != nil {
			log.Ctx(ctx).Err(copyError).
				Str("object", f.relPath).
				Str("renameErr", errRename.Error()).
				Msg("failed to move or copy git object")
			return fmt.Errorf("failed to move or copy git object: %w", copyError)
		}

		log.Ctx(ctx).Warn().
			Str("object", f.relPath).
			Str("renameErr", errRename.Error()).
			Msg("copied git object")
	}

	return nil
}

// filePriority is based on https://github.com/git/git/blob/master/tmp-objdir.c#L168
func filePriority(name string) int {
	switch {
	case !strings.HasPrefix(name, "pack"):
		return 0
	case strings.HasSuffix(name, ".keep"):
		return 1
	case strings.HasSuffix(name, ".pack"):
		return 2
	case strings.HasSuffix(name, ".rev"):
		return 3
	case strings.HasSuffix(name, ".idx"):
		return 4
	default:
		return 5
	}
}

type fileEntry struct {
	fileName string
	fullPath string
	relPath  string
	priority int
}

func parseMovePayload(payload []byte) (string, []byte, error) {
	var newContent []byte
	var newPath string
	filePathEnd := bytes.IndexByte(payload, 0)
	if filePathEnd < 0 {
		newPath = string(payload)
		newContent = nil
	} else {
		newPath = string(payload[:filePathEnd])
		newContent = payload[filePathEnd+1:]
	}

	newPath = api.CleanUploadFileName(newPath)
	if newPath == "" {
		return "", nil, api.ErrInvalidPath
	}

	return newPath, newContent, nil
}
