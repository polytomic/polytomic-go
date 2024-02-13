// This file was auto-generated from our API Definition.

package polytomic

import (
	fmt "fmt"
)

type V2CreateSyncRequest struct {
	Active         *bool                   `json:"active,omitempty" url:"active,omitempty"`
	Fields         []*V2SyncField          `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string                 `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*V2Filter             `json:"filters,omitempty" url:"filters,omitempty"`
	Identity       *V2Identity             `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           V2CreateSyncRequestMode `json:"mode,omitempty" url:"mode,omitempty"`
	Name           string                  `json:"name" url:"name"`
	OrganizationId *string                 `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OverrideFields []*V2SyncField          `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	Overrides      []*V2Override           `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string                `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *V2Schedule             `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool                   `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *V2Target               `json:"target,omitempty" url:"target,omitempty"`
}

type V2StartSyncRequest struct {
	Identities []string `json:"identities,omitempty" url:"identities,omitempty"`
}

type V2UpdateSyncRequest struct {
	Active         *bool                   `json:"active,omitempty" url:"active,omitempty"`
	Fields         []*V2SyncField          `json:"fields,omitempty" url:"fields,omitempty"`
	FilterLogic    *string                 `json:"filter_logic,omitempty" url:"filter_logic,omitempty"`
	Filters        []*V2Filter             `json:"filters,omitempty" url:"filters,omitempty"`
	Identity       *V2Identity             `json:"identity,omitempty" url:"identity,omitempty"`
	Mode           V2UpdateSyncRequestMode `json:"mode,omitempty" url:"mode,omitempty"`
	Name           string                  `json:"name" url:"name"`
	OrganizationId *string                 `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	OverrideFields []*V2SyncField          `json:"override_fields,omitempty" url:"override_fields,omitempty"`
	Overrides      []*V2Override           `json:"overrides,omitempty" url:"overrides,omitempty"`
	Policies       []string                `json:"policies,omitempty" url:"policies,omitempty"`
	Schedule       *V2Schedule             `json:"schedule,omitempty" url:"schedule,omitempty"`
	SyncAllRecords *bool                   `json:"sync_all_records,omitempty" url:"sync_all_records,omitempty"`
	Target         *V2Target               `json:"target,omitempty" url:"target,omitempty"`
}

type V2CreateSyncRequestMode string

const (
	V2CreateSyncRequestModeUpdate         V2CreateSyncRequestMode = "update"
	V2CreateSyncRequestModeUpdateOrCreate V2CreateSyncRequestMode = "updateOrCreate"
	V2CreateSyncRequestModeCreate         V2CreateSyncRequestMode = "create"
	V2CreateSyncRequestModeReplace        V2CreateSyncRequestMode = "replace"
	V2CreateSyncRequestModeAppend         V2CreateSyncRequestMode = "append"
	V2CreateSyncRequestModeSnapshot       V2CreateSyncRequestMode = "snapshot"
)

func NewV2CreateSyncRequestModeFromString(s string) (V2CreateSyncRequestMode, error) {
	switch s {
	case "update":
		return V2CreateSyncRequestModeUpdate, nil
	case "updateOrCreate":
		return V2CreateSyncRequestModeUpdateOrCreate, nil
	case "create":
		return V2CreateSyncRequestModeCreate, nil
	case "replace":
		return V2CreateSyncRequestModeReplace, nil
	case "append":
		return V2CreateSyncRequestModeAppend, nil
	case "snapshot":
		return V2CreateSyncRequestModeSnapshot, nil
	}
	var t V2CreateSyncRequestMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (v V2CreateSyncRequestMode) Ptr() *V2CreateSyncRequestMode {
	return &v
}

type V2UpdateSyncRequestMode string

const (
	V2UpdateSyncRequestModeUpdate         V2UpdateSyncRequestMode = "update"
	V2UpdateSyncRequestModeUpdateOrCreate V2UpdateSyncRequestMode = "updateOrCreate"
	V2UpdateSyncRequestModeCreate         V2UpdateSyncRequestMode = "create"
	V2UpdateSyncRequestModeReplace        V2UpdateSyncRequestMode = "replace"
	V2UpdateSyncRequestModeAppend         V2UpdateSyncRequestMode = "append"
	V2UpdateSyncRequestModeSnapshot       V2UpdateSyncRequestMode = "snapshot"
)

func NewV2UpdateSyncRequestModeFromString(s string) (V2UpdateSyncRequestMode, error) {
	switch s {
	case "update":
		return V2UpdateSyncRequestModeUpdate, nil
	case "updateOrCreate":
		return V2UpdateSyncRequestModeUpdateOrCreate, nil
	case "create":
		return V2UpdateSyncRequestModeCreate, nil
	case "replace":
		return V2UpdateSyncRequestModeReplace, nil
	case "append":
		return V2UpdateSyncRequestModeAppend, nil
	case "snapshot":
		return V2UpdateSyncRequestModeSnapshot, nil
	}
	var t V2UpdateSyncRequestMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (v V2UpdateSyncRequestMode) Ptr() *V2UpdateSyncRequestMode {
	return &v
}
