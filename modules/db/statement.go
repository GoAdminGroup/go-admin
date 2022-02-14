// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	dbsql "database/sql"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"
)

// SQL wraps the Connection and driver dialect methods.
type SQL struct {
	dialect.SQLComponent
	diver   Connection
	dialect dialect.Dialect
	conn    string
	tx      *dbsql.Tx
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
				Order:      "",
				Group:      "",
				Limit:      "",
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

// TableName return a SQL with given table and default connection.
func Table(table string) *SQL {
	sql := newSQL()
	sql.TableName = table
	sql.conn = "default"
	return sql
}

// WithDriver return a SQL with given driver.
func WithDriver(conn Connection) *SQL {
	sql := newSQL()
	sql.diver = conn
	sql.dialect = dialect.GetDialectByDriver(conn.Name())
	sql.conn = "default"
	return sql
}

// WithDriverAndConnection return a SQL with given driver and connection name.
func WithDriverAndConnection(connName string, conn Connection) *SQL {
	sql := newSQL()
	sql.diver = conn
	sql.dialect = dialect.GetDialectByDriver(conn.Name())
	sql.conn = connName
	return sql
}

// WithDriver return a SQL with given driver.
func (sql *SQL) WithDriver(conn Connection) *SQL {
	sql.diver = conn
	sql.dialect = dialect.GetDialectByDriver(conn.Name())
	return sql
}

// WithConnection set the connection name of SQL.
func (sql *SQL) WithConnection(conn string) *SQL {
	sql.conn = conn
	return sql
}

// WithTx set the database transaction object of SQL.
func (sql *SQL) WithTx(tx *dbsql.Tx) *SQL {
	sql.tx = tx
	return sql
}

// TableName set table of SQL.
func (sql *SQL) Table(table string) *SQL {
	sql.clean()
	sql.TableName = table
	return sql
}

// Select set select fields.
func (sql *SQL) Select(fields ...string) *SQL {
	sql.Fields = fields
	sql.Functions = make([]string, len(fields))
	reg, _ := regexp.Compile(`(.*?)\((.*?)\)`)
	for k, field := range fields {
		res := reg.FindAllStringSubmatch(field, -1)
		if len(res) > 0 && len(res[0]) > 2 {
			sql.Functions[k] = res[0][1]
			sql.Fields[k] = res[0][2]
		}
	}
	return sql
}

// OrderBy set order fields.
func (sql *SQL) OrderBy(fields ...string) *SQL {
	if len(fields) == 0 {
		panic("wrong order field")
	}
	for i := 0; i < len(fields); i++ {
		if i == len(fields)-2 {
			sql.Order += " " + sql.wrap(fields[i]) + " " + fields[i+1]
			return sql
		}
		sql.Order += " " + sql.wrap(fields[i]) + " and "
	}
	return sql
}

// OrderByRaw set order by.
func (sql *SQL) OrderByRaw(order string) *SQL {
	if order != "" {
		sql.Order += " " + order
	}
	return sql
}

func (sql *SQL) GroupBy(fields ...string) *SQL {
	if len(fields) == 0 {
		panic("wrong group by field")
	}
	for i := 0; i < len(fields); i++ {
		if i == len(fields)-1 {
			sql.Group += " " + sql.wrap(fields[i])
		} else {
			sql.Group += " " + sql.wrap(fields[i]) + ","
		}
	}
	return sql
}

// GroupByRaw set group by.
func (sql *SQL) GroupByRaw(group string) *SQL {
	if group != "" {
		sql.Group += " " + group
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
		panic("wrong parameter")
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
		panic("wrong parameter")
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
		res    map[string]interface{}
		err    error
		driver = sql.diver.Name()
	)

	if res, err = sql.Select("count(*)").First(); err != nil {
		return 0, err
	}

	if driver == DriverPostgresql {
		return res["count"].(int64), nil
	} else if driver == DriverMssql {
		return res[""].(int64), nil
	}

	return res["count(*)"].(int64), nil
}

// Sum sum the value of given field.
func (sql *SQL) Sum(field string) (float64, error) {
	var (
		res map[string]interface{}
		err error
		key = "sum(" + sql.wrap(field) + ")"
	)
	if res, err = sql.Select("sum(" + field + ")").First(); err != nil {
		return 0, err
	}

	if res == nil {
		return 0, nil
	}

	if r, ok := res[key].(float64); ok {
		return r, nil
	} else if r, ok := res[key].([]uint8); ok {
		return strconv.ParseFloat(string(r), 64)
	} else {
		return 0, nil
	}
}

// Max find the maximal value of given field.
func (sql *SQL) Max(field string) (interface{}, error) {
	var (
		res map[string]interface{}
		err error
		key = "max(" + sql.wrap(field) + ")"
	)
	if res, err = sql.Select("max(" + field + ")").First(); err != nil {
		return 0, err
	}

	if res == nil {
		return 0, nil
	}

	return res[key], nil
}

// Min find the minimal value of given field.
func (sql *SQL) Min(field string) (interface{}, error) {
	var (
		res map[string]interface{}
		err error
		key = "min(" + sql.wrap(field) + ")"
	)
	if res, err = sql.Select("min(" + field + ")").First(); err != nil {
		return 0, err
	}

	if res == nil {
		return 0, nil
	}

	return res[key], nil
}

// Avg find the average value of given field.
func (sql *SQL) Avg(field string) (interface{}, error) {
	var (
		res map[string]interface{}
		err error
		key = "avg(" + sql.wrap(field) + ")"
	)
	if res, err = sql.Select("avg(" + field + ")").First(); err != nil {
		return 0, err
	}

	if res == nil {
		return 0, nil
	}

	return res[key], nil
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
// Transaction method
// *******************************

// TxFn is the transaction callback function.
type TxFn func(tx *dbsql.Tx) (error, map[string]interface{})

// WithTransaction call the callback function within the transaction and
// catch the error.
func (sql *SQL) WithTransaction(fn TxFn) (res map[string]interface{}, err error) {

	tx := sql.diver.BeginTxAndConnection(sql.conn)

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err, res = fn(tx)
	return
}

// WithTransactionByLevel call the callback function within the transaction
// of given transaction level and catch the error.
func (sql *SQL) WithTransactionByLevel(level dbsql.IsolationLevel, fn TxFn) (res map[string]interface{}, err error) {

	tx := sql.diver.BeginTxWithLevelAndConnection(sql.conn, level)

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err, res = fn(tx)
	return
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

	res, err := sql.diver.QueryWith(sql.tx, sql.conn, sql.Statement, sql.Args...)

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

	return sql.diver.QueryWith(sql.tx, sql.conn, sql.Statement, sql.Args...)
}

// ShowColumns show columns info.
func (sql *SQL) ShowColumns() ([]map[string]interface{}, error) {
	defer RecycleSQL(sql)

	return sql.diver.QueryWithConnection(sql.conn, sql.dialect.ShowColumns(sql.TableName))
}

// ShowTables show table info.
func (sql *SQL) ShowTables() ([]string, error) {
	defer RecycleSQL(sql)

	models, err := sql.diver.QueryWithConnection(sql.conn, sql.dialect.ShowTables())

	if err != nil {
		return []string{}, err
	}

	tables := make([]string, 0)
	if len(models) == 0 {
		return tables, nil
	}

	key := "Tables_in_" + sql.TableName
	if sql.diver.Name() == DriverPostgresql || sql.diver.Name() == DriverSqlite {
		key = "tablename"
	} else if sql.diver.Name() == DriverMssql {
		key = "TABLE_NAME"
	} else if _, ok := models[0][key].(string); !ok {
		key = "Tables_in_" + strings.ToLower(sql.TableName)
	}

	for i := 0; i < len(models); i++ {
		// skip sqlite system tables
		if sql.diver.Name() == DriverSqlite && models[i][key].(string) == "sqlite_sequence" {
			continue
		}

		tables = append(tables, models[i][key].(string))
	}

	return tables, nil
}

// Update exec the update method of given key/value pairs.
func (sql *SQL) Update(values dialect.H) (int64, error) {
	defer RecycleSQL(sql)

	sql.Values = values

	sql.dialect.Update(&sql.SQLComponent)

	res, err := sql.diver.ExecWith(sql.tx, sql.conn, sql.Statement, sql.Args...)

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

	res, err := sql.diver.ExecWith(sql.tx, sql.conn, sql.Statement, sql.Args...)

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

	res, err := sql.diver.ExecWith(sql.tx, sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

const postgresInsertCheckTableName = "goadmin_menu|goadmin_permissions|goadmin_roles|goadmin_users"

// Insert exec the insert method of given key/value pairs.
func (sql *SQL) Insert(values dialect.H) (int64, error) {
	defer RecycleSQL(sql)

	sql.Values = values

	sql.dialect.Insert(&sql.SQLComponent)

	if sql.diver.Name() == DriverPostgresql && (strings.Contains(postgresInsertCheckTableName, sql.TableName)) {

		resMap, err := sql.diver.QueryWith(sql.tx, sql.conn, sql.Statement+" RETURNING id", sql.Args...)

		if err != nil {

			// Fixed java h2 database postgresql mode
			_, err := sql.diver.QueryWith(sql.tx, sql.conn, sql.Statement, sql.Args...)

			if err != nil {
				return 0, err
			}

			res, err := sql.diver.QueryWithConnection(sql.conn, `SELECT max("id") as "id" FROM "`+sql.TableName+`"`)

			if err != nil {
				return 0, err
			}

			if len(res) != 0 {
				return res[0]["id"].(int64), nil
			}

			return 0, err
		}

		if len(resMap) == 0 {
			return 0, errors.New("no affect row")
		}

		return resMap[0]["id"].(int64), nil
	}

	res, err := sql.diver.ExecWith(sql.tx, sql.conn, sql.Statement, sql.Args...)

	if err != nil {
		return 0, err
	}

	if affectRow, _ := res.RowsAffected(); affectRow < 1 {
		return 0, errors.New("no affect row")
	}

	return res.LastInsertId()
}

func (sql *SQL) wrap(field string) string {
	return sql.diver.GetDelimiter() + field + sql.diver.GetDelimiter2()
}

func (sql *SQL) clean() {
	sql.Functions = make([]string, 0)
	sql.Group = ""
	sql.Values = make(map[string]interface{})
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
}

// RecycleSQL clear the SQL and put into the pool.
func RecycleSQL(sql *SQL) {

	logger.LogSQL(sql.Statement, sql.Args)

	sql.clean()

	sql.conn = ""
	sql.diver = nil
	sql.tx = nil
	sql.dialect = nil

	SQLPool.Put(sql)
}
