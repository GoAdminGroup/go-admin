// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"errors"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"strconv"
	"strings"
	"sync"
)

// SQL wraps the Connection and driver dialect methods.
type SQL struct {
	dialect.SQLComponent
	diver   Connection
	dialect dialect.Dialect
	conn    string
}

// SQLPool is a object pool of SQL.
var SQLPool = sync.Pool{
	New: func() interface{} {
		return &SQL{
			SQLComponent: dialect.SQLComponent{
				Fields:     make([]string, 0),
				TableName:  "",
				Args:       make([]interface{}, 0),
				Wheres:     make([]dialect.Where, 0),
				Leftjoins:  make([]dialect.Join, 0),
				UpdateRaws: make([]dialect.RawUpdate, 0),
				WhereRaws:  "",
			},
			diver:   nil,
			dialect: nil,
		}
	},
}

// H is a shorthand of map.
type H map[string]interface{}

// newSQL get a new SQL from SQLPool.
func newSQL() *SQL {
	return SQLPool.Get().(*SQL)
}

// *******************************
// process method
// *******************************

// Table return a SQL with given table and default connection.
func Table(table string) *SQL {
	sql := newSQL()
	sql.TableName = table
	sql.diver = GetConnection()
	sql.dialect = dialect.GetDialect()
	sql.conn = "default"
	return sql
}

// WithDriver return a SQL with given driver.
func WithDriver(driver string) *SQL {
	sql := newSQL()
	sql.diver = GetConnectionByDriver(driver)
	sql.dialect = dialect.GetDialectByDriver(driver)
	sql.conn = "default"
	return sql
}

// WithDriverAndConnection return a SQL with given driver and connection name.
func WithDriverAndConnection(conn, driver string) *SQL {
	sql := newSQL()
	sql.diver = GetConnectionByDriver(driver)
	sql.dialect = dialect.GetDialectByDriver(driver)
	sql.conn = conn
	return sql
}

// WithConnection set the connection name of SQL.
func (sql *SQL) WithConnection(conn string) *SQL {
	sql.conn = conn
	return sql
}

// Table set table of SQL.
func (sql *SQL) Table(table string) *SQL {
	sql.TableName = table
	return sql
}

// Select set select fields.
func (sql *SQL) Select(fields ...string) *SQL {
	sql.Fields = fields
	return sql
}

// OrderBy set order fields.
func (sql *SQL) OrderBy(fields ...string) *SQL {
	if len(fields) == 0 {
		panic("wrong order field")
	}
	for i := 0; i < len(fields); i++ {
		if i == len(fields)-2 {
			sql.Order += " " + sql.filed(fields[i]) + " " + fields[i+1]
			return sql
		}
		sql.Order += " " + sql.filed(fields[i]) + " and "
	}
	return sql
}

// Skip set offset value.
func (sql *SQL) Skip(offset int) *SQL {
	sql.Offset = strconv.Itoa(offset)
	return sql
}

// Take set limit value.
func (sql *SQL) Take(take int) *SQL {
	sql.Limit = strconv.Itoa(take)
	return sql
}

// Where add the where operation and argument value.
func (sql *SQL) Where(field string, operation string, arg interface{}) *SQL {
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: operation,
		Qmark:     "?",
	})
	sql.Args = append(sql.Args, arg)
	return sql
}

// WhereIn add the where operation of "in" and argument values.
func (sql *SQL) WhereIn(field string, arg []interface{}) *SQL {
	if len(arg) == 0 {
		return sql
	}
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: "in",
		Qmark:     "(" + strings.Repeat("?,", len(arg)-1) + "?)",
	})
	sql.Args = append(sql.Args, arg...)
	return sql
}

// WhereNotIn add the where operation of "not in" and argument values.
func (sql *SQL) WhereNotIn(field string, arg []interface{}) *SQL {
	if len(arg) == 0 {
		return sql
	}
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: "not in",
		Qmark:     "(" + strings.Repeat("?,", len(arg)-1) + "?)",
	})
	sql.Args = append(sql.Args, arg...)
	return sql
}

// Find query the sql result with given id assuming that primary key name is "id".
func (sql *SQL) Find(arg interface{}) (map[string]interface{}, error) {
	return sql.Where("id", "=", arg).First()
}

// Count query the count of query results.
func (sql *SQL) Count() (int64, error) {
	var (
		res map[string]interface{}
		err error
	)
	if res, err = sql.Select("count(*)").First(); err != nil {
		return 0, err
	}
	return res["count(*)"].(int64), nil
}

// WhereRaw set WhereRaws and arguments.
func (sql *SQL) WhereRaw(raw string, args ...interface{}) *SQL {
	sql.WhereRaws = raw
	sql.Args = append(sql.Args, args...)
	return sql
}

// UpdateRaw set UpdateRaw.
func (sql *SQL) UpdateRaw(raw string, args ...interface{}) *SQL {
	sql.UpdateRaws = append(sql.UpdateRaws, dialect.RawUpdate{
		Expression: raw,
		Args:       args,
	})
	return sql
}

// LeftJoin add a left join info.
func (sql *SQL) LeftJoin(table string, fieldA string, operation string, fieldB string) *SQL {
	sql.Leftjoins = append(sql.Leftjoins, dialect.Join{
		FieldA:    fieldA,
		FieldB:    fieldB,
		Table:     table,
		Operation: operation,
	})
	return sql
}

// *******************************
// terminal method
// -------------------------------
// sql args order:
// update ... => where ...
// *******************************

// First query the result and return the first row.
func (sql *SQL) First() (map[string]interface{}, error) {
	defer RecycleSQL(sql)

	sql.dialect.Select(&sql.SQLComponent)

	res, err := sql.diver.QueryWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return nil, err
	}

	if len(res) < 1 {
		return nil, errors.New("out of index")
	}
	return res[0], nil
}

// All query all the result and return.
func (sql *SQL) All() ([]map[string]interface{}, error) {
	defer RecycleSQL(sql)

	sql.dialect.Select(&sql.SQLComponent)

	return sql.diver.QueryWithConnection(sql.conn, sql.Statement, sql.Args...)
}

// ShowColumns show columns info.
func (sql *SQL) ShowColumns() ([]map[string]interface{}, error) {
	defer RecycleSQL(sql)

	return sql.diver.QueryWithConnection(sql.conn, sql.dialect.ShowColumns(sql.TableName))
}

// ShowTables show table info.
func (sql *SQL) ShowTables() ([]map[string]interface{}, error) {
	defer RecycleSQL(sql)

	return sql.diver.QueryWithConnection(sql.conn, sql.dialect.ShowTables())
}

// Update exec the update method of given key/value pairs.
func (sql *SQL) Update(values dialect.H) (int64, error) {
	defer RecycleSQL(sql)

	sql.Values = values

	sql.dialect.Update(&sql.SQLComponent)

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

// Delete exec the delete method.
func (sql *SQL) Delete() error {
	defer RecycleSQL(sql)

	sql.dialect.Delete(&sql.SQLComponent)

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return errors.New("no affect row")
	}

	return nil
}

// Exec exec the exec method.
func (sql *SQL) Exec() (int64, error) {
	defer RecycleSQL(sql)

	sql.dialect.Update(&sql.SQLComponent)

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

// Insert exec the insert method of given key/value pairs.
func (sql *SQL) Insert(values dialect.H) (int64, error) {
	defer RecycleSQL(sql)

	sql.Values = values

	sql.dialect.Insert(&sql.SQLComponent)

	if sql.diver.GetName() == DriverPostgresql {
		if sql.TableName == "goadmin_menu" ||
			sql.TableName == "goadmin_permissions" ||
			sql.TableName == "goadmin_roles" ||
			sql.TableName == "goadmin_users" {
			res, err := sql.diver.QueryWithConnection(sql.conn, sql.Statement+" RETURNING id", sql.Args...)

			if err != nil {
				return 0, err
			}

			if len(res) == 0 {
				return 0, errors.New("no affect row")
			}
			return res[0]["id"].(int64), nil
		}
	}

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

func (sql *SQL) filed(filed string) string {
	return sql.diver.GetDelimiter() + filed + sql.diver.GetDelimiter()
}

// RecycleSQL clear the SQL and put into the pool.
func RecycleSQL(sql *SQL) {

	logger.LogSQL(sql.Statement, sql.Args)

	sql.Fields = make([]string, 0)
	sql.TableName = ""
	sql.Wheres = make([]dialect.Where, 0)
	sql.Leftjoins = make([]dialect.Join, 0)
	sql.Args = make([]interface{}, 0)
	sql.Order = ""
	sql.Offset = ""
	sql.Limit = ""
	sql.WhereRaws = ""
	sql.UpdateRaws = make([]dialect.RawUpdate, 0)
	sql.Statement = ""

	SQLPool.Put(sql)
}
