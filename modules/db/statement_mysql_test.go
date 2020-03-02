package db

import (
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"testing"
)

var mysqlConn Connection

func init() {
	mysqlConn = testConn(DriverMysql, config.Database{
		Host: "127.0.0.1",
		Port: "3306",
		User: "root",
		Pwd:  "root",
		Name: "goadmin",
	})
}

func TestMysqlSQL_WhereIn(t *testing.T) { testSQLWhereIn(t, mysqlConn) }
func TestMysqlSQL_Count(t *testing.T)   { testSQLCount(t, mysqlConn) }
