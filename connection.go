package polytomic

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Connection struct {
	OrganizationId string         `json:"organization_id"`
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Type           ConnectionType `json:"type"`
	Configuration  interface{}    `json:"configuration"`
	AuthURL        string         `json:"auth_url"`
	Status         string         `json:"status"`
	Policies       []string       `json:"policies,omitempty"`
}

type ConnectionType struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Operations []string `json:"operations"`
}

type CreateConnectionMutation struct {
	OrganizationId string      `json:"organization_id,omitempty"`
	Name           string      `json:"name" validate:"required"`
	Configuration  interface{} `json:"configuration" validate:"required"`
	Type           string      `json:"type" validate:"required"`
	Policies       []string    `json:"policies"`
}

type UpdateConnectionMutation struct {
	OrganizationId string      `json:"organization_id,omitempty"`
	Name           string      `json:"name" validate:"required"`
	Configuration  interface{} `json:"configuration" validate:"required"`
	Policies       []string    `json:"policies"`
}

type ConnectionMetaWrapper struct {
	Items         map[string]ConnectionMeta `json:"items"`
	RequiresOneOf []string                  `json:"requires_one_of"`
}

type ConnectionMeta struct {
	Items         []interface{} `json:"items"`
	RequiresOneOf []string      `json:"requires_one_of"`
	HasItems      bool          `json:"has_items"`
}

type ConnectionApi struct {
	client *Client
}

func (c *ConnectionApi) Create(ctx context.Context, ws CreateConnectionMutation) (*Connection, error) {
	var connection Connection
	resp := Response{Data: &connection}
	err := c.client.newRequest("/api/connections").
		BodyJSON(&ws).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (c *ConnectionApi) GetSource(ctx context.Context, connectionId uuid.UUID) (*ConnectionMetaWrapper, error) {
	var source ConnectionMetaWrapper
	resp := Response{Data: &source}
	err := c.client.newRequest(fmt.Sprintf("/api/connections/%s/source", connectionId)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return &source, nil
}

func (c *ConnectionApi) Get(ctx context.Context, connectionId uuid.UUID) (*Connection, error) {
	var connection Connection
	resp := Response{Data: &connection}
	err := c.client.newRequest(fmt.Sprintf("/api/connections/%s", connectionId)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (c *ConnectionApi) List(ctx context.Context) ([]Connection, error) {
	var connections []Connection
	resp := Response{Data: &connections}
	err := c.client.newRequest("/api/connections").
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

func (c *ConnectionApi) Update(ctx context.Context, connectionId uuid.UUID, ws UpdateConnectionMutation) (*Connection, error) {
	var connection Connection
	resp := Response{Data: &connection}
	err := c.client.newRequest(fmt.Sprintf("/api/connections/%s", connectionId)).
		Patch().
		BodyJSON(&ws).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (c *ConnectionApi) Delete(ctx context.Context, connectionId uuid.UUID) error {
	err := c.client.newRequest(fmt.Sprintf("/api/connections/%s", connectionId)).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
