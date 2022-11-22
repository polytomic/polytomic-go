package polytomic

import (
	"context"
	"fmt"
)

type BulkSyncRequest struct {
	Name                     string                 `json:"name"`
	DestConnectionID         string                 `json:"destination_connection_id"`
	SourceConnectionID       string                 `json:"source_connection_id"`
	Mode                     string                 `json:"mode"`
	Schemas                  []string               `json:"schemas,omitempty"`
	Discover                 bool                   `json:"discover"`
	Active                   bool                   `json:"active"`
	Schedule                 Schedule               `json:"schedule"`
	DestinationConfiguration map[string]interface{} `json:"destination_configuration"`
	SourceConfiguration      map[string]interface{} `json:"source_configuration"`
}

type Schedule struct {
	Frequency  *string `json:"frequency,omitempty" tfsdk:"frequency"`
	DayOfWeek  *string `json:"day_of_week,omitempty" tfsdk:"day_of_week"`
	Hour       *string `json:"hour,omitempty" tfsdk:"hour"`
	Minute     *string `json:"minut,omitempty" tfsdk:"minute"`
	Month      *string `json:"month,omitempty" tfsdk:"month"`
	DayOfMonth *string `json:"day_of_month,omitempty" tfsdk:"day_of_month"`
}

type BulkSchema struct {
	ID      string `json:"id" tfsdk:"id"`
	Name    string `json:"name" tfsdk:"name"`
	Enabled bool   `json:"enabled" tfsdk:"enabled"`
}

type Schema struct {
	ID    string   `json:"id" tfsdk:"id"`
	Name  string   `json:"name" tfsdk:"name"`
	Modes []string `json:"modes" tfsdk:"modes"`
}

type Mode struct {
	ID          string `json:"id" tfsdk:"id"`
	Label       string `json:"label" tfsdk:"label"`
	Description string `json:"description" tfsdk:"description"`
}

type BulkSource struct {
	Schemas       []string    `json:"schemas" tfsdk:"schemas"`
	Configuration interface{} `json:"configuration" tfsdk:"configuration"`
}

type BulkDestination struct {
	Modes         []Mode      `json:"modes" tfsdk:"modes"`
	Configuration interface{} `json:"configuration" tfsdk:"configuration"`
}

type BulkApi struct {
	client *Client
}

func (b *BulkApi) GetSource(ctx context.Context, connID string) (*BulkSource, error) {
	var source BulkSource
	result := topLevelResult{Result: &source}

	err := b.client.newRequest(fmt.Sprintf("/api/bulk/source/%s", connID)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (b *BulkApi) GetDestination(ctx context.Context, connID string) (*BulkDestination, error) {
	var dest BulkDestination
	result := topLevelResult{Result: &dest}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/dest/%s", connID)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BulkSyncResponse struct {
	ID                 string   `json:"id" tfsdk:"id"`
	Name               string   `json:"name" tfsdk:"name"`
	DestConnectionID   string   `json:"destination_connection_id" tfsdk:"destination_connection_id"`
	SourceConnectionID string   `json:"source_connection_id" tfsdk:"source_connection_id"`
	Mode               string   `json:"mode" tfsdk:"mode"`
	Discover           bool     `json:"discover" tfsdk:"discover"`
	Active             bool     `json:"active" tfsdk:"active"`
	Schedule           Schedule `json:"schedule" tfsdk:"schedule"`
}

func (b *BulkApi) CreateBulkSync(ctx context.Context, sync BulkSyncRequest) (*BulkSyncResponse, error) {
	var bulk BulkSyncResponse
	result := topLevelResult{Result: &bulk}
	err := b.client.newRequest("/api/bulk/syncs").
		BodyJSON(&sync).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &bulk, nil
}

func (b *BulkApi) UpdateBulkSync(ctx context.Context, id string, sync BulkSyncRequest) (*BulkSyncResponse, error) {
	var bulk BulkSyncResponse
	result := topLevelResult{Result: &bulk}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s", id)).
		Patch().
		BodyJSON(&sync).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &bulk, nil
}

func (b *BulkApi) DeleteBulkSync(ctx context.Context, id string) error {
	return b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s", id)).
		Delete().
		Fetch(ctx)
}

func (b *BulkApi) GetBulkSync(ctx context.Context, id string) (*BulkSyncResponse, error) {
	var bulk BulkSyncResponse
	result := topLevelResult{Result: &bulk}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s", id)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &bulk, nil
}

func (b *BulkApi) GetBulkSyncSchemas(ctx context.Context, id string) ([]BulkSchema, error) {
	var schemas []BulkSchema
	result := topLevelResult{Result: &schemas}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas", id)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return schemas, nil
}

func (b *BulkApi) UpdateBulkSyncSchemas(ctx context.Context, id string, schemas []BulkSchema) error {
	return b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas", id)).
		Patch().
		BodyJSON(&schemas).
		Fetch(ctx)
}

func (b *BulkApi) GetBulkSyncSchema(ctx context.Context, id, schemaID string, schema BulkSchema) (*BulkSchema, error) {
	var result BulkSchema
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas/%s", id, schemaID)).
		ToJSON(&result).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
