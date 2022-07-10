package polytomic

import "github.com/carlmjohnson/requests"

type Client struct {
	host          string
	deploymentKey string
}

func NewClient(host, key string) *Client {
	return &Client{
		host:          host,
		deploymentKey: key,
	}
}

func (c *Client) Workspaces() *WorkspaceApi {
	return &WorkspaceApi{client: c}
}

func (c *Client) Users() *UserApi {
	return &UserApi{client: c}
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
