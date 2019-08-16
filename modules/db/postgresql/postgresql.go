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
	SqlDBmap map[string]*sql.DB
	Once     sync.Once
}

func GetPostgresqlDB() *Postgresql {
	return &DB
}

func (db *Postgresql) GetName() string {
	return "postgresql"
}

func (db *Postgresql) Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	queCount := strings.Count(query, "?")
	for i := 1; i < queCount+1; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i), 1)
	}
	query = strings.Replace(query, "`", "", -1)
	// TODO: add " to the keyword
	query = strings.Replace(query, "by order ", `by "order" `, -1)
	return performer.Query(db.SqlDBmap["default"], query, args...)
}

func (db *Postgresql) Exec(query string, args ...interface{}) sql.Result {
	queCount := strings.Count(query, "?")
	for i := 1; i < queCount+1; i++ {
		query = strings.Replace(query, "?", "$"+strconv.Itoa(i), 1)
	}
	query = strings.Replace(query, "`", "", -1)
	// TODO: add " to the keyword
	query = strings.Replace(query, "by order ", `by "order" `, -1)
	fmt.Println("exec", query)
	return performer.Exec(db.SqlDBmap["default"], query, args...)
}

func (db *Postgresql) InitDB(cfgList map[string]config.Database) {
	db.Once.Do(func() {
		var (
			sqlDB *sql.DB
			err   error
		)

		for conn, cfg := range cfgList {
			connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				cfg.HOST, cfg.PORT, cfg.USER, cfg.PWD, cfg.NAME)

			sqlDB, err = sql.Open("postgres", connStr)
			if err != nil {
				panic(err)
			}

			db.SqlDBmap[conn] = sqlDB
		}
	})
}

var DB = Postgresql{
	SqlDBmap: map[string]*sql.DB{},
}
