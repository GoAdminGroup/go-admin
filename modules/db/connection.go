// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db/mssql"
	"github.com/chenhg5/go-admin/modules/db/mysql"
	"github.com/chenhg5/go-admin/modules/db/postgresql"
	"github.com/chenhg5/go-admin/modules/db/sqlite"
)

const (
	DriverMysql      = "mysql"
	DriverMssql      = "mssql"
	DriverSqlite     = "sqlite"
	DriverPostgresql = "postgresql"
)

type Connection interface {
	Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows)
	Exec(query string, args ...interface{}) sql.Result
	QueryWithConnection(conn, query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows)
	ExecWithConnection(conn, query string, args ...interface{}) sql.Result
	InitDB(cfg map[string]config.Database)
	GetName() string
}

func GetConnectionByDriver(driver string) Connection {
	switch driver {
	case "mysql":
		return mysql.GetMysqlDB()
	case "mssql":
		return mssql.GetMssqlDB()
	case "sqlite":
		return sqlite.GetSqliteDB()
	case "postgresql":
		return postgresql.GetPostgresqlDB()
	default:
		panic("driver not found!")
	}
}

func GetConnection() Connection {
	return GetConnectionByDriver(config.Get().Databases.GetDefault().Driver)
}

func Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return GetConnection().Query(query, args...)
}

func Exec(query string, args ...interface{}) sql.Result {
	return GetConnection().Exec(query, args...)
}
