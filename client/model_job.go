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

// Job Job
type Job struct {
	// The unique ID of the user.
	CanceledBy string `json:"canceled_by,omitempty"`
	// A sequence of the unique job IDs for the jobs that this job depends upon in the workflow.
	Dependencies []string `json:"dependencies"`
	// The number of the job.
	JobNumber int64 `json:"job_number,omitempty"`
	// The unique ID of the job.
	Id string `json:"id"`
	// The date and time the job started.
	StartedAt time.Time `json:"started_at"`
	// The name of the job.
	Name string `json:"name"`
	// The unique ID of the user.
	ApprovedBy string `json:"approved_by,omitempty"`
	// The project-slug for the job.
	ProjectSlug string `json:"project_slug"`
	// The current status of the job.
	Status interface{} `json:"status"`
	// The type of job.
	Type string `json:"type"`
	// The time when the job stopped.
	StoppedAt time.Time `json:"stopped_at,omitempty"`
	// The unique ID of the job.
	ApprovalRequestId string `json:"approval_request_id,omitempty"`
}