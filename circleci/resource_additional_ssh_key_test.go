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

func TestAdditionalSSHKeyResource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"resource argument for project_slug is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			assert.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "project_slug")
		},
		"resource argument for project_slug is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "project_slug")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["project_slug"].Type)
		},
		"resource argument for project_slug is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "project_slug")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["project_slug"].Required)
		},
		"changing resource argument for project_slug forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "project_slug")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["project_slug"].ForceNew)
		},
		"resource argument for host_name is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			assert.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "host_name")
		},
		"resource argument for host_name is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "host_name")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["host_name"].Type)
		},
		"resource argument for host_name is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "host_name")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["host_name"].Required)
		},
		"changing resource argument for host_name forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "host_name")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["host_name"].ForceNew)
		},
		"resource argument for fingerprint is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			assert.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "fingerprint")
		},
		"resource argument for fingerprint is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "fingerprint")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["fingerprint"].Type)
		},
		"resource argument for fingerprint is computed": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "fingerprint")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["fingerprint"].Computed)
		},
		"resource argument for private_key is named as expected": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			assert.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "private_key")
		},
		"resource argument for private_key is string": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "private_key")
			assert.Equal(t, schema.TypeString, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["private_key"].Type)
		},
		"resource argument for private_key is required": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "private_key")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["private_key"].Required)
		},
		"changing resource argument for private_key forces new": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "private_key")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["private_key"].ForceNew)
		},
		"resource argument for private_key is sensitive": func(t *testing.T) {
			require.Contains(t, Provider().ResourcesMap, "circleci_additional_ssh_key")
			require.Contains(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema, "private_key")
			assert.True(t, Provider().ResourcesMap["circleci_additional_ssh_key"].Schema["private_key"].Sensitive)
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}

type testAdditionalSSHKeyResourceConfig struct {
	projectSlug        string
	hostName           string
	privateKeyArgument string // this could be hard-coded or a Terraform function like file()
}

func (c testAdditionalSSHKeyResourceConfig) materialize() string {
	// since privateKeyArgument could be a Terraform function, it is not quoted by default here
	return fmt.Sprintf(`
    resource "circleci_additional_ssh_key" "test" {
      project_slug  = "%s"
      host_name     = "%s"
      private_key   = %s
    }
  `, c.projectSlug, c.hostName, c.privateKeyArgument)
}

func (c testAdditionalSSHKeyResourceConfig) withValidDefaultProjectSlug() testAdditionalSSHKeyResourceConfig {
	newConfig := c
	newConfig.projectSlug = testSlug
	return newConfig
}

func TestAccAdditionalSSHKeyResource(t *testing.T) {
	testCases := map[string]func(*testing.T){
		"errors when project_slug does not start with allowed values": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAdditionalSSHKeyResourceConfig{
							projectSlug:        "nope",
							hostName:           "FOO",
							privateKeyArgument: `file("testdata/ssh-keys/id_rsa")`,
						}.materialize(),

						ExpectError: regexp.MustCompile(`A project_slug must begin with 'gh/' or 'bb/' depending on your vcs provider.`),
					},
				},
			})
		},
		"errors when private_key does not parse": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAdditionalSSHKeyResourceConfig{
							hostName:           "FOO",
							privateKeyArgument: `"nope"`,
						}.withValidDefaultProjectSlug().materialize(),

						ExpectError: regexp.MustCompile(`Received error 'ssh: no key found' while trying to parse private_key`),
					},
				},
			})
		},
		"RSA key creates and deletes as expected": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					assureAdditionalSSHKeyDoesNotExist(t, testSlug, "b3:fb:7a:ff:ec:7f:8c:7c:3b:1b:24:c1:42:c6:5d:a1", "FOO")
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAdditionalSSHKeyResourceConfig{
							hostName:           "FOO",
							privateKeyArgument: `file("testdata/ssh-keys/id_rsa")`,
						}.withValidDefaultProjectSlug().materialize(),

						Check: resource.ComposeTestCheckFunc(
							// the fingerprint is of the public key
							resource.TestCheckResourceAttr("circleci_additional_ssh_key.test", "id", fmt.Sprintf("%s/b3:fb:7a:ff:ec:7f:8c:7c:3b:1b:24:c1:42:c6:5d:a1", testSlug)),
							resource.TestCheckResourceAttr("circleci_additional_ssh_key.test", "fingerprint", "b3:fb:7a:ff:ec:7f:8c:7c:3b:1b:24:c1:42:c6:5d:a1"),
						),
					},
				},
			})
		},
		"ECDSA key creates and deletes as expected": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					assureAdditionalSSHKeyDoesNotExist(t, testSlug, "aa:c6:6d:1b:4f:23:4d:9b:a7:ea:f0:ea:d7:ee:ad:45", "FOO")
				},
				ProviderFactories: testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: testAdditionalSSHKeyResourceConfig{
							hostName:           "FOO",
							privateKeyArgument: `file("testdata/ssh-keys/id_ecdsa")`,
						}.withValidDefaultProjectSlug().materialize(),

						Check: resource.ComposeTestCheckFunc(
							// the fingerprint is of the public key
							resource.TestCheckResourceAttr("circleci_additional_ssh_key.test", "id", fmt.Sprintf("%s/aa:c6:6d:1b:4f:23:4d:9b:a7:ea:f0:ea:d7:ee:ad:45", testSlug)),
							resource.TestCheckResourceAttr("circleci_additional_ssh_key.test", "fingerprint", "aa:c6:6d:1b:4f:23:4d:9b:a7:ea:f0:ea:d7:ee:ad:45"),
						),
					},
				},
			})
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}

func assureAdditionalSSHKeyDoesNotExist(t *testing.T, slug string, fingerprint string, hostname string) {
	t.Helper()

	code, _, err := deleteFromLegacyEndpoint(context.Background(), slug, hostname, fingerprint, testAccProvider.Meta())

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, code)
}

func confirmAdditionalSSHKeyDestroyed(state *terraform.State) error {
	if state.Empty() {
		return errors.New("pre-destroy state should not be empty")
	}

	for _, resource := range state.RootModule().Resources {
		if resource.Type != "circleci_additional_ssh_key" {
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

	return errors.New("did not find any resources of type circleci_additional_ssh_key in pre-destroy state")
}
