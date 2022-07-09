package polytomic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Workspace struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	SSODomain string    `json:"sso_domain,omitempty"`
	SSOOrgId  string    `json:"sso_org_id,omitempty"`
}

type WorkspaceMutation struct {
	Name      string `json:"name,omitempty"`
	SSODomain string `json:"sso_domain,omitempty"`
	SSOOrgId  string `json:"sso_org_id,omitempty"`
}

type WorkspaceApi struct {
	client *Client
}

func (a *WorkspaceApi) Create(ctx context.Context, ws WorkspaceMutation) (*Workspace, error) {
	var workspace Workspace
	result := topLevelResult{Result: &workspace}

	err := a.client.newRequest("/api/workspaces").
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *WorkspaceApi) Get(ctx context.Context, id uuid.UUID) (*Workspace, error) {
	var workspace Workspace
	result := topLevelResult{Result: &workspace}

	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s", id.String())).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *WorkspaceApi) Update(ctx context.Context, id uuid.UUID, ws WorkspaceMutation) (*Workspace, error) {
	var workspace Workspace
	result := topLevelResult{Result: &workspace}

	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s", id.String())).
		Patch().
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (a *WorkspaceApi) Delete(ctx context.Context, id uuid.UUID) error {
	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s", id.String())).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}

type topLevelResult struct {
	Result interface{} `json:"result"`
}
