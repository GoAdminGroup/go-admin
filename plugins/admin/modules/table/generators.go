package table

import (
	"database/sql"
	"errors"
	"fmt"
	tmpl "html/template"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/collection"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	errs "github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/ui"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/tools"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/action"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	selection "github.com/GoAdminGroup/go-admin/template/types/form/select"
	"github.com/GoAdminGroup/html"
	"golang.org/x/crypto/bcrypt"
)

type SystemTable struct {
	conn db.Connection
	c    *config.Config
}

func NewSystemTable(conn db.Connection, c *config.Config) *SystemTable {
	return &SystemTable{conn: conn, c: c}
}

func (s *SystemTable) GetManagerTable(ctx *context.Context) (managerTable Table) {
	managerTable = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver))

	info := managerTable.GetInfo().AddXssJsFilter().HideFilterArea()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("Name"), "username", db.Varchar).FieldFilterable()
	info.AddField(lg("Nickname"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("role"), "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "goadmin_role_users",
			JoinField: "user_id",
			Field:     "id",
		}).
		FieldJoin(types.Join{
			Table:     "goadmin_roles",
			JoinField: "id",
			Field:     "role_id",
			BaseTable: "goadmin_role_users",
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labels := template.HTML("")
			labelTpl := label().SetType("success")

			labelValues := strings.Split(model.Value, types.JoinFieldValueDelimiter)
			for key, label := range labelValues {
				if key == len(labelValues)-1 {
					labels += labelTpl.SetContent(template.HTML(label)).GetContent()
				} else {
					labels += labelTpl.SetContent(template.HTML(label)).GetContent() + "<br><br>"
				}
			}

			if labels == template.HTML("") {
				return lg("no roles")
			}

			return labels
		}).FieldFilterable()
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

				if db.CheckError(deleteUserRoleErr, db.DELETE) {
					return deleteUserRoleErr, nil
				}

				deleteUserPermissionErr := s.connection().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("user_id", ids).
					Delete()

				if db.CheckError(deleteUserPermissionErr, db.DELETE) {
					return deleteUserPermissionErr, nil
				}

				deleteUserErr := s.connection().WithTx(tx).
					Table("goadmin_users").
					WhereIn("id", ids).
					Delete()

				if db.CheckError(deleteUserErr, db.DELETE) {
					return deleteUserErr, nil
				}

				return nil, nil
			})

			return txErr
		})

	formList := managerTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
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

		_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

			avatar := values.Get("avatar")

			if values.Get("avatar__delete_flag") == "1" {
				avatar = ""
			}

			_, updateUserErr := user.WithTx(tx).Update(values.Get("username"),
				password, values.Get("name"), avatar, values.Get("avatar__change_flag") == "1")

			if db.CheckError(updateUserErr, db.UPDATE) {
				return updateUserErr, nil
			}

			delRoleErr := user.WithTx(tx).DeleteRoles()

			if db.CheckError(delRoleErr, db.DELETE) {
				return delRoleErr, nil
			}

			for i := 0; i < len(values["role_id[]"]); i++ {
				_, addRoleErr := user.WithTx(tx).AddRole(values["role_id[]"][i])
				if db.CheckError(addRoleErr, db.INSERT) {
					return addRoleErr, nil
				}
			}

			delPermissionErr := user.WithTx(tx).DeletePermissions()

			if db.CheckError(delPermissionErr, db.DELETE) {
				return delPermissionErr, nil
			}

			for i := 0; i < len(values["permission_id[]"]); i++ {
				_, addPermissionErr := user.WithTx(tx).AddPermission(values["permission_id[]"][i])
				if db.CheckError(addPermissionErr, db.INSERT) {
					return addPermissionErr, nil
				}
			}

			return nil, nil
		})

		return txErr
	})
	formList.SetInsertFn(func(values form2.Values) error {
		if values.IsEmpty("name", "username", "password") {
			return errors.New("username and password can not be empty")
		}

		password := values.Get("password")

		if password != values.Get("password_again") {
			return errors.New("password does not match")
		}

		_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

			user, createUserErr := models.User().WithTx(tx).SetConn(s.conn).New(values.Get("username"),
				encodePassword([]byte(values.Get("password"))),
				values.Get("name"),
				values.Get("avatar"))

			if db.CheckError(createUserErr, db.INSERT) {
				return createUserErr, nil
			}

			for i := 0; i < len(values["role_id[]"]); i++ {
				_, addRoleErr := user.WithTx(tx).AddRole(values["role_id[]"][i])
				if db.CheckError(addRoleErr, db.INSERT) {
					return addRoleErr, nil
				}
			}

			for i := 0; i < len(values["permission_id[]"]); i++ {
				_, addPermissionErr := user.WithTx(tx).AddPermission(values["permission_id[]"][i])
				if db.CheckError(addPermissionErr, db.INSERT) {
					return addPermissionErr, nil
				}
			}

			return nil, nil
		})
		return txErr
	})

	detail := managerTable.GetDetail()
	detail.AddField("ID", "id", db.Int)
	detail.AddField(lg("Name"), "username", db.Varchar)
	detail.AddField(lg("Avatar"), "avatar", db.Varchar).
		FieldDisplay(func(model types.FieldModel) interface{} {
			if model.Value == "" || config.GetStore().Prefix == "" {
				model.Value = config.Url("/assets/dist/img/avatar04.png")
			} else {
				model.Value = config.GetStore().URL(model.Value)
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

func (s *SystemTable) GetNormalManagerTable(ctx *context.Context) (managerTable Table) {
	managerTable = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver))

	info := managerTable.GetInfo().AddXssJsFilter().HideFilterArea()

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField(lg("Name"), "username", db.Varchar).FieldFilterable()
	info.AddField(lg("Nickname"), "name", db.Varchar).FieldFilterable()
	info.AddField(lg("role"), "name", db.Varchar).
		FieldJoin(types.Join{
			Table:     "goadmin_role_users",
			JoinField: "user_id",
			Field:     "id",
		}).
		FieldJoin(types.Join{
			Table:     "goadmin_roles",
			JoinField: "id",
			Field:     "role_id",
			BaseTable: "goadmin_role_users",
		}).
		FieldDisplay(func(model types.FieldModel) interface{} {
			labels := template.HTML("")
			labelTpl := label().SetType("success")

			labelValues := strings.Split(model.Value, types.JoinFieldValueDelimiter)
			for key, label := range labelValues {
				if key == len(labelValues)-1 {
					labels += labelTpl.SetContent(template.HTML(label)).GetContent()
				} else {
					labels += labelTpl.SetContent(template.HTML(label)).GetContent() + "<br><br>"
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

				if db.CheckError(deleteUserRoleErr, db.DELETE) {
					return deleteUserRoleErr, nil
				}

				deleteUserPermissionErr := s.connection().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("user_id", ids).
					Delete()

				if db.CheckError(deleteUserPermissionErr, db.DELETE) {
					return deleteUserPermissionErr, nil
				}

				deleteUserErr := s.connection().WithTx(tx).
					Table("goadmin_users").
					WhereIn("id", ids).
					Delete()

				if db.CheckError(deleteUserErr, db.DELETE) {
					return deleteUserErr, nil
				}

				return nil, nil
			})

			return txErr
		})

	formList := managerTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
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
			return errors.New(errs.NoPermission)
		}

		password := values.Get("password")

		if password != "" {

			if password != values.Get("password_again") {
				return errors.New("password does not match")
			}

			password = encodePassword([]byte(values.Get("password")))
		}

		avatar := values.Get("avatar")

		if values.Get("avatar__delete_flag") == "1" {
			avatar = ""
		}

		_, updateUserErr := user.Update(values.Get("username"),
			password, values.Get("name"), avatar, values.Get("avatar__change_flag") == "1")

		if db.CheckError(updateUserErr, db.UPDATE) {
			return updateUserErr
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

		if values.Has("permission", "role") {
			return errors.New(errs.NoPermission)
		}

		_, createUserErr := models.User().SetConn(s.conn).New(values.Get("username"),
			encodePassword([]byte(values.Get("password"))),
			values.Get("name"),
			values.Get("avatar"))

		if db.CheckError(createUserErr, db.INSERT) {
			return createUserErr
		}

		return nil
	})

	return
}

func (s *SystemTable) GetPermissionTable(ctx *context.Context) (permissionTable Table) {
	permissionTable = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver))

	info := permissionTable.GetInfo().AddXssJsFilter().HideFilterArea()

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

				if db.CheckError(deleteRolePermissionErr, db.DELETE) {
					return deleteRolePermissionErr, nil
				}

				deleteUserPermissionErr := s.connection().WithTx(tx).
					Table("goadmin_user_permissions").
					WhereIn("permission_id", ids).
					Delete()

				if db.CheckError(deleteUserPermissionErr, db.DELETE) {
					return deleteUserPermissionErr, nil
				}

				deletePermissionsErr := s.connection().WithTx(tx).
					Table("goadmin_permissions").
					WhereIn("id", ids).
					Delete()

				if deletePermissionsErr != nil {
					return deletePermissionsErr, nil
				}

				return nil, nil
			})

			return txErr
		})

	formList := permissionTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
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

	formList.AddField(lg("path"), "http_path", db.Varchar, form.TextArea).
		FieldPostFilterFn(func(model types.PostFieldModel) interface{} {
			return strings.TrimSpace(model.Value.Value())
		}).
		FieldHelpMsg(template.HTML(lg("a path a line, without global prefix")))
	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldDisableWhenCreate()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldDisableWhenCreate()

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

		if values.IsInsertPost() {
			return nil
		}

		_, err := s.connection().Table("goadmin_permissions").
			Where("id", "=", values.Get("id")).
			Update(dialect.H{
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})

		if db.CheckError(err, db.UPDATE) {
			return err
		}

		return nil
	})

	return
}

func (s *SystemTable) GetRolesTable(ctx *context.Context) (roleTable Table) {
	roleTable = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver))

	info := roleTable.GetInfo().AddXssJsFilter().HideFilterArea()

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

				if db.CheckError(deleteRoleUserErr, db.DELETE) {
					return deleteRoleUserErr, nil
				}

				deleteRoleMenuErr := s.connection().WithTx(tx).
					Table("goadmin_role_menu").
					WhereIn("role_id", ids).
					Delete()

				if db.CheckError(deleteRoleMenuErr, db.DELETE) {
					return deleteRoleMenuErr, nil
				}

				deleteRolePermissionErr := s.connection().WithTx(tx).
					Table("goadmin_role_permissions").
					WhereIn("role_id", ids).
					Delete()

				if db.CheckError(deleteRolePermissionErr, db.DELETE) {
					return deleteRolePermissionErr, nil
				}

				deleteRolesErr := s.connection().WithTx(tx).
					Table("goadmin_roles").
					WhereIn("id", ids).
					Delete()

				if db.CheckError(deleteRolesErr, db.DELETE) {
					return deleteRolesErr, nil
				}

				return nil, nil
			})

			return txErr
		})

	formList := roleTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
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

	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldDisableWhenCreate()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldDisableWhenCreate()

	formList.SetTable("goadmin_roles").
		SetTitle(lg("Roles Manage")).
		SetDescription(lg("Roles Manage"))

	formList.SetUpdateFn(func(values form2.Values) error {

		if models.Role().SetConn(s.conn).IsSlugExist(values.Get("slug"), values.Get("id")) {
			return errors.New("slug exists")
		}

		role := models.RoleWithId(values.Get("id")).SetConn(s.conn)

		_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {

			_, updateRoleErr := role.WithTx(tx).Update(values.Get("name"), values.Get("slug"))

			if db.CheckError(updateRoleErr, db.UPDATE) {
				return updateRoleErr, nil
			}

			delPermissionErr := role.WithTx(tx).DeletePermissions()

			if db.CheckError(delPermissionErr, db.DELETE) {
				return delPermissionErr, nil
			}

			for i := 0; i < len(values["permission_id[]"]); i++ {
				_, addPermissionErr := role.WithTx(tx).AddPermission(values["permission_id[]"][i])
				if db.CheckError(addPermissionErr, db.INSERT) {
					return addPermissionErr, nil
				}
			}

			return nil, nil
		})

		return txErr
	})

	formList.SetInsertFn(func(values form2.Values) error {

		if models.Role().SetConn(s.conn).IsSlugExist(values.Get("slug"), "") {
			return errors.New("slug exists")
		}

		_, txErr := s.connection().WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {
			role, createRoleErr := models.Role().WithTx(tx).SetConn(s.conn).New(values.Get("name"), values.Get("slug"))

			if db.CheckError(createRoleErr, db.INSERT) {
				return createRoleErr, nil
			}

			for i := 0; i < len(values["permission_id[]"]); i++ {
				_, addPermissionErr := role.WithTx(tx).AddPermission(values["permission_id[]"][i])
				if db.CheckError(addPermissionErr, db.INSERT) {
					return addPermissionErr, nil
				}
			}

			return nil, nil
		})

		return txErr
	})

	return
}

func (s *SystemTable) GetOpTable(ctx *context.Context) (opTable Table) {
	opTable = NewDefaultTable(Config{
		Driver:     config.GetDatabases().GetDefault().Driver,
		CanAdd:     false,
		Editable:   false,
		Deletable:  config.GetAllowDelOperationLog(),
		Exportable: true,
		Connection: "default",
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	})

	info := opTable.GetInfo().AddXssJsFilter().
		HideFilterArea().HideDetailButton().HideEditButton().HideNewButton()

	if !config.GetAllowDelOperationLog() {
		info = info.HideDeleteButton()
	}

	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("userID", "user_id", db.Int).FieldHide()
	info.AddField(lg("user"), "name", db.Varchar).FieldJoin(types.Join{
		Table:     config.GetAuthUserTable(),
		JoinField: "id",
		Field:     "user_id",
	}).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().
			Link().
			SetURL(config.Url("/info/manager/detail?__goadmin_detail_pk=") + strconv.Itoa(int(value.Row["user_id"].(int64)))).
			SetContent(template.HTML(value.Value)).
			OpenInNewTab().
			SetTabTitle("Manager Detail").
			GetContent()
	}).FieldFilterable()
	info.AddField(lg("path"), "path", db.Varchar).FieldFilterable()
	info.AddField(lg("method"), "method", db.Varchar).FieldFilterable()
	info.AddField(lg("ip"), "ip", db.Varchar).FieldFilterable()
	info.AddField(lg("content"), "input", db.Text).FieldWidth(230)
	info.AddField(lg("createdAt"), "created_at", db.Timestamp)

	users, _ := s.table(config.GetAuthUserTable()).Select("id", "name").All()
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

	formList := opTable.GetForm().AddXssJsFilter()

	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField(lg("userID"), "user_id", db.Int, form.Text)
	formList.AddField(lg("path"), "path", db.Varchar, form.Text)
	formList.AddField(lg("method"), "method", db.Varchar, form.Text)
	formList.AddField(lg("ip"), "ip", db.Varchar, form.Text)
	formList.AddField(lg("content"), "input", db.Varchar, form.Text)
	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldDisableWhenCreate()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldDisableWhenCreate()

	formList.SetTable("goadmin_operation_log").
		SetTitle(lg("operation log")).
		SetDescription(lg("operation log"))

	return
}

func (s *SystemTable) GetMenuTable(ctx *context.Context) (menuTable Table) {
	menuTable = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver))

	name := ctx.Query("__plugin_name")

	info := menuTable.GetInfo().AddXssJsFilter().HideFilterArea().Where("plugin_name", "=", name)

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

				if db.CheckError(deleteRoleMenuErr, db.DELETE) {
					return deleteRoleMenuErr, nil
				}

				deleteMenusErr := s.connection().WithTx(tx).
					Table("goadmin_menu").
					WhereIn("id", ids).
					Delete()

				if db.CheckError(deleteMenusErr, db.DELETE) {
					return deleteMenusErr, nil
				}

				return nil, map[string]interface{}{}
			})

			return txErr
		})

	var parentIDOptions = types.FieldOptions{
		{
			Text:  "ROOT",
			Value: "0",
		},
	}

	allMenus, _ := s.connection().Table("goadmin_menu").
		Where("parent_id", "=", 0).
		Where("plugin_name", "=", name).
		Select("id", "title").
		OrderBy("order", "asc").
		All()
	allMenuIDs := make([]interface{}, len(allMenus))

	if len(allMenuIDs) > 0 {
		for i := 0; i < len(allMenus); i++ {
			allMenuIDs[i] = allMenus[i]["id"]
		}

		secondLevelMenus, _ := s.connection().Table("goadmin_menu").
			WhereIn("parent_id", allMenuIDs).
			Where("plugin_name", "=", name).
			Select("id", "title", "parent_id").
			All()

		secondLevelMenusCol := collection.Collection(secondLevelMenus)

		for i := 0; i < len(allMenus); i++ {
			parentIDOptions = append(parentIDOptions, types.FieldOption{
				TextHTML: "&nbsp;&nbsp;┝  " + language.GetFromHtml(template.HTML(allMenus[i]["title"].(string))),
				Value:    strconv.Itoa(int(allMenus[i]["id"].(int64))),
			})
			col := secondLevelMenusCol.Where("parent_id", "=", allMenus[i]["id"].(int64))
			for i := 0; i < len(col); i++ {
				parentIDOptions = append(parentIDOptions, types.FieldOption{
					TextHTML: "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;┝  " +
						language.GetFromHtml(template.HTML(col[i]["title"].(string))),
					Value: strconv.Itoa(int(col[i]["id"].(int64))),
				})
			}
		}
	}

	formList := menuTable.GetForm().AddXssJsFilter()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField(lg("parent"), "parent_id", db.Int, form.SelectSingle).
		FieldOptions(parentIDOptions).
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
	formList.AddField("PluginName", "plugin_name", db.Varchar, form.Text).FieldDefault(name).FieldHide()
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

	formList.AddField(lg("updatedAt"), "updated_at", db.Timestamp, form.Default).FieldDisableWhenCreate()
	formList.AddField(lg("createdAt"), "created_at", db.Timestamp, form.Default).FieldDisableWhenCreate()

	formList.SetTable("goadmin_menu").
		SetTitle(lg("Menus Manage")).
		SetDescription(lg("Menus Manage"))

	return
}

func (s *SystemTable) GetSiteTable(ctx *context.Context) (siteTable Table) {
	siteTable = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver).
		SetOnlyUpdateForm().
		SetGetDataFun(func(params parameter.Parameters) (i []map[string]interface{}, i2 int) {
			return []map[string]interface{}{models.Site().SetConn(s.conn).AllToMapInterface()}, 1
		}))

	trueStr := lgWithConfigScore("true")
	falseStr := lgWithConfigScore("false")

	formList := siteTable.GetForm().AddXssJsFilter()
	formList.AddField("ID", "id", db.Varchar, form.Default).FieldDefault("1").FieldHide()
	formList.AddField(lgWithConfigScore("site off"), "site_off", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("debug"), "debug", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("env"), "env", db.Varchar, form.Default).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return s.c.Env
		})

	langOps := make(types.FieldOptions, len(language.Langs))
	for k, t := range language.Langs {
		langOps[k] = types.FieldOption{Text: lgWithConfigScore(t, "language"), Value: t}
	}
	formList.AddField(lgWithConfigScore("language"), "language", db.Varchar, form.SelectSingle).
		FieldDisplay(func(value types.FieldModel) interface{} {
			return language.FixedLanguageKey(value.Value)
		}).
		FieldOptions(langOps)
	themes := template.Themes()
	themesOps := make(types.FieldOptions, len(themes))
	for k, t := range themes {
		themesOps[k] = types.FieldOption{Text: t, Value: t}
	}

	formList.AddField(lgWithConfigScore("theme"), "theme", db.Varchar, form.SelectSingle).
		FieldOptions(themesOps).
		FieldOnChooseShow("adminlte",
			"color_scheme")
	formList.AddField(lgWithConfigScore("title"), "title", db.Varchar, form.Text).FieldMust()
	formList.AddField(lgWithConfigScore("color scheme"), "color_scheme", db.Varchar, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "skin-black", Value: "skin-black"},
			{Text: "skin-black-light", Value: "skin-black-light"},
			{Text: "skin-blue", Value: "skin-blue"},
			{Text: "skin-blue-light", Value: "skin-blue-light"},
			{Text: "skin-green", Value: "skin-green"},
			{Text: "skin-green-light", Value: "skin-green-light"},
			{Text: "skin-purple", Value: "skin-purple"},
			{Text: "skin-purple-light", Value: "skin-purple-light"},
			{Text: "skin-red", Value: "skin-red"},
			{Text: "skin-red-light", Value: "skin-red-light"},
			{Text: "skin-yellow", Value: "skin-yellow"},
			{Text: "skin-yellow-light", Value: "skin-yellow-light"},
		}).FieldHelpMsg(template.HTML(lgWithConfigScore("It will work when theme is adminlte")))
	formList.AddField(lgWithConfigScore("login title"), "login_title", db.Varchar, form.Text).FieldMust()
	formList.AddField(lgWithConfigScore("extra"), "extra", db.Varchar, form.TextArea)
	formList.AddField(lgWithConfigScore("logo"), "logo", db.Varchar, form.Code).FieldMust()
	formList.AddField(lgWithConfigScore("mini logo"), "mini_logo", db.Varchar, form.Code).FieldMust()
	if s.c.IsNotProductionEnvironment() {
		formList.AddField(lgWithConfigScore("bootstrap file path"), "bootstrap_file_path", db.Varchar, form.Text)
		formList.AddField(lgWithConfigScore("go mod file path"), "go_mod_file_path", db.Varchar, form.Text)
	}
	formList.AddField(lgWithConfigScore("session life time"), "session_life_time", db.Varchar, form.Number).
		FieldMust().
		FieldHelpMsg(template.HTML(lgWithConfigScore("must bigger than 900 seconds")))
	formList.AddField(lgWithConfigScore("custom head html"), "custom_head_html", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("custom foot Html"), "custom_foot_html", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("custom 404 html"), "custom_404_html", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("custom 403 html"), "custom_403_html", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("custom 500 Html"), "custom_500_html", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("footer info"), "footer_info", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("login logo"), "login_logo", db.Varchar, form.Code)
	formList.AddField(lgWithConfigScore("no limit login ip"), "no_limit_login_ip", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("operation log off"), "operation_log_off", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("allow delete operation log"), "allow_del_operation_log", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("hide config center entrance"), "hide_config_center_entrance", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("hide app info entrance"), "hide_app_info_entrance", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("hide tool entrance"), "hide_tool_entrance", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("hide plugin entrance"), "hide_plugin_entrance", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("animation type"), "animation_type", db.Varchar, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "", Value: ""},
			{Text: "bounce", Value: "bounce"}, {Text: "flash", Value: "flash"}, {Text: "pulse", Value: "pulse"},
			{Text: "rubberBand", Value: "rubberBand"}, {Text: "shake", Value: "shake"}, {Text: "swing", Value: "swing"},
			{Text: "tada", Value: "tada"}, {Text: "wobble", Value: "wobble"}, {Text: "jello", Value: "jello"},
			{Text: "heartBeat", Value: "heartBeat"}, {Text: "bounceIn", Value: "bounceIn"}, {Text: "bounceInDown", Value: "bounceInDown"},
			{Text: "bounceInLeft", Value: "bounceInLeft"}, {Text: "bounceInRight", Value: "bounceInRight"}, {Text: "bounceInUp", Value: "bounceInUp"},
			{Text: "fadeIn", Value: "fadeIn"}, {Text: "fadeInDown", Value: "fadeInDown"}, {Text: "fadeInDownBig", Value: "fadeInDownBig"},
			{Text: "fadeInLeft", Value: "fadeInLeft"}, {Text: "fadeInLeftBig", Value: "fadeInLeftBig"}, {Text: "fadeInRight", Value: "fadeInRight"},
			{Text: "fadeInRightBig", Value: "fadeInRightBig"}, {Text: "fadeInUp", Value: "fadeInUp"}, {Text: "fadeInUpBig", Value: "fadeInUpBig"},
			{Text: "flip", Value: "flip"}, {Text: "flipInX", Value: "flipInX"}, {Text: "flipInY", Value: "flipInY"},
			{Text: "lightSpeedIn", Value: "lightSpeedIn"}, {Text: "rotateIn", Value: "rotateIn"}, {Text: "rotateInDownLeft", Value: "rotateInDownLeft"},
			{Text: "rotateInDownRight", Value: "rotateInDownRight"}, {Text: "rotateInUpLeft", Value: "rotateInUpLeft"}, {Text: "rotateInUpRight", Value: "rotateInUpRight"},
			{Text: "slideInUp", Value: "slideInUp"}, {Text: "slideInDown", Value: "slideInDown"}, {Text: "slideInLeft", Value: "slideInLeft"}, {Text: "slideInRight", Value: "slideInRight"},
			{Text: "slideOutRight", Value: "slideOutRight"}, {Text: "zoomIn", Value: "zoomIn"}, {Text: "zoomInDown", Value: "zoomInDown"},
			{Text: "zoomInLeft", Value: "zoomInLeft"}, {Text: "zoomInRight", Value: "zoomInRight"}, {Text: "zoomInUp", Value: "zoomInUp"},
			{Text: "hinge", Value: "hinge"}, {Text: "jackInTheBox", Value: "jackInTheBox"}, {Text: "rollIn", Value: "rollIn"},
		}).FieldOnChooseHide("", "animation_duration", "animation_delay").
		FieldOptionExt(map[string]interface{}{"allowClear": true}).
		FieldHelpMsg(`see more: <a href="https://daneden.github.io/animate.css/">https://daneden.github.io/animate.css/</a>`)

	formList.AddField(lgWithConfigScore("animation duration"), "animation_duration", db.Varchar, form.Number)
	formList.AddField(lgWithConfigScore("animation delay"), "animation_delay", db.Varchar, form.Number)

	formList.AddField(lgWithConfigScore("file upload engine"), "file_upload_engine", db.Varchar, form.Text)

	formList.AddField(lgWithConfigScore("cdn url"), "asset_url", db.Varchar, form.Text).
		FieldHelpMsg(template.HTML(lgWithConfigScore("Do not modify when you have not set up all assets")))

	formList.AddField(lgWithConfigScore("info log path"), "info_log_path", db.Varchar, form.Text)
	formList.AddField(lgWithConfigScore("error log path"), "error_log_path", db.Varchar, form.Text)
	formList.AddField(lgWithConfigScore("access log path"), "access_log_path", db.Varchar, form.Text)
	formList.AddField(lgWithConfigScore("info log off"), "info_log_off", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("error log off"), "error_log_off", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("access log off"), "access_log_off", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("access assets log off"), "access_assets_log_off", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("sql log on"), "sql_log", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		})
	formList.AddField(lgWithConfigScore("log level"), "logger_level", db.Varchar, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "Debug", Value: "-1"},
			{Text: "Info", Value: "0"},
			{Text: "Warn", Value: "1"},
			{Text: "Error", Value: "2"},
		}).FieldDisplay(defaultFilterFn("0"))

	formList.AddField(lgWithConfigScore("logger rotate max size"), "logger_rotate_max_size", db.Varchar, form.Number).
		FieldDivider(lgWithConfigScore("logger rotate")).FieldDisplay(defaultFilterFn("10", "0"))
	formList.AddField(lgWithConfigScore("logger rotate max backups"), "logger_rotate_max_backups", db.Varchar, form.Number).
		FieldDisplay(defaultFilterFn("5", "0"))
	formList.AddField(lgWithConfigScore("logger rotate max age"), "logger_rotate_max_age", db.Varchar, form.Number).
		FieldDisplay(defaultFilterFn("30", "0"))
	formList.AddField(lgWithConfigScore("logger rotate compress"), "logger_rotate_compress", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: trueStr, Value: "true"},
			{Text: falseStr, Value: "false"},
		}).FieldDisplay(defaultFilterFn("false"))

	formList.AddField(lgWithConfigScore("logger rotate encoder encoding"), "logger_encoder_encoding", db.Varchar,
		form.SelectSingle).
		FieldDivider(lgWithConfigScore("logger rotate encoder")).
		FieldOptions(types.FieldOptions{
			{Text: "JSON", Value: "json"},
			{Text: "Console", Value: "console"},
		}).FieldDisplay(defaultFilterFn("console")).
		FieldOnChooseHide("Console",
			"logger_encoder_time_key", "logger_encoder_level_key", "logger_encoder_caller_key",
			"logger_encoder_message_key", "logger_encoder_stacktrace_key", "logger_encoder_name_key")

	formList.AddField(lgWithConfigScore("logger rotate encoder time key"), "logger_encoder_time_key", db.Varchar, form.Text).
		FieldDisplay(defaultFilterFn("ts"))
	formList.AddField(lgWithConfigScore("logger rotate encoder level key"), "logger_encoder_level_key", db.Varchar, form.Text).
		FieldDisplay(defaultFilterFn("level"))
	formList.AddField(lgWithConfigScore("logger rotate encoder name key"), "logger_encoder_name_key", db.Varchar, form.Text).
		FieldDisplay(defaultFilterFn("logger"))
	formList.AddField(lgWithConfigScore("logger rotate encoder caller key"), "logger_encoder_caller_key", db.Varchar, form.Text).
		FieldDisplay(defaultFilterFn("caller"))
	formList.AddField(lgWithConfigScore("logger rotate encoder message key"), "logger_encoder_message_key", db.Varchar, form.Text).
		FieldDisplay(defaultFilterFn("msg"))
	formList.AddField(lgWithConfigScore("logger rotate encoder stacktrace key"), "logger_encoder_stacktrace_key", db.Varchar, form.Text).
		FieldDisplay(defaultFilterFn("stacktrace"))

	formList.AddField(lgWithConfigScore("logger rotate encoder level"), "logger_encoder_level", db.Varchar,
		form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: lgWithConfigScore("capital"), Value: "capital"},
			{Text: lgWithConfigScore("capitalcolor"), Value: "capitalColor"},
			{Text: lgWithConfigScore("lowercase"), Value: "lowercase"},
			{Text: lgWithConfigScore("lowercasecolor"), Value: "color"},
		}).FieldDisplay(defaultFilterFn("capitalColor"))
	formList.AddField(lgWithConfigScore("logger rotate encoder time"), "logger_encoder_time", db.Varchar,
		form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "ISO8601(2006-01-02T15:04:05.000Z0700)", Value: "iso8601"},
			{Text: lgWithConfigScore("millisecond"), Value: "millis"},
			{Text: lgWithConfigScore("nanosecond"), Value: "nanos"},
			{Text: "RFC3339(2006-01-02T15:04:05Z07:00)", Value: "rfc3339"},
			{Text: "RFC3339 Nano(2006-01-02T15:04:05.999999999Z07:00)", Value: "rfc3339nano"},
		}).FieldDisplay(defaultFilterFn("iso8601"))
	formList.AddField(lgWithConfigScore("logger rotate encoder duration"), "logger_encoder_duration", db.Varchar,
		form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: lgWithConfigScore("seconds"), Value: "string"},
			{Text: lgWithConfigScore("nanosecond"), Value: "nanos"},
			{Text: lgWithConfigScore("microsecond"), Value: "ms"},
		}).FieldDisplay(defaultFilterFn("string"))
	formList.AddField(lgWithConfigScore("logger rotate encoder caller"), "logger_encoder_caller", db.Varchar,
		form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: lgWithConfigScore("full path"), Value: "full"},
			{Text: lgWithConfigScore("short path"), Value: "short"},
		}).FieldDisplay(defaultFilterFn("full"))

	formList.HideBackButton().HideContinueEditCheckBox().HideContinueNewCheckBox()
	formList.SetTabGroups(types.NewTabGroups("id", "debug", "env", "language", "theme", "color_scheme",
		"asset_url", "title", "login_title", "session_life_time", "bootstrap_file_path", "go_mod_file_path", "no_limit_login_ip",
		"operation_log_off", "allow_del_operation_log", "hide_config_center_entrance", "hide_app_info_entrance", "hide_tool_entrance",
		"hide_plugin_entrance", "animation_type",
		"animation_duration", "animation_delay", "file_upload_engine", "extra").
		AddGroup("access_log_off", "access_assets_log_off", "info_log_off", "error_log_off", "sql_log", "logger_level",
			"info_log_path", "error_log_path",
			"access_log_path", "logger_rotate_max_size", "logger_rotate_max_backups",
			"logger_rotate_max_age", "logger_rotate_compress",
			"logger_encoder_encoding", "logger_encoder_time_key", "logger_encoder_level_key", "logger_encoder_name_key",
			"logger_encoder_caller_key", "logger_encoder_message_key", "logger_encoder_stacktrace_key", "logger_encoder_level",
			"logger_encoder_time", "logger_encoder_duration", "logger_encoder_caller").
		AddGroup("logo", "mini_logo", "custom_head_html", "custom_foot_html", "footer_info", "login_logo",
			"custom_404_html", "custom_403_html", "custom_500_html")).
		SetTabHeaders(lgWithConfigScore("general"), lgWithConfigScore("log"), lgWithConfigScore("custom"))

	formList.SetTable("goadmin_site").
		SetTitle(lgWithConfigScore("site setting")).
		SetDescription(lgWithConfigScore("site setting"))

	formList.SetUpdateFn(func(values form2.Values) error {

		ses := values.Get("session_life_time")
		sesInt, _ := strconv.Atoi(ses)
		if sesInt < 900 {
			return errors.New("wrong session life time, must bigger than 900 seconds")
		}
		if err := checkJSON(values, "file_upload_engine"); err != nil {
			return err
		}

		values["logo"][0] = escape(values.Get("logo"))
		values["mini_logo"][0] = escape(values.Get("mini_logo"))
		values["custom_head_html"][0] = escape(values.Get("custom_head_html"))
		values["custom_foot_html"][0] = escape(values.Get("custom_foot_html"))
		values["custom_404_html"][0] = escape(values.Get("custom_404_html"))
		values["custom_403_html"][0] = escape(values.Get("custom_403_html"))
		values["custom_500_html"][0] = escape(values.Get("custom_500_html"))
		values["footer_info"][0] = escape(values.Get("footer_info"))
		values["login_logo"][0] = escape(values.Get("login_logo"))

		var err error
		if s.c.UpdateProcessFn != nil {
			values, err = s.c.UpdateProcessFn(values)
			if err != nil {
				return err
			}
		}

		ui.GetService(services).RemoveOrShowSiteNavButton(values["hide_config_center_entrance"][0] == "true")
		ui.GetService(services).RemoveOrShowInfoNavButton(values["hide_app_info_entrance"][0] == "true")
		ui.GetService(services).RemoveOrShowToolNavButton(values["hide_tool_entrance"][0] == "true")
		ui.GetService(services).RemoveOrShowPlugNavButton(values["hide_plugin_entrance"][0] == "true")

		// TODO: add transaction
		err = models.Site().SetConn(s.conn).Update(values.RemoveSysRemark())
		if err != nil {
			return err
		}
		return s.c.Update(values.ToMap())
	})

	formList.EnableAjax(lgWithConfigScore("modify site config"),
		lgWithConfigScore("modify site config"),
		"",
		lgWithConfigScore("modify site config success"),
		lgWithConfigScore("modify site config fail"))

	return
}

func (s *SystemTable) GetGenerateForm(ctx *context.Context) (generateTool Table) {
	generateTool = NewDefaultTable(DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver).
		SetOnlyNewForm())

	formList := generateTool.GetForm().AddXssJsFilter().
		SetHeadWidth(1).
		SetInputWidth(4).
		HideBackButton().
		HideContinueNewCheckBox()

	formList.AddField("ID", "id", db.Varchar, form.Default).FieldDefault("1").FieldHide()

	connNames := config.GetDatabases().Connections()
	ops := make(types.FieldOptions, len(connNames))
	for i, name := range connNames {
		ops[i] = types.FieldOption{Text: name, Value: name}
	}

	// General options
	// ================================

	formList.AddField(lgWithScore("connection", "tool"), "conn", db.Varchar, form.SelectSingle).
		FieldOptions(ops).
		FieldOnChooseAjax("table", "/tool/choose/conn",
			func(ctx *context.Context) (success bool, msg string, data interface{}) {
				connName := ctx.FormValue("value")
				if connName == "" {
					return false, "wrong parameter", nil
				}
				cfg := s.c.Databases[connName]
				conn := db.GetConnectionFromService(services.Get(cfg.Driver))
				tables, err := db.WithDriverAndConnection(connName, conn).Table(cfg.Name).ShowTables()
				if err != nil {
					return false, err.Error(), nil
				}
				ops := make(selection.Options, len(tables))
				for i, table := range tables {
					ops[i] = selection.Option{Text: table, ID: table}
				}
				return true, "ok", ops
			})
	formList.AddField(lgWithScore("table", "tool"), "table", db.Varchar, form.SelectSingle).
		FieldOnChooseAjax("xxxx", "/tool/choose/table",
			func(ctx *context.Context) (success bool, msg string, data interface{}) {

				var (
					tableName       = ctx.FormValue("value")
					connName        = ctx.FormValue("conn")
					driver          = s.c.Databases[connName].Driver
					conn            = db.GetConnectionFromService(services.Get(driver))
					columnsModel, _ = db.WithDriverAndConnection(connName, conn).Table(tableName).ShowColumns()

					fieldField = "Field"
					typeField  = "Type"
				)

				if driver == "postgresql" {
					fieldField = "column_name"
					typeField = "udt_name"
				}
				if driver == "sqlite" {
					fieldField = "name"
					typeField = "type"
				}
				if driver == "mssql" {
					fieldField = "column_name"
					typeField = "data_type"
				}

				headName := make([]string, len(columnsModel))
				fieldName := make([]string, len(columnsModel))
				dbTypeList := make([]string, len(columnsModel))
				formTypeList := make([]string, len(columnsModel))

				for i, model := range columnsModel {
					typeName := getType(model[typeField].(string))

					headName[i] = strings.Title(model[fieldField].(string))
					fieldName[i] = model[fieldField].(string)
					dbTypeList[i] = typeName
					formTypeList[i] = form.GetFormTypeFromFieldType(db.DT(strings.ToUpper(typeName)),
						model[fieldField].(string))
				}

				return true, "ok", [][]string{headName, fieldName, dbTypeList, formTypeList}
			}, template.HTML(utils.ParseText("choose_table_ajax", tmpls["choose_table_ajax"], nil)),
			`"conn":$('.conn').val(),`)
	formList.AddField(lgWithScore("package", "tool"), "package", db.Varchar, form.Text).FieldDefault("tables")
	formList.AddField(lgWithScore("primarykey", "tool"), "pk", db.Varchar, form.Text).FieldDefault("id")

	formList.AddField(lgWithScore("table permission", "tool"), "permission", db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: lgWithScore("yes", "tool"), Value: "y"},
			{Text: lgWithScore("no", "tool"), Value: "n"},
		}).FieldDefault("n")

	formList.AddField(lgWithScore("extra import package", "tool"), "extra_import_package", db.Varchar, form.Select).
		FieldOptions(types.FieldOptions{
			{Text: "time", Value: "time"},
			{Text: "log", Value: "log"},
			{Text: "fmt", Value: "fmt"},
			{Text: "github.com/GoAdminGroup/go-admin/modules/db/dialect", Value: "github.com/GoAdminGroup/go-admin/modules/db/dialect"},
			{Text: "github.com/GoAdminGroup/go-admin/modules/db", Value: "github.com/GoAdminGroup/go-admin/modules/db"},
			{Text: "github.com/GoAdminGroup/go-admin/modules/language", Value: "github.com/GoAdminGroup/go-admin/modules/language"},
			{Text: "github.com/GoAdminGroup/go-admin/modules/logger", Value: "github.com/GoAdminGroup/go-admin/modules/logger"},
		}).
		FieldDefault("").
		FieldOptionExt(map[string]interface{}{
			"tags": true,
		})

	formList.AddField(lgWithScore("output", "tool"), "path", db.Varchar, form.Text).
		FieldDefault("").FieldMust().FieldHelpMsg(template.HTML(lgWithScore("use absolute path", "tool")))

	formList.AddField(lgWithScore("extra code", "tool"), "extra_code", db.Varchar, form.Code).
		FieldDefault("").FieldInputWidth(11)

	// Info table generate options
	// ================================

	formList.AddField(lgWithScore("title", "tool"), "table_title", db.Varchar, form.Text)
	formList.AddField(lgWithScore("description", "tool"), "table_description", db.Varchar, form.Text)

	formList.AddRow(func(panel *types.FormPanel) {
		addSwitchForTool(panel, "filter area", "hide_filter_area", "n", 2)
		panel.AddField(lgWithScore("filter form layout", "tool"), "filter_form_layout", db.Varchar, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: form.LayoutDefault.String(), Value: form.LayoutDefault.String()},
				{Text: form.LayoutTwoCol.String(), Value: form.LayoutTwoCol.String()},
				{Text: form.LayoutThreeCol.String(), Value: form.LayoutThreeCol.String()},
				{Text: form.LayoutFourCol.String(), Value: form.LayoutFourCol.String()},
				{Text: form.LayoutFlow.String(), Value: form.LayoutFlow.String()},
			}).FieldDefault(form.LayoutDefault.String()).
			FieldRowWidth(4).FieldHeadWidth(3)
	})

	formList.AddRow(func(panel *types.FormPanel) {
		addSwitchForTool(panel, "new button", "hide_new_button", "n", 2)
		addSwitchForTool(panel, "export button", "hide_export_button", "n", 4, 3)
		addSwitchForTool(panel, "edit button", "hide_edit_button", "n", 4, 2)
	})

	formList.AddRow(func(panel *types.FormPanel) {
		addSwitchForTool(panel, "pagination", "hide_pagination", "n", 2)
		addSwitchForTool(panel, "delete button", "hide_delete_button", "n", 4, 3)
		addSwitchForTool(panel, "detail button", "hide_detail_button", "n", 4, 2)
	})

	formList.AddRow(func(panel *types.FormPanel) {
		addSwitchForTool(panel, "filter button", "hide_filter_button", "n", 2)
		addSwitchForTool(panel, "row selector", "hide_row_selector", "n", 4, 3)
		addSwitchForTool(panel, "query info", "hide_query_info", "n", 4, 2)
	})

	formList.AddTable(lgWithScore("field", "tool"), "fields", func(pa *types.FormPanel) {
		pa.AddField(lgWithScore("title", "tool"), "field_head", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("field name", "tool"), "field_name", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("field filterable", "tool"), "field_filterable", db.Varchar, form.CheckboxSingle).
			FieldOptions(types.FieldOptions{
				{Text: "", Value: "y"},
				{Text: "", Value: "n"},
			}).
			FieldDefault("n").
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"n"}
			})
		pa.AddField(lgWithScore("field sortable", "tool"), "field_sortable", db.Varchar, form.CheckboxSingle).
			FieldOptions(types.FieldOptions{
				{Text: "", Value: "y"},
				{Text: "", Value: "n"},
			}).
			FieldDefault("n").
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"n"}
			})
		pa.AddField(lgWithScore("field hide", "tool"), "field_hide", db.Varchar, form.CheckboxSingle).
			FieldOptions(types.FieldOptions{
				{Text: "", Value: "y"},
				{Text: "", Value: "n"},
			}).
			FieldDefault("n").
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"n"}
			})
		pa.AddField(lgWithScore("info field editable", "tool"), "info_field_editable", db.Varchar, form.CheckboxSingle).
			FieldOptions(types.FieldOptions{
				{Text: "", Value: "y"},
				{Text: "", Value: "n"},
			}).
			FieldDefault("n").
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"n"}
			})
		//pa.AddField(lgWithScore("db display type", "tool"), "field_display_type", db.Varchar, form.SelectSingle).
		//	FieldOptions(infoFieldDisplayTypeOptions()).
		//	FieldDisplay(func(value types.FieldModel) interface{} {
		//		return []string{""}
		//	})
		pa.AddField(lgWithScore("db type", "tool"), "field_db_type", db.Varchar, form.SelectSingle).
			FieldOptions(databaseTypeOptions()).
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"Int"}
			})
	}).FieldInputWidth(11)

	// Form generate options
	// ================================

	formList.AddField(lgWithScore("title", "tool"), "form_title", db.Varchar, form.Text)
	formList.AddField(lgWithScore("description", "tool"), "form_description", db.Varchar, form.Text)

	formList.AddRow(func(panel *types.FormPanel) {
		addSwitchForTool(panel, "continue edit checkbox", "hide_continue_edit_check_box", "n", 2)
		addSwitchForTool(panel, "reset button", "hide_reset_button", "n", 5, 3)
	})

	formList.AddRow(func(panel *types.FormPanel) {
		addSwitchForTool(panel, "continue new checkbox", "hide_continue_new_check_box", "n", 2)
		addSwitchForTool(panel, "back button", "hide_back_button", "n", 5, 3)
	})

	formList.AddTable(lgWithScore("field", "tool"), "fields_form", func(pa *types.FormPanel) {
		pa.AddField(lgWithScore("title", "tool"), "field_head_form", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("field name", "tool"), "field_name_form", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("field editable", "tool"), "field_canedit", db.Varchar, form.CheckboxSingle).
			FieldOptions(types.FieldOptions{
				{Text: "", Value: "y"},
				{Text: "", Value: "n"},
			}).
			FieldDefault("y").
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"y"}
			})
		pa.AddField(lgWithScore("field can add", "tool"), "field_canadd", db.Varchar, form.CheckboxSingle).
			FieldOptions(types.FieldOptions{
				{Text: "", Value: "y"},
				{Text: "", Value: "n"},
			}).
			FieldDefault("y").
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"y"}
			})
		pa.AddField(lgWithScore("field default", "tool"), "field_default", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("field display", "tool"), "field_display", db.Varchar, form.SelectSingle).
			FieldOptions(types.FieldOptions{
				{Text: lgWithScore("field display normal", "tool"), Value: "0"},
				{Text: lgWithScore("field diplay hide", "tool"), Value: "1"},
				{Text: lgWithScore("field diplay edit hide", "tool"), Value: "2"},
				{Text: lgWithScore("field diplay create hide", "tool"), Value: "3"},
			}).
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"0"}
			})
		pa.AddField(lgWithScore("db type", "tool"), "field_db_type_form", db.Varchar, form.SelectSingle).
			FieldOptions(databaseTypeOptions()).
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"Int"}
			})
		pa.AddField(lgWithScore("form type", "tool"), "field_form_type_form", db.Varchar, form.SelectSingle).
			FieldOptions(formTypeOptions()).FieldDisplay(func(value types.FieldModel) interface{} {
			return []string{"Text"}
		})
	}).FieldInputWidth(11)

	// Detail page generate options
	// ================================

	formList.AddField(lgWithScore("title", "tool"), "detail_title", db.Varchar, form.Text)
	formList.AddField(lgWithScore("description", "tool"), "detail_description", db.Varchar, form.Text)

	formList.AddField(lgWithScore("detail display", "tool"), "detail_display", db.Varchar, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: lgWithScore("follow list page", "tool"), Value: "0"},
			{Text: lgWithScore("inherit from list page", "tool"), Value: "1"},
			{Text: lgWithScore("independent from list page", "tool"), Value: "2"},
		}).
		FieldDefault("0").
		FieldOnChooseHide("0", "detail_title", "detail_description", "fields_detail")

	formList.AddTable(lgWithScore("field", "tool"), "fields_detail", func(pa *types.FormPanel) {
		pa.AddField(lgWithScore("title", "tool"), "detail_field_head", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("field name", "tool"), "detail_field_name", db.Varchar, form.Text).FieldHideLabel().
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{""}
			})
		pa.AddField(lgWithScore("db type", "tool"), "detail_field_db_type", db.Varchar, form.SelectSingle).
			FieldOptions(databaseTypeOptions()).
			FieldDisplay(func(value types.FieldModel) interface{} {
				return []string{"Int"}
			})
	}).FieldInputWidth(11)

	formList.SetTabGroups(types.
		NewTabGroups("conn", "table", "package", "pk", "permission", "extra_import_package", "path", "extra_code").
		AddGroup("table_title", "table_description", "hide_filter_area", "filter_form_layout",
			"hide_new_button", "hide_export_button", "hide_edit_button",
			"hide_pagination", "hide_delete_button", "hide_detail_button",
			"hide_filter_button", "hide_row_selector", "hide_query_info",
			"fields").
		AddGroup("form_title", "form_description", "hide_continue_edit_check_box", "hide_reset_button",
			"hide_continue_new_check_box", "hide_back_button",
			"fields_form").
		AddGroup("detail_display", "detail_title", "detail_description", "fields_detail")).
		SetTabHeaders(lgWithScore("basic info", "tool"), lgWithScore("table info", "tool"),
			lgWithScore("form info", "tool"), lgWithScore("detail info", "tool"))

	formList.SetTable("tool").
		SetTitle(lgWithScore("tool", "tool")).
		SetDescription(lgWithScore("tool", "tool")).
		SetHeader(template.HTML(`<h3 class="box-title">` +
			lgWithScore("generate table model", "tool") + `</h3>`))

	formList.SetInsertFn(func(values form2.Values) error {

		table := values.Get("table")

		if table == "" {
			return errors.New("table is empty")
		}

		if values.Get("permission") == "y" {
			tools.InsertPermissionOfTable(s.conn, table)
		}

		output := values.Get("path")

		if output == "" {
			return errors.New("output path is empty")
		}

		connName := values.Get("conn")

		fields := make(tools.Fields, len(values["field_head"]))

		for i := 0; i < len(values["field_head"]); i++ {
			fields[i] = tools.Field{
				Head:         values["field_head"][i],
				Name:         values["field_name"][i],
				DBType:       values["field_db_type"][i],
				Filterable:   values["field_filterable"][i] == "y",
				Sortable:     values["field_sortable"][i] == "y",
				Hide:         values["field_hide"][i] == "y",
				InfoEditable: values["info_field_editable"][i] == "y",
			}
		}

		extraImport := ""
		for _, pack := range values["extra_import_package[]"] {
			if extraImport != "" {
				extraImport += `
`
			}
			extraImport += `	"` + pack + `"`
		}

		formFields := make(tools.Fields, len(values["field_head_form"]))

		for i := 0; i < len(values["field_head_form"]); i++ {
			extraFun := ""
			if values["field_name_form"][i] == `created_at` {
				extraFun += `.FieldNowWhenInsert()`
			} else if values["field_name_form"][i] == `updated_at` {
				extraFun += `.FieldNowWhenUpdate()`
			} else if values["field_default"][i] != "" && !strings.Contains(values["field_default"][i], `"`) {
				values["field_default"][i] = `"` + values["field_default"][i] + `"`
			}
			formFields[i] = tools.Field{
				Head:       values["field_head_form"][i],
				Name:       values["field_name_form"][i],
				Default:    values["field_default"][i],
				FormType:   values["field_form_type_form"][i],
				DBType:     values["field_db_type_form"][i],
				CanAdd:     values["field_canadd"][i] == "y",
				Editable:   values["field_canedit"][i] == "y",
				FormHide:   values["field_display"][i] == "1",
				CreateHide: values["field_display"][i] == "2",
				EditHide:   values["field_display"][i] == "3",
				ExtraFun:   extraFun,
			}
		}

		detailFields := make(tools.Fields, len(values["detail_field_head"]))

		for i := 0; i < len(values["detail_field_head"]); i++ {
			detailFields[i] = tools.Field{
				Head:   values["detail_field_head"][i],
				Name:   values["detail_field_name"][i],
				DBType: values["detail_field_db_type"][i],
			}
		}

		detailDisplay, _ := strconv.ParseUint(values.Get("detail_display"), 10, 64)

		err := tools.Generate(tools.NewParamWithFields(tools.Config{
			Connection:               connName,
			Driver:                   s.c.Databases[connName].Driver,
			Package:                  values.Get("package"),
			Table:                    table,
			HideFilterArea:           values.Get("hide_filter_area") == "y",
			HideNewButton:            values.Get("hide_new_button") == "y",
			HideExportButton:         values.Get("hide_export_button") == "y",
			HideEditButton:           values.Get("hide_edit_button") == "y",
			HideDeleteButton:         values.Get("hide_delete_button") == "y",
			HideDetailButton:         values.Get("hide_detail_button") == "y",
			HideFilterButton:         values.Get("hide_filter_button") == "y",
			HideRowSelector:          values.Get("hide_row_selector") == "y",
			HidePagination:           values.Get("hide_pagination") == "y",
			HideQueryInfo:            values.Get("hide_query_info") == "y",
			HideContinueEditCheckBox: values.Get("hide_continue_edit_check_box") == "y",
			HideContinueNewCheckBox:  values.Get("hide_continue_new_check_box") == "y",
			HideResetButton:          values.Get("hide_reset_button") == "y",
			HideBackButton:           values.Get("hide_back_button") == "y",
			FilterFormLayout:         form.GetLayoutFromString(values.Get("filter_form_layout")),
			Schema:                   values.Get("schema"),
			Output:                   output,
			DetailDisplay:            uint8(detailDisplay),
			FormTitle:                values.Get("form_title"),
			FormDescription:          values.Get("form_description"),
			DetailTitle:              values.Get("detail_title"),
			DetailDescription:        values.Get("detail_description"),
			TableTitle:               values.Get("table_title"),
			TableDescription:         values.Get("table_description"),
			ExtraImport:              extraImport,
			ExtraCode:                escape(values.Get("extra_code")),
		}, fields, formFields, detailFields))

		if err != nil {
			return err
		}

		return tools.GenerateTables(output, values.Get("package"), []string{table}, false)
	})

	formList.EnableAjaxData(types.AjaxData{
		SuccessTitle: lgWithScore("generate table model", "tool"),
		ErrorTitle:   lgWithScore("generate table model", "tool"),
		SuccessText:  lgWithScore("generate table model success", "tool"),
		ErrorText:    lgWithScore("generate table model fail", "tool"),
		DisableJump:  true,
	})

	formList.SetFooterHtml(utils.ParseHTML("generator", tmpls["generator"], map[string]string{
		"prefix": "go_admin_" + config.GetAppID() + "_generator_",
	}))

	formList.SetFormNewBtnWord(template.HTML(lgWithScore("generate", "tool")))
	formList.SetWrapper(func(content tmpl.HTML) tmpl.HTML {
		headli := html.LiEl().SetClass("list-group-item", "list-head").
			SetContent(template.HTML(lgWithScore("generated tables", "tool"))).Get()
		return html.UlEl().SetClass("save_table_list", "list-group").SetContent(
			headli).Get() + content
	})

	formList.SetHideSideBar()

	return generateTool
}

// -------------------------
// helper functions
// -------------------------

func encodePassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

func label() types.LabelAttribute {
	return template.Get(config.GetTheme()).Label().SetType("success")
}

func lg(v string) string {
	return language.Get(v)
}

func defaultFilterFn(val string, def ...string) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		if len(def) > 0 {
			if value.Value == def[0] {
				return val
			}
		} else {
			if value.Value == "" {
				return val
			}
		}
		return value.Value
	}
}

func lgWithScore(v string, score ...string) string {
	return language.GetWithScope(v, score...)
}

func lgWithConfigScore(v string, score ...string) string {
	scores := append([]string{"config"}, score...)
	return language.GetWithScope(v, scores...)
}

func link(url, content string) tmpl.HTML {
	return html.AEl().
		SetAttr("href", url).
		SetContent(template.HTML(lg(content))).
		Get()
}

func escape(s string) string {
	if s == "" {
		return ""
	}
	s, err := url.QueryUnescape(s)
	if err != nil {
		logger.Error("escape error", err)
	}
	return s
}

func checkJSON(values form2.Values, key string) error {
	v := values.Get(key)
	if v != "" && !utils.IsJSON(v) {
		return errors.New("wrong " + key)
	}
	return nil
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

func addSwitchForTool(formList *types.FormPanel, head, field, def string, row ...int) {
	formList.AddField(lgWithScore(head, "tool"), field, db.Varchar, form.Switch).
		FieldOptions(types.FieldOptions{
			{Text: lgWithScore("show", "tool"), Value: "n"},
			{Text: lgWithScore("hide", "tool"), Value: "y"},
		}).FieldDefault(def)
	if len(row) > 0 {
		formList.FieldRowWidth(row[0])
	}
	if len(row) > 1 {
		formList.FieldHeadWidth(row[1])
	}
	if len(row) > 2 {
		formList.FieldInputWidth(row[2])
	}
}

func formTypeOptions() types.FieldOptions {
	opts := make(types.FieldOptions, len(form.AllType))
	for i := 0; i < len(form.AllType); i++ {
		v := form.AllType[i].Name()
		opts[i] = types.FieldOption{Text: v, Value: v}
	}
	return opts
}

func databaseTypeOptions() types.FieldOptions {
	opts := make(types.FieldOptions, len(db.IntTypeList)+
		len(db.StringTypeList)+
		len(db.FloatTypeList)+
		len(db.UintTypeList)+
		len(db.BoolTypeList))
	z := 0
	for _, t := range db.IntTypeList {
		text := string(t)
		v := strings.Title(strings.ToLower(text))
		opts[z] = types.FieldOption{Text: text, Value: v}
		z++
	}
	for _, t := range db.StringTypeList {
		text := string(t)
		v := strings.Title(strings.ToLower(text))
		opts[z] = types.FieldOption{Text: text, Value: v}
		z++
	}
	for _, t := range db.FloatTypeList {
		text := string(t)
		v := strings.Title(strings.ToLower(text))
		opts[z] = types.FieldOption{Text: text, Value: v}
		z++
	}
	for _, t := range db.UintTypeList {
		text := string(t)
		v := strings.Title(strings.ToLower(text))
		opts[z] = types.FieldOption{Text: text, Value: v}
		z++
	}
	for _, t := range db.BoolTypeList {
		text := string(t)
		v := strings.Title(strings.ToLower(text))
		opts[z] = types.FieldOption{Text: text, Value: v}
		z++
	}
	return opts
}

func getType(typeName string) string {
	r, _ := regexp.Compile(`\(.*?\)`)
	typeName = r.ReplaceAllString(typeName, "")
	r2, _ := regexp.Compile(`unsigned(.*)`)
	return strings.TrimSpace(strings.Title(strings.ToLower(r2.ReplaceAllString(typeName, ""))))
}
