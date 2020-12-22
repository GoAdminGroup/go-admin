package table

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
)

type Config struct {
	Driver         string
	DriverMode     string
	Connection     string
	CanAdd         bool
	Editable       bool
	Deletable      bool
	Exportable     bool
	PrimaryKey     PrimaryKey
	SourceURL      string
	GetDataFun     GetDataFun
	OnlyInfo       bool
	OnlyNewForm    bool
	OnlyUpdateForm bool
	OnlyDetail     bool
}

func DefaultConfig() Config {
	return Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: DefaultConnectionName,
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	}
}

func (config Config) SetPrimaryKey(name string, typ db.DatabaseType) Config {
	config.PrimaryKey.Name = name
	config.PrimaryKey.Type = typ
	return config
}

func (config Config) SetDriverMode(mode string) Config {
	config.DriverMode = mode
	return config
}

func (config Config) SetPrimaryKeyType(typ string) Config {
	config.PrimaryKey.Type = db.GetDTAndCheck(typ)
	return config
}

func (config Config) SetCanAdd(canAdd bool) Config {
	config.CanAdd = canAdd
	return config
}

func (config Config) SetSourceURL(url string) Config {
	config.SourceURL = url
	return config
}

func (config Config) SetGetDataFun(fun GetDataFun) Config {
	config.GetDataFun = fun
	return config
}

func (config Config) SetEditable(editable bool) Config {
	config.Editable = editable
	return config
}

func (config Config) SetDeletable(deletable bool) Config {
	config.Deletable = deletable
	return config
}

func (config Config) SetOnlyInfo() Config {
	config.OnlyInfo = true
	return config
}

func (config Config) SetOnlyUpdateForm() Config {
	config.OnlyUpdateForm = true
	return config
}

func (config Config) SetOnlyNewForm() Config {
	config.OnlyNewForm = true
	return config
}

func (config Config) SetOnlyDetail() Config {
	config.OnlyDetail = true
	return config
}

func (config Config) SetExportable(exportable bool) Config {
	config.Exportable = exportable
	return config
}

func (config Config) SetConnection(connection string) Config {
	config.Connection = connection
	return config
}

func DefaultConfigWithDriver(driver string) Config {
	return Config{
		Driver:     driver,
		Connection: DefaultConnectionName,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	}
}

func DefaultConfigWithDriverAndConnection(driver, conn string) Config {
	return Config{
		Driver:     driver,
		Connection: conn,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	}
}
