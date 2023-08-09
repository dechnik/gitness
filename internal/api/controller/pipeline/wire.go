// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package pipeline

import (
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"

	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/types/check"
)

// WireSet provides a wire set for this package.
var WireSet = wire.NewSet(
	ProvideController,
)

func ProvideController(db *sqlx.DB,
	uidCheck check.PathUID,
	pathStore store.PathStore,
	repoStore store.RepoStore,
	authorizer authz.Authorizer,
	pipelineStore store.PipelineStore,
	spaceStore store.SpaceStore,
) *Controller {
	return NewController(db, uidCheck, authorizer, pathStore, repoStore, pipelineStore, spaceStore)
}
