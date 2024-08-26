// This file was auto-generated from our API Definition.

package bulksync

type ExecutionsExportLogsRequest struct {
	// Send a notification to the user when the logs are ready for download.
	Notify *bool `json:"-" url:"notify,omitempty"`
}

type ExecutionsListStatusRequest struct {
	// Return the execution status of all syncs in the organization
	All *bool `json:"-" url:"all,omitempty"`
	// Return the execution status of all active syncs in the organization
	Active *bool `json:"-" url:"active,omitempty"`
	// Return the execution status of the specified sync; this may be supplied multiple times.
	SyncId []*string `json:"-" url:"sync_id,omitempty"`
}
