package models

import (
	"github.com/chenhg5/go-admin/modules/connections"
	"strconv"
	"strings"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/modules/config"
)

func GetManagerTable() (ManagerTable GlobalTable) {

	ManagerTable.Info.FieldList = []types.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Name",
			Field:    "username",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Nickname",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Role",
			Field:    "roles",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				labelModel, _ := connections.GetConnection().Query("select r.name from goadmin_role_users as u left join goadmin_roles as r on "+
					"u.role_id = r.id where user_id = ?", model.ID)
				return string(template.Get("adminlte").Label().SetContent(labelModel[0]["name"].(string)).GetContent())
			},
		},
		{
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	ManagerTable.Info.Table = "goadmin_users"
	ManagerTable.Info.Title = "Managers"
	ManagerTable.Info.Description = "Managers"

	var roles, permissions []map[string]string
	rolesModel, _ := connections.GetConnection().Query("select `id`, `slug` from goadmin_roles where id > ?", 0)
	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}
	permissionsModel, _ := connections.GetConnection().Query("select `id`, `slug` from goadmin_permissions where id > ?", 0)
	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	ManagerTable.Form.FormList = []types.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Nickname",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "头像",
			Field:    "avatar",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "file",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "密码",
			Field:    "password",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "password",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Role",
			Field:    "role_id",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options:  roles,
			ExcuFun: func(model types.RowModel) interface{} {
				roleModel, _ := connections.GetConnection().Query("select role_id from goadmin_role_users where user_id = ?", model.ID)
				var roles []string
				for _, v := range roleModel {
					roles = append(roles, strconv.FormatInt(v["role_id"].(int64), 10))
				}
				return roles
			},
		}, {
			Head:     "权限",
			Field:    "permission_id",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options:  permissions,
			ExcuFun: func(model types.RowModel) interface{} {
				permissionModel, _ := connections.GetConnection().Query("select permission_id from goadmin_user_permissions where user_id = ?", model.ID)
				var permissions []string
				for _, v := range permissionModel {
					permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
				}
				return permissions
			},
		}, {
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	ManagerTable.Form.Table = "goadmin_users"
	ManagerTable.Form.Title = "Managers"
	ManagerTable.Form.Description = "Managers"

	ManagerTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetPermissionTable() (PermissionTable GlobalTable) {

	PermissionTable.Info.FieldList = []types.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "method",
			Field:    "http_method",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "path",
			Field:    "http_path",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	PermissionTable.Info.Table = "goadmin_permissions"
	PermissionTable.Info.Title = "Permission Manage"
	PermissionTable.Info.Description = "Permission Manage"

	PermissionTable.Form.FormList = []types.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "method",
			Field:    "http_method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options: []map[string]string{
				{"value": "GET", "field": "GET"},
				{"value": "PUT", "field": "PUT"},
				{"value": "POST", "field": "POST"},
				{"value": "DELETE", "field": "DELETE"},
				{"value": "PATCH", "field": "PATCH"},
				{"value": "OPTIONS", "field": "OPTIONS"},
				{"value": "HEAD", "field": "HEAD"},
			},
			ExcuFun: func(model types.RowModel) interface{} {
				return strings.Split(model.Value, ",")
			},
		}, {
			Head:     "path",
			Field:    "http_path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "textarea",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	PermissionTable.Form.Table = "goadmin_permissions"
	PermissionTable.Form.Title = "Permission Manage"
	PermissionTable.Form.Description = "Permission Manage"

	PermissionTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetRolesTable() (RolesTable GlobalTable) {

	var permissions []map[string]string
	permissionsModel, _ := connections.GetConnection().Query("select `id`, `slug` from goadmin_permissions where id > ?", 0)
	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	RolesTable.Info.FieldList = []types.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	RolesTable.Info.Table = "goadmin_roles"
	RolesTable.Info.Title = "角色管理"
	RolesTable.Info.Description = "角色管理"

	RolesTable.Form.FormList = []types.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "权限",
			Field:    "permission_id",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "selectbox",
			Options:  permissions,
			ExcuFun: func(model types.RowModel) interface{} {
				perModel, _ := connections.GetConnection().Query("select permission_id from goadmin_role_permissions where role_id = ?", model.ID)
				var permissions []string
				for _, v := range perModel {
					permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
				}
				return permissions
			},
		}, {
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	RolesTable.Form.Table = "goadmin_roles"
	RolesTable.Form.Title = "角色管理"
	RolesTable.Form.Description = "角色管理"

	RolesTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetOpTable() (OpTable GlobalTable) {

	OpTable.Info.FieldList = []types.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "path",
			Field:    "path",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "method",
			Field:    "method",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	OpTable.Info.Table = "goadmin_operation_log"
	OpTable.Info.Title = "操作日志"
	OpTable.Info.Description = "操作日志"

	OpTable.Form.FormList = []types.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "path",
			Field:    "path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "method",
			Field:    "method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	OpTable.Form.Table = "goadmin_operation_log"
	OpTable.Form.Title = "操作日志"
	OpTable.Form.Description = "操作日志"

	OpTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetMenuTable() (MenuTable GlobalTable) {

	MenuTable.Info.FieldList = []types.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "父级",
			Field:    "parent_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Name",
			Field:    "title",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Icon",
			Field:    "icon",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Uri",
			Field:    "uri",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "角色",
			Field:    "roles",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	MenuTable.Info.Table = "goadmin_menu"
	MenuTable.Info.Title = "菜单"
	MenuTable.Info.Description = "菜单"

	var roles, parents []map[string]string
	rolesModel, _ := connections.GetConnection().Query("select `id`, `slug` from goadmin_roles where id > ?", 0)
	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}
	parentsModel, _ := connections.GetConnection().Query("select `id`, `title` from goadmin_menu where id > ? order by `order` asc", 0)
	for _, v := range parentsModel {
		parents = append(parents, map[string]string{
			"field": v["title"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}
	parents = append([]map[string]string{{
		"field": "root",
		"value": "0",
	}}, parents...)

	MenuTable.Form.FormList = []types.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "父级",
			Field:    "parent_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "select_single",
			Options:  parents,
			ExcuFun: func(model types.RowModel) interface{} {
				menuModel, _ := connections.GetConnection().Query("select parent_id from goadmin_menu where id = ?", model.ID)
				var menuItem []string
				menuItem = append(menuItem, strconv.FormatInt(menuModel[0]["parent_id"].(int64), 10))
				return menuItem
			},
		}, {
			Head:     "Name",
			Field:    "title",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Icon",
			Field:    "icon",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "iconpicker",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Uri",
			Field:    "uri",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "角色",
			Field:    "roles",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options:  roles,
			ExcuFun: func(model types.RowModel) interface{} {
				roleModel, _ := connections.GetConnection().Query("select role_id from goadmin_role_menu where menu_id = ?", model.ID)
				var roles []string
				for _, v := range roleModel {
					roles = append(roles, strconv.FormatInt(v["role_id"].(int64), 10))
				}
				return roles
			},
		}, {
			Head:     "updatedAt",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "createdAt",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	MenuTable.Form.Table = "goadmin_menu"
	MenuTable.Form.Title = "菜单"
	MenuTable.Form.Description = "菜单"

	MenuTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}
