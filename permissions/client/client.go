// This file was auto-generated from our API Definition.

package client

import (
	core "github.com/polytomic/polytomic-go/core"
	option "github.com/polytomic/polytomic-go/option"
	policies "github.com/polytomic/polytomic-go/permissions/policies"
	roles "github.com/polytomic/polytomic-go/permissions/roles"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Policies *policies.Client
	Roles    *roles.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: core.NewCaller(
			&core.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header:   options.ToHeader(),
		Policies: policies.NewClient(opts...),
		Roles:    roles.NewClient(opts...),
	}
}
