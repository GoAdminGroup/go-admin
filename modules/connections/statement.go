package connections

import (
	"errors"
)

type Where struct {
	operation string
	field     string
}

type Join struct {
	table     string
	fieldA    string
	operation string
	fieldB    string
}

type Sql struct {
	fields    []string
	table     string
	wheres    []Where
	leftjoins []Join
	args      []interface{}
}

type H map[string]interface{}

func newSql() *Sql {
	return &Sql{
		fields:    make([]string, 0),
		table:     "",
		args:      make([]interface{}, 0),
		wheres:    make([]Where, 0),
		leftjoins: make([]Join, 0),
	}
}

func Table(table string) *Sql {
	sql := newSql()
	sql.table = table
	return sql
}

func (sql *Sql) Select(fields ...string) *Sql {
	sql.fields = fields
	return sql
}

func (sql *Sql) Where(field string, operation string, arg interface{}) *Sql {
	sql.wheres = append(sql.wheres, Where{
		field:     field,
		operation: operation,
	})
	sql.args = append(sql.args, arg)
	return sql
}

func (sql *Sql) LeftJoin(table string, fieldA string, operation string, fieldB string) *Sql {
	sql.leftjoins = append(sql.leftjoins, Join{
		fieldA:    fieldA,
		fieldB:    fieldB,
		table:     table,
		operation: operation,
	})
	return sql
}

func (sql *Sql) First() (map[string]interface{}, error) {

	var statement = "select " + GetFields(sql.fields) + " from " + sql.table + GetJoins(sql.leftjoins) + GetWheres(sql.wheres)

	res, _ := GetConnection().Query(statement, sql.args...)
	if len(res) < 1 {
		return nil, errors.New("out of index")
	}
	return res[0], nil
}

func (sql *Sql) All() ([]map[string]interface{}, error) {

	var statement = "select " + GetFields(sql.fields) + " from " + sql.table + GetJoins(sql.leftjoins) + GetWheres(sql.wheres)

	res, _ := GetConnection().Query(statement, sql.args...)

	return res, nil
}

func (sql *Sql) Update(values H) error {

	fields := ""

	args := make([]interface{}, 0)
	for key, value := range values {
		fields += "`" + key + "` = ?, "
		args = append(args, value)
	}

	var statement = "update " + sql.table + " set " + fields[:len(fields)-2] + GetWheres(sql.wheres)
	sql.args = append(args, sql.args...)

	res := GetConnection().Exec(statement, sql.args...)

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return errors.New("no affect row")
	}

	return nil
}

func (sql *Sql) Insert(values H) error {

	fields := "("
	quesMark := "("

	for key, value := range values {
		fields += "`" + key + "`,"
		quesMark += "?,"
		sql.args = append(sql.args, value)
	}
	fields = fields[:len(fields)-1] + ")"
	quesMark = quesMark[:len(quesMark)-1] + ")"

	var statement = "insert into " + sql.table + fields + " values " + quesMark

	res := GetConnection().Exec(statement, sql.args...)

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return errors.New("no affect row")
	}

	return nil
}

func GetJoins(list []Join) string {
	if len(list) == 0 {
		return ""
	}
	joins := " left join "
	for _, join := range list {
		joins += join.table + " on " + join.fieldA + " " + join.operation + " " + join.fieldB + " "
	}
	return joins
}

func GetFields(list []string) string {
	if len(list) == 0 {
		return "*"
	}
	fields := ""
	for _, field := range list {
		fields += "`" + field + "`,"
	}
	return fields[:len(fields)-1]
}

func GetWheres(list []Where) string {
	if len(list) == 0 {
		return ""
	}
	wheres := " where "
	for _, where := range list {
		wheres += where.field + " " + where.operation + " ? and "
	}
	return wheres[:len(wheres)-5]
}
