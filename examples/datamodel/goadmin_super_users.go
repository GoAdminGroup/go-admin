package datamodel

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetGoadminSuperUsersTable(ctx *context.Context) table.Table {

	goadminSuperUsers := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := goadminSuperUsers.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("Username", "username", db.Varchar)
	info.AddField("Password", "password", db.Varchar)
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Avatar", "avatar", db.Varchar)
	info.AddField("Remember_token", "remember_token", db.Varchar)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Updated_at", "updated_at", db.Timestamp)

	info.SetTable("goadmin_super_users").SetTitle("GoadminSuperUsers").SetDescription("GoadminSuperUsers")

	formList := goadminSuperUsers.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("Password", "password", db.Varchar, form.Password)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Avatar", "avatar", db.Varchar, form.Text)
	formList.AddField("Remember_token", "remember_token", db.Varchar, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Updated_at", "updated_at", db.Timestamp, form.Datetime)

	formList.SetTable("goadmin_super_users").SetTitle("GoadminSuperUsers").SetDescription("GoadminSuperUsers")

	return goadminSuperUsers
}
