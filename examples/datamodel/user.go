package datamodel

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/template/types/form"
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
			TypeName: db.Tinyint,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				if model.Value == "0" {
					return "men"
				}
				if model.Value == "1" {
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
		{
			Head:     "Created At",
			Field:    "created_at",
			TypeName: db.Varchar,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "Updated At",
			Field:    "updated_at",
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

	userTable.GetInfo().GroupHeaders = []string{"profile1", "profile2"}
	userTable.GetInfo().Group = [][]string{
		{"id", "ip", "name", "gender", "city"},
		{"id", "phone", "created_at", "updated_at"},
	}

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
			TypeName: db.Tinyint,
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
			Head:     "Custom Field",
			Field:    "role",
			TypeName: db.Varchar,
			Default:  "",
			Editable: true,
			FormType: form.Text,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
			ProcessFn: func(value types.PostRowModel) {
				fmt.Println("user custom field", value)
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

	userTable.GetForm().Group = [][]string{
		{"id", "ip", "name", "gender", "city"},
		{"phone", "role", "created_at", "updated_at"},
	}
	userTable.GetForm().GroupHeaders = []string{
		"profile1", "profile2",
	}

	userTable.GetForm().Table = "users"
	userTable.GetForm().Title = "Users"
	userTable.GetForm().Description = "Users"

	return
}
