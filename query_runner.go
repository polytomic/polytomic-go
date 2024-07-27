// This file was auto-generated from our API Definition.

package polytomic

type QueryRunnerGetQueryRequest struct {
	Page *string `json:"-" url:"page,omitempty"`
}

type V4RunQueryRequest struct {
	// The query to execute against the connection.
	Query *string `json:"-" url:"query,omitempty"`
}
