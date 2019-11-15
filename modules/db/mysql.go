// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"sync"
)

// SQLTx is an in-progress database transaction.
type SQLTx struct {
	Tx *sql.Tx
}

// Mysql is a Connection of mssql.
type Mysql struct {
	DbList map[string]*sql.DB
	Once   sync.Once
}

// MysqlDB is a global variable which handles the mysql connection.
var MysqlDB = Mysql{
	DbList: map[string]*sql.DB{},
}

// GetMysqlDB return the global mssql connection.
func GetMysqlDB() *Mysql {
	return &MysqlDB
}

// GetDelimiter implements the method Connection.GetDelimiter.
func (db *Mysql) GetDelimiter() string {
	return "`"
}

// GetName implements the method Connection.GetName.
func (db *Mysql) GetName() string {
	return "mysql"
}

// InitDB implements the method Connection.InitDB.
func (db *Mysql) InitDB(cfgs map[string]config.Database) {
	db.Once.Do(func() {
		for conn, cfg := range cfgs {

			if cfg.Dsn == "" {
				cfg.Dsn = cfg.User + ":" + cfg.Pwd + "@tcp(" + cfg.Host + ":" + cfg.Port + ")/" + cfg.Name + "?charset=utf8mb4"
			}

			sqlDB, err := sql.Open("mysql", cfg.Dsn)

			if err != nil {
				if sqlDB != nil {
					_ = sqlDB.Close()
				}
				panic(err.Error())
			} else {
				// Largest set up the database connection reduce time wait
				sqlDB.SetMaxIdleConns(cfg.MaxIdleCon)
				sqlDB.SetMaxOpenConns(cfg.MaxOpenCon)

				db.DbList[conn] = sqlDB
			}
		}
	})
}

// QueryWithConnection implements the method Connection.QueryWithConnection.
func (db *Mysql) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

// ExecWithConnection implements the method Connection.ExecWithConnection.
func (db *Mysql) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

// Query implements the method Connection.Query.
func (db *Mysql) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

// Exec implements the method Connection.Exec.
func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

// BeginTransactionsWithReadUncommitted starts a transaction with level LevelReadUncommitted.
func (db *Mysql) BeginTransactionsWithReadUncommitted() *SQLTx {
	return db.BeginTransactionsWithLevel(sql.LevelReadUncommitted)
}

// BeginTransactionsWithReadCommitted starts a transaction with level LevelReadCommitted.
func (db *Mysql) BeginTransactionsWithReadCommitted() *SQLTx {
	return db.BeginTransactionsWithLevel(sql.LevelReadCommitted)
}

// BeginTransactionsWithRepeatableRead starts a transaction with level LevelRepeatableRead.
func (db *Mysql) BeginTransactionsWithRepeatableRead() *SQLTx {
	return db.BeginTransactionsWithLevel(sql.LevelRepeatableRead)
}

// BeginTransactions starts a transaction with level LevelDefault.
func (db *Mysql) BeginTransactions() *SQLTx {
	return db.BeginTransactionsWithLevel(sql.LevelDefault)
}

// BeginTransactionsWithLevel starts a transaction with given transaction isolation level.
func (db *Mysql) BeginTransactionsWithLevel(level sql.IsolationLevel) *SQLTx {
	tx, err := db.DbList["default"].BeginTx(context.Background(),
		&sql.TxOptions{Isolation: level})
	if err != nil {
		panic(err)
	}

	sqlTx := new(SQLTx)

	(*sqlTx).Tx = tx
	return sqlTx
}

// Exec is exec method within the transaction.
func (SqlTx *SQLTx) Exec(query string, args ...interface{}) (sql.Result, error) {
	rs, err := SqlTx.Tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	if rows, execError := rs.RowsAffected(); execError != nil || rows == 0 {
		return nil, errors.New("exec fail")
	}

	return rs, nil
}

// Query is query method within the transaction.
func (SqlTx *SQLTx) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
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

// TxFn is the transaction callback function.
type TxFn func(*SQLTx) (error, map[string]interface{})

// WithTransaction call the callback function within the transaction and
// catch the error.
func (db *Mysql) WithTransaction(fn TxFn) (res map[string]interface{}, err error) {

	tx := db.BeginTransactions()

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = tx.Tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Tx.Rollback()
		} else {
			// all good, commit
			err = tx.Tx.Commit()
		}
	}()

	err, res = fn(tx)
	return
}

// WithTransactionByLevel call the callback function within the transaction
// of given transaction level and catch the error.
func (db *Mysql) WithTransactionByLevel(level sql.IsolationLevel, fn TxFn) (res map[string]interface{}, err error) {

	tx := db.BeginTransactionsWithLevel(level)

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			_ = tx.Tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Tx.Rollback()
		} else {
			// all good, commit
			err = tx.Tx.Commit()
		}
	}()

	err, res = fn(tx)
	return
}
