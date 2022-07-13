package polytomic

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckApiResponse(t *testing.T) {
	tests := map[string]struct {
		status   int
		body     string
		expected string
	}{
		"non-error": {http.StatusOK, "", ""},
		"unauthorized": {
			http.StatusUnauthorized,
			`{"error":{"message":"invalid or unknown API key"}}`,
			"invalid or unknown API key (401)",
		},
		"not top-level": {
			http.StatusInternalServerError,
			"Oh hello, error",
			"unexpected error (500): Oh hello, error",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resp := &http.Response{
				Body:       io.NopCloser(bytes.NewBufferString(test.body)),
				StatusCode: test.status,
			}

			result := checkApiResponse(resp)
			if test.expected == "" {
				assert.NoError(t, result)
			} else {
				assert.EqualError(t, result, test.expected)
			}
		})
	}

}
