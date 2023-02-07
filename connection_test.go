package polytomic

import (
	"context"
	"encoding/base64"
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

	t.Run("test deployment key", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/connections/%s", connId.String())).
			Reply(http.StatusForbidden).
			BodyString(`{"error":{"message":"invalid or unknown API key"}}`)
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		_, err := client.Connections().Get(context.Background(), connId)

		apiErr := ApiError{}
		assert.ErrorAs(t, err, &apiErr)
		assert.Equal(t, "invalid or unknown API key", apiErr.Message)
		assert.True(t, gock.IsDone())
	})

	t.Run("test api key", func(t *testing.T) {
		header := "Bearer " + base64.StdEncoding.EncodeToString([]byte("key"))
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/connections/%s", connId.String())).
			MatchHeader("Authorization", header).
			Reply(http.StatusOK)

		client := NewClient("polytomic.example.com", APIKey("key"))
		client.Connections().Get(context.Background(), connId)

		assert.True(t, gock.IsDone())
	})

	t.Run("Get connection by ID", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/connections/%s", connId.String())).
			Reply(http.StatusOK).
			File("fixtures/connection.json")

		client := NewClient("polytomic.example.com", APIKey("key"))
		conn, err := client.Connections().Get(context.Background(), connId)
		assert.NoError(t, err)
		assert.Equal(t, "example", conn.Name)

		assert.True(t, gock.IsDone())
	})

	t.Run("List connections", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get("/api/connections").
			Reply(http.StatusOK).
			File("fixtures/list_connections.json")

		client := NewClient("polytomic.example.com", APIKey("key"))
		conns, err := client.Connections().List(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 2, len(conns))
	})
}
