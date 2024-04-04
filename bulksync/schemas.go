// This file was auto-generated from our API Definition.

package bulksync

import (
	polytomicgo "github.com/polytomic/polytomic-go"
)

type SchemasListRequest struct {
	Filters map[string]*string `json:"-" url:"filters,omitempty"`
}

type BulkSyncSchemasRequest struct {
	Schemas []*polytomicgo.BulkSchema `json:"schemas,omitempty" url:"schemas,omitempty"`
}

type UpdateBulkSchema struct {
	Enabled      *bool                    `json:"enabled,omitempty" url:"enabled,omitempty"`
	Fields       []*polytomicgo.BulkField `json:"fields,omitempty" url:"fields,omitempty"`
	PartitionKey *string                  `json:"partition_key,omitempty" url:"partition_key,omitempty"`
}
