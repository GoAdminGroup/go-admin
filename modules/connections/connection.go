package connections

import (
	"database/sql"
	"github.com/chenhg5/go-admin/modules/connections/mysql"
	"github.com/chenhg5/go-admin/modules/config"
)

type Connection interface {
	Query(query string, args ...interface{}) ([]map[string]interface{}, *sql.Rows)
	Exec(query string, args ...interface{}) sql.Result
	InitDB(cfg map[string]config.Database)
}

func GetConnectionByDriver(driver string) Connection {
	switch driver {
	case "mysql":
		return mysql.GetMysqlDB()
	default:
		panic("driver not found!")
	}
}

func GetConnection() Connection {
	return GetConnectionByDriver(config.GlobalCfg.DATABASE.DRIVER)
}