// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/polytomic/polytomic-go/core"
	time "time"
)

type BulkSchedule struct {
	DayOfMonth *string `json:"day_of_month,omitempty" url:"day_of_month,omitempty"`
	DayOfWeek  *string `json:"day_of_week,omitempty" url:"day_of_week,omitempty"`
	Frequency  string  `json:"frequency" url:"frequency"`
	Hour       *string `json:"hour,omitempty" url:"hour,omitempty"`
	Minute     *string `json:"minute,omitempty" url:"minute,omitempty"`
	Month      *string `json:"month,omitempty" url:"month,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSchedule) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSchedule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSchedule(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSchedule) String() string {
	if len(b._rawJSON) > 0 {
		if value, err := core.StringifyJSON(b._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

type LabelLabel = map[string]interface{}

type RestErrResponse struct {
	// Application-specific error code.
	Code *int `json:"code,omitempty" url:"code,omitempty"`
	// Application context.
	Context map[string]interface{} `json:"context,omitempty" url:"context,omitempty"`
	// Error message.
	Error *string `json:"error,omitempty" url:"error,omitempty"`
	// Status text.
	Status *string `json:"status,omitempty" url:"status,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RestErrResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RestErrResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RestErrResponse(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RestErrResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

type SchemaAssociation struct {
	Id              *string  `json:"id,omitempty" url:"id,omitempty"`
	Name            *string  `json:"name,omitempty" url:"name,omitempty"`
	ReferenceTo     []string `json:"reference_to,omitempty" url:"reference_to,omitempty"`
	ReferencedField *string  `json:"referenced_field,omitempty" url:"referenced_field,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SchemaAssociation) UnmarshalJSON(data []byte) error {
	type unmarshaler SchemaAssociation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SchemaAssociation(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SchemaAssociation) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type SyncDestinationProperties struct {
	DoesNotReportOperationCounts *bool   `json:"does_not_report_operation_counts,omitempty" url:"does_not_report_operation_counts,omitempty"`
	NewTargetLabel               *string `json:"new_target_label,omitempty" url:"new_target_label,omitempty"`
	OptionalTargetMappings       *bool   `json:"optional_target_mappings,omitempty" url:"optional_target_mappings,omitempty"`
	PrimaryMetadataObject        *string `json:"primary_metadata_object,omitempty" url:"primary_metadata_object,omitempty"`
	RequiresConfiguration        *bool   `json:"requires_configuration,omitempty" url:"requires_configuration,omitempty"`
	SupportsFieldCreation        *bool   `json:"supports_field_creation,omitempty" url:"supports_field_creation,omitempty"`
	SupportsFieldTypeSelection   *bool   `json:"supports_field_type_selection,omitempty" url:"supports_field_type_selection,omitempty"`
	SupportsTargetFilters        *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`
	TargetCreator                *bool   `json:"target_creator,omitempty" url:"target_creator,omitempty"`
	UseFieldNamesAsLabels        *bool   `json:"use_field_names_as_labels,omitempty" url:"use_field_names_as_labels,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncDestinationProperties) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncDestinationProperties
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncDestinationProperties(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncDestinationProperties) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

type V2ActivateSyncEnvelope struct {
	Data *V2ActivateSyncOutput `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ActivateSyncEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ActivateSyncEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ActivateSyncEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ActivateSyncEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ActivateSyncInput struct {
	Active bool `json:"active" url:"active"`

	_rawJSON json.RawMessage
}

func (v *V2ActivateSyncInput) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ActivateSyncInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ActivateSyncInput(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ActivateSyncInput) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ActivateSyncOutput struct {
	Active *bool   `json:"active,omitempty" url:"active,omitempty"`
	Id     *string `json:"id,omitempty" url:"id,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ActivateSyncOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ActivateSyncOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ActivateSyncOutput(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ActivateSyncOutput) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ApiKeyResponse struct {
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ApiKeyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ApiKeyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ApiKeyResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ApiKeyResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ApiKeyResponseEnvelope struct {
	Data *V2ApiKeyResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ApiKeyResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ApiKeyResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ApiKeyResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ApiKeyResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2BulkSyncDest struct {
	Configuration map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Modes         []*V2SupportedMode     `json:"modes,omitempty" url:"modes,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2BulkSyncDest) UnmarshalJSON(data []byte) error {
	type unmarshaler V2BulkSyncDest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2BulkSyncDest(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2BulkSyncDest) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2BulkSyncDestEnvelope struct {
	Data *V2BulkSyncDest `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2BulkSyncDestEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2BulkSyncDestEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2BulkSyncDestEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2BulkSyncDestEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2BulkSyncListEnvelope struct {
	Data []*V2BulkSyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2BulkSyncListEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2BulkSyncListEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2BulkSyncListEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2BulkSyncListEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2BulkSyncResponse struct {
	Active                   *bool                  `json:"active,omitempty" url:"active,omitempty"`
	DestinationConfiguration map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty"`
	DestinationConnectionId  *string                `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	Discover                 *bool                  `json:"discover,omitempty" url:"discover,omitempty"`
	Id                       *string                `json:"id,omitempty" url:"id,omitempty"`
	Mode                     *string                `json:"mode,omitempty" url:"mode,omitempty"`
	Name                     *string                `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId           *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies                 []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule                 *BulkSchedule          `json:"schedule,omitempty" url:"schedule,omitempty"`
	SourceConfiguration      map[string]interface{} `json:"source_configuration,omitempty" url:"source_configuration,omitempty"`
	SourceConnectionId       *string                `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2BulkSyncResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2BulkSyncResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2BulkSyncResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2BulkSyncResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2BulkSyncResponseEnvelope struct {
	Data *V2BulkSyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2BulkSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2BulkSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2BulkSyncResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2BulkSyncResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConfigurationValue struct {
	Items []interface{} `json:"items,omitempty" url:"items,omitempty"`
	Type  *string       `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConfigurationValue) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConfigurationValue
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConfigurationValue(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConfigurationValue) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionListResponseEnvelope struct {
	Data []*V2ConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionListResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionListResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionListResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionListResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionMeta struct {
	HasItems      *bool         `json:"has_items,omitempty" url:"has_items,omitempty"`
	Items         []interface{} `json:"items,omitempty" url:"items,omitempty"`
	RequiresOneOf []string      `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionMeta(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionMeta) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionMetaResponse struct {
	Configuration map[string]*V2ConfigurationValue `json:"configuration,omitempty" url:"configuration,omitempty"`
	Items         map[string]*V2ConnectionMeta     `json:"items,omitempty" url:"items,omitempty"`
	RequiresOneOf []string                         `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionMetaResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionMetaResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionMetaResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionMetaResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionResponseEnvelope struct {
	Data *V2ConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionResponseSchema struct {
	Configuration  map[string]interface{}  `json:"configuration,omitempty" url:"configuration,omitempty"`
	Id             *string                 `json:"id,omitempty" url:"id,omitempty"`
	Name           *string                 `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string                 `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string                `json:"policies,omitempty" url:"policies,omitempty"`
	Status         *string                 `json:"status,omitempty" url:"status,omitempty"`
	StatusError    *string                 `json:"status_error,omitempty" url:"status_error,omitempty"`
	Type           *V2ConnectionTypeSchema `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionResponseSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionResponseSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionResponseSchema(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionResponseSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionType struct {
	EnvConfig map[string]interface{} `json:"envConfig,omitempty" url:"envConfig,omitempty"`
	Id        *string                `json:"id,omitempty" url:"id,omitempty"`
	Name      *string                `json:"name,omitempty" url:"name,omitempty"`
	UseOauth  *bool                  `json:"use_oauth,omitempty" url:"use_oauth,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionType) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionType(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionType) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionTypeResponseEnvelope struct {
	Data []*V2ConnectionType `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionTypeResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionTypeResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionTypeResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionTypeResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ConnectionTypeSchema struct {
	Id   *string `json:"id,omitempty" url:"id,omitempty"`
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ConnectionTypeSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ConnectionTypeSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ConnectionTypeSchema(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ConnectionTypeSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2CreateConnectionResponseEnvelope struct {
	Data *V2CreateConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2CreateConnectionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2CreateConnectionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2CreateConnectionResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2CreateConnectionResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2CreateConnectionResponseSchema struct {
	// Code to enter in order to complete connection authentication.
	AuthCode *string `json:"auth_code,omitempty" url:"auth_code,omitempty"`
	// URL to visit to complete connection authentication.
	AuthUrl        *string                 `json:"auth_url,omitempty" url:"auth_url,omitempty"`
	Configuration  map[string]interface{}  `json:"configuration,omitempty" url:"configuration,omitempty"`
	Id             *string                 `json:"id,omitempty" url:"id,omitempty"`
	Name           *string                 `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string                 `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string                `json:"policies,omitempty" url:"policies,omitempty"`
	Status         *string                 `json:"status,omitempty" url:"status,omitempty"`
	StatusError    *string                 `json:"status_error,omitempty" url:"status_error,omitempty"`
	Type           *V2ConnectionTypeSchema `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2CreateConnectionResponseSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler V2CreateConnectionResponseSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2CreateConnectionResponseSchema(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2CreateConnectionResponseSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Event struct {
	CreatedAt      *time.Time  `json:"created_at,omitempty" url:"created_at,omitempty"`
	Event          interface{} `json:"event,omitempty" url:"event,omitempty"`
	Id             *string     `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string     `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Type           *string     `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Event) UnmarshalJSON(data []byte) error {
	type embed V2Event
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V2Event(unmarshaler.embed)
	v.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Event) MarshalJSON() ([]byte, error) {
	type embed V2Event
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*v),
		CreatedAt: core.NewOptionalDateTime(v.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V2Event) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2EventTypesEnvelope struct {
	Data []string `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2EventTypesEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2EventTypesEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2EventTypesEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2EventTypesEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2EventsEnvelope struct {
	Data []*V2Event `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2EventsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2EventsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2EventsEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2EventsEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ExecutionCounts struct {
	Error  *int `json:"error,omitempty" url:"error,omitempty"`
	Insert *int `json:"insert,omitempty" url:"insert,omitempty"`
	Total  *int `json:"total,omitempty" url:"total,omitempty"`
	Update *int `json:"update,omitempty" url:"update,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ExecutionCounts) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ExecutionCounts
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ExecutionCounts(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ExecutionCounts) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Filter struct {
	FieldId   string      `json:"field_id" url:"field_id"`
	FieldType string      `json:"field_type" url:"field_type"`
	Function  string      `json:"function" url:"function"`
	Label     *string     `json:"label,omitempty" url:"label,omitempty"`
	Value     interface{} `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Filter) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Filter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Filter(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Filter) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2GetConnectionMetaEnvelope struct {
	Data *V2ConnectionMetaResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2GetConnectionMetaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2GetConnectionMetaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2GetConnectionMetaEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2GetConnectionMetaEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2GetExecutionResponseEnvelope struct {
	Data *V2GetExecutionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2GetExecutionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2GetExecutionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2GetExecutionResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2GetExecutionResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2GetExecutionResponseSchema struct {
	CompletedAt *time.Time         `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	Counts      *V2ExecutionCounts `json:"counts,omitempty" url:"counts,omitempty"`
	CreatedAt   *time.Time         `json:"created_at,omitempty" url:"created_at,omitempty"`
	Errors      []string           `json:"errors,omitempty" url:"errors,omitempty"`
	Id          *string            `json:"id,omitempty" url:"id,omitempty"`
	StartedAt   *time.Time         `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status      *string            `json:"status,omitempty" url:"status,omitempty"`
	Type        *string            `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2GetExecutionResponseSchema) UnmarshalJSON(data []byte) error {
	type embed V2GetExecutionResponseSchema
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V2GetExecutionResponseSchema(unmarshaler.embed)
	v.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	v.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	v.StartedAt = unmarshaler.StartedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2GetExecutionResponseSchema) MarshalJSON() ([]byte, error) {
	type embed V2GetExecutionResponseSchema
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*v),
		CompletedAt: core.NewOptionalDateTime(v.CompletedAt),
		CreatedAt:   core.NewOptionalDateTime(v.CreatedAt),
		StartedAt:   core.NewOptionalDateTime(v.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V2GetExecutionResponseSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2GetIdentityResponseEnvelope struct {
	Data *V2GetIdentityResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2GetIdentityResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2GetIdentityResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2GetIdentityResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2GetIdentityResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2GetIdentityResponseSchema struct {
	Email            *string `json:"email,omitempty" url:"email,omitempty"`
	Id               *string `json:"id,omitempty" url:"id,omitempty"`
	IsOrganization   *bool   `json:"is_organization,omitempty" url:"is_organization,omitempty"`
	IsPartner        *bool   `json:"is_partner,omitempty" url:"is_partner,omitempty"`
	IsSystem         *bool   `json:"is_system,omitempty" url:"is_system,omitempty"`
	IsUser           *bool   `json:"is_user,omitempty" url:"is_user,omitempty"`
	OrganizationId   *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OrganizationName *string `json:"organization_name,omitempty" url:"organization_name,omitempty"`
	Role             *string `json:"role,omitempty" url:"role,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2GetIdentityResponseSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler V2GetIdentityResponseSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2GetIdentityResponseSchema(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2GetIdentityResponseSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Identity struct {
	Function          string    `json:"function" url:"function"`
	NewField          *bool     `json:"new_field,omitempty" url:"new_field,omitempty"`
	RemoteFieldTypeId *string   `json:"remote_field_type_id,omitempty" url:"remote_field_type_id,omitempty"`
	Source            *V2Source `json:"source,omitempty" url:"source,omitempty"`
	Target            string    `json:"target" url:"target"`

	_rawJSON json.RawMessage
}

func (v *V2Identity) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Identity
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Identity(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Identity) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2IdentityFunction struct {
	Id    *string `json:"id,omitempty" url:"id,omitempty"`
	Label *string `json:"label,omitempty" url:"label,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2IdentityFunction) UnmarshalJSON(data []byte) error {
	type unmarshaler V2IdentityFunction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2IdentityFunction(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2IdentityFunction) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ListExecutionResponseEnvelope struct {
	Data []*V2GetExecutionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ListExecutionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ListExecutionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ListExecutionResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ListExecutionResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ListPoliciesResponseEnvelope struct {
	Data []*V2PolicyResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ListPoliciesResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ListPoliciesResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ListPoliciesResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ListPoliciesResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ListSyncResponseEnvelope struct {
	Data []*V2SyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ListSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ListSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ListSyncResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ListSyncResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ListUsersEnvelope struct {
	Data []*V2User `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ListUsersEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ListUsersEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ListUsersEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ListUsersEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Mode struct {
	Description           *string `json:"description,omitempty" url:"description,omitempty"`
	Label                 *string `json:"label,omitempty" url:"label,omitempty"`
	Mode                  *string `json:"mode,omitempty" url:"mode,omitempty"`
	RequiresIdentity      *bool   `json:"requires_identity,omitempty" url:"requires_identity,omitempty"`
	SupportsFieldSyncMode *bool   `json:"supports_field_sync_mode,omitempty" url:"supports_field_sync_mode,omitempty"`
	SupportsTargetFilters *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Mode) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Mode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Mode(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Mode) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ModelField struct {
	Description *string     `json:"description,omitempty" url:"description,omitempty"`
	Example     interface{} `json:"example,omitempty" url:"example,omitempty"`
	Label       *string     `json:"label,omitempty" url:"label,omitempty"`
	Name        *string     `json:"name,omitempty" url:"name,omitempty"`
	RemoteType  *string     `json:"remote_type,omitempty" url:"remote_type,omitempty"`
	Type        *string     `json:"type,omitempty" url:"type,omitempty"`
	Unique      *bool       `json:"unique,omitempty" url:"unique,omitempty"`
	UserAdded   *bool       `json:"user_added,omitempty" url:"user_added,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ModelField) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ModelField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ModelField(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ModelField) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ModelFieldRequest struct {
	Example *string `json:"example,omitempty" url:"example,omitempty"`
	Label   string  `json:"label" url:"label"`
	Name    string  `json:"name" url:"name"`
	Type    string  `json:"type" url:"type"`

	_rawJSON json.RawMessage
}

func (v *V2ModelFieldRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ModelFieldRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ModelFieldRequest(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ModelFieldRequest) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ModelFieldResponse struct {
	Data []*V2ModelField `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ModelFieldResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ModelFieldResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ModelFieldResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ModelFieldResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ModelListResponseEnvelope struct {
	Data []*V2ModelResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ModelListResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ModelListResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ModelListResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ModelListResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ModelResponse struct {
	Configuration   map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId    *string                `json:"connection_id,omitempty" url:"connection_id,omitempty"`
	Fields          []*V2ModelField        `json:"fields,omitempty" url:"fields,omitempty"`
	Id              *string                `json:"id,omitempty" url:"id,omitempty"`
	Identifier      *string                `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels          []LabelLabel           `json:"labels,omitempty" url:"labels,omitempty"`
	Name            *string                `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId  *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies        []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Relations       []*V2Relation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns []string               `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
	Type            *string                `json:"type,omitempty" url:"type,omitempty"`
	Version         *int                   `json:"version,omitempty" url:"version,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ModelResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ModelResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ModelResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ModelResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2ModelResponseEnvelope struct {
	Data *V2ModelResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2ModelResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2ModelResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2ModelResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2ModelResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Organization struct {
	Id        *string `json:"id,omitempty" url:"id,omitempty"`
	Issuer    *string `json:"issuer,omitempty" url:"issuer,omitempty"`
	Name      *string `json:"name,omitempty" url:"name,omitempty"`
	SsoDomain *string `json:"sso_domain,omitempty" url:"sso_domain,omitempty"`
	SsoOrgId  *string `json:"sso_org_id,omitempty" url:"sso_org_id,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Organization) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Organization
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Organization(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Organization) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2OrganizationEnvelope struct {
	Data *V2Organization `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2OrganizationEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2OrganizationEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2OrganizationEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2OrganizationEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2OrganizationsEnvelope struct {
	Data []*V2Organization `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2OrganizationsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2OrganizationsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2OrganizationsEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2OrganizationsEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Override struct {
	FieldId  *string     `json:"field_id,omitempty" url:"field_id,omitempty"`
	Function *string     `json:"function,omitempty" url:"function,omitempty"`
	Override interface{} `json:"override,omitempty" url:"override,omitempty"`
	Value    interface{} `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Override) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Override
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Override(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Override) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2PolicyAction struct {
	Action  string   `json:"action" url:"action"`
	RoleIds []string `json:"role_ids,omitempty" url:"role_ids,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2PolicyAction) UnmarshalJSON(data []byte) error {
	type unmarshaler V2PolicyAction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2PolicyAction(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2PolicyAction) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2PolicyResponse struct {
	Id             *string           `json:"id,omitempty" url:"id,omitempty"`
	Name           *string           `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string           `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	PolicyActions  []*V2PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty"`
	System         *bool             `json:"system,omitempty" url:"system,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2PolicyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2PolicyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2PolicyResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2PolicyResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2PolicyResponseEnvelope struct {
	Data *V2PolicyResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2PolicyResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2PolicyResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2PolicyResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2PolicyResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Relation struct {
	From *string       `json:"from,omitempty" url:"from,omitempty"`
	To   *V2RelationTo `json:"to,omitempty" url:"to,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Relation) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Relation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Relation(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Relation) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2RelationTo struct {
	Field   *string `json:"field,omitempty" url:"field,omitempty"`
	ModelId *string `json:"model_id,omitempty" url:"model_id,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2RelationTo) UnmarshalJSON(data []byte) error {
	type unmarshaler V2RelationTo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2RelationTo(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2RelationTo) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2RoleListResponseEnvelope struct {
	Data []*V2RoleResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2RoleListResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2RoleListResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2RoleListResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2RoleListResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2RoleResponse struct {
	Id             *string `json:"id,omitempty" url:"id,omitempty"`
	Name           *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	System         *bool   `json:"system,omitempty" url:"system,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2RoleResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2RoleResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2RoleResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2RoleResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2RoleResponseEnvelope struct {
	Data *V2RoleResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2RoleResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2RoleResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2RoleResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2RoleResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2RunAfter struct {
	BulkSyncIds []string `json:"bulk_sync_ids,omitempty" url:"bulk_sync_ids,omitempty"`
	SyncIds     []string `json:"sync_ids,omitempty" url:"sync_ids,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2RunAfter) UnmarshalJSON(data []byte) error {
	type unmarshaler V2RunAfter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2RunAfter(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2RunAfter) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Schedule struct {
	DayOfMonth *string     `json:"day_of_month,omitempty" url:"day_of_month,omitempty"`
	DayOfWeek  *string     `json:"day_of_week,omitempty" url:"day_of_week,omitempty"`
	Frequency  *string     `json:"frequency,omitempty" url:"frequency,omitempty"`
	Hour       *string     `json:"hour,omitempty" url:"hour,omitempty"`
	Minute     *string     `json:"minute,omitempty" url:"minute,omitempty"`
	Month      *string     `json:"month,omitempty" url:"month,omitempty"`
	RunAfter   *V2RunAfter `json:"run_after,omitempty" url:"run_after,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Schedule) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Schedule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Schedule(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Schedule) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Source struct {
	Field   string `json:"field" url:"field"`
	ModelId string `json:"model_id" url:"model_id"`

	_rawJSON json.RawMessage
}

func (v *V2Source) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Source
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Source(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Source) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2StartSyncResponseEnvelope struct {
	Data *V2StartSyncResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2StartSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2StartSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2StartSyncResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2StartSyncResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2StartSyncResponseSchema struct {
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	Id        *string    `json:"id,omitempty" url:"id,omitempty"`
	Status    *string    `json:"status,omitempty" url:"status,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2StartSyncResponseSchema) UnmarshalJSON(data []byte) error {
	type embed V2StartSyncResponseSchema
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V2StartSyncResponseSchema(unmarshaler.embed)
	v.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2StartSyncResponseSchema) MarshalJSON() ([]byte, error) {
	type embed V2StartSyncResponseSchema
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*v),
		CreatedAt: core.NewOptionalDateTime(v.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V2StartSyncResponseSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2SupportedMode struct {
	Description           *string `json:"description,omitempty" url:"description,omitempty"`
	Id                    *string `json:"id,omitempty" url:"id,omitempty"`
	Label                 *string `json:"label,omitempty" url:"label,omitempty"`
	RequiresIdentity      *bool   `json:"requires_identity,omitempty" url:"requires_identity,omitempty"`
	SupportsFieldSyncMode *bool   `json:"supports_field_sync_mode,omitempty" url:"supports_field_sync_mode,omitempty"`
	SupportsTargetFilters *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2SupportedMode) UnmarshalJSON(data []byte) error {
	type unmarshaler V2SupportedMode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2SupportedMode(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2SupportedMode) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2SyncField struct {
	New           *bool     `json:"new,omitempty" url:"new,omitempty"`
	OverrideValue *string   `json:"override_value,omitempty" url:"override_value,omitempty"`
	Source        *V2Source `json:"source,omitempty" url:"source,omitempty"`
	SyncMode      *string   `json:"sync_mode,omitempty" url:"sync_mode,omitempty"`
	Target        string    `json:"target" url:"target"`

	_rawJSON json.RawMessage
}

func (v *V2SyncField) UnmarshalJSON(data []byte) error {
	type unmarshaler V2SyncField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2SyncField(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2SyncField) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2SyncResponse struct {
	Active         *bool          `json:"active,omitempty" url:"active,omitempty"`
	Fields         []*V2SyncField `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string        `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*V2Filter    `json:"filters,omitempty" url:"filters,omitempty"`
	Id             *string        `json:"id,omitempty" url:"id,omitempty"`
	Identity       *V2Identity    `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           *string        `json:"mode,omitempty" url:"mode,omitempty"`
	Name           *string        `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string        `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OverrideFields []*V2SyncField `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	Overrides      []*V2Override  `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string       `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *V2Schedule    `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool          `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *V2Target      `json:"target,omitempty" url:"target,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2SyncResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V2SyncResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2SyncResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2SyncResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2SyncResponseEnvelope struct {
	Data *V2SyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2SyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2SyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2SyncResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2SyncResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2SyncStatusEnvelope struct {
	Data *V2SyncStatusResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2SyncStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2SyncStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2SyncStatusEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2SyncStatusEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2SyncStatusResponse struct {
	CurrentExecution  *V2GetExecutionResponseSchema `json:"current_execution,omitempty" url:"current_execution,omitempty"`
	LastExecution     *V2GetExecutionResponseSchema `json:"last_execution,omitempty" url:"last_execution,omitempty"`
	NextExecutionTime *time.Time                    `json:"next_execution_time,omitempty" url:"next_execution_time,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2SyncStatusResponse) UnmarshalJSON(data []byte) error {
	type embed V2SyncStatusResponse
	var unmarshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V2SyncStatusResponse(unmarshaler.embed)
	v.NextExecutionTime = unmarshaler.NextExecutionTime.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2SyncStatusResponse) MarshalJSON() ([]byte, error) {
	type embed V2SyncStatusResponse
	var marshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed:             embed(*v),
		NextExecutionTime: core.NewOptionalDateTime(v.NextExecutionTime),
	}
	return json.Marshal(marshaler)
}

func (v *V2SyncStatusResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Target struct {
	Configuration map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId  string                 `json:"connection_id" url:"connection_id"`
	FilterLogic   *string                `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	NewName       *string                `json:"new_name,omitempty" url:"new_name,omitempty"`
	Object        string                 `json:"object" url:"object"`
	SearchValues  map[string]interface{} `json:"search_values,omitempty" url:"search_values,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Target) UnmarshalJSON(data []byte) error {
	type unmarshaler V2Target
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2Target(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Target) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2TargetField struct {
	Association       *bool                 `json:"association,omitempty" url:"association,omitempty"`
	Createable        *bool                 `json:"createable,omitempty" url:"createable,omitempty"`
	Description       *string               `json:"description,omitempty" url:"description,omitempty"`
	Filterable        *bool                 `json:"filterable,omitempty" url:"filterable,omitempty"`
	Id                *string               `json:"id,omitempty" url:"id,omitempty"`
	IdentityFunctions []*V2IdentityFunction `json:"identity_functions,omitempty" url:"identity_functions,omitempty"`
	Name              *string               `json:"name,omitempty" url:"name,omitempty"`
	Required          *bool                 `json:"required,omitempty" url:"required,omitempty"`
	SourceType        *string               `json:"source_type,omitempty" url:"source_type,omitempty"`
	SupportsIdentity  *bool                 `json:"supports_identity,omitempty" url:"supports_identity,omitempty"`
	Type              *string               `json:"type,omitempty" url:"type,omitempty"`
	Updateable        *bool                 `json:"updateable,omitempty" url:"updateable,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2TargetField) UnmarshalJSON(data []byte) error {
	type unmarshaler V2TargetField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2TargetField(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2TargetField) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2TargetResponse struct {
	Fields      []*V2TargetField           `json:"fields,omitempty" url:"fields,omitempty"`
	Id          *string                    `json:"id,omitempty" url:"id,omitempty"`
	Modes       []*V2Mode                  `json:"modes,omitempty" url:"modes,omitempty"`
	Name        *string                    `json:"name,omitempty" url:"name,omitempty"`
	Properties  *SyncDestinationProperties `json:"properties,omitempty" url:"properties,omitempty"`
	RefreshedAt *time.Time                 `json:"refreshed_at,omitempty" url:"refreshed_at,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2TargetResponse) UnmarshalJSON(data []byte) error {
	type embed V2TargetResponse
	var unmarshaler = struct {
		embed
		RefreshedAt *core.DateTime `json:"refreshed_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V2TargetResponse(unmarshaler.embed)
	v.RefreshedAt = unmarshaler.RefreshedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2TargetResponse) MarshalJSON() ([]byte, error) {
	type embed V2TargetResponse
	var marshaler = struct {
		embed
		RefreshedAt *core.DateTime `json:"refreshed_at,omitempty"`
	}{
		embed:       embed(*v),
		RefreshedAt: core.NewOptionalDateTime(v.RefreshedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V2TargetResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2TargetResponseEnvelope struct {
	Data *V2TargetResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2TargetResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2TargetResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2TargetResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2TargetResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2User struct {
	Email          *string `json:"email,omitempty" url:"email,omitempty"`
	Id             *string `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Role           *string `json:"role,omitempty" url:"role,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2User) UnmarshalJSON(data []byte) error {
	type unmarshaler V2User
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2User(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2User) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2UserEnvelope struct {
	Data *V2User `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2UserEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2UserEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2UserEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2UserEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2Webhook struct {
	CreatedAt      *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	Endpoint       *string    `json:"endpoint,omitempty" url:"endpoint,omitempty"`
	Id             *string    `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string    `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         *string    `json:"secret,omitempty" url:"secret,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2Webhook) UnmarshalJSON(data []byte) error {
	type embed V2Webhook
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V2Webhook(unmarshaler.embed)
	v.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2Webhook) MarshalJSON() ([]byte, error) {
	type embed V2Webhook
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*v),
		CreatedAt: core.NewOptionalDateTime(v.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V2Webhook) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2WebhookEnvelope struct {
	Data *V2Webhook `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2WebhookEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2WebhookEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2WebhookEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2WebhookEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V2WebhookListEnvelope struct {
	Data []*V2Webhook `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2WebhookListEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2WebhookListEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2WebhookListEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2WebhookListEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkField struct {
	Enabled    *bool   `json:"enabled,omitempty" url:"enabled,omitempty"`
	Id         *string `json:"id,omitempty" url:"id,omitempty"`
	Obfuscated *bool   `json:"obfuscated,omitempty" url:"obfuscated,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkField) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkField(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkField) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSchema struct {
	Enabled      *bool          `json:"enabled,omitempty" url:"enabled,omitempty"`
	Fields       []*V3BulkField `json:"fields,omitempty" url:"fields,omitempty"`
	Id           *string        `json:"id,omitempty" url:"id,omitempty"`
	PartitionKey *string        `json:"partition_key,omitempty" url:"partition_key,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSchema(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSchema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSchemaEnvelope struct {
	Data *V3BulkSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSchemaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSchemaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSchemaEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSchemaEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncExecution struct {
	CompletedAt *time.Time                   `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	CreatedAt   *time.Time                   `json:"created_at,omitempty" url:"created_at,omitempty"`
	Id          *string                      `json:"id,omitempty" url:"id,omitempty"`
	IsResync    *bool                        `json:"is_resync,omitempty" url:"is_resync,omitempty"`
	IsTest      *bool                        `json:"is_test,omitempty" url:"is_test,omitempty"`
	Schemas     []*V3BulkSyncSchemaExecution `json:"schemas,omitempty" url:"schemas,omitempty"`
	StartedAt   *time.Time                   `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status      *string                      `json:"status,omitempty" url:"status,omitempty"`
	Type        *string                      `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncExecution) UnmarshalJSON(data []byte) error {
	type embed V3BulkSyncExecution
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V3BulkSyncExecution(unmarshaler.embed)
	v.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	v.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	v.StartedAt = unmarshaler.StartedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncExecution) MarshalJSON() ([]byte, error) {
	type embed V3BulkSyncExecution
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*v),
		CompletedAt: core.NewOptionalDateTime(v.CompletedAt),
		CreatedAt:   core.NewOptionalDateTime(v.CreatedAt),
		StartedAt:   core.NewOptionalDateTime(v.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V3BulkSyncExecution) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncExecutionEnvelope struct {
	Data *V3BulkSyncExecution `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncExecutionEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSyncExecutionEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSyncExecutionEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncExecutionEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncSchemaExecution struct {
	CompletedAt   *time.Time `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	ErrorCount    *int       `json:"error_count,omitempty" url:"error_count,omitempty"`
	RecordCount   *int       `json:"record_count,omitempty" url:"record_count,omitempty"`
	Schema        *string    `json:"schema,omitempty" url:"schema,omitempty"`
	StartedAt     *time.Time `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status        *string    `json:"status,omitempty" url:"status,omitempty"`
	StatusMessage *string    `json:"status_message,omitempty" url:"status_message,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncSchemaExecution) UnmarshalJSON(data []byte) error {
	type embed V3BulkSyncSchemaExecution
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V3BulkSyncSchemaExecution(unmarshaler.embed)
	v.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	v.StartedAt = unmarshaler.StartedAt.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncSchemaExecution) MarshalJSON() ([]byte, error) {
	type embed V3BulkSyncSchemaExecution
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*v),
		CompletedAt: core.NewOptionalDateTime(v.CompletedAt),
		StartedAt:   core.NewOptionalDateTime(v.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (v *V3BulkSyncSchemaExecution) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncSource struct {
	Configuration interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Schemas       []*V3Schema `json:"schemas,omitempty" url:"schemas,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncSource) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSyncSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSyncSource(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncSource) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncSourceEnvelope struct {
	Data *V3BulkSyncSource `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncSourceEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSyncSourceEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSyncSourceEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncSourceEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncSourceSchemaEnvelope struct {
	Data *V3Schema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncSourceSchemaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSyncSourceSchemaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSyncSourceSchemaEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncSourceSchemaEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncSourceStatus struct {
	CacheStatus         *string    `json:"cache_status,omitempty" url:"cache_status,omitempty"`
	LastRefreshFinished *time.Time `json:"last_refresh_finished,omitempty" url:"last_refresh_finished,omitempty"`
	LastRefreshStarted  *time.Time `json:"last_refresh_started,omitempty" url:"last_refresh_started,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncSourceStatus) UnmarshalJSON(data []byte) error {
	type embed V3BulkSyncSourceStatus
	var unmarshaler = struct {
		embed
		LastRefreshFinished *core.DateTime `json:"last_refresh_finished,omitempty"`
		LastRefreshStarted  *core.DateTime `json:"last_refresh_started,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V3BulkSyncSourceStatus(unmarshaler.embed)
	v.LastRefreshFinished = unmarshaler.LastRefreshFinished.TimePtr()
	v.LastRefreshStarted = unmarshaler.LastRefreshStarted.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncSourceStatus) MarshalJSON() ([]byte, error) {
	type embed V3BulkSyncSourceStatus
	var marshaler = struct {
		embed
		LastRefreshFinished *core.DateTime `json:"last_refresh_finished,omitempty"`
		LastRefreshStarted  *core.DateTime `json:"last_refresh_started,omitempty"`
	}{
		embed:               embed(*v),
		LastRefreshFinished: core.NewOptionalDateTime(v.LastRefreshFinished),
		LastRefreshStarted:  core.NewOptionalDateTime(v.LastRefreshStarted),
	}
	return json.Marshal(marshaler)
}

func (v *V3BulkSyncSourceStatus) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncSourceStatusEnvelope struct {
	Data *V3BulkSyncSourceStatus `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncSourceStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSyncSourceStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSyncSourceStatusEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncSourceStatusEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncStatusEnvelope struct {
	Data *V3BulkSyncStatusResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3BulkSyncStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3BulkSyncStatusEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncStatusEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3BulkSyncStatusResponse struct {
	CurrentExecution  *V3BulkSyncExecution `json:"current_execution,omitempty" url:"current_execution,omitempty"`
	LastExecution     *V3BulkSyncExecution `json:"last_execution,omitempty" url:"last_execution,omitempty"`
	NextExecutionTime *time.Time           `json:"next_execution_time,omitempty" url:"next_execution_time,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3BulkSyncStatusResponse) UnmarshalJSON(data []byte) error {
	type embed V3BulkSyncStatusResponse
	var unmarshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed: embed(*v),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*v = V3BulkSyncStatusResponse(unmarshaler.embed)
	v.NextExecutionTime = unmarshaler.NextExecutionTime.TimePtr()
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3BulkSyncStatusResponse) MarshalJSON() ([]byte, error) {
	type embed V3BulkSyncStatusResponse
	var marshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed:             embed(*v),
		NextExecutionTime: core.NewOptionalDateTime(v.NextExecutionTime),
	}
	return json.Marshal(marshaler)
}

func (v *V3BulkSyncStatusResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3ConnectCardResponse struct {
	RedirectUrl *string `json:"redirect_url,omitempty" url:"redirect_url,omitempty"`
	Token       *string `json:"token,omitempty" url:"token,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3ConnectCardResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V3ConnectCardResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3ConnectCardResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3ConnectCardResponse) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3ConnectCardResponseEnvelope struct {
	Data *V3ConnectCardResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3ConnectCardResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3ConnectCardResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3ConnectCardResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3ConnectCardResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3ListBulkSchemaEnvelope struct {
	Data []*V3BulkSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3ListBulkSchemaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3ListBulkSchemaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3ListBulkSchemaEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3ListBulkSchemaEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3ListBulkSyncExecutionsEnvelope struct {
	Data []*V3BulkSyncExecution `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3ListBulkSyncExecutionsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3ListBulkSyncExecutionsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3ListBulkSyncExecutionsEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3ListBulkSyncExecutionsEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3PickValue struct {
	Label *string `json:"label,omitempty" url:"label,omitempty"`
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3PickValue) UnmarshalJSON(data []byte) error {
	type unmarshaler V3PickValue
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3PickValue(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3PickValue) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3Schema struct {
	Fields []*V3SchemaField `json:"fields,omitempty" url:"fields,omitempty"`
	Id     *string          `json:"id,omitempty" url:"id,omitempty"`
	Name   *string          `json:"name,omitempty" url:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3Schema) UnmarshalJSON(data []byte) error {
	type unmarshaler V3Schema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3Schema(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3Schema) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3SchemaField struct {
	Association *SchemaAssociation `json:"association,omitempty" url:"association,omitempty"`
	Id          *string            `json:"id,omitempty" url:"id,omitempty"`
	Name        *string            `json:"name,omitempty" url:"name,omitempty"`
	RemoteType  *string            `json:"remote_type,omitempty" url:"remote_type,omitempty"`
	Type        *string            `json:"type,omitempty" url:"type,omitempty"`
	Values      []*V3PickValue     `json:"values,omitempty" url:"values,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3SchemaField) UnmarshalJSON(data []byte) error {
	type unmarshaler V3SchemaField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3SchemaField(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3SchemaField) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}

type V3SchemaRecordsResponseEnvelope struct {
	Data []map[string]interface{} `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V3SchemaRecordsResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V3SchemaRecordsResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V3SchemaRecordsResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V3SchemaRecordsResponseEnvelope) String() string {
	if len(v._rawJSON) > 0 {
		if value, err := core.StringifyJSON(v._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(v); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", v)
}
