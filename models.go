// This file was auto-generated from our API Definition.

package polytomic

type CreateModelRequest struct {
	Async            *bool                     `json:"-" url:"async,omitempty"`
	AdditionalFields []*ModelModelFieldRequest `json:"additional_fields,omitempty" url:"additional_fields,omitempty"`
	Configuration    map[string]interface{}    `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId     string                    `json:"connection_id" url:"connection_id"`
	Fields           []string                  `json:"fields,omitempty" url:"fields,omitempty"`
	Identifier       *string                   `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels           []string                  `json:"labels,omitempty" url:"labels,omitempty"`
	Name             string                    `json:"name" url:"name"`
	OrganizationId   *string                   `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies         []string                  `json:"policies,omitempty" url:"policies,omitempty"`
	Relations        []*ModelRelation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns  []string                  `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
}

type UpdateModelRequest struct {
	Async            *bool                     `json:"-" url:"async,omitempty"`
	AdditionalFields []*ModelModelFieldRequest `json:"additional_fields,omitempty" url:"additional_fields,omitempty"`
	Configuration    map[string]interface{}    `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId     string                    `json:"connection_id" url:"connection_id"`
	Fields           []string                  `json:"fields,omitempty" url:"fields,omitempty"`
	Identifier       *string                   `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels           []string                  `json:"labels,omitempty" url:"labels,omitempty"`
	Name             string                    `json:"name" url:"name"`
	OrganizationId   *string                   `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies         []string                  `json:"policies,omitempty" url:"policies,omitempty"`
	Relations        []*ModelRelation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns  []string                  `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
}
