// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/polytomic/polytomic-go/core"
	time "time"
)

type ApiKeyResponse struct {
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (a *ApiKeyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ApiKeyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApiKeyResponse(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApiKeyResponse) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ApiKeyResponseEnvelope struct {
	Data *ApiKeyResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (a *ApiKeyResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ApiKeyResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApiKeyResponseEnvelope(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApiKeyResponseEnvelope) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ActivateSyncEnvelope struct {
	Data *ActivateSyncOutput `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (a *ActivateSyncEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ActivateSyncEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ActivateSyncEnvelope(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ActivateSyncEnvelope) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ActivateSyncInput struct {
	Active bool `json:"active" url:"active"`

	_rawJSON json.RawMessage
}

func (a *ActivateSyncInput) UnmarshalJSON(data []byte) error {
	type unmarshaler ActivateSyncInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ActivateSyncInput(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ActivateSyncInput) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ActivateSyncOutput struct {
	Active *bool   `json:"active,omitempty" url:"active,omitempty"`
	Id     *string `json:"id,omitempty" url:"id,omitempty"`

	_rawJSON json.RawMessage
}

func (a *ActivateSyncOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler ActivateSyncOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ActivateSyncOutput(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ActivateSyncOutput) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type ApiError struct {
	Key      *string                `json:"key,omitempty" url:"key,omitempty"`
	Message  *string                `json:"message,omitempty" url:"message,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Status   *int                   `json:"status,omitempty" url:"status,omitempty"`

	_rawJSON json.RawMessage
}

func (a *ApiError) UnmarshalJSON(data []byte) error {
	type unmarshaler ApiError
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = ApiError(value)
	a._rawJSON = json.RawMessage(data)
	return nil
}

func (a *ApiError) String() string {
	if len(a._rawJSON) > 0 {
		if value, err := core.StringifyJSON(a._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}

type BulkDiscover string

const (
	BulkDiscoverAll                BulkDiscover = "all"
	BulkDiscoverOnlyIncremental    BulkDiscover = "onlyIncremental"
	BulkDiscoverOnlyNonIncremental BulkDiscover = "onlyNonIncremental"
	BulkDiscoverNone               BulkDiscover = "none"
)

func NewBulkDiscoverFromString(s string) (BulkDiscover, error) {
	switch s {
	case "all":
		return BulkDiscoverAll, nil
	case "onlyIncremental":
		return BulkDiscoverOnlyIncremental, nil
	case "onlyNonIncremental":
		return BulkDiscoverOnlyNonIncremental, nil
	case "none":
		return BulkDiscoverNone, nil
	}
	var t BulkDiscover
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BulkDiscover) Ptr() *BulkDiscover {
	return &b
}

type BulkExecutionStatus string

const (
	BulkExecutionStatusCreated    BulkExecutionStatus = "created"
	BulkExecutionStatusScheduled  BulkExecutionStatus = "scheduled"
	BulkExecutionStatusRunning    BulkExecutionStatus = "running"
	BulkExecutionStatusExporting  BulkExecutionStatus = "exporting"
	BulkExecutionStatusCanceling  BulkExecutionStatus = "canceling"
	BulkExecutionStatusCanceled   BulkExecutionStatus = "canceled"
	BulkExecutionStatusCompleted  BulkExecutionStatus = "completed"
	BulkExecutionStatusFailed     BulkExecutionStatus = "failed"
	BulkExecutionStatusProcessing BulkExecutionStatus = "processing"
	BulkExecutionStatusErrors     BulkExecutionStatus = "errors"
)

func NewBulkExecutionStatusFromString(s string) (BulkExecutionStatus, error) {
	switch s {
	case "created":
		return BulkExecutionStatusCreated, nil
	case "scheduled":
		return BulkExecutionStatusScheduled, nil
	case "running":
		return BulkExecutionStatusRunning, nil
	case "exporting":
		return BulkExecutionStatusExporting, nil
	case "canceling":
		return BulkExecutionStatusCanceling, nil
	case "canceled":
		return BulkExecutionStatusCanceled, nil
	case "completed":
		return BulkExecutionStatusCompleted, nil
	case "failed":
		return BulkExecutionStatusFailed, nil
	case "processing":
		return BulkExecutionStatusProcessing, nil
	case "errors":
		return BulkExecutionStatusErrors, nil
	}
	var t BulkExecutionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BulkExecutionStatus) Ptr() *BulkExecutionStatus {
	return &b
}

type BulkField struct {
	Enabled    *bool   `json:"enabled,omitempty" url:"enabled,omitempty"`
	Id         *string `json:"id,omitempty" url:"id,omitempty"`
	Obfuscated *bool   `json:"obfuscated,omitempty" url:"obfuscated,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkField) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkField(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkField) String() string {
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

type BulkFilter struct {
	// Schema field ID to filter on.
	FieldId  *string        `json:"field_id,omitempty" url:"field_id,omitempty"`
	Function FilterFunction `json:"function,omitempty" url:"function,omitempty"`
	Value    interface{}    `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkFilter) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkFilter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkFilter(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkFilter) String() string {
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

type BulkItemizedSchedule struct {
	Item     BulkSelectiveMode `json:"item,omitempty" url:"item,omitempty"`
	Schedule *BulkSchedule     `json:"schedule,omitempty" url:"schedule,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkItemizedSchedule) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkItemizedSchedule
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkItemizedSchedule(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkItemizedSchedule) String() string {
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

type BulkMultiScheduleConfiguration struct {
	Schedules []*BulkItemizedSchedule `json:"schedules,omitempty" url:"schedules,omitempty"`
	Type      *string                 `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkMultiScheduleConfiguration) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkMultiScheduleConfiguration
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkMultiScheduleConfiguration(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkMultiScheduleConfiguration) String() string {
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

type BulkSchedule struct {
	DayOfMonth *string                         `json:"day_of_month,omitempty" url:"day_of_month,omitempty"`
	DayOfWeek  *string                         `json:"day_of_week,omitempty" url:"day_of_week,omitempty"`
	Frequency  ScheduleFrequency               `json:"frequency,omitempty" url:"frequency,omitempty"`
	Hour       *string                         `json:"hour,omitempty" url:"hour,omitempty"`
	Minute     *string                         `json:"minute,omitempty" url:"minute,omitempty"`
	Month      *string                         `json:"month,omitempty" url:"month,omitempty"`
	Multi      *BulkMultiScheduleConfiguration `json:"multi,omitempty" url:"multi,omitempty"`

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

type BulkSchema struct {
	DataCutoffTimestamp *time.Time    `json:"data_cutoff_timestamp,omitempty" url:"data_cutoff_timestamp,omitempty"`
	DisableDataCutoff   *bool         `json:"disable_data_cutoff,omitempty" url:"disable_data_cutoff,omitempty"`
	Enabled             *bool         `json:"enabled,omitempty" url:"enabled,omitempty"`
	Fields              []*BulkField  `json:"fields,omitempty" url:"fields,omitempty"`
	Filters             []*BulkFilter `json:"filters,omitempty" url:"filters,omitempty"`
	Id                  *string       `json:"id,omitempty" url:"id,omitempty"`
	OutputName          *string       `json:"output_name,omitempty" url:"output_name,omitempty"`
	PartitionKey        *string       `json:"partition_key,omitempty" url:"partition_key,omitempty"`
	TrackingField       *string       `json:"tracking_field,omitempty" url:"tracking_field,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSchema) UnmarshalJSON(data []byte) error {
	type embed BulkSchema
	var unmarshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSchema(unmarshaler.embed)
	b.DataCutoffTimestamp = unmarshaler.DataCutoffTimestamp.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSchema) MarshalJSON() ([]byte, error) {
	type embed BulkSchema
	var marshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed:               embed(*b),
		DataCutoffTimestamp: core.NewOptionalDateTime(b.DataCutoffTimestamp),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSchema) String() string {
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

type BulkSchemaEnvelope struct {
	Data *BulkSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSchemaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSchemaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSchemaEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSchemaEnvelope) String() string {
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

type BulkSchemaExecutionStatus string

const (
	BulkSchemaExecutionStatusCreated   BulkSchemaExecutionStatus = "created"
	BulkSchemaExecutionStatusScheduled BulkSchemaExecutionStatus = "scheduled"
	BulkSchemaExecutionStatusRunning   BulkSchemaExecutionStatus = "running"
	BulkSchemaExecutionStatusExporting BulkSchemaExecutionStatus = "exporting"
	BulkSchemaExecutionStatusCanceled  BulkSchemaExecutionStatus = "canceled"
	BulkSchemaExecutionStatusCompleted BulkSchemaExecutionStatus = "completed"
	BulkSchemaExecutionStatusFailed    BulkSchemaExecutionStatus = "failed"
)

func NewBulkSchemaExecutionStatusFromString(s string) (BulkSchemaExecutionStatus, error) {
	switch s {
	case "created":
		return BulkSchemaExecutionStatusCreated, nil
	case "scheduled":
		return BulkSchemaExecutionStatusScheduled, nil
	case "running":
		return BulkSchemaExecutionStatusRunning, nil
	case "exporting":
		return BulkSchemaExecutionStatusExporting, nil
	case "canceled":
		return BulkSchemaExecutionStatusCanceled, nil
	case "completed":
		return BulkSchemaExecutionStatusCompleted, nil
	case "failed":
		return BulkSchemaExecutionStatusFailed, nil
	}
	var t BulkSchemaExecutionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BulkSchemaExecutionStatus) Ptr() *BulkSchemaExecutionStatus {
	return &b
}

type BulkSelectiveMode string

const (
	BulkSelectiveModeNone                 BulkSelectiveMode = "none"
	BulkSelectiveModeIncrementalFields    BulkSelectiveMode = "incrementalFields"
	BulkSelectiveModeNonincrementalFields BulkSelectiveMode = "nonincrementalFields"
)

func NewBulkSelectiveModeFromString(s string) (BulkSelectiveMode, error) {
	switch s {
	case "none":
		return BulkSelectiveModeNone, nil
	case "incrementalFields":
		return BulkSelectiveModeIncrementalFields, nil
	case "nonincrementalFields":
		return BulkSelectiveModeNonincrementalFields, nil
	}
	var t BulkSelectiveMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BulkSelectiveMode) Ptr() *BulkSelectiveMode {
	return &b
}

type BulkSyncCanceledEvent struct {
	DestinationConnectionId *string `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	ExecutionId             *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	Name                    *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId          *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SourceConnectionId      *string `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`
	SyncId                  *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncCanceledEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncCanceledEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncCanceledEvent(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncCanceledEvent) String() string {
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

type BulkSyncCompletedEvent struct {
	DestinationConnectionId *string `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	ExecutionId             *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	Name                    *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId          *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SourceConnectionId      *string `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`
	SyncId                  *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	TriggerSource           *string `json:"trigger_source,omitempty" url:"trigger_source,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncCompletedEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncCompletedEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncCompletedEvent(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncCompletedEvent) String() string {
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

type BulkSyncCompletedWithErrorEvent struct {
	DestinationConnectionId *string `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	ExecutionId             *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	Name                    *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId          *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SourceConnectionId      *string `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`
	SyncId                  *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	TriggerSource           *string `json:"trigger_source,omitempty" url:"trigger_source,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncCompletedWithErrorEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncCompletedWithErrorEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncCompletedWithErrorEvent(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncCompletedWithErrorEvent) String() string {
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

type BulkSyncDest struct {
	Configuration map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Modes         []*SupportedBulkMode   `json:"modes,omitempty" url:"modes,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncDest) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncDest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncDest(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncDest) String() string {
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

type BulkSyncDestEnvelope struct {
	Data *BulkSyncDest `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncDestEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncDestEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncDestEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncDestEnvelope) String() string {
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

type BulkSyncExecution struct {
	CompletedAt *time.Time                 `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	CreatedAt   *time.Time                 `json:"created_at,omitempty" url:"created_at,omitempty"`
	Id          *string                    `json:"id,omitempty" url:"id,omitempty"`
	IsResync    *bool                      `json:"is_resync,omitempty" url:"is_resync,omitempty"`
	IsTest      *bool                      `json:"is_test,omitempty" url:"is_test,omitempty"`
	Schemas     []*BulkSyncSchemaExecution `json:"schemas,omitempty" url:"schemas,omitempty"`
	StartedAt   *time.Time                 `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status      *BulkExecutionStatus       `json:"status,omitempty" url:"status,omitempty"`
	Type        *string                    `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncExecution) UnmarshalJSON(data []byte) error {
	type embed BulkSyncExecution
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncExecution(unmarshaler.embed)
	b.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	b.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	b.StartedAt = unmarshaler.StartedAt.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncExecution) MarshalJSON() ([]byte, error) {
	type embed BulkSyncExecution
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*b),
		CompletedAt: core.NewOptionalDateTime(b.CompletedAt),
		CreatedAt:   core.NewOptionalDateTime(b.CreatedAt),
		StartedAt:   core.NewOptionalDateTime(b.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncExecution) String() string {
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

type BulkSyncExecutionEnvelope struct {
	Data *BulkSyncExecution `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncExecutionEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncExecutionEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncExecutionEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncExecutionEnvelope) String() string {
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

type BulkSyncExecutionStatus struct {
	NextExecutionTime *time.Time                       `json:"nextExecutionTime,omitempty" url:"nextExecutionTime,omitempty"`
	Schemas           []*BulkSyncSchemaExecutionStatus `json:"schemas,omitempty" url:"schemas,omitempty"`
	Status            *BulkExecutionStatus             `json:"status,omitempty" url:"status,omitempty"`
	SyncId            *string                          `json:"sync_id,omitempty" url:"sync_id,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncExecutionStatus) UnmarshalJSON(data []byte) error {
	type embed BulkSyncExecutionStatus
	var unmarshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"nextExecutionTime,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncExecutionStatus(unmarshaler.embed)
	b.NextExecutionTime = unmarshaler.NextExecutionTime.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncExecutionStatus) MarshalJSON() ([]byte, error) {
	type embed BulkSyncExecutionStatus
	var marshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"nextExecutionTime,omitempty"`
	}{
		embed:             embed(*b),
		NextExecutionTime: core.NewOptionalDateTime(b.NextExecutionTime),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncExecutionStatus) String() string {
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

type BulkSyncFailedEvent struct {
	DestinationConnectionId *string `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	Error                   *string `json:"error,omitempty" url:"error,omitempty"`
	ExecutionId             *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	Name                    *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId          *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SourceConnectionId      *string `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`
	SyncId                  *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	TriggerSource           *string `json:"trigger_source,omitempty" url:"trigger_source,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncFailedEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncFailedEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncFailedEvent(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncFailedEvent) String() string {
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

type BulkSyncListEnvelope struct {
	Data []*BulkSyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncListEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncListEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncListEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncListEnvelope) String() string {
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

type BulkSyncResponse struct {
	Active                     *bool         `json:"active,omitempty" url:"active,omitempty"`
	AutomaticallyAddNewFields  *BulkDiscover `json:"automatically_add_new_fields,omitempty" url:"automatically_add_new_fields,omitempty"`
	AutomaticallyAddNewObjects *BulkDiscover `json:"automatically_add_new_objects,omitempty" url:"automatically_add_new_objects,omitempty"`
	DataCutoffTimestamp        *time.Time    `json:"data_cutoff_timestamp,omitempty" url:"data_cutoff_timestamp,omitempty"`
	// Destination-specific bulk sync configuration. e.g. output schema name, s3 file format, etc.
	DestinationConfiguration map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty"`
	DestinationConnectionId  *string                `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	DisableRecordTimestamps  *bool                  `json:"disable_record_timestamps,omitempty" url:"disable_record_timestamps,omitempty"`
	// DEPRECATED: Use automatically_add_new_objects/automatically_add_new_fields instead
	Discover *bool   `json:"discover,omitempty" url:"discover,omitempty"`
	Id       *string `json:"id,omitempty" url:"id,omitempty"`
	Mode     *string `json:"mode,omitempty" url:"mode,omitempty"`
	// Name of the bulk sync
	Name           *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	// List of permissions policies applied to the bulk sync.
	Policies []string      `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule *BulkSchedule `json:"schedule,omitempty" url:"schedule,omitempty"`
	// Source-specific bulk sync configuration. e.g. replication slot name, sync lookback, etc.
	SourceConfiguration map[string]interface{} `json:"source_configuration,omitempty" url:"source_configuration,omitempty"`
	SourceConnectionId  *string                `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncResponse) UnmarshalJSON(data []byte) error {
	type embed BulkSyncResponse
	var unmarshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncResponse(unmarshaler.embed)
	b.DataCutoffTimestamp = unmarshaler.DataCutoffTimestamp.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncResponse) MarshalJSON() ([]byte, error) {
	type embed BulkSyncResponse
	var marshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed:               embed(*b),
		DataCutoffTimestamp: core.NewOptionalDateTime(b.DataCutoffTimestamp),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncResponse) String() string {
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

type BulkSyncResponseEnvelope struct {
	Data *BulkSyncResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncResponseEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncResponseEnvelope) String() string {
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

type BulkSyncRunningEvent struct {
	DestinationConnectionId *string `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty"`
	ExecutionId             *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	Name                    *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId          *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SourceConnectionId      *string `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty"`
	SyncId                  *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncRunningEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncRunningEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncRunningEvent(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncRunningEvent) String() string {
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

type BulkSyncSchemaExecution struct {
	CompletedAt   *time.Time                 `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	ErrorCount    *int                       `json:"error_count,omitempty" url:"error_count,omitempty"`
	RecordCount   *int                       `json:"record_count,omitempty" url:"record_count,omitempty"`
	Schema        *string                    `json:"schema,omitempty" url:"schema,omitempty"`
	StartedAt     *time.Time                 `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status        *BulkSchemaExecutionStatus `json:"status,omitempty" url:"status,omitempty"`
	StatusMessage *string                    `json:"status_message,omitempty" url:"status_message,omitempty"`
	WarningCount  *int                       `json:"warning_count,omitempty" url:"warning_count,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSchemaExecution) UnmarshalJSON(data []byte) error {
	type embed BulkSyncSchemaExecution
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncSchemaExecution(unmarshaler.embed)
	b.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	b.StartedAt = unmarshaler.StartedAt.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSchemaExecution) MarshalJSON() ([]byte, error) {
	type embed BulkSyncSchemaExecution
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*b),
		CompletedAt: core.NewOptionalDateTime(b.CompletedAt),
		StartedAt:   core.NewOptionalDateTime(b.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncSchemaExecution) String() string {
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

type BulkSyncSchemaExecutionStatus struct {
	CompletedAt *time.Time `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	ErrorCount  *int       `json:"error_count,omitempty" url:"error_count,omitempty"`
	// ID of the most recent execution for the schema.
	ExecutionId   *string                    `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	RecordCount   *int                       `json:"record_count,omitempty" url:"record_count,omitempty"`
	Schema        *string                    `json:"schema,omitempty" url:"schema,omitempty"`
	StartedAt     *time.Time                 `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status        *BulkSchemaExecutionStatus `json:"status,omitempty" url:"status,omitempty"`
	StatusMessage *string                    `json:"status_message,omitempty" url:"status_message,omitempty"`
	WarningCount  *int                       `json:"warning_count,omitempty" url:"warning_count,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSchemaExecutionStatus) UnmarshalJSON(data []byte) error {
	type embed BulkSyncSchemaExecutionStatus
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncSchemaExecutionStatus(unmarshaler.embed)
	b.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	b.StartedAt = unmarshaler.StartedAt.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSchemaExecutionStatus) MarshalJSON() ([]byte, error) {
	type embed BulkSyncSchemaExecutionStatus
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*b),
		CompletedAt: core.NewOptionalDateTime(b.CompletedAt),
		StartedAt:   core.NewOptionalDateTime(b.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncSchemaExecutionStatus) String() string {
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

type BulkSyncSource struct {
	Configuration interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Schemas       []*Schema   `json:"schemas,omitempty" url:"schemas,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSource) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncSource(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSource) String() string {
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

type BulkSyncSourceEnvelope struct {
	Data *BulkSyncSource `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSourceEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncSourceEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncSourceEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSourceEnvelope) String() string {
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

type BulkSyncSourceSchemaEnvelope struct {
	Data *Schema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSourceSchemaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncSourceSchemaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncSourceSchemaEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSourceSchemaEnvelope) String() string {
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

type BulkSyncSourceStatus struct {
	CacheStatus         *string    `json:"cache_status,omitempty" url:"cache_status,omitempty"`
	LastRefreshFinished *time.Time `json:"last_refresh_finished,omitempty" url:"last_refresh_finished,omitempty"`
	LastRefreshStarted  *time.Time `json:"last_refresh_started,omitempty" url:"last_refresh_started,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSourceStatus) UnmarshalJSON(data []byte) error {
	type embed BulkSyncSourceStatus
	var unmarshaler = struct {
		embed
		LastRefreshFinished *core.DateTime `json:"last_refresh_finished,omitempty"`
		LastRefreshStarted  *core.DateTime `json:"last_refresh_started,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncSourceStatus(unmarshaler.embed)
	b.LastRefreshFinished = unmarshaler.LastRefreshFinished.TimePtr()
	b.LastRefreshStarted = unmarshaler.LastRefreshStarted.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSourceStatus) MarshalJSON() ([]byte, error) {
	type embed BulkSyncSourceStatus
	var marshaler = struct {
		embed
		LastRefreshFinished *core.DateTime `json:"last_refresh_finished,omitempty"`
		LastRefreshStarted  *core.DateTime `json:"last_refresh_started,omitempty"`
	}{
		embed:               embed(*b),
		LastRefreshFinished: core.NewOptionalDateTime(b.LastRefreshFinished),
		LastRefreshStarted:  core.NewOptionalDateTime(b.LastRefreshStarted),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncSourceStatus) String() string {
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

type BulkSyncSourceStatusEnvelope struct {
	Data *BulkSyncSourceStatus `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncSourceStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncSourceStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncSourceStatusEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncSourceStatusEnvelope) String() string {
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

type BulkSyncStatusEnvelope struct {
	Data *BulkSyncStatusResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncStatusEnvelope(value)
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncStatusEnvelope) String() string {
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

type BulkSyncStatusResponse struct {
	CurrentExecution  *BulkSyncExecution `json:"current_execution,omitempty" url:"current_execution,omitempty"`
	LastExecution     *BulkSyncExecution `json:"last_execution,omitempty" url:"last_execution,omitempty"`
	NextExecutionTime *time.Time         `json:"next_execution_time,omitempty" url:"next_execution_time,omitempty"`

	_rawJSON json.RawMessage
}

func (b *BulkSyncStatusResponse) UnmarshalJSON(data []byte) error {
	type embed BulkSyncStatusResponse
	var unmarshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed: embed(*b),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*b = BulkSyncStatusResponse(unmarshaler.embed)
	b.NextExecutionTime = unmarshaler.NextExecutionTime.TimePtr()
	b._rawJSON = json.RawMessage(data)
	return nil
}

func (b *BulkSyncStatusResponse) MarshalJSON() ([]byte, error) {
	type embed BulkSyncStatusResponse
	var marshaler = struct {
		embed
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty"`
	}{
		embed:             embed(*b),
		NextExecutionTime: core.NewOptionalDateTime(b.NextExecutionTime),
	}
	return json.Marshal(marshaler)
}

func (b *BulkSyncStatusResponse) String() string {
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

type CommonOutputActor struct {
	Id   *string `json:"id,omitempty" url:"id,omitempty"`
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	Type *string `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CommonOutputActor) UnmarshalJSON(data []byte) error {
	type unmarshaler CommonOutputActor
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CommonOutputActor(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CommonOutputActor) String() string {
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

type ConfigurationValue struct {
	Items []interface{} `json:"items,omitempty" url:"items,omitempty"`
	Type  *string       `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConfigurationValue) UnmarshalJSON(data []byte) error {
	type unmarshaler ConfigurationValue
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConfigurationValue(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConfigurationValue) String() string {
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

type ConnectCardResponse struct {
	RedirectUrl *string `json:"redirect_url,omitempty" url:"redirect_url,omitempty"`
	Token       *string `json:"token,omitempty" url:"token,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectCardResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectCardResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectCardResponse(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectCardResponse) String() string {
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

type ConnectCardResponseEnvelope struct {
	Data *ConnectCardResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectCardResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectCardResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectCardResponseEnvelope(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectCardResponseEnvelope) String() string {
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

type ConnectionListResponseEnvelope struct {
	Data []*ConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionListResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionListResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionListResponseEnvelope(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionListResponseEnvelope) String() string {
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

type ConnectionParameterValue struct {
	Label *string     `json:"label,omitempty" url:"label,omitempty"`
	Value interface{} `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionParameterValue) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionParameterValue
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionParameterValue(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionParameterValue) String() string {
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

type ConnectionParameterValuesResp struct {
	AllowsCreation *bool                       `json:"allows_creation,omitempty" url:"allows_creation,omitempty"`
	Values         []*ConnectionParameterValue `json:"values,omitempty" url:"values,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionParameterValuesResp) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionParameterValuesResp
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionParameterValuesResp(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionParameterValuesResp) String() string {
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

type ConnectionParameterValuesResponseEnvelope struct {
	Data map[string]*ConnectionParameterValuesResp `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionParameterValuesResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionParameterValuesResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionParameterValuesResponseEnvelope(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionParameterValuesResponseEnvelope) String() string {
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

type ConnectionResponseEnvelope struct {
	Data *ConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionResponseEnvelope(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionResponseEnvelope) String() string {
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

type ConnectionResponseSchema struct {
	// API calls made to service in the last 24h (supported integrations only).
	ApiCallsLast24Hours *int                   `json:"api_calls_last_24_hours,omitempty" url:"api_calls_last_24_hours,omitempty"`
	Configuration       map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Id                  *string                `json:"id,omitempty" url:"id,omitempty"`
	Name                *string                `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId      *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies            []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Status              *string                `json:"status,omitempty" url:"status,omitempty"`
	StatusError         *string                `json:"status_error,omitempty" url:"status_error,omitempty"`
	Type                *ConnectionTypeSchema  `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionResponseSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionResponseSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionResponseSchema(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionResponseSchema) String() string {
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

type ConnectionType struct {
	EnvConfig map[string]interface{} `json:"envConfig,omitempty" url:"envConfig,omitempty"`
	Id        *string                `json:"id,omitempty" url:"id,omitempty"`
	Name      *string                `json:"name,omitempty" url:"name,omitempty"`
	UseOauth  *bool                  `json:"use_oauth,omitempty" url:"use_oauth,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionType) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionType
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionType(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionType) String() string {
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

type ConnectionTypeResponseEnvelope struct {
	Data []*ConnectionType `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionTypeResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionTypeResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionTypeResponseEnvelope(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionTypeResponseEnvelope) String() string {
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

type ConnectionTypeSchema struct {
	Id         *string  `json:"id,omitempty" url:"id,omitempty"`
	Name       *string  `json:"name,omitempty" url:"name,omitempty"`
	Operations []string `json:"operations,omitempty" url:"operations,omitempty"`

	_rawJSON json.RawMessage
}

func (c *ConnectionTypeSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler ConnectionTypeSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConnectionTypeSchema(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ConnectionTypeSchema) String() string {
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

type CreateConnectionResponseEnvelope struct {
	Data *CreateConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CreateConnectionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConnectionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateConnectionResponseEnvelope(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateConnectionResponseEnvelope) String() string {
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

type CreateConnectionResponseSchema struct {
	// Code to enter in order to complete connection authentication.
	AuthCode *string `json:"auth_code,omitempty" url:"auth_code,omitempty"`
	// URL to visit to complete connection authentication.
	AuthUrl        *string                `json:"auth_url,omitempty" url:"auth_url,omitempty"`
	Configuration  map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Id             *string                `json:"id,omitempty" url:"id,omitempty"`
	Name           *string                `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Status         *string                `json:"status,omitempty" url:"status,omitempty"`
	StatusError    *string                `json:"status_error,omitempty" url:"status_error,omitempty"`
	Type           *ConnectionTypeSchema  `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CreateConnectionResponseSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateConnectionResponseSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateConnectionResponseSchema(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateConnectionResponseSchema) String() string {
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

type CreateModelRequest struct {
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
	Relations        []*ModelRelation          `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns  []string                  `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`

	_rawJSON json.RawMessage
}

func (c *CreateModelRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateModelRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateModelRequest(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateModelRequest) String() string {
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

type Enrichment struct {
	Configuration *V2EnricherConfiguration `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId  *string                  `json:"connection_id,omitempty" url:"connection_id,omitempty"`
	// Must be provided to update an existing enrichment
	EnricherId *string `json:"enricher_id,omitempty" url:"enricher_id,omitempty"`
	// If not provided, all fields will be enabled.
	Fields   []*ModelField      `json:"fields,omitempty" url:"fields,omitempty"`
	Mappings *V2EnricherMapping `json:"mappings,omitempty" url:"mappings,omitempty"`

	_rawJSON json.RawMessage
}

func (e *Enrichment) UnmarshalJSON(data []byte) error {
	type unmarshaler Enrichment
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = Enrichment(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *Enrichment) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type Event struct {
	CreatedAt      *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	Event          *EventBody `json:"event,omitempty" url:"event,omitempty"`
	Id             *string    `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string    `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Type           *string    `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type embed Event
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = Event(unmarshaler.embed)
	e.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *Event) MarshalJSON() ([]byte, error) {
	type embed Event
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*e),
		CreatedAt: core.NewOptionalDateTime(e.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (e *Event) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EventBody struct {
	SyncRunningEvent                *SyncRunningEvent
	SyncCompletedEvent              *SyncCompletedEvent
	SyncFailedEvent                 *SyncFailedEvent
	SyncCanceledEvent               *SyncCanceledEvent
	SyncCompletedWithErrorsEvent    *SyncCompletedWithErrorsEvent
	BulkSyncRunningEvent            *BulkSyncRunningEvent
	BulkSyncCompletedEvent          *BulkSyncCompletedEvent
	BulkSyncCanceledEvent           *BulkSyncCanceledEvent
	BulkSyncCompletedWithErrorEvent *BulkSyncCompletedWithErrorEvent
	BulkSyncFailedEvent             *BulkSyncFailedEvent
}

func NewEventBodyFromSyncRunningEvent(value *SyncRunningEvent) *EventBody {
	return &EventBody{SyncRunningEvent: value}
}

func NewEventBodyFromSyncCompletedEvent(value *SyncCompletedEvent) *EventBody {
	return &EventBody{SyncCompletedEvent: value}
}

func NewEventBodyFromSyncFailedEvent(value *SyncFailedEvent) *EventBody {
	return &EventBody{SyncFailedEvent: value}
}

func NewEventBodyFromSyncCanceledEvent(value *SyncCanceledEvent) *EventBody {
	return &EventBody{SyncCanceledEvent: value}
}

func NewEventBodyFromSyncCompletedWithErrorsEvent(value *SyncCompletedWithErrorsEvent) *EventBody {
	return &EventBody{SyncCompletedWithErrorsEvent: value}
}

func NewEventBodyFromBulkSyncRunningEvent(value *BulkSyncRunningEvent) *EventBody {
	return &EventBody{BulkSyncRunningEvent: value}
}

func NewEventBodyFromBulkSyncCompletedEvent(value *BulkSyncCompletedEvent) *EventBody {
	return &EventBody{BulkSyncCompletedEvent: value}
}

func NewEventBodyFromBulkSyncCanceledEvent(value *BulkSyncCanceledEvent) *EventBody {
	return &EventBody{BulkSyncCanceledEvent: value}
}

func NewEventBodyFromBulkSyncCompletedWithErrorEvent(value *BulkSyncCompletedWithErrorEvent) *EventBody {
	return &EventBody{BulkSyncCompletedWithErrorEvent: value}
}

func NewEventBodyFromBulkSyncFailedEvent(value *BulkSyncFailedEvent) *EventBody {
	return &EventBody{BulkSyncFailedEvent: value}
}

func (e *EventBody) UnmarshalJSON(data []byte) error {
	valueSyncRunningEvent := new(SyncRunningEvent)
	if err := json.Unmarshal(data, &valueSyncRunningEvent); err == nil {
		e.SyncRunningEvent = valueSyncRunningEvent
		return nil
	}
	valueSyncCompletedEvent := new(SyncCompletedEvent)
	if err := json.Unmarshal(data, &valueSyncCompletedEvent); err == nil {
		e.SyncCompletedEvent = valueSyncCompletedEvent
		return nil
	}
	valueSyncFailedEvent := new(SyncFailedEvent)
	if err := json.Unmarshal(data, &valueSyncFailedEvent); err == nil {
		e.SyncFailedEvent = valueSyncFailedEvent
		return nil
	}
	valueSyncCanceledEvent := new(SyncCanceledEvent)
	if err := json.Unmarshal(data, &valueSyncCanceledEvent); err == nil {
		e.SyncCanceledEvent = valueSyncCanceledEvent
		return nil
	}
	valueSyncCompletedWithErrorsEvent := new(SyncCompletedWithErrorsEvent)
	if err := json.Unmarshal(data, &valueSyncCompletedWithErrorsEvent); err == nil {
		e.SyncCompletedWithErrorsEvent = valueSyncCompletedWithErrorsEvent
		return nil
	}
	valueBulkSyncRunningEvent := new(BulkSyncRunningEvent)
	if err := json.Unmarshal(data, &valueBulkSyncRunningEvent); err == nil {
		e.BulkSyncRunningEvent = valueBulkSyncRunningEvent
		return nil
	}
	valueBulkSyncCompletedEvent := new(BulkSyncCompletedEvent)
	if err := json.Unmarshal(data, &valueBulkSyncCompletedEvent); err == nil {
		e.BulkSyncCompletedEvent = valueBulkSyncCompletedEvent
		return nil
	}
	valueBulkSyncCanceledEvent := new(BulkSyncCanceledEvent)
	if err := json.Unmarshal(data, &valueBulkSyncCanceledEvent); err == nil {
		e.BulkSyncCanceledEvent = valueBulkSyncCanceledEvent
		return nil
	}
	valueBulkSyncCompletedWithErrorEvent := new(BulkSyncCompletedWithErrorEvent)
	if err := json.Unmarshal(data, &valueBulkSyncCompletedWithErrorEvent); err == nil {
		e.BulkSyncCompletedWithErrorEvent = valueBulkSyncCompletedWithErrorEvent
		return nil
	}
	valueBulkSyncFailedEvent := new(BulkSyncFailedEvent)
	if err := json.Unmarshal(data, &valueBulkSyncFailedEvent); err == nil {
		e.BulkSyncFailedEvent = valueBulkSyncFailedEvent
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, e)
}

func (e EventBody) MarshalJSON() ([]byte, error) {
	if e.SyncRunningEvent != nil {
		return json.Marshal(e.SyncRunningEvent)
	}
	if e.SyncCompletedEvent != nil {
		return json.Marshal(e.SyncCompletedEvent)
	}
	if e.SyncFailedEvent != nil {
		return json.Marshal(e.SyncFailedEvent)
	}
	if e.SyncCanceledEvent != nil {
		return json.Marshal(e.SyncCanceledEvent)
	}
	if e.SyncCompletedWithErrorsEvent != nil {
		return json.Marshal(e.SyncCompletedWithErrorsEvent)
	}
	if e.BulkSyncRunningEvent != nil {
		return json.Marshal(e.BulkSyncRunningEvent)
	}
	if e.BulkSyncCompletedEvent != nil {
		return json.Marshal(e.BulkSyncCompletedEvent)
	}
	if e.BulkSyncCanceledEvent != nil {
		return json.Marshal(e.BulkSyncCanceledEvent)
	}
	if e.BulkSyncCompletedWithErrorEvent != nil {
		return json.Marshal(e.BulkSyncCompletedWithErrorEvent)
	}
	if e.BulkSyncFailedEvent != nil {
		return json.Marshal(e.BulkSyncFailedEvent)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EventBodyVisitor interface {
	VisitSyncRunningEvent(*SyncRunningEvent) error
	VisitSyncCompletedEvent(*SyncCompletedEvent) error
	VisitSyncFailedEvent(*SyncFailedEvent) error
	VisitSyncCanceledEvent(*SyncCanceledEvent) error
	VisitSyncCompletedWithErrorsEvent(*SyncCompletedWithErrorsEvent) error
	VisitBulkSyncRunningEvent(*BulkSyncRunningEvent) error
	VisitBulkSyncCompletedEvent(*BulkSyncCompletedEvent) error
	VisitBulkSyncCanceledEvent(*BulkSyncCanceledEvent) error
	VisitBulkSyncCompletedWithErrorEvent(*BulkSyncCompletedWithErrorEvent) error
	VisitBulkSyncFailedEvent(*BulkSyncFailedEvent) error
}

func (e *EventBody) Accept(visitor EventBodyVisitor) error {
	if e.SyncRunningEvent != nil {
		return visitor.VisitSyncRunningEvent(e.SyncRunningEvent)
	}
	if e.SyncCompletedEvent != nil {
		return visitor.VisitSyncCompletedEvent(e.SyncCompletedEvent)
	}
	if e.SyncFailedEvent != nil {
		return visitor.VisitSyncFailedEvent(e.SyncFailedEvent)
	}
	if e.SyncCanceledEvent != nil {
		return visitor.VisitSyncCanceledEvent(e.SyncCanceledEvent)
	}
	if e.SyncCompletedWithErrorsEvent != nil {
		return visitor.VisitSyncCompletedWithErrorsEvent(e.SyncCompletedWithErrorsEvent)
	}
	if e.BulkSyncRunningEvent != nil {
		return visitor.VisitBulkSyncRunningEvent(e.BulkSyncRunningEvent)
	}
	if e.BulkSyncCompletedEvent != nil {
		return visitor.VisitBulkSyncCompletedEvent(e.BulkSyncCompletedEvent)
	}
	if e.BulkSyncCanceledEvent != nil {
		return visitor.VisitBulkSyncCanceledEvent(e.BulkSyncCanceledEvent)
	}
	if e.BulkSyncCompletedWithErrorEvent != nil {
		return visitor.VisitBulkSyncCompletedWithErrorEvent(e.BulkSyncCompletedWithErrorEvent)
	}
	if e.BulkSyncFailedEvent != nil {
		return visitor.VisitBulkSyncFailedEvent(e.BulkSyncFailedEvent)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EventTypesEnvelope struct {
	Data []string `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (e *EventTypesEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler EventTypesEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EventTypesEnvelope(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EventTypesEnvelope) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type EventsEnvelope struct {
	Data []*Event `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (e *EventsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler EventsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EventsEnvelope(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *EventsEnvelope) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ExecutionCounts struct {
	Delete *int `json:"delete,omitempty" url:"delete,omitempty"`
	Error  *int `json:"error,omitempty" url:"error,omitempty"`
	Insert *int `json:"insert,omitempty" url:"insert,omitempty"`
	Total  *int `json:"total,omitempty" url:"total,omitempty"`
	Update *int `json:"update,omitempty" url:"update,omitempty"`

	_rawJSON json.RawMessage
}

func (e *ExecutionCounts) UnmarshalJSON(data []byte) error {
	type unmarshaler ExecutionCounts
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ExecutionCounts(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *ExecutionCounts) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ExecutionLogResponse struct {
	Expires *time.Time `json:"expires,omitempty" url:"expires,omitempty"`
	Urls    []string   `json:"urls,omitempty" url:"urls,omitempty"`

	_rawJSON json.RawMessage
}

func (e *ExecutionLogResponse) UnmarshalJSON(data []byte) error {
	type embed ExecutionLogResponse
	var unmarshaler = struct {
		embed
		Expires *core.DateTime `json:"expires,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = ExecutionLogResponse(unmarshaler.embed)
	e.Expires = unmarshaler.Expires.TimePtr()
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *ExecutionLogResponse) MarshalJSON() ([]byte, error) {
	type embed ExecutionLogResponse
	var marshaler = struct {
		embed
		Expires *core.DateTime `json:"expires,omitempty"`
	}{
		embed:   embed(*e),
		Expires: core.NewOptionalDateTime(e.Expires),
	}
	return json.Marshal(marshaler)
}

func (e *ExecutionLogResponse) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ExecutionLogsResponseEnvelope struct {
	Data *ExecutionLogResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (e *ExecutionLogsResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ExecutionLogsResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = ExecutionLogsResponseEnvelope(value)
	e._rawJSON = json.RawMessage(data)
	return nil
}

func (e *ExecutionLogsResponseEnvelope) String() string {
	if len(e._rawJSON) > 0 {
		if value, err := core.StringifyJSON(e._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

type ExecutionStatus string

const (
	ExecutionStatusCreated     ExecutionStatus = "created"
	ExecutionStatusScheduled   ExecutionStatus = "scheduled"
	ExecutionStatusQueued      ExecutionStatus = "queued"
	ExecutionStatusWaiting     ExecutionStatus = "waiting"
	ExecutionStatusRunning     ExecutionStatus = "running"
	ExecutionStatusProcessing  ExecutionStatus = "processing"
	ExecutionStatusCanceling   ExecutionStatus = "canceling"
	ExecutionStatusCanceled    ExecutionStatus = "canceled"
	ExecutionStatusCompleted   ExecutionStatus = "completed"
	ExecutionStatusFailed      ExecutionStatus = "failed"
	ExecutionStatusInterrupted ExecutionStatus = "interrupted"
)

func NewExecutionStatusFromString(s string) (ExecutionStatus, error) {
	switch s {
	case "created":
		return ExecutionStatusCreated, nil
	case "scheduled":
		return ExecutionStatusScheduled, nil
	case "queued":
		return ExecutionStatusQueued, nil
	case "waiting":
		return ExecutionStatusWaiting, nil
	case "running":
		return ExecutionStatusRunning, nil
	case "processing":
		return ExecutionStatusProcessing, nil
	case "canceling":
		return ExecutionStatusCanceling, nil
	case "canceled":
		return ExecutionStatusCanceled, nil
	case "completed":
		return ExecutionStatusCompleted, nil
	case "failed":
		return ExecutionStatusFailed, nil
	case "interrupted":
		return ExecutionStatusInterrupted, nil
	}
	var t ExecutionStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ExecutionStatus) Ptr() *ExecutionStatus {
	return &e
}

type FieldConfiguration struct {
	// Whether the field is enabled for syncing.
	Enabled *bool   `json:"enabled,omitempty" url:"enabled,omitempty"`
	Id      *string `json:"id,omitempty" url:"id,omitempty"`
	// Whether the field should be obfuscated.
	Obfuscate *bool `json:"obfuscate,omitempty" url:"obfuscate,omitempty"`

	_rawJSON json.RawMessage
}

func (f *FieldConfiguration) UnmarshalJSON(data []byte) error {
	type unmarshaler FieldConfiguration
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FieldConfiguration(value)
	f._rawJSON = json.RawMessage(data)
	return nil
}

func (f *FieldConfiguration) String() string {
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

type FilterFunction string

const (
	FilterFunctionEquality             FilterFunction = "Equality"
	FilterFunctionInequality           FilterFunction = "Inequality"
	FilterFunctionIsNull               FilterFunction = "IsNull"
	FilterFunctionIsNotNull            FilterFunction = "IsNotNull"
	FilterFunctionTrue                 FilterFunction = "True"
	FilterFunctionFalse                FilterFunction = "False"
	FilterFunctionOnOrAfter            FilterFunction = "OnOrAfter"
	FilterFunctionOnOrBefore           FilterFunction = "OnOrBefore"
	FilterFunctionGreaterThan          FilterFunction = "GreaterThan"
	FilterFunctionGreaterThanEqual     FilterFunction = "GreaterThanEqual"
	FilterFunctionLessThan             FilterFunction = "LessThan"
	FilterFunctionLessThanEqual        FilterFunction = "LessThanEqual"
	FilterFunctionStringContains       FilterFunction = "StringContains"
	FilterFunctionStringEndsWith       FilterFunction = "StringEndsWith"
	FilterFunctionStringDoesNotContain FilterFunction = "StringDoesNotContain"
	FilterFunctionStringDoesNotEndWith FilterFunction = "StringDoesNotEndWith"
	FilterFunctionStringOneOf          FilterFunction = "StringOneOf"
	FilterFunctionStringNotOneOf       FilterFunction = "StringNotOneOf"
	FilterFunctionBetween              FilterFunction = "Between"
	FilterFunctionArrayContains        FilterFunction = "ArrayContains"
	FilterFunctionArrayDoesNotContain  FilterFunction = "ArrayDoesNotContain"
	FilterFunctionInTheLast            FilterFunction = "InTheLast"
	FilterFunctionRelativeOnOrBefore   FilterFunction = "RelativeOnOrBefore"
	FilterFunctionRelativeOnOrAfter    FilterFunction = "RelativeOnOrAfter"
	FilterFunctionStringLike           FilterFunction = "StringLike"
	FilterFunctionStringNotLike        FilterFunction = "StringNotLike"
	FilterFunctionStringMatchesTrimmed FilterFunction = "StringMatchesTrimmed"
)

func NewFilterFunctionFromString(s string) (FilterFunction, error) {
	switch s {
	case "Equality":
		return FilterFunctionEquality, nil
	case "Inequality":
		return FilterFunctionInequality, nil
	case "IsNull":
		return FilterFunctionIsNull, nil
	case "IsNotNull":
		return FilterFunctionIsNotNull, nil
	case "True":
		return FilterFunctionTrue, nil
	case "False":
		return FilterFunctionFalse, nil
	case "OnOrAfter":
		return FilterFunctionOnOrAfter, nil
	case "OnOrBefore":
		return FilterFunctionOnOrBefore, nil
	case "GreaterThan":
		return FilterFunctionGreaterThan, nil
	case "GreaterThanEqual":
		return FilterFunctionGreaterThanEqual, nil
	case "LessThan":
		return FilterFunctionLessThan, nil
	case "LessThanEqual":
		return FilterFunctionLessThanEqual, nil
	case "StringContains":
		return FilterFunctionStringContains, nil
	case "StringEndsWith":
		return FilterFunctionStringEndsWith, nil
	case "StringDoesNotContain":
		return FilterFunctionStringDoesNotContain, nil
	case "StringDoesNotEndWith":
		return FilterFunctionStringDoesNotEndWith, nil
	case "StringOneOf":
		return FilterFunctionStringOneOf, nil
	case "StringNotOneOf":
		return FilterFunctionStringNotOneOf, nil
	case "Between":
		return FilterFunctionBetween, nil
	case "ArrayContains":
		return FilterFunctionArrayContains, nil
	case "ArrayDoesNotContain":
		return FilterFunctionArrayDoesNotContain, nil
	case "InTheLast":
		return FilterFunctionInTheLast, nil
	case "RelativeOnOrBefore":
		return FilterFunctionRelativeOnOrBefore, nil
	case "RelativeOnOrAfter":
		return FilterFunctionRelativeOnOrAfter, nil
	case "StringLike":
		return FilterFunctionStringLike, nil
	case "StringNotLike":
		return FilterFunctionStringNotLike, nil
	case "StringMatchesTrimmed":
		return FilterFunctionStringMatchesTrimmed, nil
	}
	var t FilterFunction
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (f FilterFunction) Ptr() *FilterFunction {
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

type GetExecutionResponseEnvelope struct {
	Data *GetExecutionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetExecutionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler GetExecutionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetExecutionResponseEnvelope(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetExecutionResponseEnvelope) String() string {
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

type GetExecutionResponseSchema struct {
	CompletedAt *time.Time       `json:"completed_at,omitempty" url:"completed_at,omitempty"`
	Counts      *ExecutionCounts `json:"counts,omitempty" url:"counts,omitempty"`
	CreatedAt   *time.Time       `json:"created_at,omitempty" url:"created_at,omitempty"`
	Errors      []string         `json:"errors,omitempty" url:"errors,omitempty"`
	Id          *string          `json:"id,omitempty" url:"id,omitempty"`
	StartedAt   *time.Time       `json:"started_at,omitempty" url:"started_at,omitempty"`
	Status      *ExecutionStatus `json:"status,omitempty" url:"status,omitempty"`
	Type        *string          `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetExecutionResponseSchema) UnmarshalJSON(data []byte) error {
	type embed GetExecutionResponseSchema
	var unmarshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed: embed(*g),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*g = GetExecutionResponseSchema(unmarshaler.embed)
	g.CompletedAt = unmarshaler.CompletedAt.TimePtr()
	g.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	g.StartedAt = unmarshaler.StartedAt.TimePtr()
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetExecutionResponseSchema) MarshalJSON() ([]byte, error) {
	type embed GetExecutionResponseSchema
	var marshaler = struct {
		embed
		CompletedAt *core.DateTime `json:"completed_at,omitempty"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty"`
		StartedAt   *core.DateTime `json:"started_at,omitempty"`
	}{
		embed:       embed(*g),
		CompletedAt: core.NewOptionalDateTime(g.CompletedAt),
		CreatedAt:   core.NewOptionalDateTime(g.CreatedAt),
		StartedAt:   core.NewOptionalDateTime(g.StartedAt),
	}
	return json.Marshal(marshaler)
}

func (g *GetExecutionResponseSchema) String() string {
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

type GetIdentityResponseEnvelope struct {
	Data *GetIdentityResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetIdentityResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler GetIdentityResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetIdentityResponseEnvelope(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetIdentityResponseEnvelope) String() string {
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

type GetIdentityResponseSchema struct {
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

func (g *GetIdentityResponseSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler GetIdentityResponseSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetIdentityResponseSchema(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetIdentityResponseSchema) String() string {
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

type GetModelSyncSourceMetaEnvelope struct {
	Data *ModelSyncSourceMetaResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (g *GetModelSyncSourceMetaEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler GetModelSyncSourceMetaEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetModelSyncSourceMetaEnvelope(value)
	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetModelSyncSourceMetaEnvelope) String() string {
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

type JobResponse struct {
	Error  *string         `json:"error,omitempty" url:"error,omitempty"`
	JobId  *string         `json:"job_id,omitempty" url:"job_id,omitempty"`
	Result interface{}     `json:"result,omitempty" url:"result,omitempty"`
	Status *WorkTaskStatus `json:"status,omitempty" url:"status,omitempty"`
	Type   *string         `json:"type,omitempty" url:"type,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler JobResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobResponse(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobResponse) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JobResponseEnvelope struct {
	Data *JobResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (j *JobResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler JobResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*j = JobResponseEnvelope(value)
	j._rawJSON = json.RawMessage(data)
	return nil
}

func (j *JobResponseEnvelope) String() string {
	if len(j._rawJSON) > 0 {
		if value, err := core.StringifyJSON(j._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(j); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", j)
}

type JsonschemaForm = map[string]interface{}

type LabelLabel = map[string]interface{}

type ListBulkSchema struct {
	Data []*BulkSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListBulkSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler ListBulkSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListBulkSchema(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListBulkSchema) String() string {
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

type ListBulkSyncExecutionStatusEnvelope struct {
	Data []*BulkSyncExecutionStatus `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListBulkSyncExecutionStatusEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ListBulkSyncExecutionStatusEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListBulkSyncExecutionStatusEnvelope(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListBulkSyncExecutionStatusEnvelope) String() string {
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

type ListBulkSyncExecutionsEnvelope struct {
	Data []*BulkSyncExecution `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListBulkSyncExecutionsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ListBulkSyncExecutionsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListBulkSyncExecutionsEnvelope(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListBulkSyncExecutionsEnvelope) String() string {
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

type ListExecutionResponseEnvelope struct {
	Data []*GetExecutionResponseSchema `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListExecutionResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ListExecutionResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListExecutionResponseEnvelope(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListExecutionResponseEnvelope) String() string {
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

type ListPoliciesResponseEnvelope struct {
	Data []*PolicyResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListPoliciesResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ListPoliciesResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListPoliciesResponseEnvelope(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListPoliciesResponseEnvelope) String() string {
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

type ListUsersEnvelope struct {
	Data []*User `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (l *ListUsersEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ListUsersEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListUsersEnvelope(value)
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListUsersEnvelope) String() string {
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

type ModelField struct {
	Description *string     `json:"description,omitempty" url:"description,omitempty"`
	Example     interface{} `json:"example,omitempty" url:"example,omitempty"`
	Id          *string     `json:"id,omitempty" url:"id,omitempty"`
	Label       *string     `json:"label,omitempty" url:"label,omitempty"`
	Name        *string     `json:"name,omitempty" url:"name,omitempty"`
	RemoteType  *string     `json:"remote_type,omitempty" url:"remote_type,omitempty"`
	Type        *string     `json:"type,omitempty" url:"type,omitempty"`
	Unique      *bool       `json:"unique,omitempty" url:"unique,omitempty"`
	UserAdded   *bool       `json:"user_added,omitempty" url:"user_added,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelField) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelField(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelField) String() string {
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

type ModelListResponseEnvelope struct {
	Data []*ModelResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelListResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelListResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelListResponseEnvelope(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelListResponseEnvelope) String() string {
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

type ModelModelFieldRequest struct {
	Example *string `json:"example,omitempty" url:"example,omitempty"`
	Label   string  `json:"label" url:"label"`
	Name    string  `json:"name" url:"name"`
	Type    string  `json:"type" url:"type"`

	_rawJSON json.RawMessage
}

func (m *ModelModelFieldRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelModelFieldRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelModelFieldRequest(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelModelFieldRequest) String() string {
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

type ModelRelation struct {
	From *string          `json:"from,omitempty" url:"from,omitempty"`
	To   *ModelRelationTo `json:"to,omitempty" url:"to,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelRelation) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelRelation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelRelation(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelRelation) String() string {
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

type ModelRelationTo struct {
	Field   *string `json:"field,omitempty" url:"field,omitempty"`
	ModelId *string `json:"model_id,omitempty" url:"model_id,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelRelationTo) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelRelationTo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelRelationTo(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelRelationTo) String() string {
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

type ModelResponse struct {
	Configuration   map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	ConnectionId    *string                `json:"connection_id,omitempty" url:"connection_id,omitempty"`
	CreatedAt       *time.Time             `json:"created_at,omitempty" url:"created_at,omitempty"`
	CreatedBy       *CommonOutputActor     `json:"created_by,omitempty" url:"created_by,omitempty"`
	Enricher        *Enrichment            `json:"enricher,omitempty" url:"enricher,omitempty"`
	Fields          []*ModelField          `json:"fields,omitempty" url:"fields,omitempty"`
	Id              *string                `json:"id,omitempty" url:"id,omitempty"`
	Identifier      *string                `json:"identifier,omitempty" url:"identifier,omitempty"`
	Labels          []LabelLabel           `json:"labels,omitempty" url:"labels,omitempty"`
	Name            *string                `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId  *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies        []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Relations       []*Relation            `json:"relations,omitempty" url:"relations,omitempty"`
	TrackingColumns []string               `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty"`
	Type            *string                `json:"type,omitempty" url:"type,omitempty"`
	UpdatedAt       *time.Time             `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UpdatedBy       *CommonOutputActor     `json:"updated_by,omitempty" url:"updated_by,omitempty"`
	Version         *int                   `json:"version,omitempty" url:"version,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelResponse) UnmarshalJSON(data []byte) error {
	type embed ModelResponse
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
	*m = ModelResponse(unmarshaler.embed)
	m.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	m.UpdatedAt = unmarshaler.UpdatedAt.TimePtr()
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelResponse) MarshalJSON() ([]byte, error) {
	type embed ModelResponse
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

func (m *ModelResponse) String() string {
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

type ModelResponseEnvelope struct {
	Data *ModelResponse `json:"data,omitempty" url:"data,omitempty"`
	Job  *JobResponse   `json:"job,omitempty" url:"job,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelResponseEnvelope(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelResponseEnvelope) String() string {
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

type ModelSample struct {
	Records  []V2SampleRecord `json:"records,omitempty" url:"records,omitempty"`
	Warnings []string         `json:"warnings,omitempty" url:"warnings,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelSample) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelSample
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelSample(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelSample) String() string {
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

type ModelSampleResponseEnvelope struct {
	Data *ModelSample `json:"data,omitempty" url:"data,omitempty"`
	Job  *JobResponse `json:"job,omitempty" url:"job,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelSampleResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelSampleResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelSampleResponseEnvelope(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelSampleResponseEnvelope) String() string {
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

type ModelSyncSourceMetaResponse struct {
	Configuration map[string]*ConfigurationValue `json:"configuration,omitempty" url:"configuration,omitempty"`
	Items         map[string]*SourceMeta         `json:"items,omitempty" url:"items,omitempty"`
	RequiresOneOf []string                       `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty"`

	_rawJSON json.RawMessage
}

func (m *ModelSyncSourceMetaResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelSyncSourceMetaResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelSyncSourceMetaResponse(value)
	m._rawJSON = json.RawMessage(data)
	return nil
}

func (m *ModelSyncSourceMetaResponse) String() string {
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

type Organization struct {
	Id        *string `json:"id,omitempty" url:"id,omitempty"`
	Issuer    *string `json:"issuer,omitempty" url:"issuer,omitempty"`
	Name      *string `json:"name,omitempty" url:"name,omitempty"`
	SsoDomain *string `json:"sso_domain,omitempty" url:"sso_domain,omitempty"`
	SsoOrgId  *string `json:"sso_org_id,omitempty" url:"sso_org_id,omitempty"`

	_rawJSON json.RawMessage
}

func (o *Organization) UnmarshalJSON(data []byte) error {
	type unmarshaler Organization
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = Organization(value)
	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *Organization) String() string {
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

type OrganizationEnvelope struct {
	Data *Organization `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (o *OrganizationEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler OrganizationEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrganizationEnvelope(value)
	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *OrganizationEnvelope) String() string {
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

type OrganizationsEnvelope struct {
	Data []*Organization `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (o *OrganizationsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler OrganizationsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrganizationsEnvelope(value)
	o._rawJSON = json.RawMessage(data)
	return nil
}

func (o *OrganizationsEnvelope) String() string {
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

type Pagination struct {
	// URL to the next page of results, if available. This may be returned as a host relative path.
	Next *string `json:"next,omitempty" url:"next,omitempty"`
	// URL to the previous page of results, if available. This may be returned as a host relative path.
	Previous *string `json:"previous,omitempty" url:"previous,omitempty"`

	_rawJSON json.RawMessage
}

func (p *Pagination) UnmarshalJSON(data []byte) error {
	type unmarshaler Pagination
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = Pagination(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *Pagination) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PickValue struct {
	Label *string `json:"label,omitempty" url:"label,omitempty"`
	Value *string `json:"value,omitempty" url:"value,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PickValue) UnmarshalJSON(data []byte) error {
	type unmarshaler PickValue
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PickValue(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PickValue) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PolicyAction struct {
	Action  string   `json:"action" url:"action"`
	RoleIds []string `json:"role_ids,omitempty" url:"role_ids,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PolicyAction) UnmarshalJSON(data []byte) error {
	type unmarshaler PolicyAction
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PolicyAction(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PolicyAction) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PolicyResponse struct {
	Id             *string         `json:"id,omitempty" url:"id,omitempty"`
	Name           *string         `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string         `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	PolicyActions  []*PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty"`
	System         *bool           `json:"system,omitempty" url:"system,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PolicyResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PolicyResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PolicyResponse(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PolicyResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type PolicyResponseEnvelope struct {
	Data *PolicyResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (p *PolicyResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler PolicyResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PolicyResponseEnvelope(value)
	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PolicyResponseEnvelope) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

type Relation struct {
	From *string     `json:"from,omitempty" url:"from,omitempty"`
	To   *RelationTo `json:"to,omitempty" url:"to,omitempty"`

	_rawJSON json.RawMessage
}

func (r *Relation) UnmarshalJSON(data []byte) error {
	type unmarshaler Relation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = Relation(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *Relation) String() string {
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

type RelationTo struct {
	Field   *string `json:"field,omitempty" url:"field,omitempty"`
	ModelId *string `json:"model_id,omitempty" url:"model_id,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RelationTo) UnmarshalJSON(data []byte) error {
	type unmarshaler RelationTo
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RelationTo(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RelationTo) String() string {
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

type RoleListResponseEnvelope struct {
	Data []*RoleResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RoleListResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler RoleListResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RoleListResponseEnvelope(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RoleListResponseEnvelope) String() string {
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

type RoleResponse struct {
	Id             *string `json:"id,omitempty" url:"id,omitempty"`
	Name           *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	System         *bool   `json:"system,omitempty" url:"system,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RoleResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RoleResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RoleResponse(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RoleResponse) String() string {
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

type RoleResponseEnvelope struct {
	Data *RoleResponse `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (r *RoleResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler RoleResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RoleResponseEnvelope(value)
	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RoleResponseEnvelope) String() string {
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

type ScheduleFrequency string

const (
	ScheduleFrequencyManual     ScheduleFrequency = "manual"
	ScheduleFrequencyContinuous ScheduleFrequency = "continuous"
	ScheduleFrequencyHourly     ScheduleFrequency = "hourly"
	ScheduleFrequencyDaily      ScheduleFrequency = "daily"
	ScheduleFrequencyWeekly     ScheduleFrequency = "weekly"
	ScheduleFrequencyCustom     ScheduleFrequency = "custom"
	ScheduleFrequencyBuilder    ScheduleFrequency = "builder"
	ScheduleFrequencyRunafter   ScheduleFrequency = "runafter"
	ScheduleFrequencyMulti      ScheduleFrequency = "multi"
	ScheduleFrequencyDbtcloud   ScheduleFrequency = "dbtcloud"
)

func NewScheduleFrequencyFromString(s string) (ScheduleFrequency, error) {
	switch s {
	case "manual":
		return ScheduleFrequencyManual, nil
	case "continuous":
		return ScheduleFrequencyContinuous, nil
	case "hourly":
		return ScheduleFrequencyHourly, nil
	case "daily":
		return ScheduleFrequencyDaily, nil
	case "weekly":
		return ScheduleFrequencyWeekly, nil
	case "custom":
		return ScheduleFrequencyCustom, nil
	case "builder":
		return ScheduleFrequencyBuilder, nil
	case "runafter":
		return ScheduleFrequencyRunafter, nil
	case "multi":
		return ScheduleFrequencyMulti, nil
	case "dbtcloud":
		return ScheduleFrequencyDbtcloud, nil
	}
	var t ScheduleFrequency
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s ScheduleFrequency) Ptr() *ScheduleFrequency {
	return &s
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

type Schema struct {
	Fields []*SchemaField `json:"fields,omitempty" url:"fields,omitempty"`
	Id     *string        `json:"id,omitempty" url:"id,omitempty"`
	Name   *string        `json:"name,omitempty" url:"name,omitempty"`

	_rawJSON json.RawMessage
}

func (s *Schema) UnmarshalJSON(data []byte) error {
	type unmarshaler Schema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Schema(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Schema) String() string {
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

type SchemaConfiguration struct {
	DataCutoffTimestamp *time.Time `json:"data_cutoff_timestamp,omitempty" url:"data_cutoff_timestamp,omitempty"`
	// Whether data cutoff is disabled for this schema.
	DisableDataCutoff *bool `json:"disable_data_cutoff,omitempty" url:"disable_data_cutoff,omitempty"`
	// Whether the schema is enabled for syncing.
	Enabled       *bool                              `json:"enabled,omitempty" url:"enabled,omitempty"`
	Fields        []*V2SchemaConfigurationFieldsItem `json:"fields,omitempty" url:"fields,omitempty"`
	Filters       []*BulkFilter                      `json:"filters,omitempty" url:"filters,omitempty"`
	Id            *string                            `json:"id,omitempty" url:"id,omitempty"`
	PartitionKey  *string                            `json:"partition_key,omitempty" url:"partition_key,omitempty"`
	TrackingField *string                            `json:"tracking_field,omitempty" url:"tracking_field,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SchemaConfiguration) UnmarshalJSON(data []byte) error {
	type embed SchemaConfiguration
	var unmarshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed: embed(*s),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*s = SchemaConfiguration(unmarshaler.embed)
	s.DataCutoffTimestamp = unmarshaler.DataCutoffTimestamp.TimePtr()
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SchemaConfiguration) MarshalJSON() ([]byte, error) {
	type embed SchemaConfiguration
	var marshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed:               embed(*s),
		DataCutoffTimestamp: core.NewOptionalDateTime(s.DataCutoffTimestamp),
	}
	return json.Marshal(marshaler)
}

func (s *SchemaConfiguration) String() string {
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

type SchemaField struct {
	Association *SchemaAssociation `json:"association,omitempty" url:"association,omitempty"`
	Id          *string            `json:"id,omitempty" url:"id,omitempty"`
	Name        *string            `json:"name,omitempty" url:"name,omitempty"`
	// The type of the field from the remote system.
	RemoteType *string        `json:"remote_type,omitempty" url:"remote_type,omitempty"`
	Type       *UtilFieldType `json:"type,omitempty" url:"type,omitempty"`
	TypeSpec   *TypesType     `json:"type_spec,omitempty" url:"type_spec,omitempty"`
	Values     []*PickValue   `json:"values,omitempty" url:"values,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SchemaField) UnmarshalJSON(data []byte) error {
	type unmarshaler SchemaField
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SchemaField(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SchemaField) String() string {
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

type SchemaRecordsResponseEnvelope struct {
	Data []map[string]interface{} `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SchemaRecordsResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler SchemaRecordsResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SchemaRecordsResponseEnvelope(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SchemaRecordsResponseEnvelope) String() string {
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

type SourceMeta struct {
	HasItems      *bool         `json:"has_items,omitempty" url:"has_items,omitempty"`
	Items         []interface{} `json:"items,omitempty" url:"items,omitempty"`
	RequiresOneOf []string      `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SourceMeta) UnmarshalJSON(data []byte) error {
	type unmarshaler SourceMeta
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SourceMeta(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SourceMeta) String() string {
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

type SupportedBulkMode struct {
	Description           *string   `json:"description,omitempty" url:"description,omitempty"`
	Id                    *SyncMode `json:"id,omitempty" url:"id,omitempty"`
	Label                 *string   `json:"label,omitempty" url:"label,omitempty"`
	RequiresIdentity      *bool     `json:"requires_identity,omitempty" url:"requires_identity,omitempty"`
	SupportsFieldSyncMode *bool     `json:"supports_field_sync_mode,omitempty" url:"supports_field_sync_mode,omitempty"`
	SupportsTargetFilters *bool     `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SupportedBulkMode) UnmarshalJSON(data []byte) error {
	type unmarshaler SupportedBulkMode
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SupportedBulkMode(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SupportedBulkMode) String() string {
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

type SyncCanceledEvent struct {
	ExecutionId        *string          `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	OrganizationId     *string          `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Status             *ExecutionStatus `json:"status,omitempty" url:"status,omitempty"`
	SyncId             *string          `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	SyncName           *string          `json:"sync_name,omitempty" url:"sync_name,omitempty"`
	TargetConnectionId *string          `json:"target_connection_id,omitempty" url:"target_connection_id,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncCanceledEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncCanceledEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncCanceledEvent(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncCanceledEvent) String() string {
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

type SyncCompletedEvent struct {
	DeletedRecords     []string         `json:"deleted_records,omitempty" url:"deleted_records,omitempty"`
	ErrorCount         *int             `json:"error_count,omitempty" url:"error_count,omitempty"`
	ErroredRecords     []string         `json:"errored_records,omitempty" url:"errored_records,omitempty"`
	ExecutionId        *string          `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	InsertedCount      *int             `json:"inserted_count,omitempty" url:"inserted_count,omitempty"`
	InsertedRecords    []string         `json:"inserted_records,omitempty" url:"inserted_records,omitempty"`
	Name               *string          `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId     *string          `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	RecordCount        *int             `json:"record_count,omitempty" url:"record_count,omitempty"`
	Status             *ExecutionStatus `json:"status,omitempty" url:"status,omitempty"`
	SyncId             *string          `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	TargetConnectionId *string          `json:"target_connection_id,omitempty" url:"target_connection_id,omitempty"`
	TotalRecords       []string         `json:"total_records,omitempty" url:"total_records,omitempty"`
	Trigger            *string          `json:"trigger,omitempty" url:"trigger,omitempty"`
	UpdatedCount       *int             `json:"updated_count,omitempty" url:"updated_count,omitempty"`
	UpdatedRecords     []string         `json:"updated_records,omitempty" url:"updated_records,omitempty"`
	UpsertedCount      *int             `json:"upserted_count,omitempty" url:"upserted_count,omitempty"`
	WarningCount       *int             `json:"warning_count,omitempty" url:"warning_count,omitempty"`
	Warnings           []string         `json:"warnings,omitempty" url:"warnings,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncCompletedEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncCompletedEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncCompletedEvent(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncCompletedEvent) String() string {
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

type SyncCompletedWithErrorsEvent struct {
	Error              *string `json:"error,omitempty" url:"error,omitempty"`
	ExecutionId        *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	NumberOfErrors     *int    `json:"number_of_errors,omitempty" url:"number_of_errors,omitempty"`
	NumberOfWarnings   *int    `json:"number_of_warnings,omitempty" url:"number_of_warnings,omitempty"`
	OrganizationId     *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SyncId             *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	SyncName           *string `json:"sync_name,omitempty" url:"sync_name,omitempty"`
	TargetConnectionId *string `json:"target_connection_id,omitempty" url:"target_connection_id,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncCompletedWithErrorsEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncCompletedWithErrorsEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncCompletedWithErrorsEvent(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncCompletedWithErrorsEvent) String() string {
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
	MappingsNotRequired          *bool   `json:"mappings_not_required,omitempty" url:"mappings_not_required,omitempty"`
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

type SyncFailedEvent struct {
	Error              *string `json:"error,omitempty" url:"error,omitempty"`
	ExecutionId        *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	OrganizationId     *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SyncId             *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	SyncName           *string `json:"sync_name,omitempty" url:"sync_name,omitempty"`
	TargetConnectionId *string `json:"target_connection_id,omitempty" url:"target_connection_id,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncFailedEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncFailedEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncFailedEvent(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncFailedEvent) String() string {
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

type SyncMode string

const (
	SyncModeCreate         SyncMode = "create"
	SyncModeUpdate         SyncMode = "update"
	SyncModeUpdateOrCreate SyncMode = "updateOrCreate"
	SyncModeReplace        SyncMode = "replace"
	SyncModeAppend         SyncMode = "append"
	SyncModeSnapshot       SyncMode = "snapshot"
	SyncModeReplicate      SyncMode = "replicate"
	SyncModeRemove         SyncMode = "remove"
)

func NewSyncModeFromString(s string) (SyncMode, error) {
	switch s {
	case "create":
		return SyncModeCreate, nil
	case "update":
		return SyncModeUpdate, nil
	case "updateOrCreate":
		return SyncModeUpdateOrCreate, nil
	case "replace":
		return SyncModeReplace, nil
	case "append":
		return SyncModeAppend, nil
	case "snapshot":
		return SyncModeSnapshot, nil
	case "replicate":
		return SyncModeReplicate, nil
	case "remove":
		return SyncModeRemove, nil
	}
	var t SyncMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (s SyncMode) Ptr() *SyncMode {
	return &s
}

type SyncRunningEvent struct {
	ExecutionId        *string `json:"execution_id,omitempty" url:"execution_id,omitempty"`
	Name               *string `json:"name,omitempty" url:"name,omitempty"`
	OrganizationId     *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	SyncId             *string `json:"sync_id,omitempty" url:"sync_id,omitempty"`
	TargetConnectionId *string `json:"target_connection_id,omitempty" url:"target_connection_id,omitempty"`

	_rawJSON json.RawMessage
}

func (s *SyncRunningEvent) UnmarshalJSON(data []byte) error {
	type unmarshaler SyncRunningEvent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SyncRunningEvent(value)
	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SyncRunningEvent) String() string {
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

type TypesType = interface{}

type User struct {
	Email          *string `json:"email,omitempty" url:"email,omitempty"`
	Id             *string `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Role           *string `json:"role,omitempty" url:"role,omitempty"`

	_rawJSON json.RawMessage
}

func (u *User) UnmarshalJSON(data []byte) error {
	type unmarshaler User
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = User(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *User) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UserEnvelope struct {
	Data *User `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (u *UserEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler UserEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserEnvelope(value)
	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UserEnvelope) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UtilFieldType string

const (
	UtilFieldTypeUnknown  UtilFieldType = "unknown"
	UtilFieldTypeString   UtilFieldType = "string"
	UtilFieldTypeNumber   UtilFieldType = "number"
	UtilFieldTypeBoolean  UtilFieldType = "boolean"
	UtilFieldTypeDatetime UtilFieldType = "datetime"
	UtilFieldTypeArray    UtilFieldType = "array"
	UtilFieldTypeObject   UtilFieldType = "object"
)

func NewUtilFieldTypeFromString(s string) (UtilFieldType, error) {
	switch s {
	case "unknown":
		return UtilFieldTypeUnknown, nil
	case "string":
		return UtilFieldTypeString, nil
	case "number":
		return UtilFieldTypeNumber, nil
	case "boolean":
		return UtilFieldTypeBoolean, nil
	case "datetime":
		return UtilFieldTypeDatetime, nil
	case "array":
		return UtilFieldTypeArray, nil
	case "object":
		return UtilFieldTypeObject, nil
	}
	var t UtilFieldType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (u UtilFieldType) Ptr() *UtilFieldType {
	return &u
}

// Similar to a model configuration, this configures the enricher. For example, if you wanted to use Apollo to enrich people, you would send `{"object": "people"}` as the configuration. Each enricher configuration can be found in the connection configuration docs.
type V2EnricherConfiguration = map[string]interface{}

// A map of parent model Source Name to child model Source Name. For example, if your model has a field called `work_email` and the enricher accepts a field called `email`, you'd send a map of `{"work_email":"email"}`. The set of required input mappings varies based on the configuration of the enrichment. You can use the `enrichment/{connection_id}/inputfields` API to discover available input field combinations for a given configuration.
type V2EnricherMapping = map[string]string

type V2ExecutionLogType string

const (
	V2ExecutionLogTypeRecords  V2ExecutionLogType = "records"
	V2ExecutionLogTypeErrors   V2ExecutionLogType = "errors"
	V2ExecutionLogTypeWarnings V2ExecutionLogType = "warnings"
	V2ExecutionLogTypeInserts  V2ExecutionLogType = "inserts"
	V2ExecutionLogTypeUpdates  V2ExecutionLogType = "updates"
	V2ExecutionLogTypeDeletes  V2ExecutionLogType = "deletes"
)

func NewV2ExecutionLogTypeFromString(s string) (V2ExecutionLogType, error) {
	switch s {
	case "records":
		return V2ExecutionLogTypeRecords, nil
	case "errors":
		return V2ExecutionLogTypeErrors, nil
	case "warnings":
		return V2ExecutionLogTypeWarnings, nil
	case "inserts":
		return V2ExecutionLogTypeInserts, nil
	case "updates":
		return V2ExecutionLogTypeUpdates, nil
	case "deletes":
		return V2ExecutionLogTypeDeletes, nil
	}
	var t V2ExecutionLogType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (v V2ExecutionLogType) Ptr() *V2ExecutionLogType {
	return &v
}

type V2GetEnrichmentInputFieldsResponseEnvelope struct {
	Data [][]string `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V2GetEnrichmentInputFieldsResponseEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V2GetEnrichmentInputFieldsResponseEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V2GetEnrichmentInputFieldsResponseEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V2GetEnrichmentInputFieldsResponseEnvelope) String() string {
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

// A map of `fieldSource` -> `fieldName: fieldValue`. Because there may be field name conflicts between the base model and enrichments, the base model fields are placed in a map under the model ID. Fields from enrichments are placed under the enricher ID.
type V2SampleRecord = map[string]map[string]interface{}

type V2SchemaConfigurationFieldsItem struct {
	String             string
	FieldConfiguration *FieldConfiguration
}

func NewV2SchemaConfigurationFieldsItemFromString(value string) *V2SchemaConfigurationFieldsItem {
	return &V2SchemaConfigurationFieldsItem{String: value}
}

func NewV2SchemaConfigurationFieldsItemFromFieldConfiguration(value *FieldConfiguration) *V2SchemaConfigurationFieldsItem {
	return &V2SchemaConfigurationFieldsItem{FieldConfiguration: value}
}

func (v *V2SchemaConfigurationFieldsItem) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		v.String = valueString
		return nil
	}
	valueFieldConfiguration := new(FieldConfiguration)
	if err := json.Unmarshal(data, &valueFieldConfiguration); err == nil {
		v.FieldConfiguration = valueFieldConfiguration
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, v)
}

func (v V2SchemaConfigurationFieldsItem) MarshalJSON() ([]byte, error) {
	if v.String != "" {
		return json.Marshal(v.String)
	}
	if v.FieldConfiguration != nil {
		return json.Marshal(v.FieldConfiguration)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V2SchemaConfigurationFieldsItemVisitor interface {
	VisitString(string) error
	VisitFieldConfiguration(*FieldConfiguration) error
}

func (v *V2SchemaConfigurationFieldsItem) Accept(visitor V2SchemaConfigurationFieldsItemVisitor) error {
	if v.String != "" {
		return visitor.VisitString(v.String)
	}
	if v.FieldConfiguration != nil {
		return visitor.VisitFieldConfiguration(v.FieldConfiguration)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", v)
}

type V4BulkSyncExecutionLogs = map[string]interface{}

type V4BulkSyncExecutionLogsEnvelope struct {
	Data *V4BulkSyncExecutionLogs `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4BulkSyncExecutionLogsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V4BulkSyncExecutionLogsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4BulkSyncExecutionLogsEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4BulkSyncExecutionLogsEnvelope) String() string {
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

type V4ExportSyncLogsEnvelope struct {
	Data *V4ExportSyncLogsResponse `json:"data,omitempty" url:"data,omitempty"`
	Job  *JobResponse              `json:"job,omitempty" url:"job,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4ExportSyncLogsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V4ExportSyncLogsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4ExportSyncLogsEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4ExportSyncLogsEnvelope) String() string {
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

type V4ExportSyncLogsResponse struct {
	Url *string `json:"url,omitempty" url:"url,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4ExportSyncLogsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler V4ExportSyncLogsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4ExportSyncLogsResponse(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4ExportSyncLogsResponse) String() string {
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

type V4QueryResultsEnvelope struct {
	Data  *V4RunQueryResult `json:"data,omitempty" url:"data,omitempty"`
	Links *Pagination       `json:"links,omitempty" url:"links,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4QueryResultsEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V4QueryResultsEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4QueryResultsEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4QueryResultsEnvelope) String() string {
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

type V4RunQueryEnvelope struct {
	Data *V4RunQueryResult `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4RunQueryEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler V4RunQueryEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4RunQueryEnvelope(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4RunQueryEnvelope) String() string {
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

type V4RunQueryResult struct {
	// The number of rows returned by the query. This will not be returned until the query completes.
	Count *int    `json:"count,omitempty" url:"count,omitempty"`
	Error *string `json:"error,omitempty" url:"error,omitempty"`
	// The time at which the query will expire and be deleted. This will not be returned until the query completes.
	Expires *string `json:"expires,omitempty" url:"expires,omitempty"`
	// The names of the fields returned by the query. This will not be returned until the query completes.
	Fields []string `json:"fields,omitempty" url:"fields,omitempty"`
	// The ID of the query task.
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The query results, returned as an array of objects.
	Results []map[string]interface{} `json:"results,omitempty" url:"results,omitempty"`
	Status  *WorkTaskStatus          `json:"status,omitempty" url:"status,omitempty"`

	_rawJSON json.RawMessage
}

func (v *V4RunQueryResult) UnmarshalJSON(data []byte) error {
	type unmarshaler V4RunQueryResult
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = V4RunQueryResult(value)
	v._rawJSON = json.RawMessage(data)
	return nil
}

func (v *V4RunQueryResult) String() string {
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

type Webhook struct {
	CreatedAt      *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	Endpoint       *string    `json:"endpoint,omitempty" url:"endpoint,omitempty"`
	Id             *string    `json:"id,omitempty" url:"id,omitempty"`
	OrganizationId *string    `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         *string    `json:"secret,omitempty" url:"secret,omitempty"`

	_rawJSON json.RawMessage
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	type embed Webhook
	var unmarshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed: embed(*w),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*w = Webhook(unmarshaler.embed)
	w.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *Webhook) MarshalJSON() ([]byte, error) {
	type embed Webhook
	var marshaler = struct {
		embed
		CreatedAt *core.DateTime `json:"created_at,omitempty"`
	}{
		embed:     embed(*w),
		CreatedAt: core.NewOptionalDateTime(w.CreatedAt),
	}
	return json.Marshal(marshaler)
}

func (w *Webhook) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookEnvelope struct {
	Data *Webhook `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (w *WebhookEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler WebhookEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebhookEnvelope(value)
	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WebhookEnvelope) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WebhookListEnvelope struct {
	Data []*Webhook `json:"data,omitempty" url:"data,omitempty"`

	_rawJSON json.RawMessage
}

func (w *WebhookListEnvelope) UnmarshalJSON(data []byte) error {
	type unmarshaler WebhookListEnvelope
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*w = WebhookListEnvelope(value)
	w._rawJSON = json.RawMessage(data)
	return nil
}

func (w *WebhookListEnvelope) String() string {
	if len(w._rawJSON) > 0 {
		if value, err := core.StringifyJSON(w._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(w); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", w)
}

type WorkTaskStatus string

const (
	WorkTaskStatusCreated WorkTaskStatus = "created"
	WorkTaskStatusRunning WorkTaskStatus = "running"
	WorkTaskStatusDone    WorkTaskStatus = "done"
	WorkTaskStatusFailed  WorkTaskStatus = "failed"
)

func NewWorkTaskStatusFromString(s string) (WorkTaskStatus, error) {
	switch s {
	case "created":
		return WorkTaskStatusCreated, nil
	case "running":
		return WorkTaskStatusRunning, nil
	case "done":
		return WorkTaskStatusDone, nil
	case "failed":
		return WorkTaskStatusFailed, nil
	}
	var t WorkTaskStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (w WorkTaskStatus) Ptr() *WorkTaskStatus {
	return &w
}
