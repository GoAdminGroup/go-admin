package db

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"testing"

	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
)

var driverTestPgConn Connection

func InitPostgresql() {

	cmd := exec.Command("createdb -p 5433 -U postgres " + driverTestDBName)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PGPASSWORD=root")
	_ = cmd.Run()

	cmd = exec.Command("psql", "-h", "localhost", "-U", "root", "-proot", "-d", driverTestDBName,
		"-f", path.Dir(path.Dir(testCurrentPath()))+"/data/admin.pgsql")
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "PGPASSWORD=root")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	driverTestPgConn = testConnDSN(DriverPostgresql, fmt.Sprintf("host=127.0.0.1 port=5433 user=postgres "+
		"password=root dbname=%s sslmode=disable", driverTestDBName))
}

func TestPgSQL_WhereIn(t *testing.T)         { testSQLWhereIn(t, driverTestPgConn) }
func TestPgSQL_Count(t *testing.T)           { testSQLCount(t, driverTestPgConn) }
func TestPgSQL_Select(t *testing.T)          { testSQLSelect(t, driverTestPgConn) }
func TestPgSQL_OrderBy(t *testing.T)         { testSQLOrderBy(t, driverTestPgConn) }
func TestPgSQL_GroupBy(t *testing.T)         { testSQLGroupBy(t, driverTestPgConn) }
func TestPgSQL_Skip(t *testing.T)            { testSQLSkip(t, driverTestPgConn) }
func TestPgSQL_Take(t *testing.T)            { testSQLTake(t, driverTestPgConn) }
func TestPgSQL_Where(t *testing.T)           { testSQLWhere(t, driverTestPgConn) }
func TestPgSQL_WhereNotIn(t *testing.T)      { testSQLWhereNotIn(t, driverTestPgConn) }
func TestPgSQL_Find(t *testing.T)            { testSQLFind(t, driverTestPgConn) }
func TestPgSQL_Sum(t *testing.T)             { testSQLSum(t, driverTestPgConn) }
func TestPgSQL_Max(t *testing.T)             { testSQLMax(t, driverTestPgConn) }
func TestPgSQL_Min(t *testing.T)             { testSQLMin(t, driverTestPgConn) }
func TestPgSQL_Avg(t *testing.T)             { testSQLAvg(t, driverTestPgConn) }
func TestPgSQL_WhereRaw(t *testing.T)        { testSQLWhereRaw(t, driverTestPgConn) }
func TestPgSQL_UpdateRaw(t *testing.T)       { testSQLUpdateRaw(t, driverTestPgConn) }
func TestPgSQL_LeftJoin(t *testing.T)        { testSQLLeftJoin(t, driverTestPgConn) }
func TestPgSQL_WithTransaction(t *testing.T) { testSQLWithTransaction(t, driverTestPgConn) }
func TestPgSQL_WithTransactionByLevel(t *testing.T) {
	testSQLWithTransactionByLevel(t, driverTestPgConn)
}
func TestPgSQL_First(t *testing.T)       { testSQLFirst(t, driverTestPgConn) }
func TestPgSQL_All(t *testing.T)         { testSQLAll(t, driverTestPgConn) }
func TestPgSQL_ShowColumns(t *testing.T) { testSQLShowColumns(t, driverTestPgConn) }
func TestPgSQL_ShowTables(t *testing.T)  { testSQLShowTables(t, driverTestPgConn) }
func TestPgSQL_Update(t *testing.T)      { testSQLUpdate(t, driverTestPgConn) }
func TestPgSQL_Delete(t *testing.T)      { testSQLDelete(t, driverTestPgConn) }
func TestPgSQL_Exec(t *testing.T)        { testSQLExec(t, driverTestPgConn) }
func TestPgSQL_Insert(t *testing.T)      { testSQLInsert(t, driverTestPgConn) }
func TestPgSQL_Wrap(t *testing.T)        { testSQLWrap(t, driverTestPgConn) }
