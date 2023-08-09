// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package pipeline

import (
	"context"
	"fmt"

	apiauth "github.com/harness/gitness/internal/api/auth"
	"github.com/harness/gitness/internal/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

type UpdateInput struct {
	Description string `json:"description"`
	UID         string `json:"uid"`
	ConfigPath  string `json:"config_path"`
}

func (c *Controller) Update(
	ctx context.Context,
	session *auth.Session,
	spaceRef string,
	uid string,
	in *UpdateInput,
) (*types.Pipeline, error) {
	space, err := c.spaceStore.FindByRef(ctx, spaceRef)
	if err != nil {
		return nil, fmt.Errorf("could not find parent space: %w", err)
	}

	err = apiauth.CheckPipeline(ctx, c.authorizer, session, space.Path, uid, enum.PermissionPipelineEdit)
	if err != nil {
		return nil, fmt.Errorf("could not authorize: %w", err)
	}

	pipeline, err := c.pipelineStore.FindByUID(ctx, space.ID, uid)
	if err != nil {
		return nil, fmt.Errorf("could not find pipeline: %w", err)
	}

	if in.Description != "" {
		pipeline.Description = in.Description
	}
	if in.UID != "" {
		pipeline.UID = in.UID
	}
	if in.ConfigPath != "" {
		pipeline.ConfigPath = in.ConfigPath
	}

	return c.pipelineStore.Update(ctx, pipeline)
}
