package polytomic

import (
	"context"
)

type PermissionsApi struct {
	client *Client
}

type Role struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organization_id"`
	Name           string `json:"name"`
	System         bool   `json:"system"`
}

type RoleRequest struct {
	Name           string `json:"name" tfsdk:"name"`
	OrganizationID string `json:"organization_id,omitempty" tfsdk:"organization_id"`
}

func (p *PermissionsApi) CreateRole(ctx context.Context, r RoleRequest) (*Role, error) {
	var role Role
	resp := Response{Data: &role}
	err := p.client.newRequest("/api/permissions/roles").
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (p *PermissionsApi) GetRole(ctx context.Context, id string) (*Role, error) {
	var role Role
	resp := Response{Data: &role}
	err := p.client.newRequest("/api/permissions/roles/" + id).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (p *PermissionsApi) ListRoles(ctx context.Context) ([]Role, error) {
	var roles []Role
	resp := Response{Data: &roles}
	err := p.client.newRequest("/api/permissions/roles").
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (p *PermissionsApi) UpdateRole(ctx context.Context, id string, r RoleRequest) (*Role, error) {
	var role Role
	resp := Response{Data: &role}
	err := p.client.newRequest("/api/permissions/roles/" + id).
		Patch().
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (p *PermissionsApi) Delete(ctx context.Context, id string) error {
	return p.client.newRequest("/api/permissions/roles/" + id).
		Delete().
		Fetch(ctx)
}

type Policy struct {
	ID             string         `json:"id"`
	Name           string         `json:"name" `
	OrganizationID string         `json:"organization_id,omitempty"`
	System         bool           `json:"system"`
	PolicyActions  []PolicyAction `json:"policy_actions"`
}

type PolicyRequest struct {
	Name           string         `json:"name" tfsdk:"name"`
	OrganizationID string         `json:"organization_id,omitempty" tfsdk:"organization_id"`
	PolicyActions  []PolicyAction `json:"policy_actions" tfsdk:"policy_actions"`
}

type PolicyAction struct {
	Action  string   `json:"action" tfsdk:"action"`
	RoleIDs []string `json:"role_ids" tfsdk:"role_ids"`
}

func (p *PermissionsApi) CreatePolicy(ctx context.Context, r PolicyRequest) (*Policy, error) {
	var policy Policy
	resp := Response{Data: &policy}
	err := p.client.newRequest("/api/permissions/policies").
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (p *PermissionsApi) GetPolicy(ctx context.Context, id string) (*Policy, error) {
	var policy Policy
	resp := Response{Data: &policy}
	err := p.client.newRequest("/api/permissions/policies/" + id).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (p *PermissionsApi) ListPolicies(ctx context.Context) ([]Policy, error) {
	var policies []Policy
	resp := Response{Data: &policies}
	err := p.client.newRequest("/api/permissions/policies").
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (p *PermissionsApi) UpdatePolicy(ctx context.Context, id string, r PolicyRequest) (*Policy, error) {
	var policy Policy
	resp := Response{Data: &policy}
	err := p.client.newRequest("/api/permissions/policies/" + id).
		Patch().
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (p *PermissionsApi) DeletePolicy(ctx context.Context, id string) error {
	return p.client.newRequest("/api/permissions/policies/" + id).
		Delete().
		Fetch(ctx)
}
