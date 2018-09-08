package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/models"

// map下标是路由前缀，对应的值是GetTableDataFunc类型，为表单与表格的数据抽象表示
var TableFuncConfig = map[string]models.GetTableDataFunc{

	// 自定义管理部分
	"user": GetUserTable,

	"posts": GetPostsTable,
	"authors": GetAuthorsTable,
}