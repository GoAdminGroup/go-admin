// Copyright 2019 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"github.com/chenhg5/go-admin/modules/config"
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
		return GetMysqlDB()
	case "mssql":
		return GetMssqlDB()
	case "sqlite":
		return GetSqliteDB()
	case "postgresql":
		return GetPostgresqlDB()
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
