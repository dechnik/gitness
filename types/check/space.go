// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package check

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/harness/gitness/types"
)

const (
	minSpaceNameLength = 1
	maxSpaceNameLength = 64
	spaceNameRegex     = "^[a-z][a-z0-9\\-\\_]*$"
)

var (
	ErrSpaceNameLength = errors.New(fmt.Sprintf("Space name has to be between %d and %d in length.", minSpaceNameLength, maxSpaceNameLength))
	ErrSpaceNameRegex  = errors.New("Space name has start with a letter and only contain the following [a-z0-9-_].")
)

// User returns true if the User if valid.
func Space(space *types.Space) (bool, error) {
	l := len(space.Name)
	if l < minSpaceNameLength || l > maxSpaceNameLength {
		return false, ErrSpaceNameLength
	}

	if ok, _ := regexp.Match(spaceNameRegex, []byte(space.Name)); !ok {
		return false, ErrSpaceNameRegex
	}

	return true, nil
}
