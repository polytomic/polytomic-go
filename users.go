// This file was auto-generated from our API Definition.

package polytomic

type V2CreateUserRequestSchema struct {
	Email string  `json:"email" url:"email"`
	Role  *string `json:"role,omitempty" url:"role,omitempty"`
}

type UsersCreateApiKeyRequest struct {
	Force *bool `json:"-" url:"force,omitempty"`
}

type V2UpdateUserRequestSchema struct {
	Email string  `json:"email" url:"email"`
	Role  *string `json:"role,omitempty" url:"role,omitempty"`
}
