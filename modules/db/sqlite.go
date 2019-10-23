// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type Sqlite struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

func GetSqliteDB() *Sqlite {
	return &DB
}

func (db *Sqlite) GetName() string {
	return "sqlite"
}

func (db *Sqlite) GetDelimiter() string {
	return "`"
}

func (db *Sqlite) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

func (db *Sqlite) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

func (db *Sqlite) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

func (db *Sqlite) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

func (db *Sqlite) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		var (
			sqlDB *sql.DB
			err   error
		)

		for conn, cfg := range cfgList {
			sqlDB, err = sql.Open("sqlite3", cfg.File)

			if err != nil {
				panic(err)
			} else {
				db.DbList[conn] = sqlDB
			}
		}
	})
}

var DB = Sqlite{
	DbList: map[string]*sql.DB{},
}
