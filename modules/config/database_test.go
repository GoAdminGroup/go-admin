package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigDatabase(t *testing.T) {
	// sqlite config
	sqliteDbCfg := NewDatabase(
		DatabaseMaxIdleConOption(150),
		DatabaseMaxOpenConOption(50),
		DatabaseSqliteOption("data/admin.db", nil),
	)

	assert.Equal(t, sqliteDbCfg.Driver, DriverSqlite)
	assert.Equal(t, sqliteDbCfg.MaxIdleCon, 150)
	assert.Equal(t, sqliteDbCfg.MaxOpenCon, 50)
	assert.Equal(t, sqliteDbCfg.File, "data/admin.db")
	assert.Equal(t, sqliteDbCfg.Dsn, "data/admin.db")

	// mysql config
	mysqlDbCfg := NewDatabase(
		DatabaseMysqlOption("127.0.0.1", "3306", "root", "root", "go_admin", nil),
	)

	assert.Equal(t, mysqlDbCfg.Driver, DriverMysql)
	assert.Equal(t, mysqlDbCfg.MaxIdleCon, 50)
	assert.Equal(t, mysqlDbCfg.MaxOpenCon, 150)
	assert.Equal(t, mysqlDbCfg.Dsn, "root:root@tcp(127.0.0.1:3306)/go_admin?charset=utf8mb4&loc=Local&parseTime=True")

	// postgresql config
	postgresqlCfg := NewDatabase(
		DatabasePostgresqlOption("127.0.0.1", "3306", "root", "root", "go_admin", nil),
	)
	assert.Equal(t, postgresqlCfg.Driver, DriverPostgresql)
	assert.Equal(t, postgresqlCfg.Dsn, "host=root port=root user=127.0.0.1 password=3306 dbname=go_admin sslmode=disable")

	// mssql config
	mssqlCfg := NewDatabase(
		DatabaseMssql("127.0.0.1", "3306", "root", "root", "go_admin", map[string]string{}),
	)
	assert.Equal(t, mssqlCfg.Driver, DriverMssql)
	assert.Equal(t, mssqlCfg.Dsn, "user id=root;password=root;server=127.0.0.1;port=3306;database=go_admin;encrypt=disable")
}
