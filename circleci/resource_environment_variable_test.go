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
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testResourceEnvironmentVariableStateKey = "circleci_environment_variable.test"

func TestEnvironmentVariableResource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"resource argument for project_slug is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			assert.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "project_slug")
		},
		"resource argument for project_slug is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "project_slug")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_environment_variable"].Schema["project_slug"].Type)
		},
		"resource argument for project_slug is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "project_slug")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["project_slug"].Required)
		},
		"changing resource argument for project_slug forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "project_slug")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["project_slug"].ForceNew)
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
		"resource argument for value is sensitive": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_environment_variable")
			require.Contains(t, Provider().ResourcesMap["circleci_environment_variable"].Schema, "value")
			assert.True(t, Provider().ResourcesMap["circleci_environment_variable"].Schema["value"].Sensitive)
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}

func TestAccEnvironmentVariableResource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"resource creates and deletes as expected": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NotEmpty(t, os.Getenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: `
            resource "circleci_environment_variable" "test" {
              project_slug = "gh/VulcanTechnologies/terraform-provider-circleci-acceptance-test-target"
              name         = "FOO"
              value        = "BAR"
            }
            `,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testResourceEnvironmentVariableStateKey, "id", "gh/VulcanTechnologies/terraform-provider-circleci-acceptance-test-target/FOO"),
							resource.TestCheckResourceAttr(testResourceEnvironmentVariableStateKey, "name", "FOO"),
							resource.TestCheckResourceAttr(testResourceEnvironmentVariableStateKey, "value", "BAR"),
						),
					},
				},
				CheckDestroy: confirmEnvironmentVariableResourceDestroyed,
			})
		},
		"errors when project_slug does not start with allowed values": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NotEmpty(t, os.Getenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: `
            resource "circleci_environment_variable" "test" {
              project_slug = "nope"
              name         = "FOO"
              value        = "BAR"
            }
            `,
						ExpectError: regexp.MustCompile(`A project_slug must begin with 'gh/' or 'bb/' depending on your vcs provider.`),
					},
				},
			})
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}

func confirmEnvironmentVariableResourceDestroyed(state *terraform.State) error {
	if state.Empty() {
		return errors.New("state should not be empty")
	}

	resourceAttributes := state.RootModule().Resources[testResourceEnvironmentVariableStateKey].Primary.Attributes

	provider := testAccProvider.Meta().(*providerContext)
	auth := provider.authenticateContext(context.Background())
	api := provider.circleCiClient.ProjectApi

	slug := resourceAttributes["project_slug"]
	name := resourceAttributes["name"]

	_, resp, _ := api.GetEnvVar(auth, slug, name)

	switch resp.StatusCode {
	case http.StatusNotFound: //unfortunately, this could mask a permissions error, but if we've gotten this far, the permissions error should have previously surfaced
		return nil
	case http.StatusOK:
		return fmt.Errorf("the environment variable named '%s' still exists", name)
	default:
		return fmt.Errorf("received unexpeced status code %d when checking if the environment variable named '%s' still exists", resp.StatusCode, name)
	}
}
