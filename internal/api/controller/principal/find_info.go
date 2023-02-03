// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package principal

import (
	"context"
	"fmt"

	"github.com/harness/gitness/types"
)

// FindInfoByUIDNoAuth tries to find the provided principal by UID.
// Note: No authorization is required for this API.
func (c *Controller) FindInfoByUIDNoAuth(ctx context.Context,
	uid string) (*types.PrincipalInfo, error) {
	principal, err := c.principalStore.FindByUID(ctx, uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get principal info from cache: %w", err)
	}

	return principal.ToPrincipalInfo(), nil
}