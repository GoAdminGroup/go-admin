package datamodel

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetPostsTable() (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultConfig)

	postsTable.GetInfo().FieldList = []types.Field{
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
			Head:     "Title",
			Field:    "title",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Description",
			Field:    "description",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Content",
			Field:    "content",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Date",
			Field:    "date",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	postsTable.GetInfo().Table = "posts"
	postsTable.GetInfo().Title = "Posts"
	postsTable.GetInfo().Description = "Posts"

	postsTable.GetForm().FormList = []types.Form{
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
			Head:     "Title",
			Field:    "title",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "name",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Description",
			Field:    "description",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Content",
			Field:    "content",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.RichText,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Date",
			Field:    "date",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Datetime,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	postsTable.GetForm().Table = "posts"
	postsTable.GetForm().Title = "Posts"
	postsTable.GetForm().Description = "Posts"

	return
}
