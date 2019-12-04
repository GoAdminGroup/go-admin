// Copyright 2019 GoAdmin Core Team. All rights reserved.
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

// Postgresql is a Connection of mssql.
type Postgresql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

// PostgresqlDB is a global variable which handles the postgresql connection.
var PostgresqlDB = Postgresql{
	DbList: map[string]*sql.DB{},
}

// GetPostgresqlDB return the global mssql connection.
func GetPostgresqlDB() *Postgresql {
	return &Postgresql{
		DbList: map[string]*sql.DB{},
	}
}

// Name implements the method Connection.Name.
func (db *Postgresql) Name() string {
	return "postgresql"
}

// GetDelimiter implements the method Connection.GetDelimiter.
func (db *Postgresql) GetDelimiter() string {
	return `"`
}

// QueryWithConnection implements the method Connection.QueryWithConnection.
func (db *Postgresql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], filterQuery(query), args...)
}

// ExecWithConnection implements the method Connection.ExecWithConnection.
func (db *Postgresql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], filterQuery(query), args...)
}

// Query implements the method Connection.Query.
func (db *Postgresql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], filterQuery(query), args...)
}

// Exec implements the method Connection.Exec.
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

// InitDB implements the method Connection.InitDB.
func (db *Postgresql) InitDB(cfgList map[string]config.Database) Connection {
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
	return db
}

// BeginTxWithReadUncommitted starts a transaction with level LevelReadUncommitted.
func (db *Postgresql) BeginTxWithReadUncommitted() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelReadUncommitted)
}

// BeginTxWithReadCommitted starts a transaction with level LevelReadCommitted.
func (db *Postgresql) BeginTxWithReadCommitted() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelReadCommitted)
}

// BeginTxWithRepeatableRead starts a transaction with level LevelRepeatableRead.
func (db *Postgresql) BeginTxWithRepeatableRead() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelRepeatableRead)
}

// BeginTx starts a transaction with level LevelDefault.
func (db *Postgresql) BeginTx() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelDefault)
}

// BeginTxWithLevel starts a transaction with given transaction isolation level.
func (db *Postgresql) BeginTxWithLevel(level sql.IsolationLevel) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], level)
}

// BeginTxWithReadUncommittedAndConnection starts a transaction with level LevelReadUncommitted and connection.
func (db *Postgresql) BeginTxWithReadUncommittedAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelReadUncommitted)
}

// BeginTxWithReadCommittedAndConnection starts a transaction with level LevelReadCommitted and connection.
func (db *Postgresql) BeginTxWithReadCommittedAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelReadCommitted)
}

// BeginTxWithRepeatableReadAndConnection starts a transaction with level LevelRepeatableRead and connection.
func (db *Postgresql) BeginTxWithRepeatableReadAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelRepeatableRead)
}

// BeginTxAndConnection starts a transaction with level LevelDefault and connection.
func (db *Postgresql) BeginTxAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelDefault)
}

// BeginTxWithLevelAndConnection starts a transaction with given transaction isolation level and connection.
func (db *Postgresql) BeginTxWithLevelAndConnection(conn string, level sql.IsolationLevel) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], level)
}

// QueryWithTx is query method within the transaction.
func (db *Postgresql) QueryWithTx(tx *sql.Tx, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQueryWithTx(tx, filterQuery(query), args...)
}

// ExecWithTx is exec method within the transaction.
func (db *Postgresql) ExecWithTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return CommonExecWithTx(tx, filterQuery(query), args...)
}
