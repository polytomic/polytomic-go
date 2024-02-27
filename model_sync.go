// This file was auto-generated from our API Definition.

package polytomic

type V2CreateSyncRequest struct {
	Active         *bool          `json:"active,omitempty" url:"active,omitempty"`
	Fields         []*V2SyncField `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string        `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*V2Filter    `json:"filters,omitempty" url:"filters,omitempty"`
	Identity       *V2Identity    `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           string         `json:"mode" url:"mode"`
	Name           string         `json:"name" url:"name"`
	OrganizationId *string        `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OverrideFields []*V2SyncField `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	Overrides      []*V2Override  `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string       `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *V2Schedule    `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool          `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *V2Target      `json:"target,omitempty" url:"target,omitempty"`
}

type V2StartSyncRequest struct {
	Identities []string `json:"identities,omitempty" url:"identities,omitempty"`
}

type V2UpdateSyncRequest struct {
	Active         *bool          `json:"active,omitempty" url:"active,omitempty"`
	Fields         []*V2SyncField `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string        `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*V2Filter    `json:"filters,omitempty" url:"filters,omitempty"`
	Identity       *V2Identity    `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           string         `json:"mode" url:"mode"`
	Name           string         `json:"name" url:"name"`
	OrganizationId *string        `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OverrideFields []*V2SyncField `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	Overrides      []*V2Override  `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string       `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *V2Schedule    `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool          `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *V2Target      `json:"target,omitempty" url:"target,omitempty"`
}
