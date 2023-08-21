package polytomic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Organization struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	SSODomain string    `json:"sso_domain,omitempty"`
	SSOOrgId  string    `json:"sso_org_id,omitempty"`
}

type OrganizationMutation struct {
	Name      string `json:"name,omitempty"`
	SSODomain string `json:"sso_domain,omitempty"`
	SSOOrgId  string `json:"sso_org_id,omitempty"`
}

type OrganizationApi struct {
	client *Client
}

func (a *OrganizationApi) Create(ctx context.Context, ws OrganizationMutation, opts ...requestOpts) (*Organization, error) {
	var workspace Organization
	resp := Response{Data: &workspace}
	err := a.client.newRequest("/api/organizations", opts...).
		BodyJSON(&ws).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *OrganizationApi) Get(ctx context.Context, id uuid.UUID, opts ...requestOpts) (*Organization, error) {
	var workspace Organization
	resp := Response{Data: &workspace}
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s", id.String()), opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *OrganizationApi) Update(ctx context.Context, id uuid.UUID, ws OrganizationMutation, opts ...requestOpts) (*Organization, error) {
	var workspace Organization
	resp := Response{Data: &workspace}
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s", id.String()), opts...).
		Patch().
		BodyJSON(&ws).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *OrganizationApi) Delete(ctx context.Context, id uuid.UUID, opts ...requestOpts) error {
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s", id.String()), opts...).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
