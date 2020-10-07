/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2.
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// EnvironmentVariablePair struct for EnvironmentVariablePair
type EnvironmentVariablePair struct {
	// The name of the environment variable.
	Name string `json:"name"`
	// The value of the environment variable.
	Value string `json:"value"`
}
