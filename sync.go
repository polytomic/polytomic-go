package polytomic

import (
	"context"
)

type SyncApi struct {
	client *Client
}

type SyncRequest struct {
	Name           string      `json:"name"`
	OrganizationID string      `json:"organization_id,omitempty"`
	Target         Target      `json:"target"`
	Mode           string      `json:"mode"`
	Fields         []SyncField `json:"fields"`
	OverrideFields []SyncField `json:"override_fields,omitempty"`
	Filters        []Filter    `json:"filters,omitempty"`
	FilterLogic    string      `json:"filter_logic,omitempty"`
	Overrides      []Override  `json:"overrides,omitempty"`
	Schedule       Schedule    `json:"schedule"`
	Identity       *Identity   `json:"identity,omitempty"`
	SyncAllRecords bool        `json:"sync_all_records"`
	Policies       []string    `json:"policies"`
	Active         bool        `json:"active"`
}

type SyncResponse struct {
	ID             string      `json:"id" tfsdk:"id" mapstructure:"id"`
	OrganizationID string      `json:"organization_id" tfsdk:"organization_id" mapstructure:"organization_id"`
	Name           string      `json:"name" tfsdk:"name" mapstructure:"name"`
	Target         Target      `json:"target" tfsdk:"target" mapstructure:"target"`
	Mode           string      `json:"mode" tfsdk:"mode" mapstructure:"mode"`
	Fields         []SyncField `json:"fields" tfsdk:"fields" mapstructure:"fields"`
	OverrideFields []SyncField `json:"override_fields" tfsdk:"override_fields" mapstructure:"override_fields"`
	Filters        []Filter    `json:"filters" tfsdk:"filters" mapstructure:"filters"`
	FilterLogic    string      `json:"filter_logic,omitempty" tfsdk:"filter_logic" mapstructure:"filter_logic"`
	Overrides      []Override  `json:"overrides" tfsdk:"overrides" mapstructure:"overrides"`
	Schedule       Schedule    `json:"schedule" tfsdk:"schedule" mapstructure:"schedule"`
	Identity       *Identity   `json:"identity" tfsdk:"identity" mapstructure:"identity"`
	SyncAllRecords bool        `json:"sync_all_records" tfsdk:"sync_all_records" mapstructure:"sync_all_records"`
	Policies       []string    `json:"policies,omitempty" tfsdk:"policies" mapstructure:"policies"`
	Active         bool        `json:"active" tfsdk:"active" mapstructure:"active"`
}

type SyncField struct {
	Source        Source  `json:"source,omitempty" tfsdk:"source" mapstructure:"source"`
	Target        string  `json:"target,omitempty" tfsdk:"target" mapstructure:"target"`
	New           *bool   `json:"new,omitempty" tfsdk:"new" mapstructure:"new"`
	OverrideValue *string `json:"override_value,omitempty" tfsdk:"override_value" mapstructure:"override_value"`
	SyncMode      *string `json:"sync_mode,omitempty" tfsdk:"sync_mode" mapstructure:"sync_mode"`
}

type Override struct {
	FieldID  string      `json:"field_id" tfsdk:"field_id" mapstructure:"field_id"`
	Function string      `json:"function" tfsdk:"function" mapstructure:"function"`
	Value    interface{} `json:"value" tfsdk:"value" mapstructure:"value"`
	Override interface{} `json:"override" tfsdk:"override" mapstructure:"override"`
}

type Filter struct {
	FieldID   string      `json:"field_id" tfsdk:"field_id" mapstructure:"field_id"`
	FieldType string      `json:"field_type" tfsdk:"field_type" mapstructure:"field_type"`
	Function  string      `json:"function" tfsdk:"function" mapstructure:"function"`
	Value     interface{} `json:"value,omitempty" tfsdk:"value" mapstructure:"value"`
	Label     string      `json:"label" tfsdk:"label" mapstructure:"label"`
}

type Target struct {
	ConnectionID  string                 `json:"connection_id" tfsdk:"connection_id" mapstructure:"connection_id"`
	Object        *string                `json:"object" tfsdk:"object" mapstructure:"object"`
	SearchValues  map[string]interface{} `json:"search_values,omitempty" tfsdk:"search_values" mapstructure:"search_values,omitempty"`
	Configuration map[string]interface{} `json:"configuration,omitempty" tfsdk:"configuration" mapstructure:"configuration,omitempty"`
	NewName       *string                `json:"new_name,omitempty" tfsdk:"new_name" mapstructure:"new_name"`
	FilterLogic   *string                `json:"filter_logic,omitempty" tfsdk:"filter_logic" mapstructure:"filter_logic"`
}

type Identity struct {
	Source            Source `json:"source" tfsdk:"source" mapstructure:"source"`
	Target            string `json:"target" tfsdk:"target" mapstructure:"target"`
	Function          string `json:"function" tfsdk:"function" mapstructure:"function"`
	RemoteFieldTypeID string `json:"remote_field_type_id" tfsdk:"remote_field_type_id" mapstructure:"remote_field_type_id"`
	NewField          bool   `json:"new_field" tfsdk:"new_field" mapstructure:"new_field"`
}

type Source struct {
	ModelID string `json:"model_id" tfsdk:"model_id" mapstructure:"model_id"`
	Field   string `json:"field" tfsdk:"field" mapstructure:"field"`
}

func (s *SyncApi) Create(ctx context.Context, r SyncRequest) (*SyncResponse, error) {
	var sync SyncResponse
	resp := Response{Data: &sync}
	err := s.client.newRequest("/api/syncs").
		BodyJSON(&r).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &sync, nil
}

func (s *SyncApi) Get(ctx context.Context, id string) (*SyncResponse, error) {
	var sync SyncResponse
	resp := Response{Data: &sync}
	err := s.client.newRequest("/api/syncs/" + id).
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return &sync, nil
}

func (s *SyncApi) List(ctx context.Context) ([]SyncResponse, error) {
	var syncs []SyncResponse
	resp := Response{Data: &syncs}
	err := s.client.newRequest("/api/syncs").
		ToJSON(&resp).
		Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return syncs, nil
}

func (s *SyncApi) Update(ctx context.Context, id string, r SyncRequest) (*SyncResponse, error) {
	var sync SyncResponse
	resp := Response{Data: &sync}
	err := s.client.newRequest("/api/syncs/" + id).
		Patch().
		BodyJSON(&r).
		ToJSON(&resp).
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
