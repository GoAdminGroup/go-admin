package db

import (
	"database/sql"
	"testing"

	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mssql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/magiconair/properties/assert"
)

func testSQLWhereIn(t *testing.T, conn Connection) {

	item, _ := WithDriver(conn).Table("goadmin_users").WhereIn("id", []interface{}{"1", "2"}).First()
	assert.Equal(t, len(item), 2)

	_, _ = WithDriver(conn).WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {
		item, _ := WithDriver(conn).WithTx(tx).Table("goadmin_users").WhereIn("id", []interface{}{"1", "2"}).All()
		assert.Equal(t, len(item), 2)
		return nil, nil
	})
}

func testSQLCount(t *testing.T, conn Connection) {
	count, _ := WithDriver(conn).Table("goadmin_users").Count()
	assert.Equal(t, count, int64(2))
}

// TODO
func testSQLSelect(t *testing.T, conn Connection) {}

// TODO
func testSQLOrderBy(t *testing.T, conn Connection) {}

// TODO
func testSQLGroupBy(t *testing.T, conn Connection) {}

// TODO
func testSQLSkip(t *testing.T, conn Connection) {}

// TODO
func testSQLTake(t *testing.T, conn Connection) {}

// TODO
func testSQLWhere(t *testing.T, conn Connection) {}

// TODO
func testSQLWhereNotIn(t *testing.T, conn Connection) {}

// TODO
func testSQLFind(t *testing.T, conn Connection) {}

// TODO
func testSQLSum(t *testing.T, conn Connection) {}

// TODO
func testSQLMax(t *testing.T, conn Connection) {}

// TODO
func testSQLMin(t *testing.T, conn Connection) {}

// TODO
func testSQLAvg(t *testing.T, conn Connection) {}

// TODO
func testSQLWhereRaw(t *testing.T, conn Connection) {}

// TODO
func testSQLUpdateRaw(t *testing.T, conn Connection) {}

// TODO
func testSQLLeftJoin(t *testing.T, conn Connection) {}

// TODO
func testSQLWithTransaction(t *testing.T, conn Connection) {}

// TODO
func testSQLWithTransactionByLevel(t *testing.T, conn Connection) {}

// TODO
func testSQLFirst(t *testing.T, conn Connection) {}

// TODO
func testSQLAll(t *testing.T, conn Connection) {}

// TODO
func testSQLShowColumns(t *testing.T, conn Connection) {}

// TODO
func testSQLShowTables(t *testing.T, conn Connection) {}

// TODO
func testSQLUpdate(t *testing.T, conn Connection) {}

// TODO
func testSQLDelete(t *testing.T, conn Connection) {}

// TODO
func testSQLExec(t *testing.T, conn Connection) {}

// TODO
func testSQLInsert(t *testing.T, conn Connection) {}

// TODO
func testSQLWrap(t *testing.T, conn Connection) {}
