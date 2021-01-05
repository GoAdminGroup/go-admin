package config

import (
	"encoding/json"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"sort"
	"strings"
)

const (
	// DriverMysql is a const value of mysql driver.
	DriverMysql = "mysql"
	// DriverSqlite is a const value of sqlite driver.
	DriverSqlite = "sqlite"
	// DriverPostgresql is a const value of postgresql driver.
	DriverPostgresql = "postgresql"
	// DriverMssql is a const value of mssql driver.
	DriverMssql = "mssql"
)

// defining database Functional Options that modify Database instance
type DatabaseOption func(*Database)

func DatabaseMaxIdleConOption(maxIdleCon int) DatabaseOption {
	return func(db *Database) {
		db.MaxIdleCon = maxIdleCon
	}
}

func DatabaseMaxOpenConOption(maxOpenCon int) DatabaseOption {
	return func(db *Database) {
		db.MaxOpenCon = maxOpenCon
	}
}

func DatabaseMysqlOption(host, port, user, pwd, name string, param map[string]string) DatabaseOption {
	return func(db *Database) {
		db.Driver = DriverMysql
		db.Host = host
		db.Port = port
		db.User = user
		db.Pwd = pwd
		db.Name = name

		if param == nil {
			param = db.Params
		}
		if _, has := param["parseTime"]; !has {
			param["parseTime"] = "True"
		}
		if _, has := param["loc"]; !has {
			param["loc"] = "Local"
		}
		if _, has := param["charset"]; !has {
			param["charset"] = "utf8mb4"
		}

		paramArr := make([]string, 0)
		for k, v := range param {
			paramArr = append(paramArr, k+"="+v)
		}
		sort.Strings(paramArr)

		db.Params = param
		db.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			user, pwd, host, port, name, strings.Join(paramArr, "&"))
	}
}

func DatabasePostgresqlOption(host, port, user, pwd, name string, param map[string]string) DatabaseOption {
	return func(db *Database) {
		db.Driver = DriverPostgresql
		db.Host = host
		db.Port = port
		db.User = user
		db.Pwd = pwd
		db.Name = name

		if param == nil {
			param = db.Params
		}
		if _, has := param["sslmode"]; !has {
			param["sslmode"] = "disable"
		}

		paramArr := make([]string, 0)
		for k, v := range param {
			paramArr = append(paramArr, k+"="+v)
		}
		sort.Strings(paramArr)

		db.Params = param
		db.Dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s %s",
			user, pwd, host, port, name, strings.Join(paramArr, " "))
	}
}

func DatabaseSqliteOption(filepath string, param map[string]string) DatabaseOption {
	return func(db *Database) {
		db.Driver = DriverSqlite
		db.File = filepath

		if param != nil {
			param = db.Params
		}
		paramArr := make([]string, 0)
		for k, v := range param {
			paramArr = append(paramArr, k+"="+v)
		}
		sort.Strings(paramArr)

		db.Params = param
		if len(paramArr) > 0 {
			filepath += "?"
		}
		db.Dsn = filepath + strings.Join(paramArr, "&")
	}
}

func DatabaseMssql(host, port, user, pwd, dbName string, param map[string]string) DatabaseOption {
	return func(db *Database) {
		db.Driver = DriverMssql
		db.Host = host
		db.Port = port
		db.User = user
		db.Pwd = pwd
		db.Name = dbName

		if param == nil {
			param = db.Params
		}
		if _, has := param["encrypt"]; !has {
			param["encrypt"] = "disable"
		}
		paramArr := make([]string, 0)
		for k, v := range param {
			paramArr = append(paramArr, k+"="+v)
		}
		sort.Strings(paramArr)

		db.Params = param
		db.Dsn = fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s;database=%s;%s",
			user, pwd, host, port, dbName, strings.Join(paramArr, ";"))
	}
}

// Database is a type of database connection config.
//
// Because a little difference of different database driver.
// The Config has multiple options but may not be used.
// Such as the sqlite driver only use the File option which
// can be ignored when the driver is mysql.
//
// If the Dsn is configured, when driver is mysql/postgresql/
// mssql, the other configurations will be ignored, except for
// MaxIdleCon and MaxOpenCon.
type Database struct {
	Host       string            `json:"host,omitempty" yaml:"host,omitempty" ini:"host,omitempty"`
	Port       string            `json:"port,omitempty" yaml:"port,omitempty" ini:"port,omitempty"`
	User       string            `json:"user,omitempty" yaml:"user,omitempty" ini:"user,omitempty"`
	Pwd        string            `json:"pwd,omitempty" yaml:"pwd,omitempty" ini:"pwd,omitempty"`
	Name       string            `json:"name,omitempty" yaml:"name,omitempty" ini:"name,omitempty"`
	MaxIdleCon int               `json:"max_idle_con,omitempty" yaml:"max_idle_con,omitempty" ini:"max_idle_con,omitempty"`
	MaxOpenCon int               `json:"max_open_con,omitempty" yaml:"max_open_con,omitempty" ini:"max_open_con,omitempty"`
	Driver     string            `json:"driver,omitempty" yaml:"driver,omitempty" ini:"driver,omitempty"`
	DriverMode string            `json:"driver_mode,omitempty" yaml:"driver_mode,omitempty" ini:"driver_mode,omitempty"`
	File       string            `json:"file,omitempty" yaml:"file,omitempty" ini:"file,omitempty"`
	Dsn        string            `json:"dsn,omitempty" yaml:"dsn,omitempty" ini:"dsn,omitempty"`
	Params     map[string]string `json:"params,omitempty" yaml:"params,omitempty" ini:"params,omitempty"`
}

func (d Database) GetDSN() string {
	if d.Dsn != "" {
		return d.Dsn
	}

	if d.Driver == DriverMysql {
		return d.User + ":" + d.Pwd + "@tcp(" + d.Host + ":" + d.Port + ")/" +
			d.Name + d.ParamStr()
	}
	if d.Driver == DriverPostgresql {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s"+d.ParamStr(),
			d.Host, d.Port, d.User, d.Pwd, d.Name)
	}
	if d.Driver == DriverMssql {
		return fmt.Sprintf("user id=%s;password=%s;server=%s;port=%s;database=%s;"+d.ParamStr(),
			d.User, d.Pwd, d.Host, d.Port, d.Name)
	}
	if d.Driver == DriverSqlite {
		return d.File + d.ParamStr()
	}
	return ""
}

func (d Database) ParamStr() string {
	p := ""
	if d.Params == nil {
		d.Params = make(map[string]string)
	}
	if d.Driver == DriverMysql || d.Driver == DriverSqlite {
		if d.Driver == DriverMysql {
			if _, ok := d.Params["charset"]; !ok {
				d.Params["charset"] = "utf8mb4"
			}
		}
		if len(d.Params) > 0 {
			p = "?"
			for k, v := range d.Params {
				p += k + "=" + v + "&"
			}
			p = p[:len(p)-1]
		}
	}
	if d.Driver == DriverMssql {
		if _, ok := d.Params["encrypt"]; !ok {
			d.Params["encrypt"] = "disable"
		}
		for k, v := range d.Params {
			p += k + "=" + v + ";"
		}
		p = p[:len(p)-1]
	}
	if d.Driver == DriverPostgresql {
		if _, ok := d.Params["sslmode"]; !ok {
			d.Params["sslmode"] = "disable"
		}
		p = " "
		for k, v := range d.Params {
			p += k + "=" + v + " "
		}
		p = p[:len(p)-1]
	}
	return p
}

// DatabaseList is a map of Database.
type DatabaseList map[string]Database

// GetDefault get the default Database.
func (d DatabaseList) GetDefault() Database {
	return d["default"]
}

// Add add a Database to the DatabaseList.
func (d DatabaseList) Add(key string, db Database) {
	d[key] = db
}

// GroupByDriver group the Databases with the drivers.
func (d DatabaseList) GroupByDriver() map[string]DatabaseList {
	drivers := make(map[string]DatabaseList)
	for key, item := range d {
		if driverList, ok := drivers[item.Driver]; ok {
			driverList.Add(key, item)
		} else {
			drivers[item.Driver] = make(DatabaseList)
			drivers[item.Driver].Add(key, item)
		}
	}
	return drivers
}

func (d DatabaseList) JSON() string {
	return utils.JSON(d)
}

func (d DatabaseList) Copy() DatabaseList {
	var c = make(DatabaseList)
	for k, v := range d {
		c[k] = v
	}
	return c
}

func (d DatabaseList) Connections() []string {
	conns := make([]string, len(d))
	count := 0
	for key := range d {
		conns[count] = key
		count++
	}
	return conns
}

func GetDatabaseListFromJSON(m string) DatabaseList {
	var d = make(DatabaseList)
	if m == "" {
		panic("wrong config")
	}
	_ = json.Unmarshal([]byte(m), &d)
	return d
}

func NewDatabase(opts ...DatabaseOption) *Database {
	database := &Database{
		MaxIdleCon: 50,
		MaxOpenCon: 150,
		Params:     make(map[string]string),
	}

	for _, opt := range opts {
		opt(database)
	}

	return database
}
