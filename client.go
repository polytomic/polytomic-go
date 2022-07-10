package polytomic

import "github.com/carlmjohnson/requests"

type Client struct {
	host          string
	deploymentKey string
}

// NewClient returns a Polytomic API client which will make requests to the
// specified host, using the provided deployment API key.
func NewClient(host, key string) *Client {
	return &Client{
		host:          host,
		deploymentKey: key,
	}
}

// Workspaces returns an API client for modifying Polytomic workspaces.
func (c *Client) Workspaces() *WorkspaceApi {
	return &WorkspaceApi{client: c}
}

// Users returns an API client for modifying Polytomic users.
func (c *Client) Users() *UserApi {
	return &UserApi{client: c}
}

// Connections returns an API client for modifying Polytomic connections.
func (c *Client) Connections() *ConnectionApi {
	return &ConnectionApi{client: c}
}

// newRequest returns a configured request builder...
func (c *Client) newRequest(url string) *requests.Builder {
	return requests.
		URL(url).
		Host(c.host).
		BasicAuth(c.deploymentKey, "").
		UserAgent("polytomic-go")
}

type topLevelResult struct {
	Result interface{} `json:"result"`
}
