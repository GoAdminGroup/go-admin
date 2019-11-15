package auth

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPermissions(t *testing.T) {

	config.Set(config.Config{
		UrlPrefix: "admin",
	})

	user := models.UserModel{
		Permissions: []models.PermissionModel{
			{
				Name:       "/",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/"},
			}, {
				Name:       "/info/user",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/info/user"},
			},
		},
	}

	assert.Equal(t, CheckPermissions(user, "/admin/", "GET"), true)
	assert.Equal(t, CheckPermissions(user, "/admin", "GET"), true)
	assert.Equal(t, CheckPermissions(user, "/", "GET"), false)
	assert.Equal(t, CheckPermissions(user, "/admin", "POST"), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/users", "GET"), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "GET"), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "get"), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "post"), false)
}

func TestInMethodArr(t *testing.T) {
	methods := []string{"get", "post"}
	assert.Equal(t, inMethodArr(methods, "get"), true)
}
