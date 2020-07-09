package db

import (
	"testing"

	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite"
)

var driverTestSQLiteConn Connection

func InitSqlite() {
	driverTestSQLiteConn = testConn(DriverSqlite, config.Database{File: "/admin.db"})
}

func TestSQLiteSQL_WhereIn(t *testing.T)         { testSQLWhereIn(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Count(t *testing.T)           { testSQLCount(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Select(t *testing.T)          { testSQLSelect(t, driverTestSQLiteConn) }
func TestSQLiteSQL_OrderBy(t *testing.T)         { testSQLOrderBy(t, driverTestSQLiteConn) }
func TestSQLiteSQL_GroupBy(t *testing.T)         { testSQLGroupBy(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Skip(t *testing.T)            { testSQLSkip(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Take(t *testing.T)            { testSQLTake(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Where(t *testing.T)           { testSQLWhere(t, driverTestSQLiteConn) }
func TestSQLiteSQL_WhereNotIn(t *testing.T)      { testSQLWhereNotIn(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Find(t *testing.T)            { testSQLFind(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Sum(t *testing.T)             { testSQLSum(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Max(t *testing.T)             { testSQLMax(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Min(t *testing.T)             { testSQLMin(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Avg(t *testing.T)             { testSQLAvg(t, driverTestSQLiteConn) }
func TestSQLiteSQL_WhereRaw(t *testing.T)        { testSQLWhereRaw(t, driverTestSQLiteConn) }
func TestSQLiteSQL_UpdateRaw(t *testing.T)       { testSQLUpdateRaw(t, driverTestSQLiteConn) }
func TestSQLiteSQL_LeftJoin(t *testing.T)        { testSQLLeftJoin(t, driverTestSQLiteConn) }
func TestSQLiteSQL_WithTransaction(t *testing.T) { testSQLWithTransaction(t, driverTestSQLiteConn) }
func TestSQLiteSQL_WithTransactionByLevel(t *testing.T) {
	testSQLWithTransactionByLevel(t, driverTestSQLiteConn)
}
func TestSQLiteSQL_First(t *testing.T)       { testSQLFirst(t, driverTestSQLiteConn) }
func TestSQLiteSQL_All(t *testing.T)         { testSQLAll(t, driverTestSQLiteConn) }
func TestSQLiteSQL_ShowColumns(t *testing.T) { testSQLShowColumns(t, driverTestSQLiteConn) }
func TestSQLiteSQL_ShowTables(t *testing.T)  { testSQLShowTables(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Update(t *testing.T)      { testSQLUpdate(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Delete(t *testing.T)      { testSQLDelete(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Exec(t *testing.T)        { testSQLExec(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Insert(t *testing.T)      { testSQLInsert(t, driverTestSQLiteConn) }
func TestSQLiteSQL_Wrap(t *testing.T)        { testSQLWrap(t, driverTestSQLiteConn) }
