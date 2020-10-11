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
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config:      testDataSourceProject,
						ExpectError: regexp.MustCompile(`The argument "api_key" is required, but was not set.`),
					},
				},
			})
		},
		"provider reads API key from environment": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NoError(t, os.Setenv("CIRCLECI_API_KEY", circleCiAPIKey))
				},
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: testDataSourceProject,
					},
				},
			})
		},
		"provider reads API key from config": func(t *testing.T) {
			resource.Test(t, resource.TestCase{
				PreCheck: func() {
					require.NoError(t, os.Unsetenv("CIRCLECI_API_KEY"))
				},
				ProviderFactories: providerFactories,
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf("provider \"circleci\" { api_key= \"%s\"}\n", circleCiAPIKey) + testDataSourceProject,
					},
				},
			})
		},
	}

	for testCaseName, testCase := range testCases {
		t.Run(testCaseName, testCase)
	}
}
