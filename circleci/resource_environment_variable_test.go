/*
 * Copyright (c) 2020 Vulcan, Inc.
 * All rights reserved.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 */

package circleci

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironmentVariableResource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"resource argument for slug is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			assert.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "slug")
		},
		"resource argument for slug is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "slug")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_environment_variable"].Schema["slug"].Type)
		},
		"resource argument for slug is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "slug")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["slug"].Required)
		},
		"changing resource argument for slug forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "slug")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["slug"].ForceNew)
		},
		"resource argument for name is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			assert.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "name")
		},
		"resource argument for name is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "name")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_environment_variable"].Schema["name"].Type)
		},
		"resource argument for name is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "name")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["name"].Required)
		},
		"changing resource argument for name forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "name")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["name"].ForceNew)
		},
		"resource argument for value is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			assert.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "value")
		},
		"resource argument for value is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "value")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_environment_variable"].Schema["value"].Type)
		},
		"resource argument for value is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "value")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["value"].Required)
		},
		"changing resource argument for value forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "value")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["value"].ForceNew)
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}
