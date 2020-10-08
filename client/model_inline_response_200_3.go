/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2.
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// InlineResponse2003 Paginated recent job runs.
type InlineResponse2003 struct {
	// Recent job runs.
	Items []InlineResponse2003Items `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}