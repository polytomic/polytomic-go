package polytomic

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGetCallerIdentity(t *testing.T) {
	defer gock.Off()

	gock.New("https://polytomic.example.com").
		Get("/api/me").
		Reply(http.StatusOK).
		File("fixtures/identity.json")
	defer gock.Off()

	client := NewClient("polytomic.example.com", DeploymentKey("key"))
	ident, err := client.Identity().Get(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", ident.Email)
	assert.Equal(t, "admin", ident.Role)
	assert.Equal(t, "Example Org", ident.Organization)

	assert.True(t, gock.IsDone())

}
