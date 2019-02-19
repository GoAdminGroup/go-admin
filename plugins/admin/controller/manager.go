package controller

import (
	"github.com/chenhg5/go-admin/modules/auth"
	"github.com/chenhg5/go-admin/modules/db"
)

func NewManager(dataList map[string][]string) {

	if dataList["name"][0] == "" ||
		dataList["username"][0] == "" ||
		dataList["password"][0] == "" {
		panic("账号密码不能为空")
	}

	// 更新管理员表
	result := db.Exec("insert into goadmin_users (username, password, name, avatar) values (?, ?, ?, ?)",
		dataList["username"][0], auth.EncodePassword([]byte(dataList["password"][0])), dataList["name"][0], dataList["avatar"][0])

	id, _ := result.LastInsertId()

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			db.Exec("insert into goadmin_role_users (role_id, user_id) values (?, ?)",
				dataList["role_id[]"][i], id)
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			db.Exec("insert into goadmin_user_permissions (permission_id, user_id) values (?, ?)",
				dataList["permission_id[]"][i], id)
		}
	}
}

func EditManager(dataList map[string][]string) {

	if dataList["name"][0] == "" ||
		dataList["username"][0] == "" ||
		dataList["password"][0] == "" {
		panic("账号密码不能为空")
	}

	// 更新管理员表
	db.Exec("update goadmin_users set username = ?, password = ?, name = ?, avatar = ? where id = ?",
		dataList["username"][0], auth.EncodePassword([]byte(dataList["password"][0])), dataList["name"][0],
		dataList["avatar"][0], dataList["id"][0])

	// 插入管理员角色表
	for i := 0; i < len(dataList["role_id[]"]); i++ {
		if dataList["role_id[]"][i] != "" {
			checkRole, _ := db.Query("select * from goadmin_role_users where role_id = ? and user_id = ?",
				dataList["role_id[]"][i], dataList["id"][0])
			if len(checkRole) < 1 {
				db.Exec("insert into goadmin_role_users (role_id, user_id) values (?, ?)",
					dataList["role_id[]"][i], dataList["id"][0])
			}
		}
	}

	// 更新管理员权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			checkPermission, _ := db.Query("select * from goadmin_user_permissions where permission_id = ? and user_id = ?",
				dataList["permission_id[]"][i], dataList["id"][0])
			if len(checkPermission) < 1 {
				db.Exec("insert into goadmin_user_permissions (permission_id, user_id) values (?, ?)",
					dataList["permission_id[]"][i], dataList["id"][0])
			}
		}
	}
}

func NewRole(dataList map[string][]string) {
	// 更新管理员角色表
	result := db.Exec("insert into goadmin_roles (name, slug) values (?, ?)",
		dataList["name"][0], dataList["slug"][0])

	id, _ := result.LastInsertId()

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			db.Exec("insert into goadmin_role_permissions (permission_id, role_id) values (?, ?)",
				dataList["permission_id[]"][i], id)
		}
	}
}

func EditRole(dataList map[string][]string) {
	// 更新管理员角色表
	db.Exec("update goadmin_roles set name = ?, slug = ? where id = ?",
		dataList["name"][0], dataList["slug"][0], dataList["id"][0])

	// 更新管理员角色权限表
	for i := 0; i < len(dataList["permission_id[]"]); i++ {
		if dataList["permission_id[]"][i] != "" {
			checkPermission, _ := db.Query("select * from goadmin_role_permissions where permission_id = ? and role_id = ?",
				dataList["permission_id[]"][i], dataList["id"][0])
			if len(checkPermission) < 1 {
				db.Exec("insert into goadmin_role_permissions (permission_id, role_id) values (?, ?)",
					dataList["permission_id[]"][i], dataList["id"][0])
			}
		}
	}
}
