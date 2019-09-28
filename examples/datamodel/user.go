package datamodel

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
)

func GetUserTable() (userTable table.Table) {

	userTable = table.NewDefaultTable(table.Config{
		Driver:     "mysql",
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: "default",
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	userTable.GetInfo().FieldList = []types.Field{
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
			Head:     "Name",
			Field:    "name",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Gender",
			Field:    "gender",
			TypeName: db.TinyInt,
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
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "City",
			Field:    "city",
			TypeName: db.Varchar,
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
			TypeName: db.Int,
			Default:  "",
			Editable: false,
			FormType: form.Default,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Ip",
			Field:    "ip",
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
			Head:     "Gender",
			Field:    "gender",
			TypeName: db.TinyInt,
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
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "City",
			Field:    "city",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Created at",
			Field:    "created_at",
			TypeName: db.Varchar,
			Default:  "2017-01-05 23:01:17",
			Editable: true,
			FormType: form.Datetime,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "Updated at",
			Field:    "updated_at",
			TypeName: db.Varchar,
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
