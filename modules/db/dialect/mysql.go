package dialect

type mysql struct {
	commonDialect
}

func (mysql) GetName() string {
	return "mysql"
}

func (mysql) ShowColumns(table string) string {
	return "show columns in " + table
}

func (mysql) ShowTables() string {
	return "show tables"
}
