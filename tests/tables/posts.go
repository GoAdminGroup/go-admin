package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

// GetPostsTable return the model of table posts.
func GetPostsTable(ctx *context.Context) (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultConfig())

	info := postsTable.GetInfo()
	info.AddField("ID", "id", db.Int).FieldSortable()
	info.AddField("Title", "title", db.Varchar)
	info.AddField("AuthorID", "author_id", db.Varchar).FieldDisplay(func(value types.FieldModel) interface{} {
		return template.Default().
			Link().
			SetURL("/admin/info/authors/detail?__goadmin_detail_pk=100").
			SetContent("100").
			OpenInNewTab().
			SetTabTitle("Author Detail").
			GetContent()
	})
	info.AddField("Description", "description", db.Varchar)
	info.AddField("Content", "content", db.Varchar).FieldEditAble(editType.Textarea)
	info.AddField("Date", "date", db.Varchar)

	info.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	formList := postsTable.GetForm()
	formList.AddField("ID", "id", db.Int, form.Default).FieldDisplayButCanNotEditWhenUpdate().FieldDisableWhenCreate()
	formList.AddField("Title", "title", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Varchar, form.Text)
	formList.AddField("Content", "content", db.Varchar, form.RichText).FieldEnableFileUpload()
	formList.AddField("Date", "date", db.Varchar, form.Datetime)

	formList.SetTable("posts").SetTitle("Posts").SetDescription("Posts")

	return
}
