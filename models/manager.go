package models

import "goAdmin/connections/mysql"

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
			Head:     "标签",
			Field:    "label",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				labelModel, _ := mysql.Query("select r.name from goadmin_role_users as u left join goadmin_roles as r on "+
					"u.role_id = r.id where user_id = ?", model.ID)
				return labelModel[0]["name"].(string)
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
			Head:     "用户名",
			Field:    "username",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "昵称",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "密码",
			Field:    "password",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "password",
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		},
	}

	userTable.Form.Table = "goadmin_users"
	userTable.Form.Title = "管理员管理"
	userTable.Form.Description = "管理员管理"

	return
}

func GetPermissionTable() (userTable GlobalTable) {

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
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "方法",
			Field:    "http_method",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "路径",
			Field:    "http_path",
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

	userTable.Info.Table = "goadmin_permissions"
	userTable.Info.Title = "权限管理"
	userTable.Info.Description = "权限管理"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
		}, {
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "方法",
			Field:    "http_method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "select",
			Options: []map[string]string{
				{"value": "GET", "field": "GET"},
				{"value": "PUT", "field": "PUT"},
				{"value": "POST", "field": "POST"},
				{"value": "DELETE", "field": "DELETE"},
				{"value": "PATCH", "field": "PATCH"},
				{"value": "OPTIONS", "field": "OPTIONS"},
				{"value": "HEAD", "field": "HEAD"},
			},
		}, {
			Head:     "路径",
			Field:    "http_path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "textarea",
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		},
	}

	userTable.Form.Table = "goadmin_permissions"
	userTable.Form.Title = "权限管理"
	userTable.Form.Description = "权限管理"

	return
}

func GetRolesTable() (userTable GlobalTable) {

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
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "标志",
			Field:    "slug",
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

	userTable.Info.Table = "goadmin_roles"
	userTable.Info.Title = "角色管理"
	userTable.Info.Description = "角色管理"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
		}, {
			Head:     "名字",
			Field:    "name",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "标志",
			Field:    "slug",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		},
	}

	userTable.Form.Table = "goadmin_roles"
	userTable.Form.Title = "角色管理"
	userTable.Form.Description = "角色管理"

	return
}

func GetOpTable() (userTable GlobalTable) {

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
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "路径",
			Field:    "path",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "方法",
			Field:    "method",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			ExcuFun: func(model RowModel) string {
				return model.Value
			},
		},
		{
			Head:     "内容",
			Field:    "input",
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

	userTable.Info.Table = "goadmin_operation_log"
	userTable.Info.Title = "操作日志"
	userTable.Info.Description = "操作日志"

	userTable.Form.FormList = []FormStruct{
		{
			Head:     "ID",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
		}, {
			Head:     "用户ID",
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "路径",
			Field:    "path",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "方法",
			Field:    "method",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "ip",
			Field:    "ip",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "内容",
			Field:    "input",
			TypeName: "varchar",
			Default:  "",
			Editable: true,
			FormType: "text",
		}, {
			Head:     "更新时间",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		}, {
			Head:     "创建时间",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: true,
			FormType: "default",
		},
	}

	userTable.Form.Table = "goadmin_operation_log"
	userTable.Form.Title = "操作日志"
	userTable.Form.Description = "操作日志"

	return
}
