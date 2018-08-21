package models

type GetTableDataFunc func() GlobalTable

// map下标是路由前缀，对应的值是GetTableDataFunc类型，为表单与表格的数据抽象表示
var TableFuncConfig = map[string]GetTableDataFunc{

	// 管理员管理部分
	"manager":    GetManagerTable,    // 管理员管理
	"permission": GetPermissionTable, // 权限管理
	"roles":      GetRolesTable,      // 角色管理
	"op":         GetOpTable,         // 操作日志管理

	// 自定义管理部分
	"user": GetUserTable,
}

func InitGlobalTableList() map[string]GlobalTable {
	table := make(map[string]GlobalTable, len(TableFuncConfig))
	for k, v := range TableFuncConfig {
		table[k] = v()
	}
	return table
}

var GlobalTableList = InitGlobalTableList()

func RefreshGlobalTableList() {
	for k, v := range TableFuncConfig {
		GlobalTableList[k] = v()
	}
}
