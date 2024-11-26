// This file was auto-generated from our API Definition.

package bulksync

import (
	json "encoding/json"
	polytomicgo "github.com/polytomic/polytomic-go"
	core "github.com/polytomic/polytomic-go/core"
	time "time"
)

type SchemasListRequest struct {
	Filters map[string]*string `json:"-" url:"filters,omitempty"`
}

type BulkSyncSchemasRequest struct {
	Schemas []*polytomicgo.BulkSchema `json:"schemas,omitempty" url:"schemas,omitempty"`
}

type UpdateBulkSchema struct {
	DataCutoffTimestamp *time.Time                `json:"data_cutoff_timestamp,omitempty" url:"data_cutoff_timestamp,omitempty"`
	DisableDataCutoff   *bool                     `json:"disable_data_cutoff,omitempty" url:"disable_data_cutoff,omitempty"`
	Enabled             *bool                     `json:"enabled,omitempty" url:"enabled,omitempty"`
	Fields              []*polytomicgo.BulkField  `json:"fields,omitempty" url:"fields,omitempty"`
	Filters             []*polytomicgo.BulkFilter `json:"filters,omitempty" url:"filters,omitempty"`
	PartitionKey        *string                   `json:"partition_key,omitempty" url:"partition_key,omitempty"`
}

func (u *UpdateBulkSchema) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateBulkSchema
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*u = UpdateBulkSchema(body)
	return nil
}

func (u *UpdateBulkSchema) MarshalJSON() ([]byte, error) {
	type embed UpdateBulkSchema
	var marshaler = struct {
		embed
		DataCutoffTimestamp *core.DateTime `json:"data_cutoff_timestamp,omitempty"`
	}{
		embed:               embed(*u),
		DataCutoffTimestamp: core.NewOptionalDateTime(u.DataCutoffTimestamp),
	}
	return json.Marshal(marshaler)
}
