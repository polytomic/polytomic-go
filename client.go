package polytomic

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/carlmjohnson/requests"
)

const (
	DefaultHost = "app.polytomic.com"

	Version       = "2022-12-12"
	VersionHeader = "X-Polytomic-Version"
)

// Authenticator defines the function signature used to set authentication
// information on client requests.
type Authenticator requests.Config

// DeploymentKey returns an Authenticator which will use the deployment-level
// key to authenticate requests.
func DeploymentKey(deploymentKey string) Authenticator {
	return func(rb *requests.Builder) {
		rb.BasicAuth(deploymentKey, "")
	}
}

// APIKey returns an Authenticator which will use the provided API key to
// authenticate requests.
func APIKey(apiKey string) Authenticator {
	return func(rb *requests.Builder) {
		value := base64.StdEncoding.EncodeToString([]byte(apiKey))
		rb.Bearer(value)
	}
}

type Client struct {
	auth requests.Config
	host string
}

// NewClient returns a Polytomic API client which will make requests to the
// specified host, using the provided Authenticator.
func NewClient(host string, auth Authenticator) *Client {
	if parsed, err := url.Parse(host); err == nil {
		host = parsed.Host
	}
	if host == "" {
		host = DefaultHost
	}
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

// Bulk returns an API client for modifying Polytomic bulk syncs.
func (c *Client) Bulk() *BulkApi {
	return &BulkApi{client: c}
}

// Models returns an API client for modifying Polytomic models.
func (c *Client) Models() *ModelApi {
	return &ModelApi{client: c}
}

// Syncs returns an API client for modifying Polytomic syncs.
func (c *Client) Syncs() *SyncApi {
	return &SyncApi{client: c}
}

// newRequest returns a configured request builder...
func (c *Client) newRequest(url string) *requests.Builder {
	return requests.
		URL(url).
		Host(c.host).
		Config(c.auth).
		Header(VersionHeader, Version).
		UserAgent("polytomic-go").
		AddValidator(checkApiResponse)
}

// checkApiResponse provides a request.ResponseHandler which will attempt to
// parse non-2xx responses as topLevelResults.
func checkApiResponse(res *http.Response) error {
	if res.StatusCode < 400 {
		return nil
	}

	buf := &bytes.Buffer{}
	defer func() {
		res.Body = io.NopCloser(buf)
	}()
	body, err := io.ReadAll(io.TeeReader(res.Body, buf))
	defer res.Body.Close()
	if err != nil {
		return err
	}

	var apiErr Error
	if err := json.Unmarshal(body, &apiErr); err != nil {
		// Try legacy format
		var legacyErr LegacyError
		if err := json.Unmarshal(body, &legacyErr); err != nil {
			return fmt.Errorf("unexpected response: %w", err)
		}
		apiErr = Error{
			Error: legacyErr.Error.Message,
		}

	}

	return ApiError{
		message:    apiErr.Error,
		statusCode: res.StatusCode,
	}

}

type LegacyError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type Error struct {
	Error      string `json:"error"`
	StatusCode string `json:"status"`
}

type ApiError struct {
	statusCode int
	message    string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%s (%d)", e.message, e.statusCode)
}

type Response struct {
	Data interface{} `json:"data"`
}
