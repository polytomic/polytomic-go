// This file was auto-generated from our API Definition.

package permissions

type CreateRoleRequest struct {
	Name           string  `json:"name" url:"name"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
}

type UpdateRoleRequest struct {
	Name           string  `json:"name" url:"name"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
}
