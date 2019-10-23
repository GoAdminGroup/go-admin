// Copyright 2019 GoAdmin Core Team.  All rights reserved.
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

type Sql struct {
	dialect.SqlComponent
	diver   Connection
	dialect dialect.Dialect
	conn    string
}

var SqlPool = sync.Pool{
	New: func() interface{} {
		return &Sql{
			SqlComponent: dialect.SqlComponent{
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

type H map[string]interface{}

func newSql() *Sql {
	return SqlPool.Get().(*Sql)
}

// *******************************
// process method
// *******************************

func Table(table string) *Sql {
	sql := newSql()
	sql.TableName = table
	sql.diver = GetConnection()
	sql.dialect = dialect.GetDialect()
	sql.conn = "default"
	return sql
}

func WithDriver(driver string) *Sql {
	sql := newSql()
	sql.diver = GetConnectionByDriver(driver)
	sql.dialect = dialect.GetDialectByDriver(driver)
	sql.conn = "default"
	return sql
}

func WithDriverAndConnection(conn, driver string) *Sql {
	sql := newSql()
	sql.diver = GetConnectionByDriver(driver)
	sql.dialect = dialect.GetDialectByDriver(driver)
	sql.conn = conn
	return sql
}

func (sql *Sql) WithConnection(conn string) *Sql {
	sql.conn = conn
	return sql
}

func (sql *Sql) Table(table string) *Sql {
	sql.TableName = table
	return sql
}

func (sql *Sql) Select(fields ...string) *Sql {
	sql.Fields = fields
	return sql
}

func (sql *Sql) OrderBy(fields ...string) *Sql {
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

func (sql *Sql) Skip(offset int) *Sql {
	sql.Offset = strconv.Itoa(offset)
	return sql
}

func (sql *Sql) Take(take int) *Sql {
	sql.Limit = strconv.Itoa(take)
	return sql
}

func (sql *Sql) Where(field string, operation string, arg interface{}) *Sql {
	sql.Wheres = append(sql.Wheres, dialect.Where{
		Field:     field,
		Operation: operation,
		Qmark:     "?",
	})
	sql.Args = append(sql.Args, arg)
	return sql
}

func (sql *Sql) WhereIn(field string, arg []interface{}) *Sql {
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

func (sql *Sql) WhereNotIn(field string, arg []interface{}) *Sql {
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

func (sql *Sql) Find(arg interface{}) (map[string]interface{}, error) {
	return sql.Where("id", "=", arg).First()
}

func (sql *Sql) Count() (int64, error) {
	var (
		res map[string]interface{}
		err error
	)
	if res, err = sql.Select("count(*)").First(); err != nil {
		return 0, err
	}
	return res["count(*)"].(int64), nil
}

func (sql *Sql) WhereRaw(raw string, args ...interface{}) *Sql {
	sql.WhereRaws = raw
	sql.Args = append(sql.Args, args...)
	return sql
}

func (sql *Sql) UpdateRaw(raw string, args ...interface{}) *Sql {
	sql.UpdateRaws = append(sql.UpdateRaws, dialect.RawUpdate{
		Expression: raw,
		Args:       args,
	})
	return sql
}

func (sql *Sql) LeftJoin(table string, fieldA string, operation string, fieldB string) *Sql {
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

func (sql *Sql) First() (map[string]interface{}, error) {
	defer RecycleSql(sql)

	sql.dialect.Select(&sql.SqlComponent)

	res, err := sql.diver.QueryWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return nil, err
	}

	if len(res) < 1 {
		return nil, errors.New("out of index")
	}
	return res[0], nil
}

func (sql *Sql) All() ([]map[string]interface{}, error) {
	defer RecycleSql(sql)

	sql.dialect.Select(&sql.SqlComponent)

	return sql.diver.QueryWithConnection(sql.conn, sql.Statement, sql.Args...)
}

func (sql *Sql) ShowColumns() ([]map[string]interface{}, error) {
	defer RecycleSql(sql)

	return sql.diver.QueryWithConnection(sql.conn, sql.dialect.ShowColumns(sql.TableName))
}

func (sql *Sql) ShowTables() ([]map[string]interface{}, error) {
	defer RecycleSql(sql)

	return sql.diver.QueryWithConnection(sql.conn, sql.dialect.ShowTables())
}

func (sql *Sql) Update(values dialect.H) (int64, error) {
	defer RecycleSql(sql)

	sql.Values = values

	sql.dialect.Update(&sql.SqlComponent)

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

func (sql *Sql) Delete() error {
	defer RecycleSql(sql)

	sql.dialect.Delete(&sql.SqlComponent)

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return errors.New("no affect row")
	}

	return nil
}

func (sql *Sql) Exec() (int64, error) {
	defer RecycleSql(sql)

	sql.dialect.Update(&sql.SqlComponent)

	res, err := sql.diver.ExecWithConnection(sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

func (sql *Sql) Insert(values dialect.H) (int64, error) {
	defer RecycleSql(sql)

	sql.Values = values

	sql.dialect.Insert(&sql.SqlComponent)

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

func (sql *Sql) filed(filed string) string {
	return sql.diver.GetDelimiter() + filed + sql.diver.GetDelimiter()
}

func RecycleSql(sql *Sql) {

	logger.LogSql(sql.Statement, sql.Args)

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

	SqlPool.Put(sql)
}
