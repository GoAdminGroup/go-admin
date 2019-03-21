package dialect

type mssql struct {
	commonDialect
}

func (mssql) GetName() string {
	return "mssql"
}
