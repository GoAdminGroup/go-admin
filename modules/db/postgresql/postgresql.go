// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package postgresql

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db/performer"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

func GetPostgresqlDB() *Postgresql {
	return &DB
}

func (db *Postgresql) GetName() string {
	return "postgresql"
}

func (db *Postgresql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return performer.Query(db.DbList[con], filterQuery(query), args...)
}

func (db *Postgresql) ExecWithConnection(con string, query string, args ...interface{}) sql.Result {
	return performer.Exec(db.DbList[con], filterQuery(query), args...)
}

func (db *Postgresql) Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return performer.Query(db.DbList["default"], filterQuery(query), args...)
}

func (db *Postgresql) Exec(query string, args ...interface{}) sql.Result {
	return performer.Exec(db.DbList["default"], filterQuery(query), args...)
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
		var (
			sqlDB *sql.DB
			err   error
		)

		for conn, cfg := range cfgList {
			connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name)

			sqlDB, err = sql.Open("postgres", connStr)
			if err != nil {
				panic(err)
			}

			db.DbList[conn] = sqlDB
		}
	})
}

var DB = Postgresql{
	DbList: map[string]*sql.DB{},
}
