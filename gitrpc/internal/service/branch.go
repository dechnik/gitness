// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package service

import (
	"context"
	"fmt"

	"github.com/harness/gitness/gitrpc/internal/gitea"
	"github.com/harness/gitness/gitrpc/internal/types"
	"github.com/harness/gitness/gitrpc/rpc"

	"code.gitea.io/gitea/modules/git"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var listBranchesRefFields = []types.GitReferenceField{types.GitReferenceFieldRefName, types.GitReferenceFieldObjectName}

func (s ReferenceService) CreateBranch(ctx context.Context,
	request *rpc.CreateBranchRequest) (*rpc.CreateBranchResponse, error) {
	base := request.GetBase()
	if base == nil {
		return nil, types.ErrBaseCannotBeEmpty
	}

	repoPath := getFullPathForRepo(s.reposRoot, base.GetRepoUid())

	// TODO: why are we using gitea operations here?!
	repo, err := git.OpenRepository(ctx, repoPath)
	if err != nil {
		return nil, processGitErrorf(err, "failed to open repo")
	}

	sharedRepo, err := NewSharedRepo(s.tmpDir, base.GetRepoUid(), repo)
	if err != nil {
		return nil, processGitErrorf(err, "failed to create new shared repo")
	}
	defer sharedRepo.Close(ctx)

	// clone repo
	err = sharedRepo.Clone(ctx, request.GetTarget())
	if err != nil {
		return nil, processGitErrorf(err, "failed to clone shared repo with branch '%s'", request.GetBranchName())
	}

	// push to new branch (all changes should go through push flow for hooks and other safety meassures)
	err = sharedRepo.Push(ctx, base, request.GetTarget(), request.GetBranchName())
	if err != nil {
		return nil, processGitErrorf(err, "failed to push new branch '%s'", request.GetBranchName())
	}

	// get branch
	// TODO: get it from shared repo to avoid opening another gitea repo
	gitBranch, err := s.adapter.GetBranch(ctx, repoPath, request.GetBranchName())
	if err != nil {
		return nil, processGitErrorf(err, "failed to get gitea branch '%s'", request.GetBranchName())
	}

	branch, err := mapGitBranch(gitBranch)
	if err != nil {
		return nil, err
	}

	return &rpc.CreateBranchResponse{
		Branch: branch,
	}, nil
}

func (s ReferenceService) DeleteBranch(ctx context.Context,
	request *rpc.DeleteBranchRequest) (*rpc.DeleteBranchResponse, error) {
	base := request.GetBase()
	if base == nil {
		return nil, types.ErrBaseCannotBeEmpty
	}

	repoPath := getFullPathForRepo(s.reposRoot, base.GetRepoUid())

	// TODO: why are we using gitea operations here?!
	repo, err := git.OpenRepository(ctx, repoPath)
	if err != nil {
		return nil, processGitErrorf(err, "failed to open repo")
	}

	sharedRepo, err := NewSharedRepo(s.tmpDir, base.GetRepoUid(), repo)
	if err != nil {
		return nil, processGitErrorf(err, "failed to create new shared repo")
	}
	defer sharedRepo.Close(ctx)

	// clone repo (technically we don't care about which branch we clone)
	err = sharedRepo.Clone(ctx, request.GetBranchName())
	if err != nil {
		return nil, processGitErrorf(err, "failed to clone shared repo with branch '%s'", request.GetBranchName())
	}

	// get latest branch commit before we delete
	gitCommit, err := sharedRepo.GetBranchCommit(request.GetBranchName())
	if err != nil {
		return nil, processGitErrorf(err, "failed to get gitea commit for branch '%s'", request.GetBranchName())
	}

	// push to new branch (all changes should go through push flow for hooks and other safety meassures)
	// NOTE: setting sourceRef to empty will delete the remote branch when pushing:
	// https://git-scm.com/docs/git-push#Documentation/git-push.txt-ltrefspecgt82308203
	err = sharedRepo.Push(ctx, base, "", request.GetBranchName())
	if err != nil {
		return nil, processGitErrorf(err, "failed to delete branch '%s' from remote repo", request.GetBranchName())
	}

	return &rpc.DeleteBranchResponse{
		Sha: gitCommit.ID.String(),
	}, nil
}

func (s ReferenceService) ListBranches(request *rpc.ListBranchesRequest,
	stream rpc.ReferenceService_ListBranchesServer) error {
	base := request.GetBase()
	if base == nil {
		return types.ErrBaseCannotBeEmpty
	}

	ctx := stream.Context()
	repoPath := getFullPathForRepo(s.reposRoot, base.GetRepoUid())

	// get all required information from git refrences
	branches, err := s.listBranchesLoadReferenceData(ctx, repoPath, request)
	if err != nil {
		return err
	}

	// get commits if needed (single call for perf savings: 1s-4s vs 5s-20s)
	if request.GetIncludeCommit() {
		commitSHAs := make([]string, len(branches))
		for i := range branches {
			commitSHAs[i] = branches[i].Sha
		}

		var gitCommits []types.Commit
		gitCommits, err = s.adapter.GetCommits(ctx, repoPath, commitSHAs)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to get commits: %v", err)
		}

		for i := range gitCommits {
			branches[i].Commit, err = mapGitCommit(&gitCommits[i])
			if err != nil {
				return err
			}
		}
	}

	// send out all branches
	for _, branch := range branches {
		err = stream.Send(&rpc.ListBranchesResponse{
			Branch: branch,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "failed to send branch: %v", err)
		}
	}

	return nil
}

func (s ReferenceService) listBranchesLoadReferenceData(ctx context.Context,
	repoPath string, request *rpc.ListBranchesRequest) ([]*rpc.Branch, error) {
	// TODO: can we be smarter with slice allocation
	branches := make([]*rpc.Branch, 0, 16)
	handler := listBranchesWalkReferencesHandler(&branches)
	instructor, endsAfter, err := wrapInstructorWithOptionalPagination(
		gitea.DefaultInstructor, // branches only have one target type, default instructor is enough
		request.GetPage(),
		request.GetPageSize())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pagination details: %v", err)
	}

	opts := &types.WalkReferencesOptions{
		Patterns:   createReferenceWalkPatternsFromQuery(gitReferenceNamePrefixBranch, request.GetQuery()),
		Sort:       mapListBranchesSortOption(request.Sort),
		Order:      mapSortOrder(request.Order),
		Fields:     listBranchesRefFields,
		Instructor: instructor,
		// we don't do any post-filtering, restrict git to only return as many elements as pagination needs.
		MaxWalkDistance: endsAfter,
	}

	err = s.adapter.WalkReferences(ctx, repoPath, handler, opts)
	if err != nil {
		return nil, processGitErrorf(err, "failed to walk branch references")
	}

	log.Ctx(ctx).Trace().Msgf("git adapter returned %d branches", len(branches))

	return branches, nil
}

func listBranchesWalkReferencesHandler(branches *[]*rpc.Branch) types.WalkReferencesHandler {
	return func(e types.WalkReferencesEntry) error {
		fullRefName, ok := e[types.GitReferenceFieldRefName]
		if !ok {
			return fmt.Errorf("entry missing reference name")
		}
		objectSHA, ok := e[types.GitReferenceFieldObjectName]
		if !ok {
			return fmt.Errorf("entry missing object sha")
		}

		branch := &rpc.Branch{
			Name: fullRefName[len(gitReferenceNamePrefixBranch):],
			Sha:  objectSHA,
		}

		// TODO: refactor to not use slice pointers?
		*branches = append(*branches, branch)

		return nil
	}
}