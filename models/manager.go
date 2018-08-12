package models

import (
	"goAdmin/components"
	"goAdmin/connections/mysql"
	"strconv"
	"strings"
)

func GetManagerTable() (ManagerTable GlobalTable) {

	ManagerTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "角色",
			Field:    "roles",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				labelModel, _ := mysql.Query("select r.name from goadmin_role_users as u left join goadmin_roles as r on "+
					"u.role_id = r.id where user_id = ?", model.ID)
				return components.Label.GetContent(labelModel[0]["name"].(string))
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	ManagerTable.Info.Table = "goadmin_users"
	ManagerTable.Info.Title = "管理员管理"
	ManagerTable.Info.Description = "管理员管理"

	var roles, permissions []map[string]string
	rolesModel, _ := mysql.Query("select `id`, `slug` from goadmin_roles where id > ?", 0)
	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}
	permissionsModel, _ := mysql.Query("select `id`, `slug` from goadmin_permissions where id > ?", 0)
	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	ManagerTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "头像",
			Field:    "avatar",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "file",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "密码",
			Field:    "password",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "password",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "角色",
			Field:    "role_id",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options:  roles,
			ExcuFun: func(model RowModel) interface{} {
				roleModel, _ := mysql.Query("select role_id from goadmin_role_users where user_id = ?", model.ID)
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
			ExcuFun: func(model RowModel) interface{} {
				permissionModel, _ := mysql.Query("select permission_id from goadmin_user_permissions where user_id = ?", model.ID)
				var permissions []string
				for _, v := range permissionModel {
					permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
				}
				return permissions
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	ManagerTable.Form.Table = "goadmin_users"
	ManagerTable.Form.Title = "管理员管理"
	ManagerTable.Form.Description = "管理员管理"

	return
}

func GetPermissionTable() (PermissionTable GlobalTable) {

	PermissionTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "方法",
			Field:    "http_method",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "路径",
			Field:    "http_path",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	PermissionTable.Info.Table = "goadmin_permissions"
	PermissionTable.Info.Title = "权限管理"
	PermissionTable.Info.Description = "权限管理"

	PermissionTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "方法",
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
			ExcuFun: func(model RowModel) interface{} {
				return strings.Split(model.Value, ",")
			},
		}, {
			Head:     "路径",
			Field:    "http_path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "textarea",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	PermissionTable.Form.Table = "goadmin_permissions"
	PermissionTable.Form.Title = "权限管理"
	PermissionTable.Form.Description = "权限管理"

	return
}

func GetRolesTable() (RolesTable GlobalTable) {

	var permissions []map[string]string
	permissionsModel, _ := mysql.Query("select `id`, `slug` from goadmin_permissions where id > ?", 0)
	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	RolesTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	RolesTable.Info.Table = "goadmin_roles"
	RolesTable.Info.Title = "角色管理"
	RolesTable.Info.Description = "角色管理"

	RolesTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
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
			ExcuFun: func(model RowModel) interface{} {
				perModel, _ := mysql.Query("select permission_id from goadmin_role_permissions where role_id = ?", model.ID)
				var permissions []string
				for _, v := range perModel {
					permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
				}
				return permissions
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	RolesTable.Form.Table = "goadmin_roles"
	RolesTable.Form.Title = "角色管理"
	RolesTable.Form.Description = "角色管理"

	return
}

func GetOpTable() (OpTable GlobalTable) {

	OpTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "路径",
			Field:    "path",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "方法",
			Field:    "method",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	OpTable.Info.Table = "goadmin_operation_log"
	OpTable.Info.Title = "操作日志"
	OpTable.Info.Description = "操作日志"

	OpTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "路径",
			Field:    "path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "方法",
			Field:    "method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
			ExcuFun: func(model RowModel) interface{} {
				return model.Value
			},
		},
	}

	OpTable.Form.Table = "goadmin_operation_log"
	OpTable.Form.Title = "操作日志"
	OpTable.Form.Description = "操作日志"

	return
}
