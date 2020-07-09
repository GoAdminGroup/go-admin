package db

import (
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mssql"
)

var driverTestMssqlConn Connection

func InitMssql() {
	driverTestMssqlConn = testConn(DriverMssql, config.Database{
		Host: "127.0.0.1",
		Port: "1433",
		User: "sa",
		Pwd:  "Aa123456",
		Name: "goadmin",
	})
}

func TestMssqlSQL_WhereIn(t *testing.T)         { testSQLWhereIn(t, driverTestMssqlConn) }
func TestMssqlSQL_Count(t *testing.T)           { testSQLCount(t, driverTestMssqlConn) }
func TestMssqlSQL_Select(t *testing.T)          { testSQLSelect(t, driverTestMssqlConn) }
func TestMssqlSQL_OrderBy(t *testing.T)         { testSQLOrderBy(t, driverTestMssqlConn) }
func TestMssqlSQL_GroupBy(t *testing.T)         { testSQLGroupBy(t, driverTestMssqlConn) }
func TestMssqlSQL_Skip(t *testing.T)            { testSQLSkip(t, driverTestMssqlConn) }
func TestMssqlSQL_Take(t *testing.T)            { testSQLTake(t, driverTestMssqlConn) }
func TestMssqlSQL_Where(t *testing.T)           { testSQLWhere(t, driverTestMssqlConn) }
func TestMssqlSQL_WhereNotIn(t *testing.T)      { testSQLWhereNotIn(t, driverTestMssqlConn) }
func TestMssqlSQL_Find(t *testing.T)            { testSQLFind(t, driverTestMssqlConn) }
func TestMssqlSQL_Sum(t *testing.T)             { testSQLSum(t, driverTestMssqlConn) }
func TestMssqlSQL_Max(t *testing.T)             { testSQLMax(t, driverTestMssqlConn) }
func TestMssqlSQL_Min(t *testing.T)             { testSQLMin(t, driverTestMssqlConn) }
func TestMssqlSQL_Avg(t *testing.T)             { testSQLAvg(t, driverTestMssqlConn) }
func TestMssqlSQL_WhereRaw(t *testing.T)        { testSQLWhereRaw(t, driverTestMssqlConn) }
func TestMssqlSQL_UpdateRaw(t *testing.T)       { testSQLUpdateRaw(t, driverTestMssqlConn) }
func TestMssqlSQL_LeftJoin(t *testing.T)        { testSQLLeftJoin(t, driverTestMssqlConn) }
func TestMssqlSQL_WithTransaction(t *testing.T) { testSQLWithTransaction(t, driverTestMssqlConn) }
func TestMssqlSQL_WithTransactionByLevel(t *testing.T) {
	testSQLWithTransactionByLevel(t, driverTestMssqlConn)
}
func TestMssqlSQL_First(t *testing.T)       { testSQLFirst(t, driverTestMssqlConn) }
func TestMssqlSQL_All(t *testing.T)         { testSQLAll(t, driverTestMssqlConn) }
func TestMssqlSQL_ShowColumns(t *testing.T) { testSQLShowColumns(t, driverTestMssqlConn) }
func TestMssqlSQL_ShowTables(t *testing.T)  { testSQLShowTables(t, driverTestMssqlConn) }
func TestMssqlSQL_Update(t *testing.T)      { testSQLUpdate(t, driverTestMssqlConn) }
func TestMssqlSQL_Delete(t *testing.T)      { testSQLDelete(t, driverTestMssqlConn) }
func TestMssqlSQL_Exec(t *testing.T)        { testSQLExec(t, driverTestMssqlConn) }
func TestMssqlSQL_Insert(t *testing.T)      { testSQLInsert(t, driverTestMssqlConn) }
func TestMssqlSQL_Wrap(t *testing.T)        { testSQLWrap(t, driverTestMssqlConn) }
