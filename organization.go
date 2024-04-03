// This file was auto-generated from our API Definition.

package polytomic

type CreateOrganizationRequestSchema struct {
	ClientId     *string `json:"client_id,omitempty" url:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty" url:"client_secret,omitempty"`
	Issuer       *string `json:"issuer,omitempty" url:"issuer,omitempty"`
	Name         string  `json:"name" url:"name"`
	SsoDomain    *string `json:"sso_domain,omitempty" url:"sso_domain,omitempty"`
	SsoOrgId     *string `json:"sso_org_id,omitempty" url:"sso_org_id,omitempty"`
}

type UpdateOrganizationRequestSchema struct {
	ClientId     *string `json:"client_id,omitempty" url:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty" url:"client_secret,omitempty"`
	Issuer       *string `json:"issuer,omitempty" url:"issuer,omitempty"`
	Name         string  `json:"name" url:"name"`
	SsoDomain    *string `json:"sso_domain,omitempty" url:"sso_domain,omitempty"`
	SsoOrgId     *string `json:"sso_org_id,omitempty" url:"sso_org_id,omitempty"`
}
