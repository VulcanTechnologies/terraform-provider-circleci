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
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProjectDataSource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"data source argument for project_slug is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			assert.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "project_slug")
		},
		"data source argument for project_slug is string": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "project_slug")
			assert.Equal(t, schema.TypeString, Provider().DataSourcesMap["circleci_project"].Schema["project_slug"].Type)
		},
		"data source argument for project_slug is required": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "project_slug")
			assert.True(t, Provider().DataSourcesMap["circleci_project"].Schema["project_slug"].Required)
		},
		"data source argument for name is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			assert.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "name")
		},
		"data source argument for name is string": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "name")
			assert.Equal(t, schema.TypeString, Provider().DataSourcesMap["circleci_project"].Schema["name"].Type)
		},
		"data source argument for name is computed": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "name")
			assert.True(t, Provider().DataSourcesMap["circleci_project"].Schema["name"].Computed)
		},
		"data source argument for organization_name is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			assert.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "organization_name")
		},
		"data source argument for organization_name is string": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "organization_name")
			assert.Equal(t, schema.TypeString, Provider().DataSourcesMap["circleci_project"].Schema["organization_name"].Type)
		},
		"data source argument for organization_name is computed": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "organization_name")
			assert.True(t, Provider().DataSourcesMap["circleci_project"].Schema["organization_name"].Computed)
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}

type testProjectDataSourceConfig struct {
	projectSlug string
}

func (c testProjectDataSourceConfig) withValidDefaultProjectSlug() testProjectDataSourceConfig {
	newConfig := c
	newConfig.projectSlug = testSlug
	return newConfig
}

func (c testProjectDataSourceConfig) materialize() string {
	return fmt.Sprintf(`
    data "circleci_project" "test" {
      project_slug = "%s"
    }
  `, c.projectSlug)
}

func TestAccProjectDataSource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"attributes are set as expected": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testProjectDataSourceConfig{}.withValidDefaultProjectSlug().materialize(),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr("data.circleci_project.test", "id", testSlug),
							resource.TestCheckResourceAttr("data.circleci_project.test", "project_slug", testSlug),
							resource.TestCheckResourceAttr("data.circleci_project.test", "name", testRepoName),
							resource.TestCheckResourceAttr("data.circleci_project.test", "organization_name", testRepoOwner)),
					},
				},
			})
		},
		"errors when project_slug does not start with allowed values": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testProjectDataSourceConfig{
							projectSlug: "nope",
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
