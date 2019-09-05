package controller

import (
	"errors"
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func newManager(dataList map[string][]string) error {

	if dataList["name"][0] == "" ||
		dataList["username"][0] == "" ||
		dataList["password"][0] == "" {
		return errors.New("username and password can not be empty")
	}

	user := models.User().New(dataList["username"][0],
		auth.EncodePassword([]byte(dataList["password"][0])),
		dataList["name"][0],
		dataList["avatar"][0])

	for i := 0; i < len(dataList["role_id[]"]); i++ {
		user.AddRole(dataList["role_id[]"][i])
	}

	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		user.AddPermission(dataList["permission_id[]"][i])
	}

	return nil
}

func editManager(dataList map[string][]string) error {

	if dataList["name"][0] == "" ||
		dataList["username"][0] == "" ||
		dataList["password"][0] == "" {
		return errors.New("username and password can not be empty")
	}

	user := models.UserWithId(dataList["id"][0])

	user.Update(dataList["username"][0],
		auth.EncodePassword([]byte(dataList["password"][0])),
		dataList["name"][0],
		dataList["avatar"][0])

	for i := 0; i < len(dataList["role_id[]"]); i++ {
		user.AddRole(dataList["role_id[]"][i])
	}

	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		user.AddPermission(dataList["permission_id[]"][i])
	}

	return nil
}

func newRole(dataList map[string][]string) error {

	role := models.Role().New(dataList["name"][0], dataList["slug"][0])

	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		role.AddPermission(dataList["permission_id[]"][i])
	}

	return nil
}

func editRole(dataList map[string][]string) error {

	role := models.RoleWithId(dataList["id"][0])

	role.Update(dataList["name"][0], dataList["slug"][0])

	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		role.AddPermission(dataList["permission_id[]"][i])
	}

	return nil
}
