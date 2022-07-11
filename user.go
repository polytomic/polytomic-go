package polytomic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id,omitempty"`
	OrganizationId uuid.UUID `json:"organization_id"`
	Email          string    `json:"email"`
	Role           string    `json:"role"`
}

type UserMutation struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserApi struct {
	client *Client
}

func (a *UserApi) Create(ctx context.Context, organizationId uuid.UUID, ws UserMutation) (*User, error) {
	var user User
	result := topLevelResult{Result: &user}

	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users", organizationId)).
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Get(ctx context.Context, organizationId, userId uuid.UUID) (*User, error) {
	var user User
	result := topLevelResult{Result: &user}

	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users/%s", organizationId, userId)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Update(ctx context.Context, organizationId, userId uuid.UUID, ws UserMutation) (*User, error) {
	var user User
	result := topLevelResult{Result: &user}

	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users/%s", organizationId, userId)).
		Patch().
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Delete(ctx context.Context, organizationId, userId uuid.UUID) error {
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users/%s", organizationId, userId)).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
