// This file was auto-generated from our API Definition.

package polytomic

type CreateWebhooksSchema struct {
	Endpoint       string  `json:"endpoint" url:"endpoint"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         string  `json:"secret" url:"secret"`
}

type UpdateWebhooksSchema struct {
	Endpoint       string  `json:"endpoint" url:"endpoint"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         string  `json:"secret" url:"secret"`
}
