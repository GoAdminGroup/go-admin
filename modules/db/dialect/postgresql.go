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

func (c postgresql) Select(comp *SqlComponent) string {
	comp.Statement = "select " + comp.getFields(c.delimiter) + ` from "` + comp.TableName + `"` + comp.getJoins(c.delimiter) +
		comp.getWheres(c.delimiter) + comp.getOrderBy() + comp.getLimit() + comp.getOffset()
	return comp.Statement
}