// This file was auto-generated from our API Definition.

package polytomic

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
