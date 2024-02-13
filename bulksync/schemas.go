// This file was auto-generated from our API Definition.

package bulksync

import (
	polytomicgo "github.com/polytomic/polytomic-go"
)

type V3UpdateBulkSchema struct {
	Enabled      *bool                      `json:"enabled,omitempty" url:"enabled,omitempty"`
	Fields       []*polytomicgo.V3BulkField `json:"fields,omitempty" url:"fields,omitempty"`
	PartitionKey *string                    `json:"partition_key,omitempty" url:"partition_key,omitempty"`
}
