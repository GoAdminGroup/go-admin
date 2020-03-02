package db

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"testing"
)

var pgConn Connection

func init() {
	pgConn = testConn(DriverPostgresql, config.Database{
		Host: "127.0.0.1",
		Port: "5433",
		User: "postgres",
		Pwd:  "root",
		Name: "godmin",
	})
}

func TestPgSQL_WhereIn(t *testing.T) { testSQLWhereIn(t, pgConn) }
func TestPgSQL_Count(t *testing.T)   { testSQLCount(t, pgConn) }
