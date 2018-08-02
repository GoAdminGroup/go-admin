package models

func GetUserTable() (userTable GlobalTable) {

	userTable.Info.FieldList = []FieldStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "姓名",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "性别",
			Field:    "sex",
			TypeName: "tinyint",
			ExcuFun: func(model RowModel) string {
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
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "城市",
			Field:    "city",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
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
