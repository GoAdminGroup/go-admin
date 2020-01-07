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
			}, {
				Name:       "/info/user/edit",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/info/user/edit"},
			}, {
				Name:       "/info/normal_manager?id=2",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/info/normal_manager?id=2"},
			}, {
				Name:       "/info/normal_manager/edit?id=2",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/info/normal_manager/edit?id=2"},
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
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?id=2&__columns=id,roles,created_at,updated_at", "get"), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?id=2", "get"), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?id=3&__columns=id,roles,created_at,updated_at", "get"), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?__columns=id,roles,created_at,updated_at&id=3", "get"), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "post"), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user/edit?id=3", "get"), true)
	assert.Equal(t, CheckPermissions(user, "/admin/logout?j=asdf", "post"), true)
}

func TestInMethodArr(t *testing.T) {
	methods := []string{"get", "post"}
	assert.Equal(t, inMethodArr(methods, "get"), true)
}
