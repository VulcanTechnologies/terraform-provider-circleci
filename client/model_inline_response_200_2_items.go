/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2.
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

import (
	"time"
)

// InlineResponse2002Items struct for InlineResponse2002Items
type InlineResponse2002Items struct {
	// The name of the job.
	Name string `json:"name"`
	// The start of the aggregation window for job metrics.
	WindowStart time.Time `json:"window_start"`
	// The end of the aggregation window for job metrics.
	WindowEnd time.Time                 `json:"window_end"`
	Metrics   InlineResponse2002Metrics `json:"metrics"`
}