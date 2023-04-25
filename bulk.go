package polytomic

import (
	"context"
	"fmt"
)

type BulkSyncRequest struct {
	Name                     string                 `json:"name"`
	OrganizationID           string                 `json:"organization_id,omitempty"`
	DestConnectionID         string                 `json:"destination_connection_id"`
	SourceConnectionID       string                 `json:"source_connection_id"`
	Mode                     string                 `json:"mode"`
	Schemas                  []string               `json:"schemas,omitempty"`
	Discover                 bool                   `json:"discover"`
	Active                   bool                   `json:"active"`
	Schedule                 Schedule               `json:"schedule"`
	DestinationConfiguration map[string]interface{} `json:"destination_configuration"`
	SourceConfiguration      map[string]interface{} `json:"source_configuration"`
	Policies                 []string               `json:"policies"`
}

type Schedule struct {
	Frequency  *string `json:"frequency,omitempty" tfsdk:"frequency" mapstructure:"frequency"`
	DayOfWeek  *string `json:"day_of_week,omitempty" tfsdk:"day_of_week" mapstructure:"day_of_week"`
	Hour       *string `json:"hour,omitempty" tfsdk:"hour" mapstructure:"hour"`
	Minute     *string `json:"minute,omitempty" tfsdk:"minute" mapstructure:"minute"`
	Month      *string `json:"month,omitempty" tfsdk:"month" mapstructure:"month"`
	DayOfMonth *string `json:"day_of_month,omitempty" tfsdk:"day_of_month" mapstructure:"day_of_month"`
}

type BulkSchema struct {
	ID           string  `json:"id" tfsdk:"id"`
	Name         string  `json:"name" tfsdk:"name"`
	Enabled      bool    `json:"enabled" tfsdk:"enabled"`
	PartitionKey string  `json:"partition_key" tfsdk:"partition_key"`
	Fields       []Field `json:"fields,omitempty" tfsdk:"fields"`
}

type Field struct {
	ID         string `json:"id" tfsdk:"id"`
	Enabled    bool   `json:"enabled" tfsdk:"enabled"`
	Obfuscated bool   `json:"obfuscate,omitempty" tfsdk:"obfuscated"`
}

type BulkSchemaUpdate struct {
	Schemas []BulkSchema `json:"schemas"`
}

type Schema struct {
	ID   string `json:"id" tfsdk:"id"`
	Name string `json:"name" tfsdk:"name"`
}

type Mode struct {
	ID          string `json:"id" tfsdk:"id"`
	Label       string `json:"label" tfsdk:"label"`
	Description string `json:"description" tfsdk:"description"`
}

type BulkSource struct {
	Schemas       []Schema    `json:"schemas" tfsdk:"schemas"`
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
	resp := Response{Data: &source}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/source/%s", connID)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &source, nil
}

func (b *BulkApi) GetSourceSchema(ctx context.Context, connID string, schemaID string) (*BulkSchema, error) {
	var schema BulkSchema
	resp := Response{Data: &schema}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/source/%s/schema/%s", connID, schemaID)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func (b *BulkApi) GetDestination(ctx context.Context, connID string) (*BulkDestination, error) {
	var dest BulkDestination
	resp := Response{Data: &dest}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/dest/%s", connID)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

type BulkSyncResponse struct {
	ID                       string                 `json:"id" tfsdk:"id" mapstructure:"id"`
	OrganizationID           string                 `json:"organization_id " tfsdk:"organization_id" mapstructure:"organization_id"`
	Name                     string                 `json:"name" tfsdk:"name" mapstructure:"name"`
	DestConnectionID         string                 `json:"destination_connection_id" tfsdk:"destination_connection_id" mapstructure:"destination_connection_id"`
	SourceConnectionID       string                 `json:"source_connection_id" tfsdk:"source_connection_id" mapstructure:"source_connection_id"`
	Mode                     string                 `json:"mode" tfsdk:"mode" mapstructure:"mode"`
	Discover                 bool                   `json:"discover" tfsdk:"discover" mapstructure:"discover"`
	Active                   bool                   `json:"active" tfsdk:"active" mapstructure:"active"`
	Schedule                 Schedule               `json:"schedule" tfsdk:"schedule" mapstructure:"schedule"`
	DestinationConfiguration map[string]interface{} `json:"destination_configuration" tfsdk:"destination_configuration" mapstructure:"destination_configuration"`
	SourceConfiguration      map[string]interface{} `json:"source_configuration" tfsdk:"source_configuration" mapstructure:"source_configuration"`
	Policies                 []string               `json:"policies,omitempty" tfsdk:"policies" mapstructure:"policies"`
}

func (b *BulkApi) CreateBulkSync(ctx context.Context, sync BulkSyncRequest) (*BulkSyncResponse, error) {
	var bulk BulkSyncResponse
	resp := Response{Data: &bulk}
	err := b.client.newRequest("/api/bulk/syncs").
		BodyJSON(&sync).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &bulk, nil
}

func (b *BulkApi) UpdateBulkSync(ctx context.Context, id string, sync BulkSyncRequest) (*BulkSyncResponse, error) {
	var bulk BulkSyncResponse
	resp := Response{Data: &bulk}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s", id)).
		Patch().
		BodyJSON(&sync).
		ToJSON(&resp).
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
	resp := Response{Data: &bulk}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s", id)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &bulk, nil
}

func (b *BulkApi) ListBulkSyncs(ctx context.Context) ([]BulkSyncResponse, error) {
	var bulks []BulkSyncResponse
	resp := Response{Data: &bulks}
	err := b.client.newRequest("/api/bulk/syncs").
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return bulks, nil
}

func (b *BulkApi) GetBulkSyncSchemas(ctx context.Context, id string) ([]BulkSchema, error) {
	var schemas []BulkSchema
	resp := Response{Data: &schemas}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas", id)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return schemas, nil
}

func (b *BulkApi) UpdateBulkSyncSchemas(ctx context.Context, id string, schemas []BulkSchema) ([]BulkSchema, error) {
	var resultSchemas []BulkSchema
	resp := Response{Data: &resultSchemas}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas", id)).
		BodyJSON(BulkSchemaUpdate{schemas}).
		ToJSON(&resp).
		Patch().
		Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return resultSchemas, nil
}

func (b *BulkApi) GetBulkSyncSchema(ctx context.Context, id, schemaID string) (*BulkSchema, error) {
	var schema BulkSchema
	resp := Response{Data: &schema}
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas/%s", id, schemaID)).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func (b *BulkApi) UpdateBulkSyncSchema(ctx context.Context, syncID, schemaID string, schema BulkSchema) (*BulkSchema, error) {
	var resultSchema BulkSchema
	err := b.client.newRequest(fmt.Sprintf("/api/bulk/syncs/%s/schemas/%s", syncID, schemaID)).
		Patch().
		BodyJSON(&schema).
		ToJSON(&resultSchema).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &resultSchema, nil
}
