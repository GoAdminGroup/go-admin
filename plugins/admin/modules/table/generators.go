package table

import (
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
	"strconv"
	"strings"
)

func GetManagerTable() (ManagerTable Table) {
	ManagerTable = NewDefaultTable(DefaultTableConfigWithDriver(config.Get().DATABASE.GetDefault().DRIVER))
	ManagerTable.GetInfo().FieldList = []types.Field{
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
				labelModel, _ := db.Table("goadmin_role_users").
					Select("goadmin_roles.name").
					LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
					Where("user_id", "=", model.ID).
					First()

				return string(template.Get("adminlte").Label().SetContent(labelModel["name"].(string)).GetContent())
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

	ManagerTable.GetInfo().Table = "goadmin_users"
	ManagerTable.GetInfo().Title = language.Get("Managers")
	ManagerTable.GetInfo().Description = language.Get("Managers")

	var roles, permissions []map[string]string
	rolesModel, _ := db.Table("goadmin_roles").Select("id", "slug").Where("id", ">", 0).All()

	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}
	permissionsModel, _ := db.Table("goadmin_permissions").Select("id", "slug").Where("id", ">", 0).All()
	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	ManagerTable.GetForm().FormList = []types.Form{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Name"),
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Nickname"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Avatar"),
			Field:    "avatar",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.File,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("password"),
			Field:    "password",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Password,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
			Field:    "role_id",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Select,
			Options:  roles,
			ExcuFun: func(model types.RowModel) interface{} {
				roleModel, _ := db.Table("goadmin_role_users").Select("role_id").Where("user_id", "=", model.ID).All()
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
			FormType: form.Select,
			Options:  permissions,
			ExcuFun: func(model types.RowModel) interface{} {
				permissionModel, _ := db.Table("goadmin_user_permissions").Select("permission_id").Where("user_id", "=", model.ID).All()
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
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	ManagerTable.GetForm().Table = "goadmin_users"
	ManagerTable.GetForm().Title = language.Get("Managers")
	ManagerTable.GetForm().Description = language.Get("Managers")

	return
}

func GetPermissionTable() (PermissionTable Table) {
	PermissionTable = NewDefaultTable(DefaultTableConfigWithDriver(config.Get().DATABASE.GetDefault().DRIVER))
	PermissionTable.GetInfo().FieldList = []types.Field{
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

	PermissionTable.GetInfo().Table = "goadmin_permissions"
	PermissionTable.GetInfo().Title = language.Get("Permission Manage")
	PermissionTable.GetInfo().Description = language.Get("Permission Manage")

	PermissionTable.GetForm().FormList = []types.Form{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("name"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("method"),
			Field:    "http_method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Select,
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
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	PermissionTable.GetForm().Table = "goadmin_permissions"
	PermissionTable.GetForm().Title = language.Get("Permission Manage")
	PermissionTable.GetForm().Description = language.Get("Permission Manage")

	return
}

func GetRolesTable() (RolesTable Table) {
	RolesTable = NewDefaultTable(DefaultTableConfigWithDriver(config.Get().DATABASE.GetDefault().DRIVER))
	var permissions []map[string]string
	permissionsModel, _ := db.Table("goadmin_permissions").Select("id", "slug").Where("id", ">", 0).All()

	for _, v := range permissionsModel {
		permissions = append(permissions, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	RolesTable.GetInfo().FieldList = []types.Field{
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

	RolesTable.GetInfo().Table = "goadmin_roles"
	RolesTable.GetInfo().Title = language.Get("Roles Manage")
	RolesTable.GetInfo().Description = language.Get("Roles Manage")

	RolesTable.GetForm().FormList = []types.Form{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("name"),
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("permission"),
			Field:    "permission_id",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.SelectBox,
			Options:  permissions,
			ExcuFun: func(model types.RowModel) interface{} {
				perModel, _ := db.Table("goadmin_role_permissions").
					Select("permission_id").
					Where("role_id", "=", model.ID).
					All()
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
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	RolesTable.GetForm().Table = "goadmin_roles"
	RolesTable.GetForm().Title = language.Get("Roles Manage")
	RolesTable.GetForm().Description = language.Get("Roles Manage")

	return
}

func GetOpTable() (OpTable Table) {
	OpTable = NewDefaultTable(DefaultTableConfigWithDriver(config.Get().DATABASE.GetDefault().DRIVER))
	OpTable.GetInfo().FieldList = []types.Field{
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

	OpTable.GetInfo().Table = "goadmin_operation_log"
	OpTable.GetInfo().Title = language.Get("operation log")
	OpTable.GetInfo().Description = language.Get("operation log")

	OpTable.GetForm().FormList = []types.Form{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("userID"),
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("path"),
			Field:    "path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("method"),
			Field:    "method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("content"),
			Field:    "input",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	OpTable.GetForm().Table = "goadmin_operation_log"
	OpTable.GetForm().Title = language.Get("operation log")
	OpTable.GetForm().Description = language.Get("operation log")

	return
}

func GetMenuTable() (MenuTable Table) {
	MenuTable = NewDefaultTable(DefaultTableConfigWithDriver(config.Get().DATABASE.GetDefault().DRIVER))
	MenuTable.GetInfo().FieldList = []types.Field{
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
			Head:     language.Get("menu name"),
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

	MenuTable.GetInfo().Table = "goadmin_menu"
	MenuTable.GetInfo().Title = language.Get("Menus Manage")
	MenuTable.GetInfo().Description = language.Get("Menus Manage")

	var roles, parents []map[string]string
	rolesModel, _ := db.Table("goadmin_roles").Select("id", "slug").Where("id", ">", 0).All()

	for _, v := range rolesModel {
		roles = append(roles, map[string]string{
			"field": v["slug"].(string),
			"value": strconv.FormatInt(v["id"].(int64), 10),
		})
	}

	parentsModel, _ := db.Table("goadmin_menu").
		Select("id", "title").
		Where("id", ">", 0).
		OrderBy("order", "asc").
		All()

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

	MenuTable.GetForm().FormList = []types.Form{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("parent"),
			Field:    "parent_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: form.SelectSingle,
			Options:  parents,
			ExcuFun: func(model types.RowModel) interface{} {
				menuModel, _ := db.Table("goadmin_menu").Select("parent_id").Find(model.ID)

				var menuItem []string
				menuItem = append(menuItem, strconv.FormatInt(menuModel["parent_id"].(int64), 10))
				return menuItem
			},
		}, {
			Head:     language.Get("menu name"),
			Field:    "title",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("icon"),
			Field:    "icon",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.IconPicker,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("uri"),
			Field:    "uri",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
			Field:    "roles",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Select,
			Options:  roles,
			ExcuFun: func(model types.RowModel) interface{} {
				roleModel, _ := db.Table("goadmin_role_menu").
					Select("role_id").
					Where("menu_id", "=", model.ID).
					All()
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
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	MenuTable.GetForm().Table = "goadmin_menu"
	MenuTable.GetForm().Title = language.Get("Menus Manage")
	MenuTable.GetForm().Description = language.Get("Menus Manage")

	return
}
