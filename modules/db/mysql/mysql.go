// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db/converter"
	"github.com/chenhg5/go-admin/modules/db/performer"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type SqlTxStruct struct {
	Tx *sql.Tx
}

type Mysql struct {
	SqlDBmap map[string]*sql.DB
	Once     sync.Once
}

var DB = Mysql{
	SqlDBmap: map[string]*sql.DB{},
}

func GetMysqlDB() *Mysql {
	return &DB
}

func (db *Mysql) InitDB(cfglist map[string]config.Database) {
	db.Once.Do(func() {
		var (
			err   error
			SqlDB *sql.DB
		)

		for conn, cfg := range cfglist {
			SqlDB, err = sql.Open("mysql", cfg.USER+
				":"+cfg.PWD+"@tcp("+cfg.HOST+":"+cfg.PORT+")/"+cfg.NAME+"?charset=utf8mb4")

			if err != nil {
				SqlDB.Close()
				panic(err.Error())
			} else {
				// 设置数据库最大连接 减少timewait 正式环境调大
				SqlDB.SetMaxIdleConns(cfg.MAX_IDLE_CON) // 连接池连接数 = mysql最大连接数/2
				SqlDB.SetMaxOpenConns(cfg.MAX_OPEN_CON) // 最大打开连接 = mysql最大连接数

				db.SqlDBmap[conn] = SqlDB
			}
		}
	})
}

func (db *Mysql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {

	rs, err := db.SqlDBmap[con].Query(query, args...)

	if err != nil {
		if rs != nil {
			rs.Close()
		}
		panic(err)
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		if rs != nil {
			rs.Close()
		}
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		if rs != nil {
			rs.Close()
		}
		panic(err)
	}

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			converter.SetColVarType(&colVar, i, typeVal[i].DatabaseTypeName())
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			rs.Close()
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			converter.SetResultValue(&result, col[j], colVar[j], typeVal[j].DatabaseTypeName())
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		if rs != nil {
			rs.Close()
		}
		panic(err)
	}
	rs.Close()
	return results, rs
}

func (db *Mysql) ShowColumns(tableName string) ([]map[string]interface{}, *sql.Rows) {
	return db.Query("show columns in " + tableName)
}

func (db *Mysql) Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows) {
	return performer.Query(db.SqlDBmap["default"], query, args...)
}

func (db *Mysql) Exec(query string, args ...interface{}) sql.Result {
	return performer.Exec(db.SqlDBmap["default"], query, args...)
}

func (db *Mysql) BeginTransactionsWithReadUncommitted() *SqlTxStruct {
	return db.BeginTransactionsWithLevel(sql.LevelReadUncommitted)
}

func (db *Mysql) BeginTransactionsWithReadCommitted() *SqlTxStruct {
	return db.BeginTransactionsWithLevel(sql.LevelReadCommitted)
}

func (db *Mysql) BeginTransactionsWithRepeatableRead() *SqlTxStruct {
	return db.BeginTransactionsWithLevel(sql.LevelRepeatableRead)
}

func (db *Mysql) BeginTransactions() *SqlTxStruct {
	return db.BeginTransactionsWithLevel(sql.LevelDefault)
}

func (db *Mysql) BeginTransactionsWithLevel(level sql.IsolationLevel) *SqlTxStruct {
	tx, err := db.SqlDBmap["default"].BeginTx(context.Background(),
		&sql.TxOptions{Isolation: level})
	if err != nil {
		panic(err)
	}

	SqlTx := new(SqlTxStruct)

	(*SqlTx).Tx = tx
	return SqlTx
}

func (SqlTx *SqlTxStruct) Exec(query string, args ...interface{}) (sql.Result, error) {
	rs, err := SqlTx.Tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	if rows, execError := rs.RowsAffected(); execError != nil || rows == 0 {
		return nil, errors.New("exec fail")
	}

	return rs, nil
}

func (SqlTx *SqlTxStruct) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rs, err := SqlTx.Tx.Query(query, args...)

	if err != nil {
		return nil, err
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		rs.Close()
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		rs.Close()
		panic(err)
	}

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			converter.SetColVarType(&colVar, i, typeVal[i].DatabaseTypeName())
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			rs.Close()
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			converter.SetResultValue(&result, col[j], colVar[j], typeVal[j].DatabaseTypeName())
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		rs.Close()
		panic(err)
	}
	return results, nil
}

type TxFn func(*SqlTxStruct) (error, map[string]interface{})

func (db *Mysql) WithTransaction(fn TxFn) (err error, res map[string]interface{}) {

	SqlTx := db.BeginTransactions()

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			SqlTx.Tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			SqlTx.Tx.Rollback()
		} else {
			// all good, commit
			err = SqlTx.Tx.Commit()
		}
	}()

	err, res = fn(SqlTx)
	return
}

func (db *Mysql) WithTransactionByLevel(level sql.IsolationLevel, fn TxFn) (err error, res map[string]interface{}) {

	SqlTx := db.BeginTransactionsWithLevel(level)

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			SqlTx.Tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			SqlTx.Tx.Rollback()
		} else {
			// all good, commit
			err = SqlTx.Tx.Commit()
		}
	}()

	err, res = fn(SqlTx)
	return
}
