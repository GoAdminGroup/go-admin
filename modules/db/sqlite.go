// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"sync"
)

// Sqlite is a Connection of mssql.
type Sqlite struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

// DB is a global variable which handles the sqlite connection.
var DB = Sqlite{
	DbList: map[string]*sql.DB{},
}

// GetSqliteDB return the global mssql connection.
func GetSqliteDB() *Sqlite {
	return &DB
}

// GetName implements the method Connection.GetName.
func (db *Sqlite) GetName() string {
	return "sqlite"
}

// GetDelimiter implements the method Connection.GetDelimiter.
func (db *Sqlite) GetDelimiter() string {
	return "`"
}

// QueryWithConnection implements the method Connection.QueryWithConnection.
func (db *Sqlite) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

// ExecWithConnection implements the method Connection.ExecWithConnection.
func (db *Sqlite) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

// Query implements the method Connection.Query.
func (db *Sqlite) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

// Exec implements the method Connection.Exec.
func (db *Sqlite) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

// InitDB implements the method Connection.InitDB.
func (db *Sqlite) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		for conn, cfg := range cfgList {
			sqlDB, err := sql.Open("sqlite3", cfg.File)

			if err != nil {
				panic(err)
			} else {
				db.DbList[conn] = sqlDB
			}
		}
	})
}
