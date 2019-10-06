package dialect

type postgresql struct {
	commonDialect
}

func (postgresql) GetName() string {
	return "postgresql"
}

func (postgresql) ShowTables() string {
	return "SELECT tablename FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND schemaname != 'information_schema';"
}
