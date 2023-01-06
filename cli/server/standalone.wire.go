// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

//go:build wireinject && !harness
// +build wireinject,!harness

package server

import (
	"context"

	"github.com/harness/gitness/events"
	"github.com/harness/gitness/gitrpc"
	gitrpcevents "github.com/harness/gitness/gitrpc/events"
	gitrpcserver "github.com/harness/gitness/gitrpc/server"
	"github.com/harness/gitness/internal/api/controller/pullreq"
	"github.com/harness/gitness/internal/api/controller/repo"
	"github.com/harness/gitness/internal/api/controller/serviceaccount"
	"github.com/harness/gitness/internal/api/controller/space"
	"github.com/harness/gitness/internal/api/controller/user"
	controllerwebhook "github.com/harness/gitness/internal/api/controller/webhook"
	"github.com/harness/gitness/internal/auth/authn"
	"github.com/harness/gitness/internal/auth/authz"
	"github.com/harness/gitness/internal/bootstrap"
	"github.com/harness/gitness/internal/cron"
	"github.com/harness/gitness/internal/router"
	"github.com/harness/gitness/internal/server"
	"github.com/harness/gitness/internal/store"
	"github.com/harness/gitness/internal/store/database"
	"github.com/harness/gitness/internal/webhook"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/check"

	"github.com/google/wire"
)

func initSystem(ctx context.Context, config *types.Config) (*system, error) {
	wire.Build(
		newSystem,
		PackageConfigsWireSet,
		ProvideRedis,
		bootstrap.WireSet,
		database.WireSet,
		router.WireSet,
		server.WireSet,
		cron.WireSet,
		space.WireSet,
		repo.WireSet,
		pullreq.WireSet,
		controllerwebhook.WireSet,
		serviceaccount.WireSet,
		user.WireSet,
		authn.WireSet,
		authz.WireSet,
		gitrpcevents.WireSet,
		gitrpcserver.WireSet,
		gitrpc.WireSet,
		store.WireSet,
		check.WireSet,
		events.WireSet,
		webhook.WireSet,
	)
	return &system{}, nil
}