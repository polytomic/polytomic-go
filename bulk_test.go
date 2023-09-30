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

func TestGetSource(t *testing.T) {
	connId := uuid.New()
	defer gock.Off()

	t.Run("Get bulk source schemas", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/bulk/source/%s", connId.String())).
			Reply(http.StatusOK).
			File("fixtures/bulk_source.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		source, err := client.Bulk().GetSource(context.Background(), connId.String())
		assert.NoError(t, err)
		assert.Equal(t, 3, len(source.Schemas))
		assert.Equal(t, Schema{ID: "test.test", Name: "test.test"}, source.Schemas[0])

		assert.True(t, gock.IsDone())
	})
}

func TestGetSourceSchema(t *testing.T) {
	connId := uuid.New()
	schemaId := "test.test"
	defer gock.Off()

	t.Run("Get bulk source schema", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/bulk/source/%s/schema/%s", connId.String(), schemaId)).
			Reply(http.StatusOK).
			File("fixtures/bulk_source_schema.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		schema, err := client.Bulk().GetSourceSchema(context.Background(), connId.String(), schemaId)
		assert.NoError(t, err)
		assert.Equal(t, "test.test", schema.ID)
		assert.Equal(t, "test.test", schema.Name)
		assert.Equal(t, "string", schema.Fields[0].Type)

		assert.True(t, gock.IsDone())
	})
}

func TestGetDestination(t *testing.T) {
	connId := uuid.New()
	defer gock.Off()

	t.Run("Get bulk destination modes", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/bulk/dest/%s", connId.String())).
			Reply(http.StatusOK).
			File("fixtures/bulk_destination.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		dest, err := client.Bulk().GetDestination(context.Background(), connId.String())
		assert.NoError(t, err)
		assert.Equal(t, 1, len(dest.Modes))
		assert.Equal(t, "replicate", dest.Modes[0].ID)
		assert.Equal(t, "Replicate", dest.Modes[0].Label)

		assert.True(t, gock.IsDone())
	})

}

func TestCreateBulkSync(t *testing.T) {
	frequency := "manual"
	sync := BulkSyncRequest{
		SourceConnectionID: "source-connection-id",
		DestConnectionID:   "dest-connection-id",
		Mode:               "replicate",
		Active:             true,
		Schedule: BulkSchedule{
			Frequency: &frequency,
		},
		Name: "test-sync",
		Schemas: []string{
			"issue",
		},
	}

	defer gock.Off()

	t.Run("Create bulk sync", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Post("/api/bulk").
			Reply(http.StatusCreated).
			File("fixtures/bulk_sync.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		created, err := client.Bulk().CreateBulkSync(context.Background(), sync)
		assert.NoError(t, err)
		assert.Equal(t, "3b3561f3-3247-4c38-96b5-a36a86d99dee", created.ID)
	})
}

func TestGetBulkSync(t *testing.T) {
	syncId := uuid.New()
	defer gock.Off()

	t.Run("Get bulk sync", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get(fmt.Sprintf("/api/bulk/syncs/%s", syncId.String())).
			Reply(http.StatusOK).
			File("fixtures/bulk_sync.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		sync, err := client.Bulk().GetBulkSync(context.Background(), syncId.String())
		assert.NoError(t, err)
		assert.Equal(t, "3b3561f3-3247-4c38-96b5-a36a86d99dee", sync.ID)
		assert.Equal(t, "Jira2", sync.Name)
	})

}
