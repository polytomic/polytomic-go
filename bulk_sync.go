// This file was auto-generated from our API Definition.

package polytomic

type V2CreateBulkSyncRequest struct {
	Active                     *bool                  `json:"active,omitempty" url:"active,omitempty"`
	AutomaticallyAddNewFields  *bool                  `json:"automatically_add_new_fields,omitempty" url:"automatically_add_new_fields,omitempty"`
	AutomaticallyAddNewObjects *bool                  `json:"automatically_add_new_objects,omitempty" url:"automatically_add_new_objects,omitempty"`
	DestinationConfiguration   map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty"`
	DestinationConnectionId    string                 `json:"destination_connection_id" url:"destination_connection_id"`
	DisableRecordTimestamps    *bool                  `json:"disable_record_timestamps,omitempty" url:"disable_record_timestamps,omitempty"`
	// DEPRECATED: Use automatically_add_new_objects/automatically_add_new_fields instead
	Discover            *bool                  `json:"discover,omitempty" url:"discover,omitempty"`
	Mode                string                 `json:"mode" url:"mode"`
	Name                string                 `json:"name" url:"name"`
	OrganizationId      *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies            []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule            *BulkSchedule          `json:"schedule,omitempty" url:"schedule,omitempty"`
	Schemas             []string               `json:"schemas,omitempty" url:"schemas,omitempty"`
	SourceConfiguration map[string]interface{} `json:"source_configuration,omitempty" url:"source_configuration,omitempty"`
	SourceConnectionId  string                 `json:"source_connection_id" url:"source_connection_id"`
}

type BulkSyncGetRequest struct {
	RefreshSchemas *bool `json:"-" url:"refresh_schemas,omitempty"`
}

type BulkSyncGetSourceRequest struct {
	RefreshSchemas *bool `json:"-" url:"refresh_schemas,omitempty"`
	IncludeFields  *bool `json:"-" url:"include_fields,omitempty"`
}

type BulkSyncRemoveRequest struct {
	RefreshSchemas *bool `json:"-" url:"refresh_schemas,omitempty"`
}

type V2UpdateBulkSyncRequest struct {
	Active                   *bool                  `json:"active,omitempty" url:"active,omitempty"`
	DestinationConfiguration map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty"`
	DestinationConnectionId  string                 `json:"destination_connection_id" url:"destination_connection_id"`
	Discover                 *bool                  `json:"discover,omitempty" url:"discover,omitempty"`
	Mode                     string                 `json:"mode" url:"mode"`
	Name                     string                 `json:"name" url:"name"`
	OrganizationId           *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies                 []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule                 *BulkSchedule          `json:"schedule,omitempty" url:"schedule,omitempty"`
	SourceConfiguration      map[string]interface{} `json:"source_configuration,omitempty" url:"source_configuration,omitempty"`
	SourceConnectionId       string                 `json:"source_connection_id" url:"source_connection_id"`
}
