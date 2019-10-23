// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/config"
)

const (
	DriverMysql      = "mysql"
	DriverMssql      = "mssql"
	DriverSqlite     = "sqlite"
	DriverPostgresql = "postgresql"
)

type Connection interface {
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryWithConnection(conn, query string, args ...interface{}) ([]map[string]interface{}, error)
	ExecWithConnection(conn, query string, args ...interface{}) (sql.Result, error)
	InitDB(cfg map[string]config.Database)
	GetName() string
	GetDelimiter() string
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

func Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return GetConnection().Query(query, args...)
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return GetConnection().Exec(query, args...)
}
