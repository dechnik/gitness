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

package enum

import (
	"strings"
)

// RuleState represents rule state.
type RuleState string

// RuleState enumeration.
const (
	RuleStateActive   = "active"
	RuleStateMonitor  = "monitor"
	RuleStateDisabled = "disabled"
)

var ruleStates = sortEnum([]RuleState{
	RuleStateActive,
	RuleStateMonitor,
	RuleStateDisabled,
})

func (RuleState) Enum() []interface{} { return toInterfaceSlice(ruleStates) }
func (s RuleState) Sanitize() (RuleState, bool) {
	return Sanitize(s, GetAllRuleStates)
}
func GetAllRuleStates() ([]RuleState, RuleState) {
	return ruleStates, RuleStateActive
}

// RuleSort contains protection rule sorting options.
type RuleSort string

const (
	RuleSortUID     = uid
	RuleSortCreated = createdAt
	RuleSortUpdated = updatedAt
)

var ruleSorts = sortEnum([]RuleSort{
	RuleSortUID,
	RuleSortCreated,
	RuleSortUpdated,
})

func (RuleSort) Enum() []interface{} { return toInterfaceSlice(ruleSorts) }
func (s RuleSort) Sanitize() (RuleSort, bool) {
	return Sanitize(s, GetAllRuleSorts)
}
func GetAllRuleSorts() ([]RuleSort, RuleSort) {
	return ruleSorts, RuleSortCreated
}

// ParseRuleSortAttr parses the protection rule sorting option.
func ParseRuleSortAttr(s string) RuleSort {
	switch strings.ToLower(s) {
	case created, createdAt:
		return RuleSortCreated
	case updated, updatedAt:
		return RuleSortUpdated
	}

	return RuleSortUID
}
