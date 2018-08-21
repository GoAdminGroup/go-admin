package controller

import (
	"goAdmin/modules/auth"
	"goAdmin/modules/connections/mysql"
)

func NewManager(dataList map[string][]string) {

	// 更新管理员表
	result := mysql.Exec("insert into goadmin_users (username, password, name, avatar) values (?, ?, ?, ?)",
		dataList["username"][0], auth.EncodePassword([]byte(dataList["password"][0])), dataList["name"][0], dataList["avatar"][0])

	id, _ := result.LastInsertId()

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			mysql.Exec("insert into goadmin_role_users (role_id, user_id) values (?, ?)",
				dataList["role_id[]"][i], id)
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			mysql.Exec("insert into goadmin_user_permissions (permission_id, user_id) values (?, ?)",
				dataList["permission_id[]"][i], id)
		}
	}
}

func EditManager(dataList map[string][]string) {

	// 更新管理员表
	mysql.Exec("update goadmin_users set username = ?, password = ?, name = ?, avatar = ? where id = ?",
		dataList["username"][0], auth.EncodePassword([]byte(dataList["password"][0])), dataList["name"][0],
		dataList["avatar"][0], dataList["id"][0])

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			mysql.Exec("insert into goadmin_role_users (role_id, user_id) values (?, ?)",
				dataList["role_id[]"][i], dataList["id"][0])
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			mysql.Exec("insert into goadmin_user_permissions (permission_id, user_id) values (?, ?)",
				dataList["permission_id[]"][i], dataList["id"][0])
		}
	}
}

func NewRole(dataList map[string][]string) {
	// 更新管理员角色表
	result := mysql.Exec("insert into goadmin_roles (name, slug) values (?, ?)",
		dataList["name"][0], dataList["slug"][0])

	id, _ := result.LastInsertId()

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			mysql.Exec("insert into goadmin_role_permissions (permission_id, role_id) values (?, ?)",
				dataList["permission_id[]"][i], id)
		}
	}
}

func EditRole(dataList map[string][]string) {
	// 更新管理员角色表
	mysql.Exec("update goadmin_roles set name = ?, slug = ? where id = ?",
		dataList["name"][0], dataList["slug"][0], dataList["id"][0])

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			mysql.Exec("insert into goadmin_role_permissions (permission_id, role_id) values (?, ?)",
				dataList["permission_id[]"][i], dataList["id"][0])
		}
	}
}
