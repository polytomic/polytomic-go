// This file was auto-generated from our API Definition.

package client

import (
	bulksyncclient "github.com/polytomic/polytomic-go/bulksync/client"
	connections "github.com/polytomic/polytomic-go/connections"
	core "github.com/polytomic/polytomic-go/core"
	events "github.com/polytomic/polytomic-go/events"
	identity "github.com/polytomic/polytomic-go/identity"
	jobs "github.com/polytomic/polytomic-go/jobs"
	models "github.com/polytomic/polytomic-go/models"
	modelsyncclient "github.com/polytomic/polytomic-go/modelsync/client"
	option "github.com/polytomic/polytomic-go/option"
	organization "github.com/polytomic/polytomic-go/organization"
	permissionsclient "github.com/polytomic/polytomic-go/permissions/client"
	queryrunner "github.com/polytomic/polytomic-go/queryrunner"
	schemas "github.com/polytomic/polytomic-go/schemas"
	users "github.com/polytomic/polytomic-go/users"
	webhooks "github.com/polytomic/polytomic-go/webhooks"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	BulkSync     *bulksyncclient.Client
	Connections  *connections.Client
	QueryRunner  *queryrunner.Client
	ModelSync    *modelsyncclient.Client
	Schemas      *schemas.Client
	Models       *models.Client
	Events       *events.Client
	Jobs         *jobs.Client
	Identity     *identity.Client
	Organization *organization.Client
	Users        *users.Client
	Permissions  *permissionsclient.Client
	Webhooks     *webhooks.Client
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
		header:       options.ToHeader(),
		BulkSync:     bulksyncclient.NewClient(opts...),
		Connections:  connections.NewClient(opts...),
		QueryRunner:  queryrunner.NewClient(opts...),
		ModelSync:    modelsyncclient.NewClient(opts...),
		Schemas:      schemas.NewClient(opts...),
		Models:       models.NewClient(opts...),
		Events:       events.NewClient(opts...),
		Jobs:         jobs.NewClient(opts...),
		Identity:     identity.NewClient(opts...),
		Organization: organization.NewClient(opts...),
		Users:        users.NewClient(opts...),
		Permissions:  permissionsclient.NewClient(opts...),
		Webhooks:     webhooks.NewClient(opts...),
	}
}
