// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
)

type ModelsCreateRequest struct {
	Async *bool               `json:"-" url:"async,omitempty"`
	Body  *CreateModelRequest `json:"-" url:"-"`
}

func (m *ModelsCreateRequest) UnmarshalJSON(data []byte) error {
	body := new(CreateModelRequest)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	m.Body = body
	return nil
}

func (m *ModelsCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Body)
}

type ModelsGetRequest struct {
	Async *bool `json:"-" url:"async,omitempty"`
}

type ModelsGetEnrichmentSourceRequest struct {
	Params map[string][]string `json:"-" url:"params,omitempty"`
}

type GetEnrichmentInputFieldsRequest struct {
	Configuration *V2EnricherConfiguration `json:"configuration,omitempty" url:"configuration,omitempty"`
}

type ModelsPreviewRequest struct {
	Async *bool               `json:"-" url:"async,omitempty"`
	Body  *CreateModelRequest `json:"-" url:"-"`
}

func (m *ModelsPreviewRequest) UnmarshalJSON(data []byte) error {
	body := new(CreateModelRequest)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	m.Body = body
	return nil
}

func (m *ModelsPreviewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Body)
}

type ModelsRemoveRequest struct {
	Async *bool `json:"-" url:"async,omitempty"`
}

type ModelsSampleRequest struct {
	Async *bool `json:"-" url:"async,omitempty"`
}

type UpdateModelRequest struct {
	Async            *bool                     `json:"-" url:"async,omitempty"`
	AdditionalFields []*ModelModelFieldRequest `json:"additional_fields,omitempty" url:"additional_fields,omitempty"`
	Configuration    map[string]interface{}    `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId     string                    `json:"connection_id" url:"connection_id"`
	Enricher         *Enrichment               `json:"enricher,omitempty" url:"enricher,omitempty"`
	Fields           []string                  `json:"fields,omitempty" url:"fields,omitempty"`
	Identifier       *string                   `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels           []string                  `json:"labels,omitempty" url:"labels,omitempty"`
	Name             string                    `json:"name" url:"name"`
	OrganizationId   *string                   `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies         []string                  `json:"policies,omitempty" url:"policies,omitempty"`
	Refresh          *bool                     `json:"refresh,omitempty" url:"refresh,omitempty"`
	Relations        []*ModelRelation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns  []string                  `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
}
