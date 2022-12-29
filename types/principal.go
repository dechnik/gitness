// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

// Package types defines common data structures.
package types

import "github.com/harness/gitness/types/enum"

type (
	// Principal represents the identity of an acting entity (User, ServiceAccount, Service).
	Principal struct {
		// TODO: int64 ID doesn't match DB
		ID          int64              `db:"principal_id"           json:"-"`
		UID         string             `db:"principal_uid"          json:"uid"`
		Email       string             `db:"principal_email"        json:"email"`
		Type        enum.PrincipalType `db:"principal_type"         json:"type"`
		DisplayName string             `db:"principal_display_name" json:"display_name"`
		Admin       bool               `db:"principal_admin"        json:"admin"`

		// Should be part of principal or not?
		Blocked bool   `db:"principal_blocked"            json:"blocked"`
		Salt    string `db:"principal_salt"               json:"-"`

		// Other info
		Created int64 `db:"principal_created"                json:"created"`
		Updated int64 `db:"principal_updated"                json:"updated"`
	}
)

func PrincipalFromUser(user *User) *Principal {
	return &Principal{
		ID:          user.ID,
		UID:         user.UID,
		Email:       user.Email,
		Type:        enum.PrincipalTypeUser,
		DisplayName: user.DisplayName,
		Admin:       user.Admin,
		Blocked:     user.Blocked,
		Salt:        user.Salt,
		Created:     user.Created,
		Updated:     user.Updated,
	}
}

func PrincipalFromServiceAccount(sa *ServiceAccount) *Principal {
	return &Principal{
		ID:          sa.ID,
		UID:         sa.UID,
		Email:       sa.Email,
		Type:        enum.PrincipalTypeServiceAccount,
		DisplayName: sa.DisplayName,
		Admin:       false,
		Blocked:     sa.Blocked,
		Salt:        sa.Salt,
		Created:     sa.Created,
		Updated:     sa.Updated,
	}
}

func PrincipalFromService(s *Service) *Principal {
	return &Principal{
		ID:          s.ID,
		UID:         s.UID,
		Email:       s.Email,
		Type:        enum.PrincipalTypeService,
		DisplayName: s.DisplayName,
		Admin:       s.Admin,
		Blocked:     s.Blocked,
		Salt:        s.Salt,
		Created:     s.Created,
		Updated:     s.Updated,
	}
}
