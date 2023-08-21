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

func (p *PermissionsApi) CreateRole(ctx context.Context, r RoleRequest, opts ...requestOpts) (*Role, error) {
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

func (p *PermissionsApi) GetRole(ctx context.Context, id string, opts ...requestOpts) (*Role, error) {
	var role Role
	resp := Response{Data: &role}
	err := p.client.newRequest("/api/permissions/roles/"+id, opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (p *PermissionsApi) ListRoles(ctx context.Context, opts ...requestOpts) ([]Role, error) {
	var roles []Role
	resp := Response{Data: &roles}
	err := p.client.newRequest("/api/permissions/roles", opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (p *PermissionsApi) UpdateRole(ctx context.Context, id string, r RoleRequest, opts ...requestOpts) (*Role, error) {
	var role Role
	resp := Response{Data: &role}
	err := p.client.newRequest("/api/permissions/roles/"+id, opts...).
		Patch().
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (p *PermissionsApi) DeleteRole(ctx context.Context, id string, opts ...requestOpts) error {
	return p.client.newRequest("/api/permissions/roles/"+id, opts...).
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

func (p *PermissionsApi) CreatePolicy(ctx context.Context, r PolicyRequest, opts ...requestOpts) (*Policy, error) {
	var policy Policy
	resp := Response{Data: &policy}
	err := p.client.newRequest("/api/permissions/policies", opts...).
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (p *PermissionsApi) GetPolicy(ctx context.Context, id string, opts ...requestOpts) (*Policy, error) {
	var policy Policy
	resp := Response{Data: &policy}
	err := p.client.newRequest("/api/permissions/policies/"+id, opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (p *PermissionsApi) ListPolicies(ctx context.Context, opts ...requestOpts) ([]Policy, error) {
	var policies []Policy
	resp := Response{Data: &policies}
	err := p.client.newRequest("/api/permissions/policies", opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return policies, nil
}

func (p *PermissionsApi) UpdatePolicy(ctx context.Context, id string, r PolicyRequest, opts ...requestOpts) (*Policy, error) {
	var policy Policy
	resp := Response{Data: &policy}
	err := p.client.newRequest("/api/permissions/policies/"+id, opts...).
		Patch().
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &policy, nil
}

func (p *PermissionsApi) DeletePolicy(ctx context.Context, id string, opts ...requestOpts) error {
	return p.client.newRequest("/api/permissions/policies/"+id, opts...).
		Delete().
		Fetch(ctx)
}
