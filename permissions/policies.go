// This file was auto-generated from our API Definition.

package permissions

import (
	polytomicgo "github.com/polytomic/polytomic-go"
)

type CreatePolicyRequest struct {
	Name           string                      `json:"name" url:"name"`
	OrganizationId *string                     `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	PolicyActions  []*polytomicgo.PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty"`
}

type UpdatePolicyRequest struct {
	Name           string                      `json:"name" url:"name"`
	OrganizationId *string                     `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	PolicyActions  []*polytomicgo.PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty"`
}
