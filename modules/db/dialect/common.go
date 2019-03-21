package dialect

type commonDialect struct {
}

func (commonDialect) ShowColumns(table string) string {
	return "select column_name, udt_name from information_schema.columns where table_name = " + table
}

func (commonDialect) GetName() string {
	return "common"
}

func (commonDialect) ShowTables() string {
	return "show tables"
}