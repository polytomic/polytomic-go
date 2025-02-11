// This file was auto-generated by Fern from our API Definition.

package polytomic

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/polytomic/polytomic-go/core"
	time "time"
)

type CreateModelSyncRequest struct {
	Active   *bool       `json:"active,omitempty" url:"active,omitempty"`
	Enricher *Enrichment `json:"enricher,omitempty" url:"enricher,omitempty"`
	// Fields to sync from source to target.
	Fields         []*ModelSyncField `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string           `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*Filter         `json:"filters,omitempty" url:"filters,omitempty"`
	Identity       *Identity         `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           string            `json:"mode" url:"mode"`
	Name           string            `json:"name" url:"name"`
	OrganizationId *string           `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	// Values to set in the target unconditionally.
	OverrideFields []*ModelSyncField `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	// Conditional value replacement for fields.
	Overrides      []*Override `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string    `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *Schedule   `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool       `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *Target     `json:"target,omitempty" url:"target,omitempty"`
}

type ModelSyncGetSourceRequest struct {
	Params map[string][]string `json:"-" url:"params,omitempty"`
}

type ModelSyncGetSourceFieldsRequest struct {
	Params map[string][]string `json:"-" url:"params,omitempty"`
}

type ModelSyncGetTargetRequest struct {
	Type   *string `json:"-" url:"type,omitempty"`
	Search *string `json:"-" url:"search,omitempty"`
}

type ModelSyncGetTargetFieldsRequest struct {
	Target  string `json:"-" url:"target"`
	Refresh *bool  `json:"-" url:"refresh,omitempty"`
}

type ModelSyncListRequest struct {
	Active             *bool     `json:"-" url:"active,omitempty"`
	Mode               *SyncMode `json:"-" url:"mode,omitempty"`
	TargetConnectionId *string   `json:"-" url:"target_connection_id,omitempty"`
}

type StartModelSyncRequest struct {
	Identities []string `json:"identities,omitempty" url:"identities,omitempty"`
	Resync     *bool    `json:"resync,omitempty" url:"resync,omitempty"`
}

type UpdateModelSyncRequest struct {
	Active   *bool       `json:"active,omitempty" url:"active,omitempty"`
	Enricher *Enrichment `json:"enricher,omitempty" url:"enricher,omitempty"`
	// Fields to sync from source to target.
	Fields         []*ModelSyncField `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string           `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*Filter         `json:"filters,omitempty" url:"filters,omitempty"`
	Identity       *Identity         `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           string            `json:"mode" url:"mode"`
	Name           string            `json:"name" url:"name"`
	OrganizationId *string           `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	// Values to set in the target unconditionally.
	OverrideFields []*ModelSyncField `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	// Conditional value replacement for fields.
	Overrides      []*Override `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string    `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *Schedule   `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool       `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *Target     `json:"target,omitempty" url:"target,omitempty"`
}

type ConnectionMeta struct {
	HasItems      *bool         `json:"has_items,omitempty" url:"has_items,omitempty"`
	Items         []interface{} `json:"items,omitempty" url:"items,omitempty"`
	RequiresOneOf []string      `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionMeta(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionMeta) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

type ConnectionMetaResponse struct {
	Configuration map[string]*ConfigurationValue `json:"configuration,omitempty" url:"configuration,omitempty"`
	Items         map[string]*ConnectionMeta     `json:"items,omitempty" url:"items,omitempty"`
	RequiresOneOf []string                       `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionMetaResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionMetaResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionMetaResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionMetaResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Either `field` or `field_id` must be provided. If `field` is provided, `field_id` is ignored.
type Filter struct {
	Field *Source `json:"field,omitempty" url:"field,omitempty"`
	// Model or Target field name to filter on.
	FieldId   *string                   `json:"field_id,omitempty" url:"field_id,omitempty"`
	FieldType *FilterFieldReferenceType `json:"field_type,omitempty" url:"field_type,omitempty"`
	Function  FilterFunction            `json:"function,omitempty" url:"function,omitempty"`
	Label     *string                   `json:"label,omitempty" url:"label,omitempty"`
	Value     interface{}               `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (f *Filter) UnmarshalJSON(data []byte) error {
	type unmarshaler Filter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = Filter(value)
	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *Filter) String() string {
	if len(f._rawJSON) > 0 {
		if value, err := core.StringifyJSON(f._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(f); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", f)
}

type FilterFieldReferenceType string

const (
	FilterFieldReferenceTypeModel  FilterFieldReferenceType = "Model"
	FilterFieldReferenceTypeTarget FilterFieldReferenceType = "Target"
)

func NewFilterFieldReferenceTypeFromString(s string) (FilterFieldReferenceType, error) {
	switch s {
	case "Model":
		return FilterFieldReferenceTypeModel, nil
	case "Target":
		return FilterFieldReferenceTypeTarget, nil
	}
	var t FilterFieldReferenceType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (f FilterFieldReferenceType) Ptr() *FilterFieldReferenceType {
	return &f
}

type GetConnectionMetaEnvelope struct {
	Data *ConnectionMetaResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetConnectionMetaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler GetConnectionMetaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetConnectionMetaEnvelope(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetConnectionMetaEnvelope) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

type Identity struct {
	Function          SchemaIdentityFunction `json:"function,omitempty" url:"function,omitempty"`
	NewField          *bool                  `json:"new_field,omitempty" url:"new_field,omitempty"`
	RemoteFieldTypeId *string                `json:"remote_field_type_id,omitempty" url:"remote_field_type_id,omitempty"`
	Source            *Source                `json:"source,omitempty" url:"source,omitempty"`
	Target            string                 `json:"target" url:"target"`

	_rawJSON json.RawMessage
}

func (i *Identity) UnmarshalJSON(data []byte) error {
	type unmarshaler Identity
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = Identity(value)
	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *Identity) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type IdentityFunction struct {
	Id    *string `json:"id,omitempty" url:"id,omitempty"`
	Label *string `json:"label,omitempty" url:"label,omitempty"`

	_rawJSON json.RawMessage
}

func (i *IdentityFunction) UnmarshalJSON(data []byte) error {
	type unmarshaler IdentityFunction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IdentityFunction(value)
	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *IdentityFunction) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}

type JsonschemaForm = map[string]interface{}

type ListModelSyncResponseEnvelope struct {
	Data []*ModelSyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListModelSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ListModelSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListModelSyncResponseEnvelope(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListModelSyncResponseEnvelope) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type Mode struct {
	Description           *string `json:"description,omitempty" url:"description,omitempty"`
	Label                 *string `json:"label,omitempty" url:"label,omitempty"`
	Mode                  *string `json:"mode,omitempty" url:"mode,omitempty"`
	RequiresIdentity      *bool   `json:"requires_identity,omitempty" url:"requires_identity,omitempty"`
	SupportsFieldSyncMode *bool   `json:"supports_field_sync_mode,omitempty" url:"supports_field_sync_mode,omitempty"`
	SupportsTargetFilters *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`

	_rawJSON json.RawMessage
}

func (m *Mode) UnmarshalJSON(data []byte) error {
	type unmarshaler Mode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = Mode(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *Mode) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type ModelFieldResponse struct {
	Data []*ModelField `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelFieldResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelFieldResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelFieldResponse(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelFieldResponse) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type ModelSyncField struct {
	// New is set to true if the target field should be created by Polytomic. This is not supported by all backends.
	New *bool `json:"new,omitempty" url:"new,omitempty"`
	// Value to set in the target field; if provided, 'source' is ignored.
	OverrideValue *string `json:"override_value,omitempty" url:"override_value,omitempty"`
	Source        *Source `json:"source,omitempty" url:"source,omitempty"`
	// Sync mode for the field; defaults to 'updateOrCreate'. If set to 'create', the field will not be synced if it already has a value. This is not supported by all backends.
	SyncMode *string `json:"sync_mode,omitempty" url:"sync_mode,omitempty"`
	// Target field ID the source field value will be written to.
	Target string `json:"target" url:"target"`

	_rawJSON json.RawMessage
}

func (m *ModelSyncField) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelSyncField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelSyncField(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelSyncField) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type ModelSyncResponse struct {
	Active         *bool              `json:"active,omitempty" url:"active,omitempty"`
	CreatedAt      *time.Time         `json:"created_at,omitempty" url:"created_at,omitempty"`
	CreatedBy      *CommonOutputActor `json:"created_by,omitempty" url:"created_by,omitempty"`
	Fields         []*ModelSyncField  `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string            `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*Filter          `json:"filters,omitempty" url:"filters,omitempty"`
	Id             *string            `json:"id,omitempty" url:"id,omitempty"`
	Identity       *Identity          `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           *string            `json:"mode,omitempty" url:"mode,omitempty"`
	Name           *string            `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string            `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OverrideFields []*ModelSyncField  `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	Overrides      []*Override        `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string           `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *Schedule          `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool              `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *Target            `json:"target,omitempty" url:"target,omitempty"`
	UpdatedAt      *time.Time         `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UpdatedBy      *CommonOutputActor `json:"updated_by,omitempty" url:"updated_by,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelSyncResponse) UnmarshalJSON(data []byte) error {
	type embed ModelSyncResponse
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
		UpdatedAt *core.DateTime `json:"updated_at,omitempty"`
	}{
		embed: embed(*m),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*m = ModelSyncResponse(unmarshaler.embed)
	m.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	m.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelSyncResponse) MarshalJSON() ([]byte, error) {
	type embed ModelSyncResponse
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
		UpdatedAt *core.DateTime `json:"updated_at,omitempty"`
	}{
		embed:     embed(*m),
		CreatedAt: core.NewOptionalDateTime(m.CreatedAt),
		UpdatedAt: core.NewOptionalDateTime(m.UpdatedAt),
	}
	return json.Marshal(marshaler)
}

func (m *ModelSyncResponse) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

type ModelSyncResponseEnvelope struct {
	Data *ModelSyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelSyncResponseEnvelope(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelSyncResponseEnvelope) String() string {
	if len(m._rawJSON) > 0 {
		if value, err := core.StringifyJSON(m._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(m); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", m)
}

// Either `field` or `field_id` must be provided. If `field_id` is provided, `field` is ignored.
type Override struct {
	Field *Source `json:"field,omitempty" url:"field,omitempty"`
	// Field ID of the model field to override.
	FieldId  *string         `json:"field_id,omitempty" url:"field_id,omitempty"`
	Function *FilterFunction `json:"function,omitempty" url:"function,omitempty"`
	Override interface{}     `json:"override,omitempty" url:"override,omitempty"`
	Value    interface{}     `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (o *Override) UnmarshalJSON(data []byte) error {
	type unmarshaler Override
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = Override(value)
	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *Override) String() string {
	if len(o._rawJSON) > 0 {
		if value, err := core.StringifyJSON(o._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(o); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", o)
}

type RunAfter struct {
	BulkSyncIds []string `json:"bulk_sync_ids,omitempty" url:"bulk_sync_ids,omitempty"`
	SyncIds     []string `json:"sync_ids,omitempty" url:"sync_ids,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RunAfter) UnmarshalJSON(data []byte) error {
	type unmarshaler RunAfter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RunAfter(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RunAfter) String() string {
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

type Schedule struct {
	ConnectionId *string            `json:"connection_id,omitempty" url:"connection_id,omitempty"`
	DayOfMonth   *string            `json:"day_of_month,omitempty" url:"day_of_month,omitempty"`
	DayOfWeek    *string            `json:"day_of_week,omitempty" url:"day_of_week,omitempty"`
	Frequency    *ScheduleFrequency `json:"frequency,omitempty" url:"frequency,omitempty"`
	Hour         *string            `json:"hour,omitempty" url:"hour,omitempty"`
	JobId        *int               `json:"job_id,omitempty" url:"job_id,omitempty"`
	Minute       *string            `json:"minute,omitempty" url:"minute,omitempty"`
	Month        *string            `json:"month,omitempty" url:"month,omitempty"`
	RunAfter     *RunAfter          `json:"run_after,omitempty" url:"run_after,omitempty"`
	// If true, the sync will only run if the dependent syncs completed successfully.
	RunAfterSuccessOnly *bool `json:"run_after_success_only,omitempty" url:"run_after_success_only,omitempty"`

	_rawJSON json.RawMessage
}

func (s *Schedule) UnmarshalJSON(data []byte) error {
	type unmarshaler Schedule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Schedule(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Schedule) String() string {
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

type ScheduleOptionResponse struct {
	ScheduleOptions []*ScheduleScheduleOption `json:"schedule_options,omitempty" url:"schedule_options,omitempty"`

	_rawJSON json.RawMessage
}

func (s *ScheduleOptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ScheduleOptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ScheduleOptionResponse(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ScheduleOptionResponse) String() string {
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

type ScheduleOptionResponseEnvelope struct {
	Data *ScheduleOptionResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *ScheduleOptionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ScheduleOptionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ScheduleOptionResponseEnvelope(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ScheduleOptionResponseEnvelope) String() string {
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

type ScheduleScheduleOption struct {
	Configuration *JsonschemaForm    `json:"configuration,omitempty" url:"configuration,omitempty"`
	Description   *string            `json:"description,omitempty" url:"description,omitempty"`
	Frequency     *ScheduleFrequency `json:"frequency,omitempty" url:"frequency,omitempty"`
	Label         *string            `json:"label,omitempty" url:"label,omitempty"`

	_rawJSON json.RawMessage
}

func (s *ScheduleScheduleOption) UnmarshalJSON(data []byte) error {
	type unmarshaler ScheduleScheduleOption
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ScheduleScheduleOption(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *ScheduleScheduleOption) String() string {
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

type SchemaIdentityFunction string

const (
	SchemaIdentityFunctionEquality   SchemaIdentityFunction = "Equality"
	SchemaIdentityFunctionISubstring SchemaIdentityFunction = "ISubstring"
	SchemaIdentityFunctionOneOf      SchemaIdentityFunction = "OneOf"
)

func NewSchemaIdentityFunctionFromString(s string) (SchemaIdentityFunction, error) {
	switch s {
	case "Equality":
		return SchemaIdentityFunctionEquality, nil
	case "ISubstring":
		return SchemaIdentityFunctionISubstring, nil
	case "OneOf":
		return SchemaIdentityFunctionOneOf, nil
	}
	var t SchemaIdentityFunction
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SchemaIdentityFunction) Ptr() *SchemaIdentityFunction {
	return &s
}

type Source struct {
	Field   string `json:"field" url:"field"`
	ModelId string `json:"model_id" url:"model_id"`

	_rawJSON json.RawMessage
}

func (s *Source) UnmarshalJSON(data []byte) error {
	type unmarshaler Source
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Source(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Source) String() string {
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

type StartModelSyncResponseEnvelope struct {
	Data *StartModelSyncResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *StartModelSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler StartModelSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StartModelSyncResponseEnvelope(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *StartModelSyncResponseEnvelope) String() string {
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

type StartModelSyncResponseSchema struct {
	CreatedAt *time.Time       `json:"created_at,omitempty" url:"created_at,omitempty"`
	Id        *string          `json:"id,omitempty" url:"id,omitempty"`
	Status    *ExecutionStatus `json:"status,omitempty" url:"status,omitempty"`

	_rawJSON json.RawMessage
}

func (s *StartModelSyncResponseSchema) UnmarshalJSON(data []byte) error {
	type embed StartModelSyncResponseSchema
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = StartModelSyncResponseSchema(unmarshaler.embed)
	s.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *StartModelSyncResponseSchema) MarshalJSON() ([]byte, error) {
	type embed StartModelSyncResponseSchema
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*s),
		CreatedAt: core.NewOptionalDateTime(s.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (s *StartModelSyncResponseSchema) String() string {
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

type SupportedMode struct {
	Id *SyncMode `json:"id,omitempty" url:"id,omitempty"`
	// True if the sync mode requires an identity field mapping.
	RequiresIdentity *bool `json:"requires_identity,omitempty" url:"requires_identity,omitempty"`
	// True if the target supports per-field sync modes.
	SupportsPerFieldMode *bool `json:"supports_per_field_mode,omitempty" url:"supports_per_field_mode,omitempty"`
	// True if the sync mode supports target filters.
	SupportsTargetFilters *bool `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SupportedMode) UnmarshalJSON(data []byte) error {
	type unmarshaler SupportedMode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SupportedMode(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SupportedMode) String() string {
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
	DoesNotReportOperationCounts  *bool   `json:"does_not_report_operation_counts,omitempty" url:"does_not_report_operation_counts,omitempty"`
	MappingsNotRequired           *bool   `json:"mappings_not_required,omitempty" url:"mappings_not_required,omitempty"`
	NewTargetLabel                *string `json:"new_target_label,omitempty" url:"new_target_label,omitempty"`
	OptionalTargetMappings        *bool   `json:"optional_target_mappings,omitempty" url:"optional_target_mappings,omitempty"`
	PrimaryMetadataObject         *string `json:"primary_metadata_object,omitempty" url:"primary_metadata_object,omitempty"`
	RequiresConfiguration         *bool   `json:"requires_configuration,omitempty" url:"requires_configuration,omitempty"`
	SupportsFieldCreation         *bool   `json:"supports_field_creation,omitempty" url:"supports_field_creation,omitempty"`
	SupportsFieldTypeSelection    *bool   `json:"supports_field_type_selection,omitempty" url:"supports_field_type_selection,omitempty"`
	SupportsIdentityFieldCreation *bool   `json:"supports_identity_field_creation,omitempty" url:"supports_identity_field_creation,omitempty"`
	SupportsTargetFilters         *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`
	TargetCreator                 *bool   `json:"target_creator,omitempty" url:"target_creator,omitempty"`
	UseFieldNamesAsLabels         *bool   `json:"use_field_names_as_labels,omitempty" url:"use_field_names_as_labels,omitempty"`

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

type SyncStatusEnvelope struct {
	Data *SyncStatusResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncStatusEnvelope(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncStatusEnvelope) String() string {
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

type SyncStatusResponse struct {
	CurrentExecution  *GetExecutionResponseSchema `json:"current_execution,omitempty" url:"current_execution,omitempty"`
	LastExecution     *GetExecutionResponseSchema `json:"last_execution,omitempty" url:"last_execution,omitempty"`
	NextExecutionTime *time.Time                  `json:"next_execution_time,omitempty" url:"next_execution_time,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncStatusResponse) UnmarshalJSON(data []byte) error {
	type embed SyncStatusResponse
	var unmarshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SyncStatusResponse(unmarshaler.embed)
	s.NextExecutionTime = unmarshaler.NextExecutionTime.TimePtr()
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncStatusResponse) MarshalJSON() ([]byte, error) {
	type embed SyncStatusResponse
	var marshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed:             embed(*s),
		NextExecutionTime: core.NewOptionalDateTime(s.NextExecutionTime),
	}
	return json.Marshal(marshaler)
}

func (s *SyncStatusResponse) String() string {
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

type Target struct {
	Configuration map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId  string                 `json:"connection_id" url:"connection_id"`
	FilterLogic   *string                `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	NewName       *string                `json:"new_name,omitempty" url:"new_name,omitempty"`
	Object        string                 `json:"object" url:"object"`
	SearchValues  map[string]interface{} `json:"search_values,omitempty" url:"search_values,omitempty"`

	_rawJSON json.RawMessage
}

func (t *Target) UnmarshalJSON(data []byte) error {
	type unmarshaler Target
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Target(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *Target) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TargetField struct {
	Association       *bool               `json:"association,omitempty" url:"association,omitempty"`
	Createable        *bool               `json:"createable,omitempty" url:"createable,omitempty"`
	Description       *string             `json:"description,omitempty" url:"description,omitempty"`
	Filterable        *bool               `json:"filterable,omitempty" url:"filterable,omitempty"`
	Id                *string             `json:"id,omitempty" url:"id,omitempty"`
	IdentityFunctions []*IdentityFunction `json:"identity_functions,omitempty" url:"identity_functions,omitempty"`
	Name              *string             `json:"name,omitempty" url:"name,omitempty"`
	Required          *bool               `json:"required,omitempty" url:"required,omitempty"`
	SourceType        *string             `json:"source_type,omitempty" url:"source_type,omitempty"`
	SupportsIdentity  *bool               `json:"supports_identity,omitempty" url:"supports_identity,omitempty"`
	Type              *string             `json:"type,omitempty" url:"type,omitempty"`
	Updateable        *bool               `json:"updateable,omitempty" url:"updateable,omitempty"`

	_rawJSON json.RawMessage
}

func (t *TargetField) UnmarshalJSON(data []byte) error {
	type unmarshaler TargetField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TargetField(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TargetField) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TargetObject struct {
	// The identifier of the target object.
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The supported sync modes and their properties for the target object.
	Modes []*SupportedMode `json:"modes,omitempty" url:"modes,omitempty"`
	// The name of the target object.
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (t *TargetObject) UnmarshalJSON(data []byte) error {
	type unmarshaler TargetObject
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TargetObject(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TargetObject) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TargetResponse struct {
	Fields      []*TargetField             `json:"fields,omitempty" url:"fields,omitempty"`
	Id          *string                    `json:"id,omitempty" url:"id,omitempty"`
	Modes       []*Mode                    `json:"modes,omitempty" url:"modes,omitempty"`
	Name        *string                    `json:"name,omitempty" url:"name,omitempty"`
	Properties  *SyncDestinationProperties `json:"properties,omitempty" url:"properties,omitempty"`
	RefreshedAt *time.Time                 `json:"refreshed_at,omitempty" url:"refreshed_at,omitempty"`

	_rawJSON json.RawMessage
}

func (t *TargetResponse) UnmarshalJSON(data []byte) error {
	type embed TargetResponse
	var unmarshaler = struct {
		embed
		RefreshedAt *core.DateTime `json:"refreshed_at,omitempty"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = TargetResponse(unmarshaler.embed)
	t.RefreshedAt = unmarshaler.RefreshedAt.TimePtr()
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TargetResponse) MarshalJSON() ([]byte, error) {
	type embed TargetResponse
	var marshaler = struct {
		embed
		RefreshedAt *core.DateTime `json:"refreshed_at,omitempty"`
	}{
		embed:       embed(*t),
		RefreshedAt: core.NewOptionalDateTime(t.RefreshedAt),
	}
	return json.Marshal(marshaler)
}

func (t *TargetResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TargetResponseEnvelope struct {
	Data *TargetResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (t *TargetResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler TargetResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TargetResponseEnvelope(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TargetResponseEnvelope) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type V4TargetObjectsResponseEnvelope struct {
	Data []*TargetObject `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4TargetObjectsResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V4TargetObjectsResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4TargetObjectsResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4TargetObjectsResponseEnvelope) String() string {
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
