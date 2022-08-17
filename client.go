package polytomic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/carlmjohnson/requests"
)

const (
	DefaultHost = "app.polytomic.com"
)

// Authenticator defines the function signature used to set authentication
// information on client requests.
type Authenticator requests.Config

// DeploymentKey returns an Authenticator which will use the deployment-level
// key to authenticate requests.
func DeploymentKey(deploymentKey string) Authenticator {
	return func(rb *requests.Builder) {
		rb.
			BasicAuth(deploymentKey, "")
	}
}

// APIKey returns an Authenticator which will use the provided API key to
// authenticate requests.
func APIKey(apiKey string) Authenticator {
	return func(rb *requests.Builder) {
		rb.Bearer(apiKey)
	}
}

type Client struct {
	auth requests.Config
	host string
}

// NewClient returns a Polytomic API client which will make requests to the
// specified host, using the provided Authenticator.
func NewClient(host string, auth Authenticator) *Client {
	return &Client{
		auth: auth,
		host: host,
	}
}

// Organizations returns an API client for modifying Polytomic organizations.
func (c *Client) Organizations() *OrganizationApi {
	return &OrganizationApi{client: c}
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
		Config(c.auth).
		UserAgent("polytomic-go").
		AddValidator(checkApiResponse)
}

type topLevelResult struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

// checkApiResponse provides a request.ResponseHandler which will attempt to
// parse non-2xx responses as topLevelResults.
func checkApiResponse(res *http.Response) error {
	if res.StatusCode < 400 {
		return nil
	}

	result := topLevelResult{}

	buf := &bytes.Buffer{}
	defer func() {
		res.Body = io.NopCloser(buf)
	}()
	body, err := io.ReadAll(io.TeeReader(res.Body, buf))
	defer res.Body.Close()
	if err != nil {
		return err
	}

	if json.Unmarshal(body, &result) != nil || result.Error == nil {
		// this wasn't a top level result
		return errors.New(fmt.Sprintf("unexpected error (%d): %s", res.StatusCode, string(body)))
	}
	if errmap, ok := result.Error.(map[string]interface{}); ok {
		if msg, ok := errmap["message"]; ok {
			if msgStr, ok := msg.(string); ok {
				return ApiError{message: msgStr, statusCode: res.StatusCode}
			}
		}
	}
	return nil
}

type ApiError struct {
	statusCode int
	message    string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%s (%d)", e.message, e.statusCode)
}
