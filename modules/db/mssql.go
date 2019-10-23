// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/denisenkom/go-mssqldb"
	"net/url"
	"sync"
)

type Mssql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

var MssqlDB = Mssql{
	DbList: map[string]*sql.DB{},
}

func GetMssqlDB() *Mssql {
	return &MssqlDB
}

func (db *Mssql) GetDelimiter() string {
	return "`"
}

func (db *Mssql) GetName() string {
	return "mssql"
}

func (db *Mssql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

func (db *Mssql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

func (db *Mssql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

func (db *Mssql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

func (db *Mssql) InitDB(cfglist map[string]config.Database) {
	db.Once.Do(func() {
		var (
			err   error
			SqlDB *sql.DB
		)

		for conn, cfg := range cfglist {

			u := &url.URL{
				Scheme: "mssql",
				User:   url.UserPassword(cfg.User, cfg.Pwd),
				Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			}

			SqlDB, err = sql.Open("mssql", u.String())

			if SqlDB == nil {
				panic("invalid connection")
			}

			if err != nil {
				_ = SqlDB.Close()
				panic(err.Error())
			} else {
				// 设置数据库最大连接 减少timewait 正式环境调大
				SqlDB.SetMaxIdleConns(cfg.MaxIdleCon) // 连接池连接数 = mysql最大连接数/2
				SqlDB.SetMaxOpenConns(cfg.MaxOpenCon) // 最大打开连接 = mysql最大连接数

				db.DbList[conn] = SqlDB
			}
		}
	})
}
