package dialect

import "github.com/chenhg5/go-admin/modules/config"

type Dialect interface {
	// GetName get dialect's name
	GetName() string

	// ShowColumns show columns
	ShowColumns(table string) string

	// ShowColumns show columns
	ShowTables() string
}

func GetDialect() Dialect {
	return GetDialectByDriver(config.Get().DATABASE[0].DRIVER)
}

func GetDialectByDriver(driver string) Dialect {
	switch driver {
	case "mysql":
		return mysql{}
	case "postgresql":
		return postgresql{}
	case "sqlite":
		return sqlite{}
	default:
		return commonDialect{}
	}
}
