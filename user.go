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

func (a *UserApi) Create(ctx context.Context, organizationId uuid.UUID, ws UserMutation, opts ...requestOpts) (*User, error) {
	var user User
	resp := Response{Data: &user}
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users", organizationId), opts...).
		BodyJSON(&ws).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Get(ctx context.Context, organizationId, userId uuid.UUID, opts ...requestOpts) (*User, error) {
	var user User
	resp := Response{Data: &user}
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users/%s", organizationId, userId), opts...).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Update(ctx context.Context, organizationId, userId uuid.UUID, ws UserMutation, opts ...requestOpts) (*User, error) {
	var user User
	resp := Response{Data: &user}
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users/%s", organizationId, userId), opts...).
		Patch().
		BodyJSON(&ws).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (a *UserApi) Delete(ctx context.Context, organizationId, userId uuid.UUID, opts ...requestOpts) error {
	err := a.client.newRequest(fmt.Sprintf("/api/organizations/%s/users/%s", organizationId, userId), opts...).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
