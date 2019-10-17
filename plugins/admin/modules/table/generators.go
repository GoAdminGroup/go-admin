package table

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	template2 "html/template"
	"strconv"
	"strings"
)

func GetManagerTable() (ManagerTable Table) {
	ManagerTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := ManagerTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("Name"), "username", db.Varchar)
	info.AddField(lg("Nickname"), "name", db.Varchar)
	info.AddField(lg("role"), "roles", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labelModels, _ := db.Table("goadmin_role_users").
				Select("goadmin_roles.name").
				LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
				Where("user_id", "=", model.ID).
				All()

			labels := template2.HTML("")
			labelTpl := template.Get(config.Get().Theme).Label()

			for key, label := range labelModels {
				if key == len(labelModels)-1 {
					labels += labelTpl.SetContent(template2.HTML(label["name"].(string))).GetContent()
				} else {
					labels += labelTpl.SetContent(template2.HTML(label["name"].(string))).GetContent() + "<br><br>"
				}
			}

			return string(labels)
		})
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_users").SetTitle(lg("Managers")).SetDescription(lg("Managers"))

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

	formList := ManagerTable.GetForm()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField(lg("Name"), "username", db.Varchar, form.Text)
	formList.AddField(lg("Nickname"), "name", db.Varchar, form.Text)
	formList.AddField(lg("Avatar"), "avatar", db.Varchar, form.File)
	formList.AddField(lg("password"), "password", db.Varchar, form.Password)
	formList.AddField(lg("role"), "role_id", db.Varchar, form.Select).
		FieldOptions(roles).FieldDisplay(func(model types.FieldModel) interface{} {
		roleModel, _ := db.Table("goadmin_role_users").Select("role_id").
			Where("user_id", "=", model.ID).All()
		var roles []string
		for _, v := range roleModel {
			roles = append(roles, strconv.FormatInt(v["role_id"].(int64), 10))
		}
		return roles
	})
	formList.AddField(lg("permission"), "permission_id", db.Varchar, form.Select).
		FieldOptions(permissions).FieldDisplay(func(model types.FieldModel) interface{} {
		permissionModel, _ := db.Table("goadmin_user_permissions").
			Select("permission_id").Where("user_id", "=", model.ID).All()
		var permissions []string
		for _, v := range permissionModel {
			permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
		}
		return permissions
	})
	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_users").SetTitle(lg("Managers")).SetDescription(lg("Managers"))

	return
}

func GetPermissionTable() (PermissionTable Table) {
	PermissionTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := PermissionTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("permission"), "name", db.Varchar)
	info.AddField(lg("slug"), "slug", db.Varchar)
	info.AddField(lg("method"), "http_method", db.Varchar)
	info.AddField(lg("path"), "http_path", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
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
		})
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_permissions").
		SetTitle(lg("Permission Manage")).
		SetDescription(lg("Permission Manage"))

	formList := PermissionTable.GetForm()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField(lg("permission"), "name", db.Varchar, form.Text)
	formList.AddField(lg("slug"), "slug", db.Varchar, form.Text)
	formList.AddField(lg("method"), "http_method", db.Varchar, form.Select).
		FieldOptions([]map[string]string{
			{"value": "GET", "field": "GET"},
			{"value": "PUT", "field": "PUT"},
			{"value": "POST", "field": "POST"},
			{"value": "DELETE", "field": "DELETE"},
			{"value": "PATCH", "field": "PATCH"},
			{"value": "OPTIONS", "field": "OPTIONS"},
			{"value": "HEAD", "field": "HEAD"},
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			return strings.Split(model.Value, ",")
		}).
		FieldPostFilterFn(func(model types.PostFieldModel) string {
			return strings.Join(model.Value, ",")
		})

	formList.AddField(lg("path"), "http_path", db.Varchar, form.TextArea)
	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_permissions").
		SetTitle(lg("Permission Manage")).
		SetDescription(lg("Permission Manage"))

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

	info := RolesTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("role"), "name", db.Varchar)
	info.AddField(lg("slug"), "slug", db.Varchar)
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_roles").
		SetTitle(lg("Roles Manage")).
		SetDescription(lg("Roles Manage"))

	formList := RolesTable.GetForm()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField(lg("role"), "name", db.Varchar, form.Text)
	formList.AddField(lg("slug"), "slug", db.Varchar, form.Text)
	formList.AddField(lg("permission"), "permission_id", db.Varchar, form.SelectBox).
		FieldOptions(permissions).FieldDisplay(func(model types.FieldModel) interface{} {
		perModel, _ := db.Table("goadmin_role_permissions").
			Select("permission_id").
			Where("role_id", "=", model.ID).
			All()
		var permissions []string
		for _, v := range perModel {
			permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
		}
		return permissions
	})

	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_roles").
		SetTitle(lg("Roles Manage")).
		SetDescription(lg("Roles Manage"))

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

	info := OpTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("userID"), "user_id", db.Int)
	info.AddField(lg("path"), "path", db.Varchar)
	info.AddField(lg("method"), "method", db.Varchar)
	info.AddField(lg("ip"), "ip", db.Varchar)
	info.AddField(lg("content"), "input", db.Varchar)
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_operation_log").
		SetTitle(lg("operation log")).
		SetDescription(lg("operation log"))

	formList := OpTable.GetForm()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField(lg("userID"), "user_id", db.Int, form.Text)
	formList.AddField(lg("path"), "path", db.Varchar, form.Text)
	formList.AddField(lg("method"), "method", db.Varchar, form.Text)
	formList.AddField(lg("ip"), "ip", db.Varchar, form.Text)
	formList.AddField(lg("content"), "input", db.Varchar, form.Text)
	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_operation_log").
		SetTitle(lg("operation log")).
		SetDescription(lg("operation log"))

	return
}

func GetMenuTable() (MenuTable Table) {
	MenuTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := MenuTable.GetInfo()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("parent"), "parent_id", db.Int)
	info.AddField(lg("menu name"), "title", db.Varchar)
	info.AddField(lg("icon"), "icon", db.Varchar)
	info.AddField(lg("uri"), "uri", db.Varchar)
	info.AddField(lg("role"), "roles", db.Varchar)
	info.AddField(lg("header"), "header", db.Varchar)
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_menu").
		SetTitle(lg("Menus Manage")).
		SetDescription(lg("Menus Manage"))

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

	formList := MenuTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField(lg("parent"), "parent_id", db.Int, form.SelectSingle).
		FieldOptions(parents).FieldDisplay(func(model types.FieldModel) interface{} {
		menuModel, _ := db.Table("goadmin_menu").Select("parent_id").Find(model.ID)

		var menuItem []string
		menuItem = append(menuItem, strconv.FormatInt(menuModel["parent_id"].(int64), 10))
		return menuItem
	})
	formList.AddField(lg("menu name"), "title", db.Varchar, form.Text)
	formList.AddField(lg("header"), "header", db.Varchar, form.Text)
	formList.AddField(lg("icon"), "icon", db.Varchar, form.IconPicker)
	formList.AddField(lg("uri"), "uri", db.Varchar, form.Text)
	formList.AddField(lg("role"), "roles", db.Int, form.Select).
		FieldOptions(roles).FieldDisplay(func(model types.FieldModel) interface{} {
		roleModel, _ := db.Table("goadmin_role_menu").
			Select("role_id").
			Where("menu_id", "=", model.ID).
			All()
		var roles []string
		for _, v := range roleModel {
			roles = append(roles, strconv.FormatInt(v["role_id"].(int64), 10))
		}
		return roles
	})

	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_menu").
		SetTitle(lg("Menus Manage")).
		SetDescription(lg("Menus Manage"))

	return
}

func lg(v string) string {
	return language.Get(v)
}
