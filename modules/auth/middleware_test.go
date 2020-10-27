package auth

import (
	"net/url"
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/stretchr/testify/assert"
)

func TestCheckPermissions(t *testing.T) {

	config.Initialize(&config.Config{
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
			}, {
				Name:       "/info/user_list?user_type=10",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/info/user_list?user_type=10"},
			}, {
				Name:       "/info/user_list?user_type=20",
				Slug:       "/",
				HttpMethod: []string{"GET"},
				HttpPath:   []string{"/info/user_list?user_type=20"},
			}, {
				Name:       "/delete/user",
				Slug:       "/",
				HttpMethod: []string{"POST"},
				HttpPath:   []string{"/delete/user"},
			},
		},
	}

	param := make(url.Values)

	assert.Equal(t, CheckPermissions(user, "/admin/", "GET", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin", "GET", param), true)
	assert.Equal(t, CheckPermissions(user, "/", "GET", param), false)
	assert.Equal(t, CheckPermissions(user, "/admin", "POST", param), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/users", "GET", param), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "GET", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "get", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?__goadmin_edit_pk=2&__columns=id,roles,created_at,updated_at", "get", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?__goadmin_edit_pk=2", "get", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?__goadmin_edit_pk=3&__columns=id,roles,created_at,updated_at", "get", param), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/normal_manager/edit?__columns=id,roles,created_at,updated_at&id=3", "get", param), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user", "post", param), false)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user/edit?id=3", "get", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/logout?j=asdf", "post", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user_list?user_type=20", "get", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/info/user_list?__goadmin_edit_pk=3&user_type=20", "get", param), true)
	assert.Equal(t, CheckPermissions(user, "/admin/delete/user", "post", param), true)
}
