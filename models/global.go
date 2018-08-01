package models

// map下标是路由前缀，对应的值是GlobalTable类型，为表单与表格的数据抽象表示
var GlobalTableList = map[string]GlobalTable{
	"user": GetUserTable(),
	"manager": GetManagerTable(),
}
