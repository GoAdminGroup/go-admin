package table

import (
	"database/sql"
	"errors"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

func GetManagerTable() (ManagerTable Table) {
	ManagerTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := ManagerTable.GetInfo().AddXssJsFilter()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("Name"), "username", db.Varchar).FieldFilterable()
	info.AddField(lg("Nickname"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("role"), "roles", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labelModels, _ := db.Table("goadmin_role_users").
				Select("goadmin_roles.name").
				LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
				Where("user_id", "=", model.ID).
				All()

			labels := template.HTML("")
			labelTpl := template.Get(config.Get().Theme).Label()

			for key, label := range labelModels {
				if key == len(labelModels)-1 {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent()
				} else {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent() + "<br><br>"
				}
			}

			return string(labels)
		}).FieldFilterable()
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_users").
		SetTitle(lg("Managers")).
		SetDescription(lg("Managers")).
		SetDeleteFn(func(idArr []string) error {

			var ids = interfaces(idArr)

			_, txErr := dbsql().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteUserRoleErr := dbsql().WithTx(tx).
					Table("goadmin_role_users").
					WhereIn("user_id", ids).
					Delete()

				if deleteUserRoleErr != nil && notNoAffectRow(deleteUserRoleErr) {
					return deleteUserRoleErr, map[string]interface{}{}
				}

				deleteUserPermissionErr := dbsql().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("user_id", ids).
					Delete()

				if deleteUserPermissionErr != nil && notNoAffectRow(deleteUserPermissionErr) {
					return deleteUserPermissionErr, map[string]interface{}{}
				}

				deleteUserErr := dbsql().WithTx(tx).
					Table("goadmin_users").
					WhereIn("id", ids).
					Delete()

				if deleteUserErr != nil {
					return deleteUserErr, map[string]interface{}{}
				}

				return nil, map[string]interface{}{}
			})

			return txErr
		})

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

	formList := ManagerTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField(lg("Name"), "username", db.Varchar, form.Text)
	formList.AddField(lg("Nickname"), "name", db.Varchar, form.Text)
	formList.AddField(lg("Avatar"), "avatar", db.Varchar, form.File)
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
	formList.AddField(lg("password"), "password", db.Varchar, form.Password).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return ""
		})
	formList.AddField(lg("confirm password"), "password_again", db.Varchar, form.Password).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return ""
		})

	formList.SetTable("goadmin_users").SetTitle(lg("Managers")).SetDescription(lg("Managers"))
	formList.SetUpdateFn(func(values form2.Values) error {

		if values.IsEmpty("name", "username") {
			return errors.New("username and password can not be empty")
		}

		user := models.UserWithId(values.Get("id"))

		password := values.Get("password")

		if password != "" {

			if password != values.Get("password_again") {
				panic("password does not match")
			}

			password = encodePassword([]byte(values.Get("password")))
		}

		user.Update(values.Get("username"), password, values.Get("name"), values.Get("avatar"))

		user.DeleteRoles()
		for i := 0; i < len(values["role_id[]"]); i++ {
			user.AddRole(values["role_id[]"][i])
		}

		user.DeletePermissions()
		for i := 0; i < len(values["permission_id[]"]); i++ {
			user.AddPermission(values["permission_id[]"][i])
		}

		return nil
	})
	formList.SetInsertFn(func(values form2.Values) error {
		if values.IsEmpty("name", "username", "password") {
			return errors.New("username and password can not be empty")
		}

		user := models.User().New(values.Get("username"),
			encodePassword([]byte(values.Get("password"))),
			values.Get("name"),
			values.Get("avatar"))

		// TODO: Add transaction support.

		for i := 0; i < len(values["role_id[]"]); i++ {
			user.AddRole(values["role_id[]"][i])
		}

		for i := 0; i < len(values["permission_id[]"]); i++ {
			user.AddPermission(values["permission_id[]"][i])
		}
		return nil
	})

	return
}

func encodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

func GetPermissionTable() (PermissionTable Table) {
	PermissionTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := PermissionTable.GetInfo().AddXssJsFilter()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("permission"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("slug"), "slug", db.Varchar).FieldFilterable()
	info.AddField(lg("method"), "http_method", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		if value.Value == "" {
			return "All methods"
		}
		return value.Value
	})
	info.AddField(lg("path"), "http_path", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			pathArr := strings.Split(model.Value, "\n")
			res := ""
			for i := 0; i < len(pathArr); i++ {
				if i == len(pathArr)-1 {
					res += string(template.Get(config.Get().Theme).Label().SetContent(template.HTML(pathArr[i])).GetContent())
				} else {
					res += string(template.Get(config.Get().Theme).Label().SetContent(template.HTML(pathArr[i])).GetContent()) + "<br><br>"
				}
			}
			return res
		})
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_permissions").
		SetTitle(lg("Permission Manage")).
		SetDescription(lg("Permission Manage")).
		SetDeleteFn(func(idArr []string) error {

			var ids = interfaces(idArr)

			_, txErr := dbsql().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteRolePermissionErr := dbsql().WithTx(tx).
					Table("goadmin_role_permissions").
					WhereIn("permission_id", ids).
					Delete()

				if deleteRolePermissionErr != nil && notNoAffectRow(deleteRolePermissionErr) {
					return deleteRolePermissionErr, map[string]interface{}{}
				}

				deleteUserPermissionErr := dbsql().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("permission_id", ids).
					Delete()

				if deleteUserPermissionErr != nil && notNoAffectRow(deleteUserPermissionErr) {
					return deleteUserPermissionErr, map[string]interface{}{}
				}

				deletePermissionsErr := dbsql().WithTx(tx).
					Table("goadmin_permissions").
					WhereIn("id", ids).
					Delete()

				if deletePermissionsErr != nil {
					return deletePermissionsErr, map[string]interface{}{}
				}

				return nil, map[string]interface{}{}
			})

			return txErr
		})

	formList := PermissionTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
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
		}).
		FieldHelpMsg(template.HTML(lg("all method if empty")))

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

	info := RolesTable.GetInfo().AddXssJsFilter()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("role"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("slug"), "slug", db.Varchar).FieldFilterable()
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_roles").
		SetTitle(lg("Roles Manage")).
		SetDescription(lg("Roles Manage")).
		SetDeleteFn(func(idArr []string) error {

			var ids = interfaces(idArr)

			_, txErr := dbsql().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteRoleUserErr := dbsql().WithTx(tx).
					Table("goadmin_role_users").
					WhereIn("role_id", ids).
					Delete()

				if deleteRoleUserErr != nil && notNoAffectRow(deleteRoleUserErr) {
					return deleteRoleUserErr, map[string]interface{}{}
				}

				deleteRoleMenuErr := dbsql().WithTx(tx).
					Table("goadmin_role_menu").
					WhereIn("role_id", ids).
					Delete()

				if deleteRoleMenuErr != nil && notNoAffectRow(deleteRoleMenuErr) {
					return deleteRoleMenuErr, map[string]interface{}{}
				}

				deleteRolePermissionErr := dbsql().WithTx(tx).
					Table("goadmin_role_permissions").
					WhereIn("role_id", ids).
					Delete()

				if deleteRolePermissionErr != nil && notNoAffectRow(deleteRolePermissionErr) {
					return deleteRolePermissionErr, map[string]interface{}{}
				}

				deleteRolesErr := dbsql().WithTx(tx).
					Table("goadmin_roles").
					WhereIn("id", ids).
					Delete()

				if deleteRolesErr != nil {
					return deleteRolesErr, map[string]interface{}{}
				}

				return nil, map[string]interface{}{}
			})

			return txErr
		})

	formList := RolesTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
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

	formList.SetUpdateFn(func(values form2.Values) error {

		role := models.RoleWithId(values.Get("id"))

		role.Update(values.Get("name"), values.Get("slug"))

		role.DeletePermissions()
		for i := 0; i < len(values["permission_id[]"]); i++ {
			role.AddPermission(values["permission_id[]"][i])
		}

		return nil
	})

	formList.SetInsertFn(func(values form2.Values) error {

		role := models.Role().New(values.Get("name"), values.Get("slug"))

		for i := 0; i < len(values["permission_id[]"]); i++ {
			role.AddPermission(values["permission_id[]"][i])
		}

		return nil
	})

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

	info := OpTable.GetInfo().AddXssJsFilter()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("userID"), "user_id", db.Int).FieldFilterable()
	info.AddField(lg("path"), "path", db.Varchar).FieldFilterable()
	info.AddField(lg("method"), "method", db.Varchar).FieldFilterable()
	info.AddField(lg("ip"), "ip", db.Varchar).FieldFilterable()
	info.AddField(lg("content"), "input", db.Varchar)
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_operation_log").
		SetTitle(lg("operation log")).
		SetDescription(lg("operation log"))

	formList := OpTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
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

	info := MenuTable.GetInfo().AddXssJsFilter()

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
		SetDescription(lg("Menus Manage")).
		SetDeleteFn(func(idArr []string) error {

			var ids = interfaces(idArr)

			_, txErr := dbsql().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteRoleMenuErr := dbsql().WithTx(tx).
					Table("goadmin_role_menu").
					WhereIn("menu_id", ids).
					Delete()

				if deleteRoleMenuErr != nil && notNoAffectRow(deleteRoleMenuErr) {
					return deleteRoleMenuErr, map[string]interface{}{}
				}

				deleteMenusErr := dbsql().WithTx(tx).
					Table("goadmin_menu").
					WhereIn("id", ids).
					Delete()

				if deleteMenusErr != nil {
					return deleteMenusErr, map[string]interface{}{}
				}

				return nil, map[string]interface{}{}
			})

			return txErr
		})

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

	formList := MenuTable.GetForm().AddXssJsFilter()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
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

func dbsql() *db.SQL {
	return db.WithDriver(config.Get().Databases.GetDefault().Driver)
}

func interfaces(arr []string) []interface{} {
	var iarr = make([]interface{}, len(arr))

	for key, v := range arr {
		iarr[key] = v
	}

	return iarr
}

func notNoAffectRow(s error) bool {
	return s.Error() != "no affect row"
}
