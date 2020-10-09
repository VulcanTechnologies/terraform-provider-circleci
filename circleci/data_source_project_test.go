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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const githubOrg = "VulcanTechnologies"
const githubRepo = "terraform-provider-circleci-acceptance-test-target"

var testCircleCiSlug = fmt.Sprintf("gh/%s/%s", githubOrg, githubRepo)
var testDataSourceProject = fmt.Sprintf("data \"circleci_project\" \"test\" { slug=\"%s\" }", testCircleCiSlug)

const testDataSourceStateKey = "data.circleci_project.test"

func TestProjectDataSource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"data source argument for slug is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			assert.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "slug")
		},
		"data source argument for slug is string": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "slug")
			assert.Equal(t, schema.TypeString, Provider().DataSourcesMap["circleci_project"].Schema["slug"].Type)
		},
		"data source argument for slug is required": func(t *testing.T) {
			require.Contains(t, Provider().DataSourcesMap, "circleci_project")
			require.Contains(t, Provider().DataSourcesMap["circleci_project"].Schema, "slug")
			assert.True(t, Provider().DataSourcesMap["circleci_project"].Schema["slug"].Required)
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

func TestAccProjectDataSource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"attributes are set as expected": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NotEmpty(t, os.Getenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: testDataSourceProject,
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(testDataSourceStateKey, "id", testCircleCiSlug),
							resource.TestCheckResourceAttr(testDataSourceStateKey, "name", githubRepo),
							resource.TestCheckResourceAttr(testDataSourceStateKey, "organization_name", githubOrg)),
					},
				},
			})
		},
		"errors when slug does not start with allowed values": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NotEmpty(t, os.Getenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config:      `data "circleci_project" "test" { slug="nope"}`,
						ExpectError: regexp.MustCompile(`A slug must begin with 'gh/' or 'bb/' depending on your vcs provider.`),
					},
				},
			})
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}
