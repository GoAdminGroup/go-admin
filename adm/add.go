package main

import (
	"runtime"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/auth"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"gopkg.in/ini.v1"
)

func addUser(cfgFile string) {

	clear(runtime.GOOS)
	cliInfo()

	var (
		driverName, host, port, dbFile, user, password,
		database, schema, name, nickname, userPassword string
	)

	if cfgFile != "" {
		cfgModel, err := ini.Load(cfgFile)

		if err != nil {
			panic(newError("wrong config file path"))
		}

		languageCfg, err := cfgModel.GetSection("language")

		if err == nil {
			setDefaultLangSet(languageCfg.Key("language").Value())
		}

		userCfg, err := cfgModel.GetSection("user")

		if err == nil {
			name = userCfg.Key("name").Value()
			nickname = userCfg.Key("nickname").Value()
			userPassword = userCfg.Key("password").Value()
		}

		dbCfgModel, exist := cfgModel.GetSection("database")

		if exist == nil {
			driverName = dbCfgModel.Key("driver").Value()
			host = dbCfgModel.Key("host").Value()
			user = dbCfgModel.Key("username").Value()
			port = dbCfgModel.Key("port").Value()
			dbFile = dbCfgModel.Key("file").Value()
			password = dbCfgModel.Key("password").Value()
			database = dbCfgModel.Key("database").Value()
		}
	}

	conn := askForDBConnection(&dbInfo{
		File:       dbFile,
		DriverName: driverName,
		Host:       host,
		Port:       port,
		User:       user,
		Password:   password,
		Schema:     schema,
		Database:   database,
	})

	if name == "" {
		name = promptWithDefault("user login name", "")
	}

	if nickname == "" {
		nickname = promptWithDefault("user nickname", "")
	}

	if userPassword == "" {
		userPassword = promptWithDefault("user password", "")
	}

	checkExist, err := db.WithDriver(conn).Table("goadmin_users").
		Where("name", "=", name).
		First()

	if db.CheckError(err, db.QUERY) {
		panic(err)
	}

	if checkExist != nil {
		panic(newError("user record exists"))
	}

	_, err = db.WithDriver(conn).Table("goadmin_users").
		Insert(dialect.H{
			"name":     name,
			"username": nickname,
			"password": auth.EncodePassword([]byte(userPassword)),
		})

	if db.CheckError(err, db.INSERT) {
		panic(err)
	}

	printSuccessInfo("Add admin user success~~🍺🍺")
}

func addPermission(cfgFile string) {

	clear(runtime.GOOS)
	cliInfo()

	var (
		driverName, host, port, dbFile, user, password,
		database, schema, tablesStr string
		tables []string
	)

	if cfgFile != "" {
		cfgModel, err := ini.Load(cfgFile)

		if err != nil {
			panic(newError("wrong config file path"))
		}

		languageCfg, err := cfgModel.GetSection("language")

		if err == nil {
			setDefaultLangSet(languageCfg.Key("language").Value())
		}

		userCfg, err := cfgModel.GetSection("permission")

		if err == nil {
			tablesStr = userCfg.Key("tables").Value()
		}

		if tablesStr != "" {
			tables = strings.Split(strings.TrimSpace(tablesStr), ",")
		}

		dbCfgModel, exist := cfgModel.GetSection("database")

		if exist == nil {
			driverName = dbCfgModel.Key("driver").Value()
			host = dbCfgModel.Key("host").Value()
			user = dbCfgModel.Key("username").Value()
			port = dbCfgModel.Key("port").Value()
			dbFile = dbCfgModel.Key("file").Value()
			password = dbCfgModel.Key("password").Value()
			database = dbCfgModel.Key("database").Value()
		}
	}

	conn := askForDBConnection(&dbInfo{
		File:       dbFile,
		DriverName: driverName,
		Host:       host,
		Port:       port,
		User:       user,
		Password:   password,
		Schema:     schema,
		Database:   database,
	})

	if tablesStr == "" {
		tablesStr = promptWithDefault("tables to generate, use comma to split", "")
	}

	if tablesStr == "" {
		panic(newError("empty tables"))
	}

	tables = strings.Split(strings.TrimSpace(tablesStr), ",")

	for _, table := range tables {
		insertPermissionOfTable(conn, table)
	}

	printSuccessInfo("Add table permissions success~~🍺🍺")
}

func insertPermissionOfTable(conn db.Connection, table string) {
	table = strings.ToLower(table)
	insertPermissionInfoDB(conn, table+" "+getWord("Query"), table+"_query", "GET", "/info/"+table)
	insertPermissionInfoDB(conn, table+" "+getWord("Show Edit Form Page"), table+"_show_edit", "GET",
		"/info/"+table+"/edit")
	insertPermissionInfoDB(conn, table+" "+getWord("Show Create Form Page"), table+"_show_create", "GET",
		"/info/"+table+"/new")
	insertPermissionInfoDB(conn, table+" "+getWord("Edit"), table+"_edit", "POST",
		"/edit/"+table)
	insertPermissionInfoDB(conn, table+" "+getWord("Create"), table+"_create", "POST",
		"/new/"+table)
	insertPermissionInfoDB(conn, table+" "+getWord("Delete"), table+"_delete", "POST",
		"/delete/"+table)
	insertPermissionInfoDB(conn, table+" "+getWord("Export"), table+"_export", "POST",
		"/export/"+table)
}

func insertPermissionInfoDB(conn db.Connection, name, slug, httpMethod, httpPath string) {
	checkExist, err := db.WithDriver(conn).Table("goadmin_permissions").
		Where("slug", "=", slug).
		First()

	if db.CheckError(err, db.QUERY) {
		panic(err)
	}

	if checkExist != nil {
		return
	}

	_, err = db.WithDriver(conn).Table("goadmin_permissions").
		Insert(dialect.H{
			"name":        name,
			"slug":        slug,
			"http_method": httpMethod,
			"http_path":   httpPath,
		})

	if db.CheckError(err, db.INSERT) {
		panic(err)
	}
}
