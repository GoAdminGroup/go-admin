package models

import (
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/connections"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"strconv"
	"strings"
)

func GetManagerTable() (ManagerTable Table) {

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
			Head:     language.Get("Name"),
			Field:    "username",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("Nickname"),
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("role"),
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
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	ManagerTable.Info.Table = "goadmin_users"
	ManagerTable.Info.Title = language.Get("Managers")
	ManagerTable.Info.Description = language.Get("Managers")

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
			Head:     language.Get("Name"),
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Nickname"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Avatar"),
			Field:    "avatar",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "file",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("password"),
			Field:    "password",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "password",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
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
			Head:     language.Get("permission"),
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
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
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
	ManagerTable.Form.Title = language.Get("Managers")
	ManagerTable.Form.Description = language.Get("Managers")

	ManagerTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetPermissionTable() (PermissionTable Table) {

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
			Head:     language.Get("name"),
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("method"),
			Field:    "http_method",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("path"),
			Field:    "http_path",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	PermissionTable.Info.Table = "goadmin_permissions"
	PermissionTable.Info.Title = language.Get("Permission Manage")
	PermissionTable.Info.Description = language.Get("Permission Manage")

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
			Head:     language.Get("name"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("method"),
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
			Head:     language.Get("path"),
			Field:    "http_path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "textarea",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
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
	PermissionTable.Form.Title = language.Get("Permission Manage")
	PermissionTable.Form.Description = language.Get("Permission Manage")

	PermissionTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetRolesTable() (RolesTable Table) {

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
			Head:     language.Get("name"),
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	RolesTable.Info.Table = "goadmin_roles"
	RolesTable.Info.Title = language.Get("Roles Manage")
	RolesTable.Info.Description = language.Get("Roles Manage")

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
			Head:     language.Get("name"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("permission"),
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
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
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
	RolesTable.Form.Title = language.Get("Roles Manage")
	RolesTable.Form.Description = language.Get("Roles Manage")

	RolesTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetOpTable() (OpTable Table) {

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
			Head:     language.Get("userID"),
			Field:    "user_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("path"),
			Field:    "path",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("method"),
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
			Head:     language.Get("content"),
			Field:    "input",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	OpTable.Info.Table = "goadmin_operation_log"
	OpTable.Info.Title = language.Get("operation log")
	OpTable.Info.Description = language.Get("operation log")

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
			Head:     language.Get("userID"),
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("path"),
			Field:    "path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("method"),
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
			Head:     language.Get("content"),
			Field:    "input",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
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
	OpTable.Form.Title = language.Get("operation log")
	OpTable.Form.Description = language.Get("operation log")

	OpTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}

func GetMenuTable() (MenuTable Table) {

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
			Head:     language.Get("parent"),
			Field:    "parent_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("name"),
			Field:    "title",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("icon"),
			Field:    "icon",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("uri"),
			Field:    "uri",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("role"),
			Field:    "roles",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	MenuTable.Info.Table = "goadmin_menu"
	MenuTable.Info.Title = language.Get("Menus Manage")
	MenuTable.Info.Description = language.Get("Menus Manage")

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
			Head:     language.Get("parent"),
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
			Head:     language.Get("name"),
			Field:    "title",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("icon"),
			Field:    "icon",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "iconpicker",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("uri"),
			Field:    "uri",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
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
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
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
	MenuTable.Form.Title = language.Get("Menus Manage")
	MenuTable.Form.Description = language.Get("Menus Manage")

	MenuTable.ConnectionDriver = config.Get().DATABASE[0].DRIVER

	return
}
