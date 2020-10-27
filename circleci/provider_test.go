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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var testSlug string
var testVcsSlug string
var testRepoOwner string
var testRepoName string

var testAccProvider *schema.Provider

var testAccProviders = map[string]func() (*schema.Provider, error){
	"circleci": func() (*schema.Provider, error) {
		if testAccProvider == nil {
			testAccProvider = Provider()
		}
		return testAccProvider, nil
	},
}

func TestMain(m *testing.M) {
	testSlug = os.Getenv("TEST_TARGET_SLUG")

	if strings.Count(testSlug, "/") != 2 {
		fmt.Printf("cannot parse the environment variable key 'TEST_TARGET_SLUG' with value '%s' as a valid project slug", testSlug)
		os.Exit(1)
	}

	split := strings.Split(testSlug, "/")
	testVcsSlug = split[0]
	testRepoOwner = split[1]
	testRepoName = split[2]

	os.Exit(m.Run())
}

func TestProvider(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"provider internal validate does not error": func(t *testing.T) {
			assert.NoError(t, Provider().InternalValidate())
		},
		"provider argument for API key is named as expected": func(t *testing.T) {
			assert.Contains(t, Provider().Schema, "api_key")
		},
		"provider API key argument is string": func(t *testing.T) {
			require.Contains(t, Provider().Schema, "api_key")
			assert.Equal(t, schema.TypeString, Provider().Schema["api_key"].Type)
		},
		"provider API key argument is required": func(t *testing.T) {
			require.Contains(t, Provider().Schema, "api_key")
			assert.True(t, Provider().Schema["api_key"].Required)
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}

type materializer interface {
	materialize() string
}

type testProviderConfig struct {
	apiKey string
}

func (c testProviderConfig) materialize() string {
	if c.apiKey == "" {
		return fmt.Sprintf(`
      provider "circleci" {}
  `)
	} else {
		return fmt.Sprintf(`
      provider "circleci" {
        api_key = "%s"
      }
  `, c.apiKey)
	}
}

func (c testProviderConfig) materializeWithAdditionalConfig(m materializer) string {
	providerConfig := c.materialize()
	additionalConfig := m.materialize()
	return fmt.Sprintf("%s\n%s", providerConfig, additionalConfig)
}

func TestAccProvider(t *testing.T) {
	circleCiAPIKey := os.Getenv("CIRCLECI_API_KEY")

	if circleCiAPIKey == "" {
		t.Fatalf("The environment variable '%s' must be set for provider acceptance tests.", "CIRCLECI_API_KEY")
	}

	defer os.Setenv("CIRCLECI_API_KEY", circleCiAPIKey)

	testCases := map[string]func(*testing.T){
		"provider errors when API key not provided": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NoError(t, os.Unsetenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config:      testProviderConfig{}.materializeWithAdditionalConfig(testProjectDataSourceConfig{}.withValidDefaultProjectSlug()),
						ExpectError: regexp.MustCompile(`The argument "api_key" is required, but no definition was found.`),
					},
				},
			})
		},
		"provider reads API key from environment": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NoError(t, os.Setenv("CIRCLECI_API_KEY", circleCiAPIKey))
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testProviderConfig{}.materializeWithAdditionalConfig(testProjectDataSourceConfig{}.withValidDefaultProjectSlug()),
					},
				},
			})
		},
		"provider reads API key from config": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NoError(t, os.Unsetenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testProviderConfig{apiKey: circleCiAPIKey}.materializeWithAdditionalConfig(testProjectDataSourceConfig{}.withValidDefaultProjectSlug()),
					},
				},
			})
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}
