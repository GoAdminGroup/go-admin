package main

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
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

func askForDBInfo(info *dbInfo) db.Connection {

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate,
		"type to filter", "type to filter, enter to select", -1)
	survey.MultiSelectQuestionTemplate = strings.Replace(survey.MultiSelectQuestionTemplate,
		"enter to select", "space to select", -1)

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate,
		"Use arrows to move, type to filter, enter to select",
		getWord("Use arrows to move, type to filter, enter to select"), -1)

	survey.MultiSelectQuestionTemplate = strings.Replace(survey.MultiSelectQuestionTemplate,
		"Use arrows to move, space to select, type to filter",
		getWord("Use arrows to move, space to select, type to filter"), -1)

	if info.DriverName == "" {
		var qs = []*survey.Question{
			{
				Name: "driver",
				Prompt: &survey.Select{
					Message: getWord("choose a driver"),
					Options: []string{"mysql", "postgresql", "sqlite", "mssql"},
					Default: "mysql",
				},
			},
		}

		var result = make(map[string]interface{})

		err := survey.Ask(qs, &result)
		checkError(err)
		info.DriverName = result["driver"].(core.OptionAnswer).Value
	}

	var (
		cfg  map[string]config.Database
		conn = db.GetConnectionByDriver(info.DriverName)
	)

	if info.DriverName != "sqlite" {

		defaultPort := "3306"
		defaultUser := "root"

		if info.DriverName == "postgresql" {
			defaultPort = "5432"
			defaultUser = "postgres"
		}

		if info.DriverName == "mssql" {
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

		if info.Schema == "" && info.DriverName == "postgresql" {
			info.Schema = promptWithDefault("sql schema", "public")
		}

		if info.Database == "" {
			info.Database = prompt("sql database name")
		}

		if conn == nil {
			panic("invalid db connection")
		}
		cfg = map[string]config.Database{
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
			info.File = prompt("sql file")
		}

		if conn == nil {
			panic("invalid db connection")
		}
		cfg = map[string]config.Database{
			"default": {
				Driver: info.DriverName,
				File:   info.File,
			},
		}
	}

	conn.InitDB(cfg)

	return conn
}
