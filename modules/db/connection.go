// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/service"
)

const (
	// DriverMysql is a const value of mysql driver.
	DriverMysql = "mysql"
	// DriverSqlite is a const value of sqlite driver.
	DriverSqlite = "sqlite"
	// DriverPostgresql is a const value of postgresql driver.
	DriverPostgresql = "postgresql"
	// DriverMssql is a const value of mssql driver.
	DriverMssql = "mssql"
)

var (
	driverConnFuncMap      = sync.Map{}
	driverGetAggExpFuncMap = sync.Map{}
)

func init() {
	driverConnFuncMap.Store("mysql", GetMysqlDB)
	driverConnFuncMap.Store("mssql", GetMssqlDB)
	driverConnFuncMap.Store("sqlite", GetSqliteDB)
	driverConnFuncMap.Store("postgresql", GetPostgresqlDB)
}

//. LoadCustomDriver 加载自定义 Connection 实现
func LoadCustomDriver(driver string, connFunc ConnFunc, getAggExpFunc GetAggregationExpressionFunc) error {
	if _, ok := driverConnFuncMap.Load(driver); ok {
		return fmt.Errorf("driver exists!")
	}
	driverConnFuncMap.Store(driver, connFunc)
	driverGetAggExpFuncMap.Store(driver, getAggExpFunc)
	return nil
}

type ConnFunc func() Connection
type GetAggregationExpressionFunc func(driver, field, headField, delimiter string) string

// Connection is a connection handler of database.
type Connection interface {
	// Query is the query method of sql.
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)

	// Exec is the exec method of sql.
	Exec(query string, args ...interface{}) (sql.Result, error)

	// QueryWithConnection is the query method with given connection of sql.
	QueryWithConnection(conn, query string, args ...interface{}) ([]map[string]interface{}, error)
	QueryWithTx(tx *sql.Tx, query string, args ...interface{}) ([]map[string]interface{}, error)
	QueryWith(tx *sql.Tx, conn, query string, args ...interface{}) ([]map[string]interface{}, error)

	// ExecWithConnection is the exec method with given connection of sql.
	ExecWithConnection(conn, query string, args ...interface{}) (sql.Result, error)
	ExecWithTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error)
	ExecWith(tx *sql.Tx, conn, query string, args ...interface{}) (sql.Result, error)

	BeginTxWithReadUncommitted() *sql.Tx
	BeginTxWithReadCommitted() *sql.Tx
	BeginTxWithRepeatableRead() *sql.Tx
	BeginTx() *sql.Tx
	BeginTxWithLevel(level sql.IsolationLevel) *sql.Tx

	BeginTxWithReadUncommittedAndConnection(conn string) *sql.Tx
	BeginTxWithReadCommittedAndConnection(conn string) *sql.Tx
	BeginTxWithRepeatableReadAndConnection(conn string) *sql.Tx
	BeginTxAndConnection(conn string) *sql.Tx
	BeginTxWithLevelAndConnection(conn string, level sql.IsolationLevel) *sql.Tx

	// InitDB initialize the database connections.
	InitDB(cfg map[string]config.Database) Connection

	// GetName get the connection name.
	Name() string

	Close() []error

	GetDelimiter() string
	GetDelimiter2() string
	GetDelimiters() []string

	GetDB(key string) *sql.DB

	GetConfig(name string) config.Database

	CreateDB(name string, beans ...interface{}) error
}

// GetConnectionByDriver return the Connection by given driver name.
func GetConnectionByDriver(driver string) Connection {
	connFunc, ok := driverConnFuncMap.Load(driver)
	if !ok {
		panic("driver not found!")
	}
	return connFunc.(func() Connection)()
}

func GetConnectionFromService(srv interface{}) Connection {
	if v, ok := srv.(Connection); ok {
		return v
	}
	panic("wrong service")
}

func GetConnection(srvs service.List) Connection {
	if v, ok := srvs.Get(config.GetDatabases().GetDefault().Driver).(Connection); ok {
		return v
	}
	panic("wrong service")
}

func GetAggregationExpression(driver, field, headField, delimiter string) string {
	switch driver {
	case "postgresql":
		return fmt.Sprintf("string_agg(%s::character varying, '%s') as %s", field, delimiter, headField)
	case "mysql":
		return fmt.Sprintf("group_concat(%s separator '%s') as %s", field, delimiter, headField)
	case "sqlite":
		return fmt.Sprintf("group_concat(%s, '%s') as %s", field, delimiter, headField)
	case "mssql":
		return fmt.Sprintf("string_agg(%s, '%s') as [%s]", field, delimiter, headField)
	default:
		if f, ok := driverGetAggExpFuncMap.Load(driver); ok {
			return f.(GetAggregationExpressionFunc)(driver, field, headField, delimiter)
		}
		panic("wrong driver")
	}
}

const (
	INSERT = 0
	DELETE = 1
	UPDATE = 2
	QUERY  = 3
)

var ignoreErrors = [...][]string{
	// insert
	{
		"LastInsertId is not supported",
		"There is no generated identity value",
		"LastInsertId is not supported by this driver",
	},
	// delete
	{
		"no affect",
	},
	// update
	{
		"LastInsertId is not supported",
		"There is no generated identity value",
		"no affect",
		"LastInsertId is not supported by this driver",
	},
	// query
	{
		"LastInsertId is not supported",
		"There is no generated identity value",
		"no affect",
		"out of index",
		"LastInsertId is not supported by this driver",
	},
}

func CheckError(err error, t int) bool {
	if err == nil {
		return false
	}
	for _, msg := range ignoreErrors[t] {
		if strings.Contains(err.Error(), msg) {
			return false
		}
	}
	return true
}
