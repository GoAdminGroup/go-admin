package datamodel

import (
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetPostsTable() (postsTable table.Table) {

	postsTable = table.NewDefaultTable(table.DefaultTableConfig)

	postsTable.GetInfo().FieldList = []types.Field{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Title",
			Field:    "title",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Description",
			Field:    "description",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Content",
			Field:    "content",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Date",
			Field:    "date",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
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
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: form.Default,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Title",
			Field:    "title",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Description",
			Field:    "description",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Content",
			Field:    "content",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.RichText,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Date",
			Field:    "date",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	postsTable.GetForm().Table = "posts"
	postsTable.GetForm().Title = "Posts"
	postsTable.GetForm().Description = "Posts"

	return
}
