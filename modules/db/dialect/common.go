package dialect

type commonDialect struct {
}

func (commonDialect) Insert(comp *SqlComponent) string {
	comp.prepareInsert()
	return comp.Statement
}

func (commonDialect) Delete(comp *SqlComponent) string {
	comp.Statement = "delete from " + comp.TableName + comp.getWheres()
	return comp.Statement
}

func (commonDialect) Update(comp *SqlComponent) string {
	comp.prepareUpdate()
	return comp.Statement
}

func (commonDialect) Count(comp *SqlComponent) string {
	comp.prepareUpdate()
	return comp.Statement
}

func (commonDialect) Select(comp *SqlComponent) string {
	comp.Statement = "select " + comp.getFields() + " from " + comp.TableName + comp.getJoins() + comp.getWheres() +
		comp.getOrderBy() + comp.getLimit() + comp.getOffset()
	return comp.Statement
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
