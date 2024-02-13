// This file was auto-generated from our API Definition.

package bulksync

type V3StartBulkSyncRequest struct {
	Resync  *bool    `json:"resync,omitempty" url:"resync,omitempty"`
	Schemas []string `json:"schemas,omitempty" url:"schemas,omitempty"`
	Test    *bool    `json:"test,omitempty" url:"test,omitempty"`
}
