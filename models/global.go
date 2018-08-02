package models

// map下标是路由前缀，对应的值是GlobalTable类型，为表单与表格的数据抽象表示
var GlobalTableList = map[string]GlobalTable{
	// 管理员管理部分
	"manager":    GetManagerTable(),    // 管理员管理
	"permission": GetPermissionTable(), // 权限管理
	"roles":      GetRolesTable(),      // 角色管理
	"op":         GetOpTable(),         // 操作日志管理

	// 自定义管理部分
	"user": GetUserTable(),
}
