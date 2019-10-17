package datamodel

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetUserTable() (userTable table.Table) {

	userTable = table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: table.DefaultConnectionName,
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := userTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Gender", "gender", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "men"
		}
		if model.Value == "1" {
			return "women"
		}
		return "unknown"
	})
	info.AddField("Phone", "phone", db.Varchar)
	info.AddField("City", "city", db.Varchar)
	info.AddField("createdAt", "created_at", db.Timestamp)
	info.AddField("updatedAt", "updated_at", db.Timestamp)

	info.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList := userTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField("Ip", "ip", db.Varchar, form.Text)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Gender", "gender", db.Tinyint, form.Radio).
		FieldOptions([]map[string]string{
			{
				"field":    "gender",
				"label":    "male",
				"value":    "0",
				"selected": "true",
			}, {
				"field":    "gender",
				"label":    "female",
				"value":    "1",
				"selected": "false",
			},
		})
	formList.AddField("Phone", "phone", db.Varchar, form.Text)
	formList.AddField("City", "city", db.Varchar, form.Text)
	formList.AddField("Custom Field", "role", db.Varchar, form.Text).
		FieldPostFilterFn(func(value types.PostFieldModel) string {
			fmt.Println("user custom field", value)
			return ""
		})

	formList.AddField("updatedAt", "updated_at", db.Timestamp, form.Default).FieldNotAllowAdd()
	formList.AddField("createdAt", "created_at", db.Timestamp, form.Default).FieldNotAllowAdd()

	userTable.GetForm().SetTabGroups(types.
		NewTabGroups("id", "ip", "name", "gender", "city").
		AddGroup("phone", "role", "created_at", "updated_at")).
		SetTabHeaders("profile1", "profile2")

	formList.SetTable("users").SetTitle("Users").SetDescription("Users")

	formList.SetPostHook(func(values form2.Values) {
		fmt.Println("userTable.GetForm().PostHook", values)
	})

	return
}
