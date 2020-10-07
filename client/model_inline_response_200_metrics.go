/*
 * CircleCI API
 *
 * This describes the resources that make up the CircleCI API v2.
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// InlineResponse200Metrics Metrics relating to a workflow's runs.
type InlineResponse200Metrics struct {
	// The ratio of successful runs / total runs.
	SuccessRate float32 `json:"success_rate"`
	// The total number of runs.
	TotalRuns int64 `json:"total_runs"`
	// The number of failed runs.
	FailedRuns int64 `json:"failed_runs"`
	// The number of successful runs.
	SuccessfulRuns int64 `json:"successful_runs"`
	// The average number of workflow runs per day.
	Throughput float32 `json:"throughput"`
	// The mean time to recovery (mean time between failures and their next success) in seconds.
	Mttr int64 `json:"mttr"`
	// The total credits consumed by the workflow in the aggregation window.
	TotalCreditsUsed int64                                   `json:"total_credits_used"`
	DurationMetrics  InlineResponse200MetricsDurationMetrics `json:"duration_metrics"`
}
