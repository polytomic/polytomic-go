// This file was auto-generated from our API Definition.

package polytomic

type V2CreateModelRequest struct {
	AdditionalFields []*V2ModelFieldRequest `json:"additional_fields,omitempty" url:"additional_fields,omitempty"`
	Configuration    map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId     string                 `json:"connection_id" url:"connection_id"`
	Fields           []string               `json:"fields,omitempty" url:"fields,omitempty"`
	Identifier       *string                `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels           []string               `json:"labels,omitempty" url:"labels,omitempty"`
	Name             string                 `json:"name" url:"name"`
	OrganizationId   *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies         []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Relations        []*V2Relation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns  []string               `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
}

type V2UpdateModelRequest struct {
	AdditionalFields []*V2ModelFieldRequest `json:"additional_fields,omitempty" url:"additional_fields,omitempty"`
	Configuration    map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId     string                 `json:"connection_id" url:"connection_id"`
	Fields           []string               `json:"fields,omitempty" url:"fields,omitempty"`
	Identifier       *string                `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels           []string               `json:"labels,omitempty" url:"labels,omitempty"`
	Name             string                 `json:"name" url:"name"`
	OrganizationId   *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies         []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Relations        []*V2Relation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns  []string               `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
}
