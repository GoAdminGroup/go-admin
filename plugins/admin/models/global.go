package models

type GetTableDataFunc func() GlobalTable

// map下标是路由前缀，对应的值是GetTableDataFunc类型，为表单与表格的数据抽象表示
var TableFuncConfig = map[string]GetTableDataFunc{}

var GlobalTableList = map[string]GlobalTable{}

func InitGlobalTableList() {
	table := make(map[string]GlobalTable, len(TableFuncConfig))
	for k, v := range TableFuncConfig {
		table[k] = v()
	}
	GlobalTableList = table
}

func RefreshGlobalTableList() {
	for k, v := range TableFuncConfig {
		GlobalTableList[k] = v()
	}
}

func SetTableFuncConfig(cfglist map[string]GetTableDataFunc) {
	for key, value := range cfglist {
		TableFuncConfig[key] = value
	}
}