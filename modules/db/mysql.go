// Copyright 2019 GoAdmin Core Team.  All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type SqlTx struct {
	Tx *sql.Tx
}

type Mysql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

var MysqlDB = Mysql{
	DbList: map[string]*sql.DB{},
}

func GetMysqlDB() *Mysql {
	return &MysqlDB
}

func (db *Mysql) GetDelimiter() string {
	return "`"
}

func (db *Mysql) GetName() string {
	return "mysql"
}

func (db *Mysql) InitDB(cfgs map[string]config.Database) {
	db.Once.Do(func() {
		var (
			err   error
			SqlDB *sql.DB
		)

		for conn, cfg := range cfgs {
			SqlDB, err = sql.Open("mysql", cfg.User+
				":"+cfg.Pwd+"@tcp("+cfg.Host+":"+cfg.Port+")/"+cfg.Name+"?charset=utf8mb4")

			if err != nil {
				if SqlDB != nil {
					_ = SqlDB.Close()
				}
				panic(err.Error())
			} else {
				// Largest set up the database connection reduce time wait
				SqlDB.SetMaxIdleConns(cfg.MaxIdleCon)
				SqlDB.SetMaxOpenConns(cfg.MaxOpenCon)

				db.DbList[conn] = SqlDB
			}
		}
	})
}

func (db *Mysql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

func (db *Mysql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

func (db *Mysql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

func (db *Mysql) BeginTransactionsWithReadUncommitted() *SqlTx {
	return db.BeginTransactionsWithLevel(sql.LevelReadUncommitted)
}

func (db *Mysql) BeginTransactionsWithReadCommitted() *SqlTx {
	return db.BeginTransactionsWithLevel(sql.LevelReadCommitted)
}

func (db *Mysql) BeginTransactionsWithRepeatableRead() *SqlTx {
	return db.BeginTransactionsWithLevel(sql.LevelRepeatableRead)
}

func (db *Mysql) BeginTransactions() *SqlTx {
	return db.BeginTransactionsWithLevel(sql.LevelDefault)
}

func (db *Mysql) BeginTransactionsWithLevel(level sql.IsolationLevel) *SqlTx {
	tx, err := db.DbList["default"].BeginTx(context.Background(),
		&sql.TxOptions{Isolation: level})
	if err != nil {
		panic(err)
	}

	SqlTx := new(SqlTx)

	(*SqlTx).Tx = tx
	return SqlTx
}

func (SqlTx *SqlTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	rs, err := SqlTx.Tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	if rows, execError := rs.RowsAffected(); execError != nil || rows == 0 {
		return nil, errors.New("exec fail")
	}

	return rs, nil
}

func (SqlTx *SqlTx) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rs, err := SqlTx.Tx.Query(query, args...)

	if err != nil {
		return nil, err
	}

	col, colErr := rs.Columns()

	if colErr != nil {
		if closeErr := rs.Close(); closeErr != nil {
			panic(closeErr)
		}
		panic(colErr)
	}

	typeVal, err := rs.ColumnTypes()
	if err != nil {
		if closeErr := rs.Close(); closeErr != nil {
			panic(closeErr)
		}
		panic(err)
	}

	results := make([]map[string]interface{}, 0)

	for rs.Next() {
		var colVar = make([]interface{}, len(col))
		for i := 0; i < len(col); i++ {
			SetColVarType(&colVar, i, typeVal[i].DatabaseTypeName())
		}
		result := make(map[string]interface{})
		if scanErr := rs.Scan(colVar...); scanErr != nil {
			if closeErr := rs.Close(); closeErr != nil {
				panic(closeErr)
			}
			panic(scanErr)
		}
		for j := 0; j < len(col); j++ {
			SetResultValue(&result, col[j], colVar[j], typeVal[j].DatabaseTypeName())
		}
		results = append(results, result)
	}
	if err := rs.Err(); err != nil {
		if closeErr := rs.Close(); closeErr != nil {
			panic(closeErr)
		}
		panic(err)
	}
	return results, nil
}

type TxFn func(*SqlTx) (error, map[string]interface{})

func (db *Mysql) WithTransaction(fn TxFn) (err error, res map[string]interface{}) {

	SqlTx := db.BeginTransactions()

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = SqlTx.Tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = SqlTx.Tx.Rollback()
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
			_ = SqlTx.Tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = SqlTx.Tx.Rollback()
		} else {
			// all good, commit
			err = SqlTx.Tx.Commit()
		}
	}()

	err, res = fn(SqlTx)
	return
}
