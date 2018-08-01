package models

func GetManagerTable() (userTable GlobalTable) {

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
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
	}

	userTable.Info.Table = "goadmin_users"
	userTable.Info.Title = "管理员管理"
	userTable.Info.Description = "管理员管理"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		},
	}

	userTable.Form.Table = "goadmin_users"
	userTable.Form.Title = "管理员管理"
	userTable.Form.Description = "管理员管理"

	return
}
