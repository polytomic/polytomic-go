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

func (a *OrganizationApi) Create(ctx context.Context, ws OrganizationMutation) (*Organization, error) {
	var workspace Organization
	result := topLevelResult{Result: &workspace}

	err := a.client.newRequest("/api/organizations").
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *OrganizationApi) Get(ctx context.Context, id uuid.UUID) (*Organization, error) {
	var workspace Organization
	result := topLevelResult{Result: &workspace}

	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s", id.String())).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *OrganizationApi) Update(ctx context.Context, id uuid.UUID, ws OrganizationMutation) (*Organization, error) {
	var workspace Organization
	result := topLevelResult{Result: &workspace}

	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s", id.String())).
		Patch().
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *OrganizationApi) Delete(ctx context.Context, id uuid.UUID) error {
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s", id.String())).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
