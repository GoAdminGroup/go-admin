package db

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestSQL_WhereIn(t *testing.T) {

	post := GetConnectionByDriver(DriverPostgresql).InitDB(map[string]config.Database{
		"default": {
			Host:       "127.0.0.1",
			Port:       "5433",
			User:       "postgres",
			Pwd:        "root",
			Name:       "godmin",
			MaxIdleCon: 50,
			MaxOpenCon: 150,
			Driver:     DriverPostgresql,
		},
	})

	item, _ := WithDriver(post).Table("users").WhereIn("id", []interface{}{"3"}).First()
	assert.Equal(t, item != nil, true)

	_, _ = WithDriver(post).WithTransaction(func(tx *sql.Tx) (e error, i map[string]interface{}) {
		item, _ := WithDriver(post).WithTx(tx).Table("users").WhereIn("id", []interface{}{"3"}).First()
		assert.Equal(t, item != nil, true)
		return nil, nil
	})
}
