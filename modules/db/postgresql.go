// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/GoAdminGroup/go-admin/modules/config"
)

type Postgresql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

func GetPostgresqlDB() *Postgresql {
	return &PostgresqlDB
}

func (db *Postgresql) GetName() string {
	return "postgresql"
}

func (db *Postgresql) GetDelimiter() string {
	return `"`
}

func (db *Postgresql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], filterQuery(query), args...)
}

func (db *Postgresql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], filterQuery(query), args...)
}

func (db *Postgresql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], filterQuery(query), args...)
}

func (db *Postgresql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], filterQuery(query), args...)
}

func filterQuery(query string) string {
	queCount := strings.Count(query, "?")
	for i := 1; i < queCount+1; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i), 1)
	}
	query = strings.Replace(query, "`", "", -1)
	// TODO: add " to the keyword
	return strings.Replace(query, "by order ", `by "order" `, -1)
}

func (db *Postgresql) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		for conn, cfg := range cfgList {

			if cfg.Dsn == "" {
				cfg.Dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
					cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name)
			}

			sqlDB, err := sql.Open("postgres", cfg.Dsn)
			if err != nil {
				panic(err)
			}

			db.DbList[conn] = sqlDB
		}
	})
}

var PostgresqlDB = Postgresql{
	DbList: map[string]*sql.DB{},
}
