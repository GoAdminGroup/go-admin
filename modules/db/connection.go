// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/config"
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

// Connection is a connection handler of database.
type Connection interface {
	// Query is the query method of sql.
	Query(query string, args ...interface{}) ([]map[string]interface{}, error)

	// Exec is the exec method of sql.
	Exec(query string, args ...interface{}) (sql.Result, error)

	// QueryWithConnection is the query method with given connection of sql.
	QueryWithConnection(conn, query string, args ...interface{}) ([]map[string]interface{}, error)

	// ExecWithConnection is the exec method with given connection of sql.
	ExecWithConnection(conn, query string, args ...interface{}) (sql.Result, error)

	// InitDB initialize the database connections.
	InitDB(cfg map[string]config.Database)

	// GetName get the connection name.
	GetName() string

	// GetDelimiter get the default delimiter.
	GetDelimiter() string
}

// GetConnectionByDriver return the Connection by given driver name.
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

// GetConnection return the default Connection.
func GetConnection() Connection {
	return GetConnectionByDriver(config.Get().Databases.GetDefault().Driver)
}

// Query call the Query method of default Connection.
func Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return GetConnection().Query(query, args...)
}

// Exec call the Exec method of default Connection.
func Exec(query string, args ...interface{}) (sql.Result, error) {
	return GetConnection().Exec(query, args...)
}
