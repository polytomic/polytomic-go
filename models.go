package polytomic

import (
	"context"
)

type ModelApi struct {
	client *Client
}

type Model struct {
	ID              string                 `json:"id" tfsdk:"id"`
	Name            string                 `json:"name" tfsdk:"name"`
	Type            string                 `json:"type" tfsdk:"type"`
	Version         int                    `json:"version" tfsdk:"version"`
	ConnectionID    string                 `json:"connection_id" tfsdk:"connection_id"`
	Configuration   map[string]interface{} `json:"configuration" tfsdk:"configuration"`
	Fields          []ModelField           `json:"fields" tfsdk:"fields"`
	Relations       []Relation             `json:"relations" tfsdk:"relations"`
	Identifier      string                 `json:"identifier" tfsdk:"identifier"`
	TrackingColumns []string               `json:"tracking_columns" tfsdk:"tracking_columns"`
}

type ModelField struct {
	Name        string `json:"name"`
	RemoteType  string `json:"remote_type"`
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Example     any    `json:"example"`
	Unique      bool   `json:"unique"`
	UserAdded   bool   `json:"user_added"`
}

type Relation struct {
	From       string     `json:"from" tfsdk:"from"`
	RelationTo RelationTo `json:"to" tfsdk:"to"`
}

type RelationTo struct {
	ModelID string `json:"model_id" tfsdk:"model_id"`
	Field   string `json:"field" tfsdk:"field"`
}

type ModelRequest struct {
	ConnectionID     string                 `json:"connection_id" tfsdk:"connection_id"`
	Name             string                 `json:"name" tfsdk:"name"`
	Configuration    map[string]interface{} `json:"configuration" tfsdk:"configuration"`
	Fields           []string               `json:"fields" tfsdk:"fields"`
	AdditionalFields []ModelFieldRequest    `json:"additional_fields" tfsdk:"additional_fields"`
	Relations        []Relation             `json:"relations" tfsdk:"relations"`
	Identifier       string                 `json:"identifier" tfsdk:"identifier"`
	TrackingColumns  []string               `json:"tracking_columns" tfsdk:"tracking_columns"`
}

type ModelFieldRequest struct {
	Name  string `json:"name" tfsdk:"name"`
	Label string `json:"label" tfsdk:"label"`
	Type  string `json:"type" tfsdk:"type"`
}

func (m *ModelApi) Create(ctx context.Context, r ModelRequest) (*Model, error) {
	var model Model
	resp := Response{Data: &model}
	err := m.client.newRequest("/api/models").
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *ModelApi) Get(ctx context.Context, id string) (*Model, error) {
	var model Model
	resp := Response{Data: &model}
	err := m.client.newRequest("/api/models/" + id).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *ModelApi) List(ctx context.Context) ([]Model, error) {
	var models []Model
	resp := Response{Data: &models}
	err := m.client.newRequest("/api/models").
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (m *ModelApi) Update(ctx context.Context, id string, r ModelRequest) (*Model, error) {
	var model Model
	resp := Response{Data: &model}
	err := m.client.newRequest("/api/models/" + id).
		Patch().
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (m *ModelApi) Delete(ctx context.Context, id string) error {
	return m.client.newRequest("/api/models/" + id).
		Delete().
		Fetch(ctx)
}
