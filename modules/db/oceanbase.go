package db

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/config"
)

// OceanBase is a Connection of OceanBase.
type OceanBase struct {
	Base
}

// GetOceanBaseDB return the global OceanBase connection.
func GetOceanBaseDB() *OceanBase {
	return &OceanBase{
		Base: Base{
			DbList: make(map[string]*sql.DB),
		},
	}
}

// Name implements the method Connection.Name.
func (db *OceanBase) Name() string {
	return "OceanBase"
}

// GetDelimiter implements the method Connection.GetDelimiter.
func (db *OceanBase) GetDelimiter() string {
	return "`"
}

// GetDelimiter2 implements the method Connection.GetDelimiter2.
func (db *OceanBase) GetDelimiter2() string {
	return "`"
}

// GetDelimiters implements the method Connection.GetDelimiters.
func (db *OceanBase) GetDelimiters() []string {
	return []string{"`", "`"}
}

// InitDB implements the method Connection.InitDB.
func (db *OceanBase) InitDB(cfgs map[string]config.Database) Connection {
	db.Configs = cfgs
	db.Once.Do(func() {
		for conn, cfg := range cfgs {

			sqlDB, err := sql.Open("mysql", cfg.GetDSN())

			if err != nil {
				if sqlDB != nil {
					_ = sqlDB.Close()
				}
				panic(err)
			}

			// Largest set up the database connection reduce time wait
			sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
			sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)
			sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)

			db.DbList[conn] = sqlDB

			if err := sqlDB.Ping(); err != nil {
				panic(err)
			}
		}
	})
	return db
}

// QueryWithConnection implements the method Connection.QueryWithConnection.
func (db *OceanBase) QueryWithConnection(con string, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList[con], query, args...)
}

// ExecWithConnection implements the method Connection.ExecWithConnection.
func (db *OceanBase) ExecWithConnection(con string, query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList[con], query, args...)
}

// Query implements the method Connection.Query.
func (db *OceanBase) Query(query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQuery(db.DbList["default"], query, args...)
}

// Exec implements the method Connection.Exec.
func (db *OceanBase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return CommonExec(db.DbList["default"], query, args...)
}

// QueryWithTx is query method within the transaction.
func (db *OceanBase) QueryWithTx(tx *sql.Tx, query string, args ...interface{}) ([]map[string]interface{}, error) {
	return CommonQueryWithTx(tx, query, args...)
}

// ExecWithTx is exec method within the transaction.
func (db *OceanBase) ExecWithTx(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return CommonExecWithTx(tx, query, args...)
}

func (db *OceanBase) QueryWith(tx *sql.Tx, conn, query string, args ...interface{}) ([]map[string]interface{}, error) {
	if tx != nil {
		return db.QueryWithTx(tx, query, args...)
	}
	return db.QueryWithConnection(conn, query, args...)
}

func (db *OceanBase) ExecWith(tx *sql.Tx, conn, query string, args ...interface{}) (sql.Result, error) {
	if tx != nil {
		return db.ExecWithTx(tx, query, args...)
	}
	return db.ExecWithConnection(conn, query, args...)
}

// BeginTxWithReadUncommitted starts a transaction with level LevelReadUncommitted.
func (db *OceanBase) BeginTxWithReadUncommitted() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelReadUncommitted)
}

// BeginTxWithReadCommitted starts a transaction with level LevelReadCommitted.
func (db *OceanBase) BeginTxWithReadCommitted() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelReadCommitted)
}

// BeginTxWithRepeatableRead starts a transaction with level LevelRepeatableRead.
func (db *OceanBase) BeginTxWithRepeatableRead() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelRepeatableRead)
}

// BeginTx starts a transaction with level LevelDefault.
func (db *OceanBase) BeginTx() *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], sql.LevelDefault)
}

// BeginTxWithLevel starts a transaction with given transaction isolation level.
func (db *OceanBase) BeginTxWithLevel(level sql.IsolationLevel) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList["default"], level)
}

// BeginTxWithReadUncommittedAndConnection starts a transaction with level LevelReadUncommitted and connection.
func (db *OceanBase) BeginTxWithReadUncommittedAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelReadUncommitted)
}

// BeginTxWithReadCommittedAndConnection starts a transaction with level LevelReadCommitted and connection.
func (db *OceanBase) BeginTxWithReadCommittedAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelReadCommitted)
}

// BeginTxWithRepeatableReadAndConnection starts a transaction with level LevelRepeatableRead and connection.
func (db *OceanBase) BeginTxWithRepeatableReadAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelRepeatableRead)
}

// BeginTxAndConnection starts a transaction with level LevelDefault and connection.
func (db *OceanBase) BeginTxAndConnection(conn string) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], sql.LevelDefault)
}

// BeginTxWithLevelAndConnection starts a transaction with given transaction isolation level and connection.
func (db *OceanBase) BeginTxWithLevelAndConnection(conn string, level sql.IsolationLevel) *sql.Tx {
	return CommonBeginTxWithLevel(db.DbList[conn], level)
}
