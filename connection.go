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
}

type UpdateConnectionMutation struct {
	OrganizationId string      `json:"organization_id,omitempty"`
	Name           string      `json:"name" validate:"required"`
	Configuration  interface{} `json:"configuration" validate:"required"`
}

type ConnectionApi struct {
	client *Client
}

func (a *ConnectionApi) Create(ctx context.Context, ws CreateConnectionMutation) (*Connection, error) {
	var connection Connection
	result := topLevelResult{Result: &connection}

	err := a.client.newRequest("/api/connections").
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (a *ConnectionApi) Get(ctx context.Context, connectionId uuid.UUID) (*Connection, error) {
	var connection Connection
	result := topLevelResult{Result: &connection}

	err := a.client.newRequest(fmt.Sprintf("/api/connections/%s", connectionId)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (a *ConnectionApi) Update(ctx context.Context, connectionId uuid.UUID, ws UpdateConnectionMutation) (*Connection, error) {
	var connection Connection
	result := topLevelResult{Result: &connection}

	err := a.client.newRequest(fmt.Sprintf("/api/connections/%s", connectionId)).
		Patch().
		BodyJSON(&ws).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

func (a *ConnectionApi) Delete(ctx context.Context, connectionId uuid.UUID) error {
	err := a.client.newRequest(fmt.Sprintf("/api/connections/%s", connectionId)).
		Delete().
		Fetch(ctx)
	if err != nil {
		return err
	}

	return nil
}
