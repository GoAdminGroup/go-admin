// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"net/url"
	"sync"
)

// Mssql is a Connection of mssql.
type Mssql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

// MssqlDB is a global variable which handles the mssql connection.
var MssqlDB = Mssql{
	DbList: map[string]*sql.DB{},
}

// GetMssqlDB return the global mssql connection.
func GetMssqlDB() *Mssql {
	return &MssqlDB
}

// GetDelimiter implements the method Connection.GetDelimiter.
func (db *Mssql) GetDelimiter() string {
	return "`"
}

// GetName implements the method Connection.GetName.
func (db *Mssql) GetName() string {
	return "mssql"
}

// QueryWithConnection implements the method Connection.QueryWithConnection.
func (db *Mssql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

// ExecWithConnection implements the method Connection.ExecWithConnection.
func (db *Mssql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

// Query implements the method Connection.Query.
func (db *Mssql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

// Exec implements the method Connection.Exec.
func (db *Mssql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

// InitDB implements the method Connection.InitDB.
func (db *Mssql) InitDB(cfglist map[string]config.Database) {
	db.Once.Do(func() {
		for conn, cfg := range cfglist {

			if cfg.Dsn == "" {
				u := &url.URL{
					Scheme: "mssql",
					User:   url.UserPassword(cfg.User, cfg.Pwd),
					Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
				}
				cfg.Dsn = u.String()
			}

			sqlDB, err := sql.Open("mssql", cfg.Dsn)

			if sqlDB == nil {
				panic("invalid connection")
			}

			if err != nil {
				_ = sqlDB.Close()
				panic(err.Error())
			} else {
				sqlDB.SetMaxIdleConns(cfg.MaxIdleCon)
				sqlDB.SetMaxOpenConns(cfg.MaxOpenCon)

				db.DbList[conn] = sqlDB
			}
		}
	})
}
