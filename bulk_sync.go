// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
	fmt "fmt"
)

type CreateBulkSyncRequest struct {
	Active                     *bool                  `json:"active,omitempty" url:"active,omitempty"`
	AutomaticallyAddNewFields  *BulkDiscover          `json:"automatically_add_new_fields,omitempty" url:"automatically_add_new_fields,omitempty"`
	AutomaticallyAddNewObjects *BulkDiscover          `json:"automatically_add_new_objects,omitempty" url:"automatically_add_new_objects,omitempty"`
	DestinationConfiguration   map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty"`
	DestinationConnectionId    string                 `json:"destination_connection_id" url:"destination_connection_id"`
	DisableRecordTimestamps    *bool                  `json:"disable_record_timestamps,omitempty" url:"disable_record_timestamps,omitempty"`
	// DEPRECATED: Use automatically_add_new_objects/automatically_add_new_fields instead
	Discover       *bool         `json:"discover,omitempty" url:"discover,omitempty"`
	Mode           *SyncMode     `json:"mode,omitempty" url:"mode,omitempty"`
	Name           string        `json:"name" url:"name"`
	OrganizationId *string       `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string      `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *BulkSchedule `json:"schedule,omitempty" url:"schedule,omitempty"`
	// List of schemas to sync; if omitted, all schemas will be selected for syncing.
	Schemas             []*V2CreateBulkSyncRequestSchemasItem `json:"schemas,omitempty" url:"schemas,omitempty"`
	SourceConfiguration map[string]interface{}                `json:"source_configuration,omitempty" url:"source_configuration,omitempty"`
	SourceConnectionId  string                                `json:"source_connection_id" url:"source_connection_id"`
}

type BulkSyncGetRequest struct {
	RefreshSchemas *bool `json:"-" url:"refresh_schemas,omitempty"`
}

type BulkSyncGetSourceRequest struct {
	IncludeFields *bool `json:"-" url:"include_fields,omitempty"`
}

type BulkSyncRemoveRequest struct {
	RefreshSchemas *bool `json:"-" url:"refresh_schemas,omitempty"`
}

type StartBulkSyncRequest struct {
	Resync  *bool    `json:"resync,omitempty" url:"resync,omitempty"`
	Schemas []string `json:"schemas,omitempty" url:"schemas,omitempty"`
	Test    *bool    `json:"test,omitempty" url:"test,omitempty"`
}

type UpdateBulkSyncRequest struct {
	Active                     *bool                  `json:"active,omitempty" url:"active,omitempty"`
	AutomaticallyAddNewFields  *BulkDiscover          `json:"automatically_add_new_fields,omitempty" url:"automatically_add_new_fields,omitempty"`
	AutomaticallyAddNewObjects *BulkDiscover          `json:"automatically_add_new_objects,omitempty" url:"automatically_add_new_objects,omitempty"`
	DestinationConfiguration   map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty"`
	DestinationConnectionId    string                 `json:"destination_connection_id" url:"destination_connection_id"`
	DisableRecordTimestamps    *bool                  `json:"disable_record_timestamps,omitempty" url:"disable_record_timestamps,omitempty"`
	// DEPRECATED: Use automatically_add_new_objects/automatically_add_new_fields instead
	Discover       *bool         `json:"discover,omitempty" url:"discover,omitempty"`
	Mode           *SyncMode     `json:"mode,omitempty" url:"mode,omitempty"`
	Name           string        `json:"name" url:"name"`
	OrganizationId *string       `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string      `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *BulkSchedule `json:"schedule,omitempty" url:"schedule,omitempty"`
	// List of schemas to sync; if omitted, all schemas will be selected for syncing.
	Schemas             []*V2UpdateBulkSyncRequestSchemasItem `json:"schemas,omitempty" url:"schemas,omitempty"`
	SourceConfiguration map[string]interface{}                `json:"source_configuration,omitempty" url:"source_configuration,omitempty"`
	SourceConnectionId  string                                `json:"source_connection_id" url:"source_connection_id"`
}

type V2CreateBulkSyncRequestSchemasItem struct {
	String              string
	SchemaConfiguration *SchemaConfiguration
}

func NewV2CreateBulkSyncRequestSchemasItemFromString(value string) *V2CreateBulkSyncRequestSchemasItem {
	return &V2CreateBulkSyncRequestSchemasItem{String: value}
}

func NewV2CreateBulkSyncRequestSchemasItemFromSchemaConfiguration(value *SchemaConfiguration) *V2CreateBulkSyncRequestSchemasItem {
	return &V2CreateBulkSyncRequestSchemasItem{SchemaConfiguration: value}
}

func (v *V2CreateBulkSyncRequestSchemasItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		v.String = valueString
		return nil
	}
	valueSchemaConfiguration := new(SchemaConfiguration)
	if err := json.Unmarshal(data, &valueSchemaConfiguration); err == nil {
		v.SchemaConfiguration = valueSchemaConfiguration
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, v)
}

func (v V2CreateBulkSyncRequestSchemasItem) MarshalJSON() ([]byte, error) {
	if v.String != "" {
		return json.Marshal(v.String)
	}
	if v.SchemaConfiguration != nil {
		return json.Marshal(v.SchemaConfiguration)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V2CreateBulkSyncRequestSchemasItemVisitor interface {
	VisitString(string) error
	VisitSchemaConfiguration(*SchemaConfiguration) error
}

func (v *V2CreateBulkSyncRequestSchemasItem) Accept(visitor V2CreateBulkSyncRequestSchemasItemVisitor) error {
	if v.String != "" {
		return visitor.VisitString(v.String)
	}
	if v.SchemaConfiguration != nil {
		return visitor.VisitSchemaConfiguration(v.SchemaConfiguration)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V2UpdateBulkSyncRequestSchemasItem struct {
	String              string
	SchemaConfiguration *SchemaConfiguration
}

func NewV2UpdateBulkSyncRequestSchemasItemFromString(value string) *V2UpdateBulkSyncRequestSchemasItem {
	return &V2UpdateBulkSyncRequestSchemasItem{String: value}
}

func NewV2UpdateBulkSyncRequestSchemasItemFromSchemaConfiguration(value *SchemaConfiguration) *V2UpdateBulkSyncRequestSchemasItem {
	return &V2UpdateBulkSyncRequestSchemasItem{SchemaConfiguration: value}
}

func (v *V2UpdateBulkSyncRequestSchemasItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		v.String = valueString
		return nil
	}
	valueSchemaConfiguration := new(SchemaConfiguration)
	if err := json.Unmarshal(data, &valueSchemaConfiguration); err == nil {
		v.SchemaConfiguration = valueSchemaConfiguration
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, v)
}

func (v V2UpdateBulkSyncRequestSchemasItem) MarshalJSON() ([]byte, error) {
	if v.String != "" {
		return json.Marshal(v.String)
	}
	if v.SchemaConfiguration != nil {
		return json.Marshal(v.SchemaConfiguration)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V2UpdateBulkSyncRequestSchemasItemVisitor interface {
	VisitString(string) error
	VisitSchemaConfiguration(*SchemaConfiguration) error
}

func (v *V2UpdateBulkSyncRequestSchemasItem) Accept(visitor V2UpdateBulkSyncRequestSchemasItemVisitor) error {
	if v.String != "" {
		return visitor.VisitString(v.String)
	}
	if v.SchemaConfiguration != nil {
		return visitor.VisitSchemaConfiguration(v.SchemaConfiguration)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}
