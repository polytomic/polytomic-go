// This file was auto-generated from our API Definition.

package executions

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	polytomicgo "github.com/polytomic/polytomic-go"
	core "github.com/polytomic/polytomic-go/core"
	option "github.com/polytomic/polytomic-go/option"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
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
		header: options.ToHeader(),
	}
}

func (c *Client) List(
	ctx context.Context,
	syncId string,
	opts ...option.RequestOption,
) (*polytomicgo.ListExecutionResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions", syncId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.ListExecutionResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) Get(
	ctx context.Context,
	syncId string,
	id string,
	opts ...option.RequestOption,
) (*polytomicgo.GetExecutionResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v", syncId, id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.GetExecutionResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetErrorLogUrLs(
	ctx context.Context,
	syncId string,
	id string,
	opts ...option.RequestOption,
) (*polytomicgo.ExecutionLogsResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/errors", syncId, id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.ExecutionLogsResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetErrorLogs(
	ctx context.Context,
	syncId string,
	id string,
	filename string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/errors/%v", syncId, id, filename)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetInsertLogUrLs(
	ctx context.Context,
	syncId string,
	id string,
	opts ...option.RequestOption,
) (*polytomicgo.ExecutionLogsResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/inserts", syncId, id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.ExecutionLogsResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetInsertLogs(
	ctx context.Context,
	syncId string,
	id string,
	filename string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/inserts/%v", syncId, id, filename)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetRecordLogUrLs(
	ctx context.Context,
	syncId string,
	id string,
	opts ...option.RequestOption,
) (*polytomicgo.ExecutionLogsResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/records", syncId, id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.ExecutionLogsResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetRecordLogs(
	ctx context.Context,
	syncId string,
	id string,
	filename string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/records/%v", syncId, id, filename)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetUpdateLogUrLs(
	ctx context.Context,
	syncId string,
	id string,
	opts ...option.RequestOption,
) (*polytomicgo.ExecutionLogsResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/updates", syncId, id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.ExecutionLogsResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetUpdateLogs(
	ctx context.Context,
	syncId string,
	id string,
	filename string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/updates/%v", syncId, id, filename)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

func (c *Client) GetWarningLogUrLs(
	ctx context.Context,
	syncId string,
	id string,
	opts ...option.RequestOption,
) (*polytomicgo.ExecutionLogsResponseEnvelope, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/warnings", syncId, id)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *polytomicgo.ExecutionLogsResponseEnvelope
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetWarningLogs(
	ctx context.Context,
	syncId string,
	id string,
	filename string,
	opts ...option.RequestOption,
) error {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://app.polytomic.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/syncs/%v/executions/%v/warnings/%v", syncId, id, filename)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 401:
			value := new(polytomicgo.UnauthorizedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			MaxAttempts:  options.MaxAttempts,
			Headers:      headers,
			Client:       options.HTTPClient,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}
