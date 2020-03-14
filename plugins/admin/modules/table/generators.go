package table

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/collection"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/GoAdminGroup/html"
	"golang.org/x/crypto/bcrypt"
	tmpl "html/template"
	"strconv"
	"strings"
	"time"
)

type SystemTable struct {
	conn db.Connection
}

func NewSystemTable(conn db.Connection) *SystemTable {
	return &SystemTable{conn: conn}
}

func (s *SystemTable) GetManagerTable(ctx *context.Context) (ManagerTable Table) {
	ManagerTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := ManagerTable.GetInfo().AddXssJsFilter().HideFilterArea()

	labelModels, _ := s.table("goadmin_role_users").
		Select("goadmin_roles.name", "user_id").
		LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
		All()
	labelCollection := collection.Collection(labelModels)

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("Name"), "username", db.Varchar).FieldFilterable()
	info.AddField(lg("Nickname"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("role"), "roles", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			uid, _ := strconv.Atoi(model.ID)
			labelCol := labelCollection.Where("user_id", int64(uid))

			labels := template.HTML("")
			labelTpl := label().SetType("success")

			for key, label := range labelCol {
				if key == len(labelCol)-1 {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent()
				} else {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent() + "<br><br>"
				}
			}

			if labels == template.HTML("") {
				return lg("no roles")
			}

			return labels
		})
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_users").
		SetTitle(lg("Managers")).
		SetDescription(lg("Managers")).
		SetDeleteFn(func(idArr []string) error {

			var ids = interfaces(idArr)

			_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteUserRoleErr := s.connection().WithTx(tx).
					Table("goadmin_role_users").
					WhereIn("user_id", ids).
					Delete()

				if deleteUserRoleErr != nil && notNoAffectRow(deleteUserRoleErr) {
					return deleteUserRoleErr, map[string]interface{}{}
				}

				deleteUserPermissionErr := s.connection().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("user_id", ids).
					Delete()

				if deleteUserPermissionErr != nil && notNoAffectRow(deleteUserPermissionErr) {
					return deleteUserPermissionErr, map[string]interface{}{}
				}

				deleteUserErr := s.connection().WithTx(tx).
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

	formList := ManagerTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField(lg("Name"), "username", db.Varchar, form.Text).
		FieldHelpMsg(template.HTML(lg("use for login"))).FieldMust()
	formList.AddField(lg("Nickname"), "name", db.Varchar, form.Text).
		FieldHelpMsg(template.HTML(lg("use to display"))).FieldMust()
	formList.AddField(lg("Avatar"), "avatar", db.Varchar, form.File)
	formList.AddField(lg("role"), "role_id", db.Varchar, form.Select).
		FieldOptionsFromTable("goadmin_roles", "slug", "id").
		FieldDisplay(func(model types.FieldModel) interface{} {
			var roles []string

			if model.ID == "" {
				return roles
			}
			roleModel, _ := s.table("goadmin_role_users").Select("role_id").
				Where("user_id", "=", model.ID).All()
			for _, v := range roleModel {
				roles = append(roles, strconv.FormatInt(v["role_id"].(int64), 10))
			}
			return roles
		}).FieldHelpMsg(template.HTML(lg("no corresponding options?")) +
		link("/admin/info/roles/new", "Create here."))

	formList.AddField(lg("permission"), "permission_id", db.Varchar, form.Select).
		FieldOptionsFromTable("goadmin_permissions", "slug", "id").
		FieldDisplay(func(model types.FieldModel) interface{} {
			var permissions []string

			if model.ID == "" {
				return permissions
			}
			permissionModel, _ := s.table("goadmin_user_permissions").
				Select("permission_id").Where("user_id", "=", model.ID).All()
			for _, v := range permissionModel {
				permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
			}
			return permissions
		}).FieldHelpMsg(template.HTML(lg("no corresponding options?")) +
		link("/admin/info/permission/new", "Create here."))

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

		user := models.UserWithId(values.Get("id")).SetConn(s.conn)

		password := values.Get("password")

		if password != "" {

			if password != values.Get("password_again") {
				return errors.New("password does not match")
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

		password := values.Get("password")

		if password != values.Get("password_again") {
			return errors.New("password does not match")
		}

		user := models.User().SetConn(s.conn).New(values.Get("username"),
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

	detail := ManagerTable.GetDetail()
	detail.AddField("ID", "id", db.Int)
	detail.AddField(lg("Name"), "username", db.Varchar)
	detail.AddField(lg("Avatar"), "avatar", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			if model.Value == "" || config.Get().Store.Prefix == "" {
				model.Value = config.Get().Url("/assets/dist/img/avatar04.png")
			} else {
				model.Value = config.Get().Store.URL(model.Value)
			}
			return template.Default().Image().
				SetSrc(template.HTML(model.Value)).
				SetHeight("120").SetWidth("120").WithModal().GetContent()
		})
	detail.AddField(lg("Nickname"), "name", db.Varchar)
	detail.AddField(lg("role"), "roles", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labelModels, _ := s.table("goadmin_role_users").
				Select("goadmin_roles.name").
				LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
				Where("user_id", "=", model.ID).
				All()

			labels := template.HTML("")
			labelTpl := label().SetType("success")

			for key, label := range labelModels {
				if key == len(labelModels)-1 {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent()
				} else {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent() + "<br><br>"
				}
			}

			if labels == template.HTML("") {
				return lg("no roles")
			}

			return labels
		})
	detail.AddField(lg("permission"), "roles", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			permissionModel, _ := s.table("goadmin_user_permissions").
				Select("goadmin_permissions.name").
				LeftJoin("goadmin_permissions", "goadmin_permissions.id", "=", "goadmin_user_permissions.permission_id").
				Where("user_id", "=", model.ID).
				All()

			permissions := template.HTML("")
			permissionTpl := label().SetType("success")

			for key, label := range permissionModel {
				if key == len(permissionModel)-1 {
					permissions += permissionTpl.SetContent(template.HTML(label["name"].(string))).GetContent()
				} else {
					permissions += permissionTpl.SetContent(template.HTML(label["name"].(string))).GetContent() + "<br><br>"
				}
			}

			return permissions
		})
	detail.AddField(lg("createdAt"), "created_at", db.Timestamp)
	detail.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	return
}

func (s *SystemTable) GetNormalManagerTable(ctx *context.Context) (ManagerTable Table) {
	ManagerTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := ManagerTable.GetInfo().AddXssJsFilter().HideFilterArea()

	labelModels, _ := s.table("goadmin_role_users").
		Select("goadmin_roles.name", "user_id").
		LeftJoin("goadmin_roles", "goadmin_roles.id", "=", "goadmin_role_users.role_id").
		All()
	labelCollection := collection.Collection(labelModels)

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("Name"), "username", db.Varchar).FieldFilterable()
	info.AddField(lg("Nickname"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("role"), "roles", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labelCol := labelCollection.Where("user_id", model.ID)

			labels := template.HTML("")
			labelTpl := label().SetType("success")

			for key, label := range labelCol {
				if key == len(labelModels)-1 {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent()
				} else {
					labels += labelTpl.SetContent(template.HTML(label["name"].(string))).GetContent() + "<br><br>"
				}
			}

			if labels == template.HTML("") {
				return lg("no roles")
			}

			return labels
		})
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)
	info.AddField(lg("updatedAt"), "updated_at", db.Timestamp)

	info.SetTable("goadmin_users").
		SetTitle(lg("Managers")).
		SetDescription(lg("Managers")).
		SetDeleteFn(func(idArr []string) error {

			var ids = interfaces(idArr)

			_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteUserRoleErr := s.connection().WithTx(tx).
					Table("goadmin_role_users").
					WhereIn("user_id", ids).
					Delete()

				if deleteUserRoleErr != nil && notNoAffectRow(deleteUserRoleErr) {
					return deleteUserRoleErr, map[string]interface{}{}
				}

				deleteUserPermissionErr := s.connection().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("user_id", ids).
					Delete()

				if deleteUserPermissionErr != nil && notNoAffectRow(deleteUserPermissionErr) {
					return deleteUserPermissionErr, map[string]interface{}{}
				}

				deleteUserErr := s.connection().WithTx(tx).
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

	formList := ManagerTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField(lg("Name"), "username", db.Varchar, form.Text).FieldHelpMsg(template.HTML(lg("use for login"))).FieldMust()
	formList.AddField(lg("Nickname"), "name", db.Varchar, form.Text).FieldHelpMsg(template.HTML(lg("use to display"))).FieldMust()
	formList.AddField(lg("Avatar"), "avatar", db.Varchar, form.File)
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

		user := models.UserWithId(values.Get("id")).SetConn(s.conn)

		if values.Has("permission", "role") {
			return errors.New("no permission")
		}

		password := values.Get("password")

		if password != "" {

			if password != values.Get("password_again") {
				return errors.New("password does not match")
			}

			password = encodePassword([]byte(values.Get("password")))
		}

		user.Update(values.Get("username"), password, values.Get("name"), values.Get("avatar"))

		return nil
	})
	formList.SetInsertFn(func(values form2.Values) error {
		if values.IsEmpty("name", "username", "password") {
			return errors.New("username and password can not be empty")
		}

		password := values.Get("password")

		if password != values.Get("password_again") {
			return errors.New("password does not match")
		}

		if values.Has("permission", "role") {
			return errors.New("no permission")
		}

		models.User().SetConn(s.conn).New(values.Get("username"),
			encodePassword([]byte(values.Get("password"))),
			values.Get("name"),
			values.Get("avatar"))

		return nil
	})

	return
}

func (s *SystemTable) GetPermissionTable(ctx *context.Context) (PermissionTable Table) {
	PermissionTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := PermissionTable.GetInfo().AddXssJsFilter().HideFilterArea()

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
					res += string(label().SetContent(template.HTML(pathArr[i])).GetContent())
				} else {
					res += string(label().SetContent(template.HTML(pathArr[i])).GetContent()) + "<br><br>"
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

			_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteRolePermissionErr := s.connection().WithTx(tx).
					Table("goadmin_role_permissions").
					WhereIn("permission_id", ids).
					Delete()

				if deleteRolePermissionErr != nil && notNoAffectRow(deleteRolePermissionErr) {
					return deleteRolePermissionErr, map[string]interface{}{}
				}

				deleteUserPermissionErr := s.connection().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("permission_id", ids).
					Delete()

				if deleteUserPermissionErr != nil && notNoAffectRow(deleteUserPermissionErr) {
					return deleteUserPermissionErr, map[string]interface{}{}
				}

				deletePermissionsErr := s.connection().WithTx(tx).
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
	formList.AddField(lg("permission"), "name", db.Varchar, form.Text).FieldMust()
	formList.AddField(lg("slug"), "slug", db.Varchar, form.Text).FieldHelpMsg(template.HTML(lg("should be unique"))).FieldMust()
	formList.AddField(lg("method"), "http_method", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Value: "GET", Text: "GET"},
			{Value: "PUT", Text: "PUT"},
			{Value: "POST", Text: "POST"},
			{Value: "DELETE", Text: "DELETE"},
			{Value: "PATCH", Text: "PATCH"},
			{Value: "OPTIONS", Text: "OPTIONS"},
			{Value: "HEAD", Text: "HEAD"},
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			return strings.Split(model.Value, ",")
		}).
		FieldPostFilterFn(func(model types.PostFieldModel) interface{} {
			return strings.Join(model.Value, ",")
		}).
		FieldHelpMsg(template.HTML(lg("all method if empty")))

	formList.AddField(lg("path"), "http_path", db.Varchar, form.TextArea).FieldHelpMsg(template.HTML(lg("a path a line")))
	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_permissions").
		SetTitle(lg("Permission Manage")).
		SetDescription(lg("Permission Manage")).
		SetPostValidator(func(values form2.Values) error {

			if values.IsEmpty("slug", "http_path", "name") {
				return errors.New("slug or http_path or name should not be empty")
			}

			if models.Permission().SetConn(s.conn).IsSlugExist(values.Get("slug"), values.Get("id")) {
				return errors.New("slug exists")
			}
			return nil
		}).SetPostHook(func(values form2.Values) error {
		_, err := s.connection().Table("goadmin_permissions").
			Where("id", "=", values.Get("id")).Update(dialect.H{
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
		return err
	})

	return
}

func (s *SystemTable) GetRolesTable(ctx *context.Context) (RolesTable Table) {
	RolesTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := RolesTable.GetInfo().AddXssJsFilter().HideFilterArea()

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

			_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteRoleUserErr := s.connection().WithTx(tx).
					Table("goadmin_role_users").
					WhereIn("role_id", ids).
					Delete()

				if deleteRoleUserErr != nil && notNoAffectRow(deleteRoleUserErr) {
					return deleteRoleUserErr, map[string]interface{}{}
				}

				deleteRoleMenuErr := s.connection().WithTx(tx).
					Table("goadmin_role_menu").
					WhereIn("role_id", ids).
					Delete()

				if deleteRoleMenuErr != nil && notNoAffectRow(deleteRoleMenuErr) {
					return deleteRoleMenuErr, map[string]interface{}{}
				}

				deleteRolePermissionErr := s.connection().WithTx(tx).
					Table("goadmin_role_permissions").
					WhereIn("role_id", ids).
					Delete()

				if deleteRolePermissionErr != nil && notNoAffectRow(deleteRolePermissionErr) {
					return deleteRolePermissionErr, map[string]interface{}{}
				}

				deleteRolesErr := s.connection().WithTx(tx).
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
	formList.AddField(lg("role"), "name", db.Varchar, form.Text).FieldMust()
	formList.AddField(lg("slug"), "slug", db.Varchar, form.Text).FieldHelpMsg(template.HTML(lg("should be unique"))).FieldMust()
	formList.AddField(lg("permission"), "permission_id", db.Varchar, form.SelectBox).
		FieldOptionsFromTable("goadmin_permissions", "name", "id").
		FieldDisplay(func(model types.FieldModel) interface{} {
			var permissions = make([]string, 0)

			if model.ID == "" {
				return permissions
			}
			perModel, _ := s.table("goadmin_role_permissions").
				Select("permission_id").
				Where("role_id", "=", model.ID).
				All()
			for _, v := range perModel {
				permissions = append(permissions, strconv.FormatInt(v["permission_id"].(int64), 10))
			}
			return permissions
		}).FieldHelpMsg(template.HTML(lg("no corresponding options?")) +
		link("/admin/info/permission/new", "Create here."))

	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	formList.SetTable("goadmin_roles").
		SetTitle(lg("Roles Manage")).
		SetDescription(lg("Roles Manage"))

	formList.SetUpdateFn(func(values form2.Values) error {

		if models.Role().SetConn(s.conn).IsSlugExist(values.Get("slug"), values.Get("id")) {
			return errors.New("slug exists")
		}

		role := models.RoleWithId(values.Get("id")).SetConn(s.conn)

		role.Update(values.Get("name"), values.Get("slug"))

		role.DeletePermissions()
		for i := 0; i < len(values["permission_id[]"]); i++ {
			role.AddPermission(values["permission_id[]"][i])
		}

		return nil
	})

	formList.SetInsertFn(func(values form2.Values) error {

		if models.Role().SetConn(s.conn).IsSlugExist(values.Get("slug"), "") {
			return errors.New("slug exists")
		}

		role := models.Role().SetConn(s.conn).New(values.Get("name"), values.Get("slug"))

		for i := 0; i < len(values["permission_id[]"]); i++ {
			role.AddPermission(values["permission_id[]"][i])
		}

		return nil
	})

	return
}

func (s *SystemTable) GetOpTable(ctx *context.Context) (OpTable Table) {
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

	info := OpTable.GetInfo().AddXssJsFilter().HideFilterArea()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("userID"), "user_id", db.Int).FieldFilterable()
	info.AddField(lg("path"), "path", db.Varchar).FieldFilterable()
	info.AddField(lg("method"), "method", db.Varchar).FieldFilterable()
	info.AddField(lg("ip"), "ip", db.Varchar).FieldFilterable()
	info.AddField(lg("content"), "input", db.Varchar).FieldWidth(230)
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)

	users, _ := s.table("goadmin_users").Select("id", "name").All()
	options := make(types.FieldOptions, len(users))
	for k, user := range users {
		options[k].Value = fmt.Sprintf("%v", user["id"])
		options[k].Text = fmt.Sprintf("%v", user["name"])
	}
	info.AddSelectBox(language.Get("user"), options, action.FieldFilter("user_id"))
	info.AddSelectBox(language.Get("method"), types.FieldOptions{
		{Value: "GET", Text: "GET"},
		{Value: "POST", Text: "POST"},
		{Value: "OPTIONS", Text: "OPTIONS"},
		{Value: "PUT", Text: "PUT"},
		{Value: "HEAD", Text: "HEAD"},
		{Value: "DELETE", Text: "DELETE"},
	}, action.FieldFilter("method"))

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

func (s *SystemTable) GetMenuTable(ctx *context.Context) (MenuTable Table) {
	MenuTable = NewDefaultTable(DefaultConfigWithDriver(config.Get().Databases.GetDefault().Driver))

	info := MenuTable.GetInfo().AddXssJsFilter().HideFilterArea()

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

			_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

				deleteRoleMenuErr := s.connection().WithTx(tx).
					Table("goadmin_role_menu").
					WhereIn("menu_id", ids).
					Delete()

				if deleteRoleMenuErr != nil && notNoAffectRow(deleteRoleMenuErr) {
					return deleteRoleMenuErr, map[string]interface{}{}
				}

				deleteMenusErr := s.connection().WithTx(tx).
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

	formList := MenuTable.GetForm().AddXssJsFilter()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit().FieldNotAllowAdd()
	formList.AddField(lg("parent"), "parent_id", db.Int, form.SelectSingle).
		FieldOptionsFromTable("goadmin_menu", "title", "id", func(sql *db.SQL) *db.SQL {
			return sql.Where("parent_id", "=", 0).OrderBy("order", "asc")
		}).
		FieldOptionsTableProcessFn(func(options types.FieldOptions) types.FieldOptions {
			return append([]types.FieldOption{{
				Text:  "root",
				Value: "0",
			}}, options...)
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			var menuItem []string

			if model.ID == "" {
				return menuItem
			}

			menuModel, _ := s.table("goadmin_menu").Select("parent_id").Find(model.ID)
			menuItem = append(menuItem, strconv.FormatInt(menuModel["parent_id"].(int64), 10))
			return menuItem
		})
	formList.AddField(lg("menu name"), "title", db.Varchar, form.Text).FieldMust()
	formList.AddField(lg("header"), "header", db.Varchar, form.Text)
	formList.AddField(lg("icon"), "icon", db.Varchar, form.IconPicker)
	formList.AddField(lg("uri"), "uri", db.Varchar, form.Text)
	formList.AddField(lg("role"), "roles", db.Int, form.Select).
		FieldOptionsFromTable("goadmin_roles", "slug", "id").
		FieldDisplay(func(model types.FieldModel) interface{} {
			var roles []string

			if model.ID == "" {
				return roles
			}

			roleModel, _ := s.table("goadmin_role_menu").
				Select("role_id").
				Where("menu_id", "=", model.ID).
				All()

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

// -------------------------
// helper functions
// -------------------------

func encodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash[:])
}

func label() types.LabelAttribute {
	return template.Get(config.Get().Theme).Label().SetType("success")
}

func lg(v string) string {
	return language.Get(v)
}

func link(url, content string) tmpl.HTML {
	return html.AEl().
		SetAttr("href", url).
		SetContent(template.HTML(lg(content))).
		Get()
}

func (s *SystemTable) table(table string) *db.SQL {
	return s.connection().Table(table)
}

func (s *SystemTable) connection() *db.SQL {
	return db.WithDriver(s.conn)
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
