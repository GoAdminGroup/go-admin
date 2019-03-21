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
