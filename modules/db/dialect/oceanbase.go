package dialect

type oceanbase struct {
	commonDialect
}

func (oceanbase) GetName() string {
	return "oceanbase"
}

func (oceanbase) ShowColumns(table string) string {
	return "show columns in " + table
}

func (oceanbase) ShowTables() string {
	return "show tables"
}
