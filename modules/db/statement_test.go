package db

import (
	"database/sql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mssql"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/magiconair/properties/assert"
	"testing"
)

func testSQLWhereIn(t *testing.T, conn Connection) {

	item, _ := WithDriver(conn).Table("users").WhereIn("id", []interface{}{"3"}).First()
	assert.Equal(t, item != nil, true)

	_, _ = WithDriver(conn).WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {
		item, _ := WithDriver(conn).WithTx(tx).Table("users").WhereIn("id", []interface{}{"3"}).First()
		assert.Equal(t, item != nil, true)
		return nil, nil
	})
}

func testSQLCount(t *testing.T, conn Connection) {
	count, _ := WithDriver(conn).Table("goadmin_users").Count()
	assert.Equal(t, count, int64(2))
}
