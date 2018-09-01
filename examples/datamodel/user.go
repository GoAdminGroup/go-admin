package datamodel

import (
	"goAdmin/template/adminlte/components"
	"goAdmin/plugins/admin/models"
)

func GetUserTable() (userTable models.GlobalTable) {

	userTable.Info.FieldList = []components.FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Sortable: true,
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			Sortable: false,
			ExcuFun: func(model components.RowModel) interface{} {
				if model.Value == "1" {
					return "男"
				}
				if model.Value == "2" {
					return "女"
				}
				return "未知"
			},
		},
		{
			Head:     "省份",
			Field:    "province",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		},
		{
			Head:     "城市",
			Field:    "city",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "users"
	userTable.Info.Title = "用户表"
	userTable.Info.Description = "用户表"

	userTable.Form.FormList = []components.FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "头像",
			Field:    "avatar",
			TypeName: "int64",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "省份",
			Field:    "province",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		}, {
			Head:     "城市",
			Field:    "city",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
			ExcuFun: func(model components.RowModel) interface{} {
				return model.Value
			},
		},
	}

	userTable.Form.Table = "users"
	userTable.Form.Title = "用户表"
	userTable.Form.Description = "用户表"

	return
}
