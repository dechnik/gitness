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

package container

import (
	"context"

	"github.com/harness/gitness/infraprovider"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

type Orchestrator interface {
	// StartGitspace starts the gitspace container using the specified image or default image, clones the code,
	// runs SSH server and installs the IDE inside the container. It returns a map of the ports used by the Gitspace.
	StartGitspace(
		ctx context.Context,
		gitspaceConfig *types.GitspaceConfig,
		devcontainerConfig *types.DevcontainerConfig,
		infra *infraprovider.Infrastructure,
	) (map[enum.IDEType]string, error)

	// StopGitspace stops and removes the gitspace container.
	StopGitspace(ctx context.Context, gitspaceConfig *types.GitspaceConfig, infra *infraprovider.Infrastructure) error

	// Status checks if the infra is reachable and ready to begin container creation.
	Status(ctx context.Context, infra *infraprovider.Infrastructure) error
}
