// This file was auto-generated from our API Definition.

package polytomic

import (
	json "encoding/json"
	core "github.com/polytomic/polytomic-go/core"
)

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
