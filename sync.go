package polytomic

import (
	"context"
)

type SyncApi struct {
	client *Client
}

type SyncRequest struct {
	Name           string      `json:"name"`
	Target         Target      `json:"target"`
	Mode           string      `json:"mode"`
	Fields         []SyncField `json:"fields"`
	OverrideFields []SyncField `json:"override_fields"`
	Filters        []Filter    `json:"filters"`
	FilterLogic    string      `json:"filter_logic"`
	Overrides      []Override  `json:"overrides"`
	Schedule       Schedule    `json:"schedule"`
	Identity       *Identity   `json:"identity"`
	SyncAllRecords bool        `json:"sync_all_records"`
}

type SyncResponse struct {
	ID             string      `json:"id" tfsdk:"id"`
	Name           string      `json:"name" tfsdk:"name"`
	Target         Target      `json:"target" tfsdk:"target"`
	Mode           string      `json:"mode" tfsdk:"mode"`
	Fields         []SyncField `json:"fields" tfsdk:"fields"`
	OverrideFields []SyncField `json:"override_fields" tfsdk:"override_fields"`
	Filters        []Filter    `json:"filters" tfsdk:"filters"`
	FilterLogic    string      `json:"filter_logic" tfsdk:"filter_logic"`
	Overrides      []Override  `json:"overrides" tfsdk:"overrides"`
	Schedule       Schedule    `json:"schedule" tfsdk:"schedule"`
	Identity       *Identity   `json:"identity" tfsdk:"identity"`
	SyncAllRecords bool        `json:"sync_all_records" tfsdk:"sync_all_records"`
}

type SyncField struct {
	Source        Source `json:"source" tfsdk:"source"`
	Target        string `json:"target" tfsdk:"target"`
	New           bool   `json:"new" tfsdk:"new"`
	OverrideValue string `json:"override_value" tfsdk:"override_value"`
	SyncMode      string `json:"sync_mode" tfsdk:"sync_mode"`
}

type Override struct {
	FieldID  string `json:"field_id" tfsdk:"field_id"`
	Function string `json:"function" tfsdk:"function"`
	Value    string `json:"value" tfsdk:"value"`
	Override string `json:"override" tfsdk:"override"`
}

type Filter struct {
	FieldID   string `json:"field_id" tfsdk:"field_id"`
	FieldType string `json:"field_type" tfsdk:"field_type"`
	Function  string `json:"function" tfsdk:"function"`
	Value     string `json:"value" tfsdk:"value"`
	Label     string `json:"label" tfsdk:"label"`
}

type Target struct {
	ConnectionID  string            `json:"connection_id" tfsdk:"connection_id"`
	Object        string            `json:"object" tfsdk:"object"`
	SearchValues  map[string]string `json:"search_values" tfsdk:"search_values"`
	Configuration map[string]string `json:"configuration" tfsdk:"configuration"`
	NewName       string            `json:"new_name" tfsdk:"new_name"`
	FilterLogic   string            `json:"filter_logic" tfsdk:"filter_logic"`
}

type Identity struct {
	Source            Source `json:"source" tfsdk:"source"`
	Target            string `json:"target" tfsdk:"target"`
	Function          string `json:"function" tfsdk:"function"`
	RemoteFieldTypeID string `json:"remote_field_type_id" tfsdk:"remote_field_type_id"`
	NewField          bool   `json:"new_field" tfsdk:"new_field"`
}

type Source struct {
	ModelID string `json:"model_id" tfsdk:"model_id"`
	Field   string `json:"field" tfsdk:"field"`
}

func (s *SyncApi) Create(ctx context.Context, r SyncRequest) (*SyncResponse, error) {
	var sync SyncResponse
	err := s.client.newRequest("/api/syncs").
		BodyJSON(&r).
		ToJSON(&sync).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &sync, nil
}

func (s *SyncApi) Get(ctx context.Context, id string) (*SyncResponse, error) {
	var sync SyncResponse
	err := s.client.newRequest("/api/syncs/" + id).
		ToJSON(&sync).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &sync, nil
}

func (s *SyncApi) Update(ctx context.Context, id string, r ModelRequest) (*SyncResponse, error) {
	var sync SyncResponse
	err := s.client.newRequest("/api/syncs/" + id).
		Patch().
		BodyJSON(&r).
		ToJSON(&sync).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &sync, nil
}

func (s *SyncApi) Delete(ctx context.Context, id string) error {
	return s.client.newRequest("/api/syncs/" + id).
		Delete().
		Fetch(ctx)
}
