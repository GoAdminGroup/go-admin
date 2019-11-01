package dialect

type sqlite struct {
	commonDialect
}

func (sqlite) GetName() string {
	return "sqlite"
}

func (sqlite) ShowColumns(table string) string {
	return "PRAGMA table_info(" + table + ");"
}

func (sqlite) ShowTables() string {
	return "SELECT name as tablename FROM sqlite_master WHERE type ='table'"
}
