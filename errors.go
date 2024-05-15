// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
	core "github.com/polytomic/polytomic-go/core"
)

// Bad Request
type BadRequestError struct {
	*core.APIError
	Body *ApiError
}

func (b *BadRequestError) UnmarshalJSON(data []byte) error {
	var body *ApiError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	b.StatusCode = 400
	b.Body = body
	return nil
}

func (b *BadRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Body)
}

func (b *BadRequestError) Unwrap() error {
	return b.APIError
}

// Forbidden
type ForbiddenError struct {
	*core.APIError
	Body *ApiError
}

func (f *ForbiddenError) UnmarshalJSON(data []byte) error {
	var body *ApiError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	f.StatusCode = 403
	f.Body = body
	return nil
}

func (f *ForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(f.Body)
}

func (f *ForbiddenError) Unwrap() error {
	return f.APIError
}

// Internal Server Error
type InternalServerError struct {
	*core.APIError
	Body *ApiError
}

func (i *InternalServerError) UnmarshalJSON(data []byte) error {
	var body *ApiError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	i.StatusCode = 500
	i.Body = body
	return nil
}

func (i *InternalServerError) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Body)
}

func (i *InternalServerError) Unwrap() error {
	return i.APIError
}

// Not Found
type NotFoundError struct {
	*core.APIError
	Body *ApiError
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body *ApiError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}

func (n *NotFoundError) Unwrap() error {
	return n.APIError
}

// Unauthorized
type UnauthorizedError struct {
	*core.APIError
	Body *RestErrResponse
}

func (u *UnauthorizedError) UnmarshalJSON(data []byte) error {
	var body *RestErrResponse
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 401
	u.Body = body
	return nil
}

func (u *UnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnauthorizedError) Unwrap() error {
	return u.APIError
}

// Unprocessable Entity
type UnprocessableEntityError struct {
	*core.APIError
	Body *ApiError
}

func (u *UnprocessableEntityError) UnmarshalJSON(data []byte) error {
	var body *ApiError
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	u.StatusCode = 422
	u.Body = body
	return nil
}

func (u *UnprocessableEntityError) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Body)
}

func (u *UnprocessableEntityError) Unwrap() error {
	return u.APIError
}
