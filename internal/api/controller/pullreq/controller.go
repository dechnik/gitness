// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package pullreq

import (
	"context"
	"errors"
	"fmt"

	"github.com/harness/gitness/gitrpc"
	apiauth "github.com/harness/gitness/internal/api/auth"
	repoctrl "github.com/harness/gitness/internal/api/controller/repo"
	"github.com/harness/gitness/internal/api/usererror"
	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"

	"github.com/jmoiron/sqlx"
)

type Controller struct {
	db                   *sqlx.DB
	authorizer           authz.Authorizer
	pullreqStore         store.PullReqStore
	pullreqActivityStore store.PullReqActivityStore
	repoStore            store.RepoStore
	principalStore       store.PrincipalStore
	gitRPCClient         gitrpc.Interface
}

func NewController(
	db *sqlx.DB,
	authorizer authz.Authorizer,
	pullreqStore store.PullReqStore,
	pullreqActivityStore store.PullReqActivityStore,
	repoStore store.RepoStore,
	principalStore store.PrincipalStore,
	gitRPCClient gitrpc.Interface,
) *Controller {
	return &Controller{
		db:                   db,
		authorizer:           authorizer,
		pullreqStore:         pullreqStore,
		pullreqActivityStore: pullreqActivityStore,
		repoStore:            repoStore,
		principalStore:       principalStore,
		gitRPCClient:         gitRPCClient,
	}
}

func (c *Controller) verifyBranchExistence(ctx context.Context,
	repo *types.Repository, branch string) error {
	if branch == "" {
		return usererror.BadRequest("branch name can't be empty")
	}

	_, err := c.gitRPCClient.GetRef(ctx,
		&gitrpc.GetRefParams{
			ReadParams: repoctrl.CreateRPCReadParams(repo),
			Name:       branch,
			Type:       gitrpc.RefTypeBranch})
	if errors.Is(err, gitrpc.ErrNotFound) {
		return usererror.BadRequest(
			fmt.Sprintf("branch %s does not exist in the repository %s", branch, repo.UID))
	}
	if err != nil {
		return fmt.Errorf(
			"failed to check existence of the branch %s in the repository %s: %w",
			branch, repo.UID, err)
	}

	return nil
}

func (c *Controller) getRepoCheckAccess(ctx context.Context,
	session *auth.Session, repoRef string, reqPermission enum.Permission) (*types.Repository, error) {
	if repoRef == "" {
		return nil, usererror.BadRequest("A valid repository reference must be provided.")
	}

	repo, err := c.repoStore.FindRepoFromRef(ctx, repoRef)
	if err != nil {
		return nil, err
	}

	if err = apiauth.CheckRepo(ctx, c.authorizer, session, repo, reqPermission, false); err != nil {
		return nil, err
	}

	return repo, nil
}

func (c *Controller) getCommentCheckEditAccess(ctx context.Context,
	session *auth.Session, pr *types.PullReq, commentID int64) (*types.PullReqActivity, error) {
	if commentID <= 0 {
		return nil, usererror.BadRequest("A valid comment ID must be provided.")
	}

	comment, err := c.pullreqActivityStore.Find(ctx, commentID)
	if err != nil || comment == nil {
		return nil, fmt.Errorf("failed to find comment by ID: %w", err)
	}

	if comment.Deleted != nil || comment.RepoID != pr.TargetRepoID || comment.PullReqID != pr.ID {
		return nil, store.ErrResourceNotFound
	}

	if comment.Kind == enum.PullReqActivityKindSystem {
		return nil, usererror.BadRequest("Can't update a comment created by the system.")
	}

	if comment.CreatedBy != session.Principal.ID {
		return nil, usererror.BadRequest("Only own comments may be updated.")
	}

	return comment, nil
}

// writeActivity updates the PR's activity sequence number (using the optimistic locking mechanism),
// sets the correct Order value and writes the activity to the database.
// Even if the writing fails, the updating of the sequence number can succeed.
func (c *Controller) writeActivity(ctx context.Context, pr *types.PullReq, act *types.PullReqActivity) error {
	prUpd, err := c.pullreqStore.UpdateActivitySeq(ctx, pr)
	if err != nil {
		return fmt.Errorf("failed to get pull request activity number: %w", err)
	}

	*pr = *prUpd // update the pull request object

	act.Order = prUpd.ActivitySeq

	err = c.pullreqActivityStore.Create(ctx, act)
	if err != nil {
		return fmt.Errorf("failed to create pull request activity: %w", err)
	}

	return nil
}

// writeReplyActivity updates the parent activity's reply sequence number (using the optimistic locking mechanism),
// sets the correct Order and SubOrder values and writes the activity to the database.
// Even if the writing fails, the updating of the sequence number can succeed.
func (c *Controller) writeReplyActivity(ctx context.Context, parent, act *types.PullReqActivity) error {
	parentUpd, err := c.pullreqActivityStore.UpdateReplySeq(ctx, parent)
	if err != nil {
		return fmt.Errorf("failed to get pull request activity number: %w", err)
	}

	*parent = *parentUpd // update the parent pull request activity object

	act.Order = parentUpd.Order
	act.SubOrder = parentUpd.ReplySeq

	err = c.pullreqActivityStore.Create(ctx, act)
	if err != nil {
		return fmt.Errorf("failed to create pull request activity: %w", err)
	}

	return nil
}