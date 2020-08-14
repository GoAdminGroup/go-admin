package main

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"gopkg.in/ini.v1"
)

type dbInfo struct {
	DriverName string
	Host       string
	Port       string
	File       string
	User       string
	Password   string
	Schema     string
	Database   string
}

func initSurvey() {
	survey.SelectQuestionTemplate = strings.ReplaceAll(survey.SelectQuestionTemplate,
		"type to filter", "type to filter, enter to select")
	survey.MultiSelectQuestionTemplate = strings.ReplaceAll(survey.MultiSelectQuestionTemplate,
		"enter to select", "space to select")

	survey.SelectQuestionTemplate = strings.ReplaceAll(survey.SelectQuestionTemplate,
		"Use arrows to move, type to filter, enter to select",
		getWord("Use arrows to move, type to filter, enter to select"))

	survey.MultiSelectQuestionTemplate = strings.ReplaceAll(survey.MultiSelectQuestionTemplate,
		"Use arrows to move, space to select, type to filter",
		getWord("Use arrows to move, space to select, type to filter"))
}

func getDBInfoFromINIConfig(cfg *ini.File, connection string) *dbInfo {

	section := "database"

	if connection != "default" && connection != "" {
		section = section + "." + connection
	}

	dbCfgModel, exist := cfg.GetSection(section)

	if exist == nil {
		return &dbInfo{
			DriverName: dbCfgModel.Key("driver").Value(),
			Host:       dbCfgModel.Key("host").Value(),
			User:       dbCfgModel.Key("username").Value(),
			Port:       dbCfgModel.Key("port").Value(),
			File:       dbCfgModel.Key("file").Value(),
			Password:   dbCfgModel.Key("password").Value(),
			Database:   dbCfgModel.Key("database").Value(),
		}
	}

	return &dbInfo{}
}

func askForDBConfig(info *dbInfo) config.DatabaseList {

	initSurvey()

	if info.DriverName == "" {
		info.DriverName = singleSelect(getWord("choose a driver"),
			[]string{db.DriverMysql, db.DriverPostgresql, db.DriverSqlite, db.DriverMssql}, db.DriverMysql)
	}

	if info.DriverName != db.DriverSqlite {

		defaultPort := "3306"
		defaultUser := "root"

		if info.DriverName == db.DriverPostgresql {
			defaultPort = "5432"
			defaultUser = "postgres"
		}

		if info.DriverName == db.DriverMssql {
			defaultPort = "1433"
			defaultUser = "sa"
		}

		if info.Host == "" {
			info.Host = promptWithDefault("sql address", "127.0.0.1")
		}

		if info.Port == "" {
			info.Port = promptWithDefault("sql port", defaultPort)
		}

		if info.User == "" {
			info.User = promptWithDefault("sql username", defaultUser)
		}

		if info.Password == "" {
			info.Password = promptPassword()
		}

		if info.Schema == "" && info.DriverName == db.DriverPostgresql {
			info.Schema = promptWithDefault("sql schema", "public")
		}

		if info.Database == "" {
			info.Database = prompt("sql database name")
		}

		return map[string]config.Database{
			"default": {
				Host:       info.Host,
				Port:       info.Port,
				User:       info.User,
				Pwd:        info.Password,
				Name:       info.Database,
				MaxIdleCon: 5,
				MaxOpenCon: 10,
				Driver:     info.DriverName,
				File:       "",
			},
		}
	} else {

		if info.File == "" {
			info.File = promptWithDefault("sql file", "./admin.db")
		}

		return map[string]config.Database{
			"default": {
				Driver: info.DriverName,
				File:   info.File,
			},
		}
	}
}

func askForDBConnection(info *dbInfo) db.Connection {

	var (
		cfg  = askForDBConfig(info)
		conn = db.GetConnectionByDriver(info.DriverName)
	)

	conn.InitDB(cfg)

	return conn
}
