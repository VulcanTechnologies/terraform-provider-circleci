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

// CheckoutKey struct for CheckoutKey
type CheckoutKey struct {
	// A public SSH key.
	PublicKey string `json:"public-key"`
	// The type of checkout key. This may be either `deploy-key` or `github-user-key`.
	Type string `json:"type"`
	// An SSH key fingerprint.
	Fingerprint string `json:"fingerprint"`
	// A boolean value that indicates if this key is preferred.
	Preferred bool `json:"preferred"`
	// The date and time the checkout key was created.
	CreatedAt time.Time `json:"created-at"`
}
