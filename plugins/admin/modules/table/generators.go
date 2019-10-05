package table

import (
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
	template2 "html/template"
	"strconv"
	"strings"
)

func GetManagerTable() (ManagerTable Table) {
	ManagerTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))
	ManagerTable.GetInfo().FieldList = []types.Field{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: db.Int,
			Sortable: true,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("Name"),
			Field:    "username",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("Nickname"),
			Field:    "name",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("role"),
			Field:    "roles",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				labelModel, _ := db.Table("goadmin_role_users").
					Select("goadmin_roles.name").
					LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
					Where("user_id", "=", model.ID).
					First()

				return string(template.Get("adminlte").Label().SetContent(template2.HTML(labelModel["name"].(string))).GetContent())
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Int,
			Default:  "",
			Editable: false,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Name"),
			Field:    "username",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Nickname"),
			Field:    "name",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("Avatar"),
			Field:    "avatar",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.File,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("password"),
			Field:    "password",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Password,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
			Field:    "role_id",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Select,
			Options:  roles,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Select,
			Options:  permissions,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
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
	PermissionTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))
	PermissionTable.GetInfo().FieldList = []types.Field{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: db.Int,
			Sortable: true,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("permission"),
			Field:    "name",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("method"),
			Field:    "http_method",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("path"),
			Field:    "http_path",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				pathArr := strings.Split(model.Value, "\n")
				res := ""
				for i := 0; i < len(pathArr); i++ {
					if i == len(pathArr)-1 {
						res += string(template.Get(config.Get().Theme).Label().SetContent(template2.HTML(pathArr[i])).GetContent())
					} else {
						res += string(template.Get(config.Get().Theme).Label().SetContent(template2.HTML(pathArr[i])).GetContent()) + "<br><br>"
					}
				}
				return res
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Int,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("permission"),
			Field:    "name",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("method"),
			Field:    "http_method",
			TypeName: db.Varchar,
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
			FilterFn: func(model types.RowModel) interface{} {
				return strings.Split(model.Value, ",")
			},
			PostFilterFn: func(model types.PostRowModel) string {
				return strings.Join(model.Value, ",")
			},
		}, {
			Head:     language.Get("path"),
			Field:    "http_path",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.TextArea,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
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
	RolesTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))
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
			TypeName: db.Int,
			Sortable: true,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("role"),
			Field:    "name",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Int,
			Default:  "",
			Editable: false,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
			Field:    "name",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("slug"),
			Field:    "slug",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("permission"),
			Field:    "permission_id",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.SelectBox,
			Options:  permissions,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
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
	OpTable = NewDefaultTable(Config{
		Driver:     config.Get().Databases.GetDefault().Driver,
		CanAdd:     false,
		Editable:   false,
		Deletable:  false,
		Exportable: false,
		Connection: "default",
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	})
	OpTable.GetInfo().FieldList = []types.Field{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: db.Int,
			Sortable: true,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("userID"),
			Field:    "user_id",
			TypeName: db.Int,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("path"),
			Field:    "path",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("method"),
			Field:    "method",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "ip",
			Field:    "ip",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("content"),
			Field:    "input",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Int,
			Default:  "",
			Editable: false,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("userID"),
			Field:    "user_id",
			TypeName: db.Int,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("path"),
			Field:    "path",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("method"),
			Field:    "method",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "ip",
			Field:    "ip",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("content"),
			Field:    "input",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
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
	MenuTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))
	MenuTable.GetInfo().FieldList = []types.Field{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: db.Int,
			Sortable: true,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("parent"),
			Field:    "parent_id",
			TypeName: db.Int,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("menu name"),
			Field:    "title",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("icon"),
			Field:    "icon",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("uri"),
			Field:    "uri",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("role"),
			Field:    "roles",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("header"),
			Field:    "header",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     language.Get("updatedAt"),
			Field:    "updated_at",
			TypeName: db.Timestamp,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Int,
			Default:  "",
			Editable: false,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("parent"),
			Field:    "parent_id",
			TypeName: db.Int,
			Default:  "",
			Editable: true,
			FormType: form.SelectSingle,
			Options:  parents,
			FilterFn: func(model types.RowModel) interface{} {
				menuModel, _ := db.Table("goadmin_menu").Select("parent_id").Find(model.ID)

				var menuItem []string
				menuItem = append(menuItem, strconv.FormatInt(menuModel["parent_id"].(int64), 10))
				return menuItem
			},
		}, {
			Head:     language.Get("menu name"),
			Field:    "title",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("header"),
			Field:    "header",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("icon"),
			Field:    "icon",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.IconPicker,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("uri"),
			Field:    "uri",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("role"),
			Field:    "roles",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Select,
			Options:  roles,
			FilterFn: func(model types.RowModel) interface{} {
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
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     language.Get("createdAt"),
			Field:    "created_at",
			TypeName: db.Timestamp,
			Default:  "",
			Editable: true,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	MenuTable.GetForm().Table = "goadmin_menu"
	MenuTable.GetForm().Title = language.Get("Menus Manage")
	MenuTable.GetForm().Description = language.Get("Menus Manage")

	return
}
