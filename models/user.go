package models

import "goAdmin/components"

func GetUserTable() (userTable GlobalTable) {

	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			ExcuFun: func(value string) string {
				return value
			},
		},
		{
			Head:     "头像",
			Field:    "avatar",
			TypeName: "varchar",
			ExcuFun: func(value string) string {
				return components.Image.GetContent(value)
			},
		},
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(value string) string {
				return value
			},
		},
		{
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			ExcuFun: func(value string) string {
				if value == "1" {
					return "男"
				}
				if value == "2" {
					return "女"
				}
				return "未知"
			},
		},
		{
			Head:     "省份",
			Field:    "province",
			TypeName: "varchar",
			ExcuFun: func(value string) string {
				return value
			},
		},
		{
			Head:     "城市",
			Field:    "city",
			TypeName: "varchar",
			ExcuFun: func(value string) string {
				return value
			},
		},
	}

	userTable.Info.Table = "users"
	userTable.Info.Title = "用户表"
	userTable.Info.Description = "用户表"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
		}, {
			Head:     "头像",
			Field:    "avatar",
			TypeName: "int64",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "省份",
			Field:    "province",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "城市",
			Field:    "city",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		},
	}

	userTable.Form.Table = "users"
	userTable.Form.Title = "用户表"
	userTable.Form.Description = "用户表"

	return
}
