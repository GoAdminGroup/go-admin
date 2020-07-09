package db

import (
	"fmt"
	"os/exec"
	"testing"

	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
)

var driverTestMysqlConn Connection

const (
	driverTestDBName = "go-admin-statement-test"
)

func InitMysql() {
	c := testConnDSN(DriverMysql, fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s", driverTestDBName))
	_, err := c.Exec(fmt.Sprintf("create database if not exists `%s`", driverTestDBName))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("mysql", "-u", "root", "-proot", driverTestDBName,
		"-e", "source "+testCurrentPath()+"/../../data/admin.sql")
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	driverTestMysqlConn = testConnDSN(DriverMysql, fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/%s", driverTestDBName))
}

func TestMysqlSQL_WhereIn(t *testing.T)         { testSQLWhereIn(t, driverTestMysqlConn) }
func TestMysqlSQL_Count(t *testing.T)           { testSQLCount(t, driverTestMysqlConn) }
func TestMysqlSQL_Select(t *testing.T)          { testSQLSelect(t, driverTestMysqlConn) }
func TestMysqlSQL_OrderBy(t *testing.T)         { testSQLOrderBy(t, driverTestMysqlConn) }
func TestMysqlSQL_GroupBy(t *testing.T)         { testSQLGroupBy(t, driverTestMysqlConn) }
func TestMysqlSQL_Skip(t *testing.T)            { testSQLSkip(t, driverTestMysqlConn) }
func TestMysqlSQL_Take(t *testing.T)            { testSQLTake(t, driverTestMysqlConn) }
func TestMysqlSQL_Where(t *testing.T)           { testSQLWhere(t, driverTestMysqlConn) }
func TestMysqlSQL_WhereNotIn(t *testing.T)      { testSQLWhereNotIn(t, driverTestMysqlConn) }
func TestMysqlSQL_Find(t *testing.T)            { testSQLFind(t, driverTestMysqlConn) }
func TestMysqlSQL_Sum(t *testing.T)             { testSQLSum(t, driverTestMysqlConn) }
func TestMysqlSQL_Max(t *testing.T)             { testSQLMax(t, driverTestMysqlConn) }
func TestMysqlSQL_Min(t *testing.T)             { testSQLMin(t, driverTestMysqlConn) }
func TestMysqlSQL_Avg(t *testing.T)             { testSQLAvg(t, driverTestMysqlConn) }
func TestMysqlSQL_WhereRaw(t *testing.T)        { testSQLWhereRaw(t, driverTestMysqlConn) }
func TestMysqlSQL_UpdateRaw(t *testing.T)       { testSQLUpdateRaw(t, driverTestMysqlConn) }
func TestMysqlSQL_LeftJoin(t *testing.T)        { testSQLLeftJoin(t, driverTestMysqlConn) }
func TestMysqlSQL_WithTransaction(t *testing.T) { testSQLWithTransaction(t, driverTestMysqlConn) }
func TestMysqlSQL_WithTransactionByLevel(t *testing.T) {
	testSQLWithTransactionByLevel(t, driverTestMysqlConn)
}
func TestMysqlSQL_First(t *testing.T)       { testSQLFirst(t, driverTestMysqlConn) }
func TestMysqlSQL_All(t *testing.T)         { testSQLAll(t, driverTestMysqlConn) }
func TestMysqlSQL_ShowColumns(t *testing.T) { testSQLShowColumns(t, driverTestMysqlConn) }
func TestMysqlSQL_ShowTables(t *testing.T)  { testSQLShowTables(t, driverTestMysqlConn) }
func TestMysqlSQL_Update(t *testing.T)      { testSQLUpdate(t, driverTestMysqlConn) }
func TestMysqlSQL_Delete(t *testing.T)      { testSQLDelete(t, driverTestMysqlConn) }
func TestMysqlSQL_Exec(t *testing.T)        { testSQLExec(t, driverTestMysqlConn) }
func TestMysqlSQL_Insert(t *testing.T)      { testSQLInsert(t, driverTestMysqlConn) }
func TestMysqlSQL_Wrap(t *testing.T)        { testSQLWrap(t, driverTestMysqlConn) }
