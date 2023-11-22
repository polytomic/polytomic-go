package polytomic

import (
	"context"

	"github.com/google/uuid"
)

type CallerIdentity struct {
	ID             uuid.UUID `json:"id,omitempty"`
	Email          string    `json:"email"`
	Role           string    `json:"role"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Organization   string    `json:"organization_name"`
	IsUser         bool      `json:"is_user"`
	IsOrganization bool      `json:"is_organization"`
	IsPartner      bool      `json:"is_partner"`
	IsSystem       bool      `json:"is_system"`
}

type IdentityApi struct {
	client *Client
}

func (a *IdentityApi) Get(ctx context.Context, opts ...requestOpts) (*CallerIdentity, error) {
	var caller CallerIdentity
	resp := Response{Data: &caller}
	err := a.client.newRequest("/api/me", opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &caller, nil
}
