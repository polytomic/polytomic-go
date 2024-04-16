// This file was auto-generated from our API Definition.

package polytomic

type ConnectCardRequest struct {
	Connection     *string  `json:"connection,omitempty" url:"connection,omitempty"`
	Name           string   `json:"name" url:"name"`
	OrganizationId *string  `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	RedirectUrl    string   `json:"redirect_url" url:"redirect_url"`
	Type           *string  `json:"type,omitempty" url:"type,omitempty"`
	Whitelist      []string `json:"whitelist,omitempty" url:"whitelist,omitempty"`
}

type CreateConnectionRequestSchema struct {
	Configuration  map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Name           string                 `json:"name" url:"name"`
	OrganizationId *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string               `json:"policies,omitempty" url:"policies,omitempty"`
	// URL to redirect to after completing OAuth flow.
	RedirectUrl *string `json:"redirect_url,omitempty" url:"redirect_url,omitempty"`
	Type        string  `json:"type" url:"type"`
	// Validate connection configuration.
	Validate *bool `json:"validate,omitempty" url:"validate,omitempty"`
}

type ConnectionsRemoveRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`
}

type UpdateConnectionRequestSchema struct {
	Configuration  map[string]interface{} `json:"configuration,omitempty" url:"configuration,omitempty"`
	Name           string                 `json:"name" url:"name"`
	OrganizationId *string                `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Policies       []string               `json:"policies,omitempty" url:"policies,omitempty"`
	Reconnect      *bool                  `json:"reconnect,omitempty" url:"reconnect,omitempty"`
	Type           *string                `json:"type,omitempty" url:"type,omitempty"`
	// Validate connection configuration.
	Validate *bool `json:"validate,omitempty" url:"validate,omitempty"`
}
