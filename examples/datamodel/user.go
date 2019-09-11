package datamodel

import (
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetUserTable() (userTable table.Table) {

	c := table.DefaultConfig
	c.Exportable = true

	userTable = table.NewDefaultTable(table.DefaultConfig)

	userTable.GetInfo().FieldList = []types.Field{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Gender",
			Field:    "gender",
			TypeName: "tinyint",
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				if model.Value == "1" {
					return "man"
				}
				if model.Value == "2" {
					return "women"
				}
				return "unknown"
			},
		},
		{
			Head:     "Phone",
			Field:    "phone",
			TypeName: "varchar",
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "City",
			Field:    "city",
			TypeName: "varchar",
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.GetInfo().Table = "users"
	userTable.GetInfo().Title = "Users"
	userTable.GetInfo().Description = "Users"

	userTable.GetForm().FormList = []types.Form{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Name",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Gender",
			Field:    "gender",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			Options: []map[string]string{
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
			},
			FormType: form.Radio,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Phone",
			Field:    "phone",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "City",
			Field:    "city",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Created at",
			Field:    "created_at",
			TypeName: "varchar",
			Default:  "2017-01-05 23:01:17",
			Editable: true,
			FormType: form.Datetime,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Updated at",
			Field:    "updated_at",
			TypeName: "varchar",
			Default:  "2017-01-05 23:01:17",
			Editable: true,
			FormType: form.Datetime,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.GetForm().Table = "users"
	userTable.GetForm().Title = "Users"
	userTable.GetForm().Description = "Users"

	return
}
