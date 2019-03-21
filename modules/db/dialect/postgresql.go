package dialect

type postgresql struct {
	commonDialect
}

func (postgresql) GetName() string {
	return "postgresql"
}

func (postgresql) ShowTables() string {
	return "select tablename from pg_catalog.pg_tables"
}