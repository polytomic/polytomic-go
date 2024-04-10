// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
	fmt "fmt"
	time "time"

	core "github.com/polytomic/polytomic-go/core"
)

type ApiKeyResponse struct {
	Value *string `json:"value,omitempty" url:"value,omitempty" tfsdk:"value"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *ApiKeyResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *ActivateSyncOutput `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Active bool `json:"active" url:"active" tfsdk:"active"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Active *bool   `json:"active,omitempty" url:"active,omitempty" tfsdk:"active"`
	Id     *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Message  *string     `json:"message,omitempty" url:"message,omitempty"`
	Metadata interface{} `json:"metadata,omitempty" url:"metadata,omitempty"`
	Status   *int        `json:"status,omitempty" url:"status,omitempty"`

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
	Enabled    *bool   `json:"enabled,omitempty" url:"enabled,omitempty" tfsdk:"enabled"`
	Id         *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Obfuscated *bool   `json:"obfuscated,omitempty" url:"obfuscated,omitempty" tfsdk:"obfuscated"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type BulkItemizedSchedule struct {
	Item     BulkSelectiveMode `json:"item,omitempty" url:"item,omitempty" tfsdk:"item"`
	Schedule *BulkSchedule     `json:"schedule,omitempty" url:"schedule,omitempty" tfsdk:"schedule"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Schedules []*BulkItemizedSchedule `json:"schedules,omitempty" url:"schedules,omitempty" tfsdk:"schedules"`
	Type      string                  `json:"type" url:"type" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	DayOfMonth *string                         `json:"day_of_month,omitempty" url:"day_of_month,omitempty" tfsdk:"day_of_month"`
	DayOfWeek  *string                         `json:"day_of_week,omitempty" url:"day_of_week,omitempty" tfsdk:"day_of_week"`
	Frequency  ScheduleFrequency               `json:"frequency,omitempty" url:"frequency,omitempty" tfsdk:"frequency"`
	Hour       *string                         `json:"hour,omitempty" url:"hour,omitempty" tfsdk:"hour"`
	Minute     *string                         `json:"minute,omitempty" url:"minute,omitempty" tfsdk:"minute"`
	Month      *string                         `json:"month,omitempty" url:"month,omitempty" tfsdk:"month"`
	Multi      *BulkMultiScheduleConfiguration `json:"multi,omitempty" url:"multi,omitempty" tfsdk:"multi"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Enabled      *bool        `json:"enabled,omitempty" url:"enabled,omitempty" tfsdk:"enabled"`
	Fields       []*BulkField `json:"fields,omitempty" url:"fields,omitempty" tfsdk:"fields"`
	Id           *string      `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	OutputName   *string      `json:"output_name,omitempty" url:"output_name,omitempty" tfsdk:"output_name"`
	PartitionKey *string      `json:"partition_key,omitempty" url:"partition_key,omitempty" tfsdk:"partition_key"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (b *BulkSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSchema
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSchema(value)
	b._rawJSON = json.RawMessage(data)
	return nil
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
	Data *BulkSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type BulkSyncDest struct {
	Configuration map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	Modes         []*SupportedMode       `json:"modes,omitempty" url:"modes,omitempty" tfsdk:"modes"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *BulkSyncDest `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	CompletedAt *time.Time                 `json:"completed_at,omitempty" url:"completed_at,omitempty" tfsdk:"completed_at"`
	CreatedAt   *time.Time                 `json:"created_at,omitempty" url:"created_at,omitempty" tfsdk:"created_at"`
	Id          *string                    `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	IsResync    *bool                      `json:"is_resync,omitempty" url:"is_resync,omitempty" tfsdk:"is_resync"`
	IsTest      *bool                      `json:"is_test,omitempty" url:"is_test,omitempty" tfsdk:"is_test"`
	Schemas     []*BulkSyncSchemaExecution `json:"schemas,omitempty" url:"schemas,omitempty" tfsdk:"schemas"`
	StartedAt   *time.Time                 `json:"started_at,omitempty" url:"started_at,omitempty" tfsdk:"started_at"`
	Status      *BulkExecutionStatus       `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`
	Type        *string                    `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (b *BulkSyncExecution) UnmarshalJSON(data []byte) error {
	type embed BulkSyncExecution
	var unmarshaler = struct {
		embed       `tfsdk:"embed"`
		CompletedAt *core.DateTime `json:"completed_at,omitempty" tfsdk:"completed_at"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
		StartedAt   *core.DateTime `json:"started_at,omitempty" tfsdk:"started_at"`
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
		embed       `tfsdk:"embed"`
		CompletedAt *core.DateTime `json:"completed_at,omitempty" tfsdk:"completed_at"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
		StartedAt   *core.DateTime `json:"started_at,omitempty" tfsdk:"started_at"`
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
	Data *BulkSyncExecution `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type BulkSyncListEnvelope struct {
	Data []*BulkSyncResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Active                   *bool                  `json:"active,omitempty" url:"active,omitempty" tfsdk:"active"`
	DestinationConfiguration map[string]interface{} `json:"destination_configuration,omitempty" url:"destination_configuration,omitempty" tfsdk:"destination_configuration"`
	DestinationConnectionId  *string                `json:"destination_connection_id,omitempty" url:"destination_connection_id,omitempty" tfsdk:"destination_connection_id"`
	Discover                 *bool                  `json:"discover,omitempty" url:"discover,omitempty" tfsdk:"discover"`
	Id                       *string                `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Mode                     *string                `json:"mode,omitempty" url:"mode,omitempty" tfsdk:"mode"`
	Name                     *string                `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId           *string                `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Policies                 []string               `json:"policies,omitempty" url:"policies,omitempty" tfsdk:"policies"`
	Schedule                 *BulkSchedule          `json:"schedule,omitempty" url:"schedule,omitempty" tfsdk:"schedule"`
	SourceConfiguration      map[string]interface{} `json:"source_configuration,omitempty" url:"source_configuration,omitempty" tfsdk:"source_configuration"`
	SourceConnectionId       *string                `json:"source_connection_id,omitempty" url:"source_connection_id,omitempty" tfsdk:"source_connection_id"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (b *BulkSyncResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BulkSyncResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BulkSyncResponse(value)
	b._rawJSON = json.RawMessage(data)
	return nil
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
	Data *BulkSyncResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type BulkSyncSchemaExecution struct {
	CompletedAt   *time.Time                 `json:"completed_at,omitempty" url:"completed_at,omitempty" tfsdk:"completed_at"`
	ErrorCount    *int                       `json:"error_count,omitempty" url:"error_count,omitempty" tfsdk:"error_count"`
	RecordCount   *int                       `json:"record_count,omitempty" url:"record_count,omitempty" tfsdk:"record_count"`
	Schema        *string                    `json:"schema,omitempty" url:"schema,omitempty" tfsdk:"schema"`
	StartedAt     *time.Time                 `json:"started_at,omitempty" url:"started_at,omitempty" tfsdk:"started_at"`
	Status        *BulkSchemaExecutionStatus `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`
	StatusMessage *string                    `json:"status_message,omitempty" url:"status_message,omitempty" tfsdk:"status_message"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (b *BulkSyncSchemaExecution) UnmarshalJSON(data []byte) error {
	type embed BulkSyncSchemaExecution
	var unmarshaler = struct {
		embed       `tfsdk:"embed"`
		CompletedAt *core.DateTime `json:"completed_at,omitempty" tfsdk:"completed_at"`
		StartedAt   *core.DateTime `json:"started_at,omitempty" tfsdk:"started_at"`
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
		embed       `tfsdk:"embed"`
		CompletedAt *core.DateTime `json:"completed_at,omitempty" tfsdk:"completed_at"`
		StartedAt   *core.DateTime `json:"started_at,omitempty" tfsdk:"started_at"`
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

type BulkSyncSource struct {
	Configuration interface{} `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	Schemas       []*Schema   `json:"schemas,omitempty" url:"schemas,omitempty" tfsdk:"schemas"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *BulkSyncSource `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *Schema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	CacheStatus         *string    `json:"cache_status,omitempty" url:"cache_status,omitempty" tfsdk:"cache_status"`
	LastRefreshFinished *time.Time `json:"last_refresh_finished,omitempty" url:"last_refresh_finished,omitempty" tfsdk:"last_refresh_finished"`
	LastRefreshStarted  *time.Time `json:"last_refresh_started,omitempty" url:"last_refresh_started,omitempty" tfsdk:"last_refresh_started"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (b *BulkSyncSourceStatus) UnmarshalJSON(data []byte) error {
	type embed BulkSyncSourceStatus
	var unmarshaler = struct {
		embed               `tfsdk:"embed"`
		LastRefreshFinished *core.DateTime `json:"last_refresh_finished,omitempty" tfsdk:"last_refresh_finished"`
		LastRefreshStarted  *core.DateTime `json:"last_refresh_started,omitempty" tfsdk:"last_refresh_started"`
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
		embed               `tfsdk:"embed"`
		LastRefreshFinished *core.DateTime `json:"last_refresh_finished,omitempty" tfsdk:"last_refresh_finished"`
		LastRefreshStarted  *core.DateTime `json:"last_refresh_started,omitempty" tfsdk:"last_refresh_started"`
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
	Data *BulkSyncSourceStatus `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *BulkSyncStatusResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	CurrentExecution  *BulkSyncExecution `json:"current_execution,omitempty" url:"current_execution,omitempty" tfsdk:"current_execution"`
	LastExecution     *BulkSyncExecution `json:"last_execution,omitempty" url:"last_execution,omitempty" tfsdk:"last_execution"`
	NextExecutionTime *time.Time         `json:"next_execution_time,omitempty" url:"next_execution_time,omitempty" tfsdk:"next_execution_time"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (b *BulkSyncStatusResponse) UnmarshalJSON(data []byte) error {
	type embed BulkSyncStatusResponse
	var unmarshaler = struct {
		embed             `tfsdk:"embed"`
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty" tfsdk:"next_execution_time"`
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
		embed             `tfsdk:"embed"`
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty" tfsdk:"next_execution_time"`
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

type ConfigurationValue struct {
	Items []interface{} `json:"items,omitempty" url:"items,omitempty" tfsdk:"items"`
	Type  *string       `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	RedirectUrl *string `json:"redirect_url,omitempty" url:"redirect_url,omitempty" tfsdk:"redirect_url"`
	Token       *string `json:"token,omitempty" url:"token,omitempty" tfsdk:"token"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *ConnectCardResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*ConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	HasItems      *bool         `json:"has_items,omitempty" url:"has_items,omitempty" tfsdk:"has_items"`
	Items         []interface{} `json:"items,omitempty" url:"items,omitempty" tfsdk:"items"`
	RequiresOneOf []string      `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty" tfsdk:"requires_one_of"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Configuration map[string]*ConfigurationValue `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	Items         map[string]*ConnectionMeta     `json:"items,omitempty" url:"items,omitempty" tfsdk:"items"`
	RequiresOneOf []string                       `json:"requires_one_of,omitempty" url:"requires_one_of,omitempty" tfsdk:"requires_one_of"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Label *string     `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`
	Value interface{} `json:"value,omitempty" url:"value,omitempty" tfsdk:"value"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	AllowsCreation *bool                       `json:"allows_creation,omitempty" url:"allows_creation,omitempty" tfsdk:"allows_creation"`
	Values         []*ConnectionParameterValue `json:"values,omitempty" url:"values,omitempty" tfsdk:"values"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data map[string]*ConnectionParameterValuesResp `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *ConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Configuration  map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	Id             *string                `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name           *string                `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId *string                `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Policies       []string               `json:"policies,omitempty" url:"policies,omitempty" tfsdk:"policies"`
	Status         *string                `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`
	StatusError    *string                `json:"status_error,omitempty" url:"status_error,omitempty" tfsdk:"status_error"`
	Type           *ConnectionTypeSchema  `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	EnvConfig map[string]interface{} `json:"envConfig,omitempty" url:"envConfig,omitempty" tfsdk:"env_config"`
	Id        *string                `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name      *string                `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	UseOauth  *bool                  `json:"use_oauth,omitempty" url:"use_oauth,omitempty" tfsdk:"use_oauth"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*ConnectionType `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Id         *string  `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name       *string  `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	Operations []string `json:"operations,omitempty" url:"operations,omitempty" tfsdk:"operations"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *CreateConnectionResponseSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	AuthCode *string `json:"auth_code,omitempty" url:"auth_code,omitempty" tfsdk:"auth_code"`
	// URL to visit to complete connection authentication.
	AuthUrl        *string                `json:"auth_url,omitempty" url:"auth_url,omitempty" tfsdk:"auth_url"`
	Configuration  map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	Id             *string                `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name           *string                `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId *string                `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Policies       []string               `json:"policies,omitempty" url:"policies,omitempty" tfsdk:"policies"`
	Status         *string                `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`
	StatusError    *string                `json:"status_error,omitempty" url:"status_error,omitempty" tfsdk:"status_error"`
	Type           *ConnectionTypeSchema  `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type Event struct {
	CreatedAt      *time.Time  `json:"created_at,omitempty" url:"created_at,omitempty" tfsdk:"created_at"`
	Event          interface{} `json:"event,omitempty" url:"event,omitempty" tfsdk:"event"`
	Id             *string     `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	OrganizationId *string     `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Type           *string     `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type embed Event
	var unmarshaler = struct {
		embed     `tfsdk:"embed"`
		CreatedAt *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
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
		embed     `tfsdk:"embed"`
		CreatedAt *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
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

type EventTypesEnvelope struct {
	Data []string `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*Event `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Error  *int `json:"error,omitempty" url:"error,omitempty" tfsdk:"error"`
	Insert *int `json:"insert,omitempty" url:"insert,omitempty" tfsdk:"insert"`
	Total  *int `json:"total,omitempty" url:"total,omitempty" tfsdk:"total"`
	Update *int `json:"update,omitempty" url:"update,omitempty" tfsdk:"update"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Expires *time.Time `json:"expires,omitempty" url:"expires,omitempty" tfsdk:"expires"`
	Urls    []string   `json:"urls,omitempty" url:"urls,omitempty" tfsdk:"urls"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (e *ExecutionLogResponse) UnmarshalJSON(data []byte) error {
	type embed ExecutionLogResponse
	var unmarshaler = struct {
		embed   `tfsdk:"embed"`
		Expires *core.DateTime `json:"expires,omitempty" tfsdk:"expires"`
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
		embed   `tfsdk:"embed"`
		Expires *core.DateTime `json:"expires,omitempty" tfsdk:"expires"`
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
	Data *ExecutionLogResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type Filter struct {
	FieldId   string      `json:"field_id" url:"field_id" tfsdk:"field_id"`
	FieldType string      `json:"field_type" url:"field_type" tfsdk:"field_type"`
	Function  string      `json:"function" url:"function" tfsdk:"function"`
	Label     *string     `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`
	Value     interface{} `json:"value,omitempty" url:"value,omitempty" tfsdk:"value"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type GetConnectionMetaEnvelope struct {
	Data *ConnectionMetaResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	CompletedAt *time.Time       `json:"completed_at,omitempty" url:"completed_at,omitempty" tfsdk:"completed_at"`
	Counts      *ExecutionCounts `json:"counts,omitempty" url:"counts,omitempty" tfsdk:"counts"`
	CreatedAt   *time.Time       `json:"created_at,omitempty" url:"created_at,omitempty" tfsdk:"created_at"`
	Errors      []string         `json:"errors,omitempty" url:"errors,omitempty" tfsdk:"errors"`
	Id          *string          `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	StartedAt   *time.Time       `json:"started_at,omitempty" url:"started_at,omitempty" tfsdk:"started_at"`
	Status      *ExecutionStatus `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`
	Type        *string          `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (g *GetExecutionResponseSchema) UnmarshalJSON(data []byte) error {
	type embed GetExecutionResponseSchema
	var unmarshaler = struct {
		embed       `tfsdk:"embed"`
		CompletedAt *core.DateTime `json:"completed_at,omitempty" tfsdk:"completed_at"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
		StartedAt   *core.DateTime `json:"started_at,omitempty" tfsdk:"started_at"`
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
		embed       `tfsdk:"embed"`
		CompletedAt *core.DateTime `json:"completed_at,omitempty" tfsdk:"completed_at"`
		CreatedAt   *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
		StartedAt   *core.DateTime `json:"started_at,omitempty" tfsdk:"started_at"`
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
	Data *GetIdentityResponseSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Email            *string `json:"email,omitempty" url:"email,omitempty" tfsdk:"email"`
	Id               *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	IsOrganization   *bool   `json:"is_organization,omitempty" url:"is_organization,omitempty" tfsdk:"is_organization"`
	IsPartner        *bool   `json:"is_partner,omitempty" url:"is_partner,omitempty" tfsdk:"is_partner"`
	IsSystem         *bool   `json:"is_system,omitempty" url:"is_system,omitempty" tfsdk:"is_system"`
	IsUser           *bool   `json:"is_user,omitempty" url:"is_user,omitempty" tfsdk:"is_user"`
	OrganizationId   *string `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	OrganizationName *string `json:"organization_name,omitempty" url:"organization_name,omitempty" tfsdk:"organization_name"`
	Role             *string `json:"role,omitempty" url:"role,omitempty" tfsdk:"role"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type Identity struct {
	Function          string  `json:"function" url:"function" tfsdk:"function"`
	NewField          *bool   `json:"new_field,omitempty" url:"new_field,omitempty" tfsdk:"new_field"`
	RemoteFieldTypeId *string `json:"remote_field_type_id,omitempty" url:"remote_field_type_id,omitempty" tfsdk:"remote_field_type_id"`
	Source            *Source `json:"source,omitempty" url:"source,omitempty" tfsdk:"source"`
	Target            string  `json:"target" url:"target" tfsdk:"target"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Id    *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Label *string `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Error  *string         `json:"error,omitempty" url:"error,omitempty" tfsdk:"error"`
	JobId  *string         `json:"job_id,omitempty" url:"job_id,omitempty" tfsdk:"job_id"`
	Result interface{}     `json:"result,omitempty" url:"result,omitempty" tfsdk:"result"`
	Status *WorkTaskStatus `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`
	Type   *string         `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *JobResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*BulkSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type ListBulkSyncExecutionsEnvelope struct {
	Data []*BulkSyncExecution `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*GetExecutionResponseSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*ModelSyncResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*PolicyResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*User `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Description           *string `json:"description,omitempty" url:"description,omitempty" tfsdk:"description"`
	Label                 *string `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`
	Mode                  *string `json:"mode,omitempty" url:"mode,omitempty" tfsdk:"mode"`
	RequiresIdentity      *bool   `json:"requires_identity,omitempty" url:"requires_identity,omitempty" tfsdk:"requires_identity"`
	SupportsFieldSyncMode *bool   `json:"supports_field_sync_mode,omitempty" url:"supports_field_sync_mode,omitempty" tfsdk:"supports_field_sync_mode"`
	SupportsTargetFilters *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty" tfsdk:"supports_target_filters"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Description *string     `json:"description,omitempty" url:"description,omitempty" tfsdk:"description"`
	Example     interface{} `json:"example,omitempty" url:"example,omitempty" tfsdk:"example"`
	Label       *string     `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`
	Name        *string     `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	RemoteType  *string     `json:"remote_type,omitempty" url:"remote_type,omitempty" tfsdk:"remote_type"`
	Type        *string     `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`
	Unique      *bool       `json:"unique,omitempty" url:"unique,omitempty" tfsdk:"unique"`
	UserAdded   *bool       `json:"user_added,omitempty" url:"user_added,omitempty" tfsdk:"user_added"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*ModelField `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*ModelResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Example *string `json:"example,omitempty" url:"example,omitempty" tfsdk:"example"`
	Label   string  `json:"label" url:"label" tfsdk:"label"`
	Name    string  `json:"name" url:"name" tfsdk:"name"`
	Type    string  `json:"type" url:"type" tfsdk:"type"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	From *string          `json:"from,omitempty" url:"from,omitempty" tfsdk:"from"`
	To   *ModelRelationTo `json:"to,omitempty" url:"to,omitempty" tfsdk:"to"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Field   *string `json:"field,omitempty" url:"field,omitempty" tfsdk:"field"`
	ModelId *string `json:"model_id,omitempty" url:"model_id,omitempty" tfsdk:"model_id"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Configuration   map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	ConnectionId    *string                `json:"connection_id,omitempty" url:"connection_id,omitempty" tfsdk:"connection_id"`
	Fields          []*ModelField          `json:"fields,omitempty" url:"fields,omitempty" tfsdk:"fields"`
	Id              *string                `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Identifier      *string                `json:"identifier,omitempty" url:"identifier,omitempty" tfsdk:"identifier"`
	Labels          []LabelLabel           `json:"labels,omitempty" url:"labels,omitempty" tfsdk:"labels"`
	Name            *string                `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId  *string                `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Policies        []string               `json:"policies,omitempty" url:"policies,omitempty" tfsdk:"policies"`
	Relations       []*Relation            `json:"relations,omitempty" url:"relations,omitempty" tfsdk:"relations"`
	TrackingColumns []string               `json:"tracking_columns,omitempty" url:"tracking_columns,omitempty" tfsdk:"tracking_columns"`
	Type            *string                `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`
	Version         *int                   `json:"version,omitempty" url:"version,omitempty" tfsdk:"version"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (m *ModelResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelResponse(value)
	m._rawJSON = json.RawMessage(data)
	return nil
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
	Data *ModelResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`
	Job  *JobResponse   `json:"job,omitempty" url:"job,omitempty" tfsdk:"job"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type ModelSyncField struct {
	New           *bool   `json:"new,omitempty" url:"new,omitempty" tfsdk:"new"`
	OverrideValue *string `json:"override_value,omitempty" url:"override_value,omitempty" tfsdk:"override_value"`
	Source        *Source `json:"source,omitempty" url:"source,omitempty" tfsdk:"source"`
	SyncMode      *string `json:"sync_mode,omitempty" url:"sync_mode,omitempty" tfsdk:"sync_mode"`
	Target        string  `json:"target" url:"target" tfsdk:"target"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Active         *bool             `json:"active,omitempty" url:"active,omitempty" tfsdk:"active"`
	Fields         []*ModelSyncField `json:"fields,omitempty" url:"fields,omitempty" tfsdk:"fields"`
	FilterLogic    *string           `json:"filter_logic,omitempty" url:"filter_logic,omitempty" tfsdk:"filter_logic"`
	Filters        []*Filter         `json:"filters,omitempty" url:"filters,omitempty" tfsdk:"filters"`
	Id             *string           `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Identity       *Identity         `json:"identity,omitempty" url:"identity,omitempty" tfsdk:"identity"`
	Mode           *string           `json:"mode,omitempty" url:"mode,omitempty" tfsdk:"mode"`
	Name           *string           `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId *string           `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	OverrideFields []*ModelSyncField `json:"override_fields,omitempty" url:"override_fields,omitempty" tfsdk:"override_fields"`
	Overrides      []*Override       `json:"overrides,omitempty" url:"overrides,omitempty" tfsdk:"overrides"`
	Policies       []string          `json:"policies,omitempty" url:"policies,omitempty" tfsdk:"policies"`
	Schedule       *Schedule         `json:"schedule,omitempty" url:"schedule,omitempty" tfsdk:"schedule"`
	SyncAllRecords *bool             `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty" tfsdk:"sync_all_records"`
	Target         *Target           `json:"target,omitempty" url:"target,omitempty" tfsdk:"target"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (m *ModelSyncResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ModelSyncResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = ModelSyncResponse(value)
	m._rawJSON = json.RawMessage(data)
	return nil
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
	Data *ModelSyncResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type Organization struct {
	Id        *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Issuer    *string `json:"issuer,omitempty" url:"issuer,omitempty" tfsdk:"issuer"`
	Name      *string `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	SsoDomain *string `json:"sso_domain,omitempty" url:"sso_domain,omitempty" tfsdk:"sso_domain"`
	SsoOrgId  *string `json:"sso_org_id,omitempty" url:"sso_org_id,omitempty" tfsdk:"sso_org_id"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *Organization `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*Organization `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type Override struct {
	FieldId  *string     `json:"field_id,omitempty" url:"field_id,omitempty" tfsdk:"field_id"`
	Function *string     `json:"function,omitempty" url:"function,omitempty" tfsdk:"function"`
	Override interface{} `json:"override,omitempty" url:"override,omitempty" tfsdk:"override"`
	Value    interface{} `json:"value,omitempty" url:"value,omitempty" tfsdk:"value"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type PickValue struct {
	Label *string `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`
	Value *string `json:"value,omitempty" url:"value,omitempty" tfsdk:"value"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Action  string   `json:"action" url:"action" tfsdk:"action"`
	RoleIds []string `json:"role_ids,omitempty" url:"role_ids,omitempty" tfsdk:"role_ids"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Id             *string         `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name           *string         `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId *string         `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	PolicyActions  []*PolicyAction `json:"policy_actions,omitempty" url:"policy_actions,omitempty" tfsdk:"policy_actions"`
	System         *bool           `json:"system,omitempty" url:"system,omitempty" tfsdk:"system"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *PolicyResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	From *string     `json:"from,omitempty" url:"from,omitempty" tfsdk:"from"`
	To   *RelationTo `json:"to,omitempty" url:"to,omitempty" tfsdk:"to"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Field   *string `json:"field,omitempty" url:"field,omitempty" tfsdk:"field"`
	ModelId *string `json:"model_id,omitempty" url:"model_id,omitempty" tfsdk:"model_id"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Code *int `json:"code,omitempty" url:"code,omitempty" tfsdk:"code"`
	// Application context.
	Context map[string]interface{} `json:"context,omitempty" url:"context,omitempty" tfsdk:"context"`
	// Error message.
	Error *string `json:"error,omitempty" url:"error,omitempty" tfsdk:"error"`
	// Status text.
	Status *string `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*RoleResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Id             *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name           *string `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	System         *bool   `json:"system,omitempty" url:"system,omitempty" tfsdk:"system"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *RoleResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	BulkSyncIds []string `json:"bulk_sync_ids,omitempty" url:"bulk_sync_ids,omitempty" tfsdk:"bulk_sync_ids"`
	SyncIds     []string `json:"sync_ids,omitempty" url:"sync_ids,omitempty" tfsdk:"sync_ids"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	ConnectionId *string   `json:"connection_id,omitempty" url:"connection_id,omitempty" tfsdk:"connection_id"`
	DayOfMonth   *string   `json:"day_of_month,omitempty" url:"day_of_month,omitempty" tfsdk:"day_of_month"`
	DayOfWeek    *string   `json:"day_of_week,omitempty" url:"day_of_week,omitempty" tfsdk:"day_of_week"`
	Frequency    *string   `json:"frequency,omitempty" url:"frequency,omitempty" tfsdk:"frequency"`
	Hour         *string   `json:"hour,omitempty" url:"hour,omitempty" tfsdk:"hour"`
	JobId        *int      `json:"job_id,omitempty" url:"job_id,omitempty" tfsdk:"job_id"`
	Minute       *string   `json:"minute,omitempty" url:"minute,omitempty" tfsdk:"minute"`
	Month        *string   `json:"month,omitempty" url:"month,omitempty" tfsdk:"month"`
	RunAfter     *RunAfter `json:"run_after,omitempty" url:"run_after,omitempty" tfsdk:"run_after"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	ScheduleOptions []*ScheduleScheduleOption `json:"schedule_options,omitempty" url:"schedule_options,omitempty" tfsdk:"schedule_options"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *ScheduleOptionResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Configuration *JsonschemaForm    `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	Description   *string            `json:"description,omitempty" url:"description,omitempty" tfsdk:"description"`
	Frequency     *ScheduleFrequency `json:"frequency,omitempty" url:"frequency,omitempty" tfsdk:"frequency"`
	Label         *string            `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Fields []*SchemaField `json:"fields,omitempty" url:"fields,omitempty" tfsdk:"fields"`
	Id     *string        `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name   *string        `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Id              *string  `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name            *string  `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	ReferenceTo     []string `json:"reference_to,omitempty" url:"reference_to,omitempty" tfsdk:"reference_to"`
	ReferencedField *string  `json:"referenced_field,omitempty" url:"referenced_field,omitempty" tfsdk:"referenced_field"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type SchemaField struct {
	Association *SchemaAssociation `json:"association,omitempty" url:"association,omitempty" tfsdk:"association"`
	Id          *string            `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Name        *string            `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	RemoteType  *string            `json:"remote_type,omitempty" url:"remote_type,omitempty" tfsdk:"remote_type"`
	Type        *string            `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`
	Values      []*PickValue       `json:"values,omitempty" url:"values,omitempty" tfsdk:"values"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type SchemaRecordsResponseEnvelope struct {
	Data []map[string]interface{} `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Field   string `json:"field" url:"field" tfsdk:"field"`
	ModelId string `json:"model_id" url:"model_id" tfsdk:"model_id"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *StartModelSyncResponseSchema `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	CreatedAt *time.Time       `json:"created_at,omitempty" url:"created_at,omitempty" tfsdk:"created_at"`
	Id        *string          `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Status    *ExecutionStatus `json:"status,omitempty" url:"status,omitempty" tfsdk:"status"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (s *StartModelSyncResponseSchema) UnmarshalJSON(data []byte) error {
	type embed StartModelSyncResponseSchema
	var unmarshaler = struct {
		embed     `tfsdk:"embed"`
		CreatedAt *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
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
		embed     `tfsdk:"embed"`
		CreatedAt *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
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
	Description           *string `json:"description,omitempty" url:"description,omitempty" tfsdk:"description"`
	Id                    *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Label                 *string `json:"label,omitempty" url:"label,omitempty" tfsdk:"label"`
	RequiresIdentity      *bool   `json:"requires_identity,omitempty" url:"requires_identity,omitempty" tfsdk:"requires_identity"`
	SupportsFieldSyncMode *bool   `json:"supports_field_sync_mode,omitempty" url:"supports_field_sync_mode,omitempty" tfsdk:"supports_field_sync_mode"`
	SupportsTargetFilters *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty" tfsdk:"supports_target_filters"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	DoesNotReportOperationCounts *bool   `json:"does_not_report_operation_counts,omitempty" url:"does_not_report_operation_counts,omitempty" tfsdk:"does_not_report_operation_counts"`
	NewTargetLabel               *string `json:"new_target_label,omitempty" url:"new_target_label,omitempty" tfsdk:"new_target_label"`
	OptionalTargetMappings       *bool   `json:"optional_target_mappings,omitempty" url:"optional_target_mappings,omitempty" tfsdk:"optional_target_mappings"`
	PrimaryMetadataObject        *string `json:"primary_metadata_object,omitempty" url:"primary_metadata_object,omitempty" tfsdk:"primary_metadata_object"`
	RequiresConfiguration        *bool   `json:"requires_configuration,omitempty" url:"requires_configuration,omitempty" tfsdk:"requires_configuration"`
	SupportsFieldCreation        *bool   `json:"supports_field_creation,omitempty" url:"supports_field_creation,omitempty" tfsdk:"supports_field_creation"`
	SupportsFieldTypeSelection   *bool   `json:"supports_field_type_selection,omitempty" url:"supports_field_type_selection,omitempty" tfsdk:"supports_field_type_selection"`
	SupportsTargetFilters        *bool   `json:"supports_target_filters,omitempty" url:"supports_target_filters,omitempty" tfsdk:"supports_target_filters"`
	TargetCreator                *bool   `json:"target_creator,omitempty" url:"target_creator,omitempty" tfsdk:"target_creator"`
	UseFieldNamesAsLabels        *bool   `json:"use_field_names_as_labels,omitempty" url:"use_field_names_as_labels,omitempty" tfsdk:"use_field_names_as_labels"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *SyncStatusResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	CurrentExecution  *GetExecutionResponseSchema `json:"current_execution,omitempty" url:"current_execution,omitempty" tfsdk:"current_execution"`
	LastExecution     *GetExecutionResponseSchema `json:"last_execution,omitempty" url:"last_execution,omitempty" tfsdk:"last_execution"`
	NextExecutionTime *time.Time                  `json:"next_execution_time,omitempty" url:"next_execution_time,omitempty" tfsdk:"next_execution_time"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (s *SyncStatusResponse) UnmarshalJSON(data []byte) error {
	type embed SyncStatusResponse
	var unmarshaler = struct {
		embed             `tfsdk:"embed"`
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty" tfsdk:"next_execution_time"`
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
		embed             `tfsdk:"embed"`
		NextExecutionTime *core.DateTime `json:"next_execution_time,omitempty" tfsdk:"next_execution_time"`
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
	Configuration map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty" tfsdk:"configuration"`
	ConnectionId  string                 `json:"connection_id" url:"connection_id" tfsdk:"connection_id"`
	FilterLogic   *string                `json:"filter_logic,omitempty" url:"filter_logic,omitempty" tfsdk:"filter_logic"`
	NewName       *string                `json:"new_name,omitempty" url:"new_name,omitempty" tfsdk:"new_name"`
	Object        string                 `json:"object" url:"object" tfsdk:"object"`
	SearchValues  map[string]interface{} `json:"search_values,omitempty" url:"search_values,omitempty" tfsdk:"search_values"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Association       *bool               `json:"association,omitempty" url:"association,omitempty" tfsdk:"association"`
	Createable        *bool               `json:"createable,omitempty" url:"createable,omitempty" tfsdk:"createable"`
	Description       *string             `json:"description,omitempty" url:"description,omitempty" tfsdk:"description"`
	Filterable        *bool               `json:"filterable,omitempty" url:"filterable,omitempty" tfsdk:"filterable"`
	Id                *string             `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	IdentityFunctions []*IdentityFunction `json:"identity_functions,omitempty" url:"identity_functions,omitempty" tfsdk:"identity_functions"`
	Name              *string             `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	Required          *bool               `json:"required,omitempty" url:"required,omitempty" tfsdk:"required"`
	SourceType        *string             `json:"source_type,omitempty" url:"source_type,omitempty" tfsdk:"source_type"`
	SupportsIdentity  *bool               `json:"supports_identity,omitempty" url:"supports_identity,omitempty" tfsdk:"supports_identity"`
	Type              *string             `json:"type,omitempty" url:"type,omitempty" tfsdk:"type"`
	Updateable        *bool               `json:"updateable,omitempty" url:"updateable,omitempty" tfsdk:"updateable"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type TargetResponse struct {
	Fields      []*TargetField             `json:"fields,omitempty" url:"fields,omitempty" tfsdk:"fields"`
	Id          *string                    `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	Modes       []*Mode                    `json:"modes,omitempty" url:"modes,omitempty" tfsdk:"modes"`
	Name        *string                    `json:"name,omitempty" url:"name,omitempty" tfsdk:"name"`
	Properties  *SyncDestinationProperties `json:"properties,omitempty" url:"properties,omitempty" tfsdk:"properties"`
	RefreshedAt *time.Time                 `json:"refreshed_at,omitempty" url:"refreshed_at,omitempty" tfsdk:"refreshed_at"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (t *TargetResponse) UnmarshalJSON(data []byte) error {
	type embed TargetResponse
	var unmarshaler = struct {
		embed       `tfsdk:"embed"`
		RefreshedAt *core.DateTime `json:"refreshed_at,omitempty" tfsdk:"refreshed_at"`
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
		embed       `tfsdk:"embed"`
		RefreshedAt *core.DateTime `json:"refreshed_at,omitempty" tfsdk:"refreshed_at"`
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
	Data *TargetResponse `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type User struct {
	Email          *string `json:"email,omitempty" url:"email,omitempty" tfsdk:"email"`
	Id             *string `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Role           *string `json:"role,omitempty" url:"role,omitempty" tfsdk:"role"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data *User `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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

type Webhook struct {
	CreatedAt      *time.Time `json:"created_at,omitempty" url:"created_at,omitempty" tfsdk:"created_at"`
	Endpoint       *string    `json:"endpoint,omitempty" url:"endpoint,omitempty" tfsdk:"endpoint"`
	Id             *string    `json:"id,omitempty" url:"id,omitempty" tfsdk:"id"`
	OrganizationId *string    `json:"organization_id,omitempty" url:"organization_id,omitempty" tfsdk:"organization_id"`
	Secret         *string    `json:"secret,omitempty" url:"secret,omitempty" tfsdk:"secret"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
}

func (w *Webhook) UnmarshalJSON(data []byte) error {
	type embed Webhook
	var unmarshaler = struct {
		embed     `tfsdk:"embed"`
		CreatedAt *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
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
		embed     `tfsdk:"embed"`
		CreatedAt *core.DateTime `json:"created_at,omitempty" tfsdk:"created_at"`
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
	Data *Webhook `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	Data []*Webhook `json:"data,omitempty" url:"data,omitempty" tfsdk:"data"`

	_rawJSON json.RawMessage `tfsdk:"__raw_json"`
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
	WorkTaskStatusRunning WorkTaskStatus = "running"
	WorkTaskStatusDone    WorkTaskStatus = "done"
	WorkTaskStatusFailed  WorkTaskStatus = "failed"
)

func NewWorkTaskStatusFromString(s string) (WorkTaskStatus, error) {
	switch s {
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
