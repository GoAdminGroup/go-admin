package controller

import (
	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
)

func newManager(dataList form.Values) {

	if dataList.IsEmpty("name", "username", "password") {
		panic("username and password can not be empty")
	}

	user := models.User().New(dataList.Get("username"),
		auth.EncodePassword([]byte(dataList.Get("password"))),
		dataList.Get("name"),
		dataList.Get("avatar"))

	for i := 0; i < len(dataList["role_id[]"]); i++ {
		user.AddRole(dataList["role_id[]"][i])
	}

	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		user.AddPermission(dataList["permission_id[]"][i])
	}

}

func editManager(dataList form.Values) {

	if dataList.IsEmpty("name", "username") {
		panic("username and password can not be empty")
	}

	user := models.UserWithId(dataList.Get("id"))

	password := dataList.Get("password")

	if password != "" {

		if password != dataList.Get("password_again") {
			panic("password does not match")
		}

		password = auth.EncodePassword([]byte(dataList.Get("password")))
	}

	user.Update(dataList.Get("username"), password, dataList.Get("name"), dataList.Get("avatar"))

	user.DeleteRoles()
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		user.AddRole(dataList["role_id[]"][i])
	}

	user.DeletePermissions()
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		user.AddPermission(dataList["permission_id[]"][i])
	}

}

func newRole(dataList form.Values) {

	role := models.Role().New(dataList.Get("name"), dataList.Get("slug"))

	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		role.AddPermission(dataList["permission_id[]"][i])
	}
}

func editRole(dataList form.Values) {

	role := models.RoleWithId(dataList.Get("id"))

	role.Update(dataList.Get("name"), dataList.Get("slug"))

	role.DeletePermissions()
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		role.AddPermission(dataList["permission_id[]"][i])
	}
}
