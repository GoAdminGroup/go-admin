package datamodel

import (
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetAuthorsTable() (authorsTable table.Table) {

	authorsTable = table.NewDefaultTable(table.DefaultTableConfig)

	// connect your custom connection
	// authorsTable = table.NewDefaultTable(table.DefaultTableConfigWithDriverAndConnection("mysql", "admin"))

	authorsTable.GetInfo().FieldList = []types.Field{
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
			Head:     "First Name",
			Field:    "first_name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Last Name",
			Field:    "last_name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Email",
			Field:    "email",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Birthdate",
			Field:    "birthdate",
			TypeName: "date",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Added",
			Field:    "added",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	authorsTable.GetInfo().Table = "authors"
	authorsTable.GetInfo().Title = "Authors"
	authorsTable.GetInfo().Description = "Authors"

	authorsTable.GetForm().FormList = []types.Form{
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
			Head:     "First Name",
			Field:    "first_name",
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
			Head:     "Email",
			Field:    "email",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Birthdate",
			Field:    "birthdate",
			TypeName: "date",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Added",
			Field:    "added",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	authorsTable.GetForm().Table = "authors"
	authorsTable.GetForm().Title = "Authors"
	authorsTable.GetForm().Description = "Authors"

	return
}
