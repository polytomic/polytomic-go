// This file was auto-generated from our API Definition.

package polytomic

type V2CreateWebhooksSchema struct {
	Endpoint       string  `json:"endpoint" url:"endpoint"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         string  `json:"secret" url:"secret"`
}

type V2UpdateWebhooksSchema struct {
	Endpoint       string  `json:"endpoint" url:"endpoint"`
	OrganizationId *string `json:"organization_id,omitempty" url:"organization_id,omitempty"`
	Secret         string  `json:"secret" url:"secret"`
}
