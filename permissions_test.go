package polytomic

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestRoles(t *testing.T) {
	t.Run("Get role", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get("/api/permissions/roles/00000000-0000-0000-0000-000000000001").
			Reply(200).
			File("fixtures/role.json")
		defer gock.Off()

		// client := NewClient("app.polytomic-local.com:8443", DeploymentKey("secret-key"))
		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		role, err := client.Permissions().GetRole(context.Background(), "00000000-0000-0000-0000-000000000001")
		assert.NoError(t, err)
		assert.Equal(t, "Admin", role.Name)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", role.ID)

		assert.True(t, gock.IsDone())
	})

	t.Run("List roles", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get("/api/permissions/roles").
			Reply(200).
			File("fixtures/list-roles.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		roles, err := client.Permissions().ListRoles(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 2, len(roles))
		assert.Equal(t, "Admin", roles[0].Name)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", roles[0].ID)
	})

	t.Run("Create role", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Post("/api/permissions/roles").
			Reply(200).
			File("fixtures/create-role.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))
		role, err := client.Permissions().CreateRole(context.Background(), RoleRequest{
			Name:           "test role",
			OrganizationID: "00000000-0000-0000-0000-000000000000",
		})
		assert.NoError(t, err)
		assert.Equal(t, "test role", role.Name)
	})
}

func TestPolicies(t *testing.T) {
	t.Run("Get policy", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get("/api/permissions/policies/00000000-0000-0000-0001-000000000001").
			Reply(200).
			File("fixtures/policy.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))

		policy, err := client.Permissions().GetPolicy(context.Background(), "00000000-0000-0000-0001-000000000001")
		assert.NoError(t, err)
		assert.Equal(t, "00000000-0000-0000-0001-000000000001", policy.ID)
		assert.Equal(t, "Model syncs", policy.Name)

	})

	t.Run("List policies", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Get("/api/permissions/policies").
			Reply(200).
			File("fixtures/list-policies.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))

		policies, err := client.Permissions().ListPolicies(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 3, len(policies))
		assert.Equal(t, "00000000-0000-0000-0001-000000000001", policies[0].ID)
		assert.Equal(t, "Model syncs", policies[0].Name)
	})

	t.Run("Create policy", func(t *testing.T) {
		gock.New("https://polytomic.example.com").
			Post("/api/permissions/policies").
			Reply(200).
			File("fixtures/create-policy.json")
		defer gock.Off()

		client := NewClient("polytomic.example.com", DeploymentKey("key"))

		policy, err := client.Permissions().CreatePolicy(context.Background(), PolicyRequest{
			Name:           "test policy",
			OrganizationID: "00000000-0000-0000-0000-000000000000",
			PolicyActions: []PolicyAction{
				{
					Action: "read",
					RoleIDs: []string{
						"00000000-0000-0000-0000-000000000001",
					},
				},
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, "test policy", policy.Name)
		assert.Equal(t, "10000000-0000-0000-0000-000000000000", policy.ID)
		assert.Equal(t, "read", policy.PolicyActions[0].Action)
		assert.Equal(t, "00000000-0000-0000-0000-000000000001", policy.PolicyActions[0].RoleIDs[0])
	})

}
