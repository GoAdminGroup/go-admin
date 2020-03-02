package db

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
	"testing"
)

var sqliteConn Connection

func init() {
	pgConn = testConn(DriverSqlite, config.Database{File: "/admin.db"})
}

func TestSQLiteSQL_WhereIn(t *testing.T) { testSQLWhereIn(t, sqliteConn) }
func TestSQLiteSQL_Count(t *testing.T)   { testSQLCount(t, sqliteConn) }
