package db

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mssql"
	"testing"
)

var msConn Connection

func init() {
	msConn = testConn(DriverMssql, config.Database{
		Host: "127.0.0.1",
		Port: "1433",
		User: "sa",
		Pwd:  "Aa123456",
		Name: "goadmin",
	})
}

func TestMssqlSQL_WhereIn(t *testing.T) { testSQLWhereIn(t, msConn) }
func TestMssqlSQL_Count(t *testing.T)   { testSQLCount(t, msConn) }
