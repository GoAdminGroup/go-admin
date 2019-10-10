package datamodel

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetPostsTable() (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultConfig)

	info := postsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable(true)
	info.AddField("Title", "title", db.Varchar)
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Content", "content", db.Varchar)
	info.AddField("Date", "date", db.Varchar)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := postsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldEditable(false)
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Varchar, form.Text)
	formList.AddField("Date", "date", db.Varchar, form.Datetime)

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return
}
