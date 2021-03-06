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
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

type testEnvironmentVariableResourceConfig struct {
	projectSlug string
	name        string
	value       string
}

func (c testEnvironmentVariableResourceConfig) materialize() string {
	return fmt.Sprintf(`
    resource "circleci_environment_variable" "test" {
      project_slug = "%s"
      name         = "%s"
      value        = "%s"
    }
  `, c.projectSlug, c.name, c.value)
}

func (c testEnvironmentVariableResourceConfig) withValidDefaultProjectSlug() testEnvironmentVariableResourceConfig {
	newConfig := c
	newConfig.projectSlug = testSlug
	return newConfig
}

func TestAccEnvironmentVariableResource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"resource creates and deletes as expected": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					assureNoExistingEnvironmentVariables(t, testSlug)
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testEnvironmentVariableResourceConfig{
							name:  "FOO",
							value: "BAR",
						}.withValidDefaultProjectSlug().materialize(),

						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "id", fmt.Sprintf("%s/FOO", testSlug)),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "project_slug", testSlug),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "name", "FOO"),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "value", "BAR"),
							confirmEnvironmentVariableResourceExists,
						),
					},
				},
				CheckDestroy: confirmEnvironmentVariableResourceDestroyed,
			})
		},
		"resource handles argument changes": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					assureNoExistingEnvironmentVariables(t, testSlug)
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testEnvironmentVariableResourceConfig{
							name:  "FOO",
							value: "BAR",
						}.withValidDefaultProjectSlug().materialize(),

						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "id", fmt.Sprintf("%s/FOO", testSlug)),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "project_slug", testSlug),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "name", "FOO"),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "value", "BAR"),
							confirmEnvironmentVariableResourceExists,
						),
					},
					{
						Config: testEnvironmentVariableResourceConfig{
							name:  "SPAM",
							value: "BAR",
						}.withValidDefaultProjectSlug().materialize(),

						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "id", fmt.Sprintf("%s/SPAM", testSlug)),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "project_slug", testSlug),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "name", "SPAM"),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "value", "BAR"),
							confirmEnvironmentVariableResourceExists,
						),
					},
					{
						Config: testEnvironmentVariableResourceConfig{
							name:  "SPAM",
							value: "EGGS",
						}.withValidDefaultProjectSlug().materialize(),

						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "id", fmt.Sprintf("%s/SPAM", testSlug)),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "project_slug", testSlug),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "name", "SPAM"),
							resource.TestCheckResourceAttr("circleci_environment_variable.test", "value", "EGGS"),
							confirmEnvironmentVariableResourceExists,
						),
					},
				},
				CheckDestroy: confirmEnvironmentVariableResourceDestroyed,
			})
		},
		"errors when project_slug does not start with allowed values": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testEnvironmentVariableResourceConfig{
							projectSlug: "nope",
							name:        "FOO",
							value:       "BAR",
						}.materialize(),

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

func assureNoExistingEnvironmentVariables(t *testing.T, slug string) {
	t.Helper()

	provider := testAccProvider.Meta().(*providerContext)
	auth := provider.authenticateContext(context.Background())
	api := provider.circleCiClient.ProjectApi

	envVars, resp, err := api.ListEnvVars(auth, slug)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	require.Empty(t, envVars.Items)
}

func environmentVariableExistsInProject(slug, name string) (bool, error) {

	provider := testAccProvider.Meta().(*providerContext)
	auth := provider.authenticateContext(context.Background())
	api := provider.circleCiClient.ProjectApi

	_, resp, _ := api.GetEnvVar(auth, slug, name)

	switch resp.StatusCode {
	case http.StatusNotFound: //unfortunately, this could mask a permissions error, but if we've gotten this far, the permissions error should have previously surfaced
		return false, nil
	case http.StatusOK:
		return true, nil
	default:
		return false, fmt.Errorf("received unexpeced status code %d when checking if the environment variable named '%s' still exists", resp.StatusCode, name)
	}
}

func confirmEnvironmentVariableResourceDestroyed(state *terraform.State) error {
	if state.Empty() {
		return errors.New("pre-destroy state should not be empty")
	}

	for _, resource := range state.RootModule().Resources {
		if resource.Type != "circleci_environment_variable" {
			continue
		}

		resourceAttributes := resource.Primary.Attributes
		slug := resourceAttributes["project_slug"]
		name := resourceAttributes["name"]

		exists, err := environmentVariableExistsInProject(slug, name)

		if err != nil {
			return err
		}

		if exists {
			return fmt.Errorf("the environment variable named '%s' still exists", name)
		}

		return nil
	}

	return errors.New("did not find any resources of type circleci_environment_variable in pre-destroy state")
}

func confirmEnvironmentVariableResourceExists(state *terraform.State) error {
	if state.Empty() {
		return errors.New("state should not be empty")
	}

	for _, resource := range state.RootModule().Resources {
		if resource.Type != "circleci_environment_variable" {
			continue
		}

		resourceAttributes := resource.Primary.Attributes
		slug := resourceAttributes["project_slug"]
		name := resourceAttributes["name"]

		exists, err := environmentVariableExistsInProject(slug, name)

		if err != nil {
			return err
		}

		if !exists {
			return fmt.Errorf("the environment variable named '%s' does not exist", name)
		}

		return nil
	}

	return errors.New("did not find any resources of type circleci_environment_variable in state")
}
