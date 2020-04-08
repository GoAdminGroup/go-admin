package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/go-ini/ini"
	"github.com/mgutz/ansi"
	"github.com/schollz/progressbar"
)

var systemGoAdminTables = []string{
	"goadmin_menu",
	"goadmin_operation_log",
	"goadmin_permissions",
	"goadmin_role_menu",
	"goadmin_roles",
	"goadmin_session",
	"goadmin_users",
	"goadmin_role_permissions",
	"goadmin_role_users",
	"goadmin_user_permissions",
}

func generating(cfgFile string) {

	clear(runtime.GOOS)
	cliInfo()

	var (
		driverName, host, port, dbFile, user, password,
		connection, packageName, outputPath, database, schema string

		chooseTables = make([]string, 0)
	)

	if cfgFile != "" {
		cfgModel, err := ini.Load(cfgFile)

		if err != nil {
			panic("wrong config file path")
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
			t := dbCfgModel.Key("tables").Value()
			if t != "" {
				chooseTables = strings.Split(t, ",")
			}
		}

		modelCfgModel, exist2 := cfgModel.GetSection("model")

		if exist2 == nil {
			connection = modelCfgModel.Key("connection").Value()
			packageName = modelCfgModel.Key("package").Value()
			outputPath = modelCfgModel.Key("output").Value()
		}

	}

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "type to filter", "type to filter, enter to select", -1)
	survey.MultiSelectQuestionTemplate = strings.Replace(survey.MultiSelectQuestionTemplate, "enter to select", "space to select", -1)

	if driverName == "" {
		var qs = []*survey.Question{
			{
				Name: "driver",
				Prompt: &survey.Select{
					Message: "choose a driver",
					Options: []string{"mysql", "postgresql", "sqlite", "mssql"},
					Default: "mysql",
				},
			},
		}

		var result = make(map[string]interface{})

		err := survey.Ask(qs, &result)
		checkError(err)
		driverName = result["driver"].(core.OptionAnswer).Value
	}

	var (
		cfg  map[string]config.Database
		conn = db.GetConnectionByDriver(driverName)
	)

	if driverName != "sqlite" {

		defaultPort := "3306"
		defaultUser := "root"

		if driverName == "postgresql" {
			defaultPort = "5432"
			defaultUser = "postgres"
		}

		if driverName == "mssql" {
			defaultPort = "1433"
			defaultUser = "sa"
		}

		if host == "" {
			host = promptWithDefault("sql address", "127.0.0.1")
		}

		if port == "" {
			port = promptWithDefault("sql port", defaultPort)
		}

		if user == "" {
			user = promptWithDefault("sql username", defaultUser)
		}

		if password == "" {
			password = promptPassword()
		}

		if schema == "" && driverName == "postgresql" {
			schema = promptWithDefault("sql schema", "public")
		}

		if database == "" {
			database = prompt("sql database name")
		}

		if conn == nil {
			exitWithError("invalid db connection")
			panic("invalid db connection")
		}
		cfg = map[string]config.Database{
			"default": {
				Host:       host,
				Port:       port,
				User:       user,
				Pwd:        password,
				Name:       database,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     driverName,
				File:       "",
			},
		}
	} else {

		if dbFile == "" {
			dbFile = prompt("sql file")
		}

		if database == "" {
			database = prompt("sql database name")
		}

		if conn == nil {
			exitWithError("invalid db connection")
			panic("invalid db connection")
		}
		cfg = map[string]config.Database{
			"default": {
				Driver: driverName,
				File:   dbFile,
			},
		}
	}

	// step 1. test connection
	conn.InitDB(cfg)

	// step 2. show tables
	if len(chooseTables) == 0 {
		tables, err := db.WithDriver(conn).ShowTables()

		if err != nil {
			panic(err)
		}

		tables = filterTables(tables)

		if len(tables) == 0 {
			exitWithError(`no tables, you should build a table of your own business first.

see: http://www.go-admin.cn/en/docs/#/plugins/admin`)
		}
		tables = append([]string{"[select all]"}, tables...)

		survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "<enter> to select", "<space> to select", -1)

		chooseTables = selects(tables)
		if len(chooseTables) == 0 {
			exitWithError("no table is selected")
		}
		if modules.InArray(chooseTables, "[select all]") {
			chooseTables = tables[1:]
		}
	}

	if packageName == "" {
		packageName = promptWithDefault("set package name", "main")
	}

	if connection == "" {
		connection = promptWithDefault("set connection name", "default")
	}

	if outputPath == "" {
		outputPath = promptWithDefault("set file output path", "./")
	}

	fmt.Println(ansi.Color("✔", "green") + " generating: ")
	fmt.Println()

	fieldField := "Field"
	typeField := "Type"
	if driverName == "postgresql" {
		fieldField = "column_name"
		typeField = "udt_name"
	}
	if driverName == "sqlite" {
		fieldField = "name"
		typeField = "type"
	}
	if driverName == "mssql" {
		fieldField = "column_name"
		typeField = "data_type"
	}

	bar := progressbar.New(len(chooseTables))
	for i := 0; i < len(chooseTables); i++ {
		_ = bar.Add(1)
		time.Sleep(10 * time.Millisecond)
		generateFile(chooseTables[i], schema, conn, fieldField, typeField, packageName, connection, driverName, outputPath)
	}
	generateTables(outputPath, chooseTables, packageName)

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color("generate success~~🍺🍺", "green"))
	fmt.Println()
	fmt.Println("see the docs: " + ansi.Color("http://doc.go-admin.cn/en/#/introduce/plugins/admin", "blue"))
	fmt.Println()
	fmt.Println()
}

func camelcase(s string) string {
	arr := strings.Split(s, "_")
	var res = ""
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			res += arr[i]
		} else {
			res += strings.Title(arr[i])
		}
	}
	return res
}

func clear(osName string) {

	if osName == "linux" || osName == "darwin" {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}

	if osName == "windows" {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		_ = cmd.Run()
	}
}

func filterTables(models []string) []string {
	tables := make([]string, 0)

	for i := 0; i < len(models); i++ {
		// skip goadmin system tables
		if isSystemTable(models[i]) {
			continue
		}
		tables = append(tables, models[i])
	}

	return tables
}

func isSystemTable(name string) bool {
	for _, v := range systemGoAdminTables {
		if name == v {
			return true
		}
	}

	return false
}

func prompt(label string) string {

	var qs = []*survey.Question{
		{
			Name:     label,
			Prompt:   &survey.Input{Message: label},
			Validate: survey.Required,
		},
	}

	var result = make(map[string]interface{})

	err := survey.Ask(qs, &result)

	checkError(err)

	return result[label].(string)
}

func promptWithDefault(label string, defaultValue string) string {

	var qs = []*survey.Question{
		{
			Name:     label,
			Prompt:   &survey.Input{Message: label, Default: defaultValue},
			Validate: survey.Required,
		},
	}

	var result = make(map[string]interface{})

	err := survey.Ask(qs, &result)

	checkError(err)

	return result[label].(string)
}

func promptPassword() string {

	password := ""
	prompt := &survey.Password{
		Message: "sql password",
	}
	err := survey.AskOne(prompt, &password, nil)

	checkError(err)

	return password
}

func selects(tables []string) []string {

	chooseTables := make([]string, 0)
	prompt := &survey.MultiSelect{
		Message:  "choose table to generate",
		Options:  tables,
		PageSize: 10,
	}
	err := survey.AskOne(prompt, &chooseTables, nil)

	checkError(err)

	return chooseTables
}

func generateFile(table, schema string, conn db.Connection, fieldField, typeField, packageName, connection, driver, outputPath string) {

	tableCamel := camelcase(table)

	dbTable := table
	if schema != "" {
		dbTable = schema + "." + table
	}

	columnsModel, _ := db.WithDriver(conn).Table(dbTable).ShowColumns()

	var newTable = `table.NewDefaultTable(table.DefaultConfigWithDriver("` + driver + `"))`
	if connection != "default" {
		newTable = `table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("` + driver + `", "` + connection + `"))`
	}

	content := `package ` + packageName + `

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func Get` + strings.Title(tableCamel) + `Table(ctx *context.Context) table.Table {

    ` + tableCamel + `Table := ` + newTable + `

	info := ` + tableCamel + `Table.GetInfo()
	
	`

	for _, model := range columnsModel {
		if model[fieldField].(string) == "id" {
			content += `info.AddField("` + strings.Title(model[fieldField].(string)) +
				`","` + model[fieldField].(string) +
				`", db.` + getType(model[typeField].(string)) + `).FieldFilterable()
	`
		} else {
			content += `info.AddField("` + strings.Title(model[fieldField].(string)) +
				`","` + model[fieldField].(string) +
				`", db.` + getType(model[typeField].(string)) + `)
	`
		}
	}

	content += `
	info.SetTable("` + dbTable + `").SetTitle("` + strings.Title(table) + `").SetDescription("` + strings.Title(table) + `")

	formList := ` + tableCamel + `Table.GetForm()
	
	`

	for _, model := range columnsModel {

		typeName := getType(model[typeField].(string))
		formType := form.GetFormTypeFromFieldType(db.DT(strings.ToUpper(typeName)), model[fieldField].(string))

		if model[fieldField].(string) == "id" {
			content += `formList.AddField("` + strings.Title(model[fieldField].(string)) + `","` +
				model[fieldField].(string) + `",db.` + typeName + `,` + formType + `).FieldNotAllowAdd()
	`
		} else {
			content += `formList.AddField("` + strings.Title(model[fieldField].(string)) + `","` +
				model[fieldField].(string) + `",db.` + typeName + `,` + formType + `)
	`
		}
	}

	content += `
	formList.SetTable("` + dbTable + `").SetTitle("` + strings.Title(table) + `").SetDescription("` + strings.Title(table) + `")

	return ` + tableCamel + `Table
}`

	c, err := format.Source([]byte(content))
	checkError(err)

	checkError(ioutil.WriteFile(outputPath+"/"+table+".go", c, 0644))
}

func generateTables(outputPath string, tables []string, packageName string) {

	tableStr := ""
	commentStr := ""
	const (
		commentStrEnd = "example end"
		tablesEnd     = "generators end"
	)

	for i := 0; i < len(tables); i++ {
		tableStr += `
	"` + tables[i] + `": Get` + strings.Title(camelcase(tables[i])) + `Table,`
		commentStr += `// "` + tables[i] + `" => http://localhost:9033/admin/info/` + tables[i] + `
`
	}
	commentStr += `//
// ` + commentStrEnd + `
`
	tableStr += `

	// ` + tablesEnd

	tablesContentByte, err := ioutil.ReadFile(outputPath + "/tables.go")
	tablesContent := string(tablesContentByte)

	content := ""

	if err == nil && tablesContent != "" && strings.Index(tablesContent, "/") != -1 {
		tablesContent = strings.Replace(tablesContent, commentStrEnd+`
//`, commentStr[3:]+"//", -1)
		content = strings.Replace(tablesContent, "// "+tablesEnd, tableStr, -1)
	} else {
		content = `package ` + packageName + `

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
` + commentStr + `//
var Generators = map[string]table.Generator{` + tableStr + `
}
`
	}

	c, err := format.Source([]byte(content))

	checkError(err)
	checkError(ioutil.WriteFile(outputPath+"/tables.go", c, 0644))
}

func getType(typeName string) string {
	r, _ := regexp.Compile(`\(.*?\)`)
	typeName = r.ReplaceAllString(typeName, "")
	r2, _ := regexp.Compile(`unsigned(.*)`)
	return strings.TrimSpace(strings.Title(strings.ToLower(r2.ReplaceAllString(typeName, ""))))
}
