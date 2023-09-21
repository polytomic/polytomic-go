package polytomic

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/carlmjohnson/requests"
)

const (
	DefaultHost = "app.polytomic.com"

	DefaultVersion       = "2023-04-25"
	VersionHeader        = "X-Polytomic-Version"
	IdempotencyKeyHeader = "Idempotency-Key"
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
	auth    requests.Config
	host    string
	ua      string
	version string
}

type clientOpts func(*clientOptions)

type clientOptions struct {
	userAgent string
	version   string
}

// WithUserAgent sets the User-Agent header on all requests.
func WithUserAgent(ua string) clientOpts {
	return func(o *clientOptions) {
		o.userAgent = ua
	}
}

func WithAPIVersion(version string) clientOpts {
	return func(o *clientOptions) {
		o.version = version
	}
}

// NewClient returns a Polytomic API client which will make requests to the
// specified host, using the provided Authenticator.
func NewClient(host string, auth Authenticator, opts ...clientOpts) *Client {
	if strings.HasPrefix(host, "https://") || strings.HasPrefix(host, "http://") {
		if parsed, err := url.Parse(host); err == nil {
			host = parsed.Host
		}
	}
	if host == "" {
		host = DefaultHost
	}

	o := &clientOptions{}
	for _, opt := range opts {
		opt(o)
	}

	if o.version == "" {
		o.version = DefaultVersion
	}

	if o.userAgent == "" {
		o.userAgent = "polytomic-go/" + o.version
	}

	return &Client{
		auth:    auth,
		host:    host,
		ua:      o.userAgent,
		version: o.version,
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

// Permissions returns an API client for modifying Polytomic permissions.
func (c *Client) Permissions() *PermissionsApi {
	return &PermissionsApi{client: c}
}

type requestOpts func(*requestOptions)

type requestOptions struct {
	IdempotencyKey       string
	ForceDelete          bool
	SkipConfigValidation bool
	Client               *http.Client
}

// WithIdempotencyKey sets the Idempotency-Key header on the request.
func WithIdempotencyKey(key string) requestOpts {
	return func(o *requestOptions) {
		o.IdempotencyKey = key
	}
}

func WithForceDelete() requestOpts {
	return func(o *requestOptions) {
		o.ForceDelete = true
	}
}

func SkipConfigValidation() requestOpts {
	return func(o *requestOptions) {
		o.SkipConfigValidation = true
	}
}

func WithClient(client *http.Client) requestOpts {
	return func(o *requestOptions) {
		o.Client = client
	}
}

// newRequest returns a configured request builder...
func (c *Client) newRequest(url string, opts ...requestOpts) *requests.Builder {
	r := requests.
		URL(url).
		Host(c.host).
		Config(c.auth).
		Header(VersionHeader, c.version).
		UserAgent(c.ua).
		AddValidator(checkApiResponse)

	options := requestOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.IdempotencyKey != "" {
		r.Header(IdempotencyKeyHeader, options.IdempotencyKey)
	}

	if options.ForceDelete {
		r.Param("force", "true")
	}

	if options.Client != nil {
		r.Client(options.Client)
	}

	return r
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
	if err := json.Unmarshal(body, &apiErr); err != nil || apiErr.Message == "" {
		// Try legacy format
		var legacyErr LegacyError
		if err := json.Unmarshal(body, &legacyErr); err != nil {
			return fmt.Errorf("unexpected response: %s", string(body))
		}
		apiErr = Error{
			Message: legacyErr.Error.Message,
		}
	}

	return ApiError{
		StatusCode: res.StatusCode,
		Message:    apiErr.Message,
		Metadata:   apiErr.Metadata,
	}

}

type LegacyError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

type Error struct {
	Message    string        `json:"message"`
	StatusCode int           `json:"status_code"`
	Metadata   []interface{} `json:"metadata"`
}

type ApiError struct {
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message"`
	Metadata   []interface{} `json:"metadata"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%s (%d): %v", e.Message, e.StatusCode, e.Metadata)
}

type Response struct {
	Data interface{} `json:"data"`
}
