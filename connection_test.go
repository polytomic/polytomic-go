package polytomic

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetConnectionReturnsServerError(t *testing.T) {
	connId := uuid.New()
	defer gock.Off()

	gock.New("https://polytomic.example.com").
		Get(fmt.Sprintf("/api/connections/%s", connId.String())).
		Reply(http.StatusForbidden).
		BodyString(`{"error":{"message":"invalid or unknown API key"}}`)

	client := NewClient("polytomic.example.com", "key")
	_, err := client.Connections().Get(context.Background(), connId)

	apiErr := ApiError{}
	assert.ErrorAs(t, err, &apiErr)
	assert.Equal(t, "invalid or unknown API key", apiErr.message)

	assert.True(t, gock.IsDone())
}
