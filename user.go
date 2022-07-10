package polytomic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id,omitempty"`
	WorkspaceId uuid.UUID `json:"workspace_id"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
}

type UserMutation struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserApi struct {
	client *Client
}

func (a *UserApi) Create(ctx context.Context, workspaceId uuid.UUID, ws UserMutation) (*User, error) {
	var user User
	result := topLevelResult{Result: &user}

	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s/users", workspaceId)).
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Get(ctx context.Context, workspaceId, userId uuid.UUID) (*User, error) {
	var user User
	result := topLevelResult{Result: &user}

	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s/users/%s", workspaceId, userId)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Update(ctx context.Context, workspaceId, userId uuid.UUID, ws UserMutation) (*User, error) {
	var user User
	result := topLevelResult{Result: &user}

	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s/users/%s", workspaceId, userId)).
		Patch().
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Delete(ctx context.Context, workspaceId, userId uuid.UUID) error {
	err := a.client.newRequest(fmt.Sprintf("/api/workspaces/%s/users/%s", workspaceId, userId)).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
