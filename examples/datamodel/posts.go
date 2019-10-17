package datamodel

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPostsTable() (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultConfig)

	info := postsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Content", "content", db.Varchar)
	info.AddField("Date", "date", db.Varchar)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := postsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldNotAllowEdit()
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Varchar, form.Text)
	formList.AddField("Date", "date", db.Varchar, form.Datetime)

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return
}
