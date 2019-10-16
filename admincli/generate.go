package main

import (
	"fmt"
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
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"github.com/mgutz/ansi"
	"github.com/schollz/progressbar"
)

func generating() {

	clear(runtime.GOOS)
	cliInfo()

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "space to select", "<enter> to select", -1)

	var qs = []*survey.Question{
		{
			Name: "driver",
			Prompt: &survey.Select{
				Message: "choose a driver",
				Options: []string{"mysql", "postgresql", "sqlite"},
				Default: "mysql",
			},
		},
	}

	var result = make(map[string]interface{})

	err := survey.Ask(qs, &result)
	checkError(err)
	driver := result["driver"].(core.OptionAnswer)

	var (
		cfg  map[string]config.Database
		name string
		conn = db.GetConnectionByDriver(driver.Value)
	)

	if driver.Value != "sqlite" {

		defaultPort := "3306"
		defaultUser := "root"

		if driver.Value == "postgresql" {
			defaultPort = "5432"
			defaultUser = "postgres"
		}

		host := promptWithDefault("sql address", "127.0.0.1")
		port := promptWithDefault("sql port", defaultPort)
		user := promptWithDefault("sql username", defaultUser)
		password := promptPassword()

		name = prompt("sql database name")

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
				Name:       name,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     driver.Value,
				File:       "",
			},
		}
	} else {
		file := prompt("sql file")

		name = prompt("sql database name")

		if conn == nil {
			exitWithError("invalid db connection")
			panic("invalid db connection")
		}
		cfg = map[string]config.Database{
			"default": {
				Driver: driver.Value,
				File:   file,
			},
		}
	}

	// step 1. test connection
	conn.InitDB(cfg)

	// step 2. show tables
	tableModels, _ := db.WithDriver(conn.GetName()).ShowTables()

	tables := getTablesFromSqlResult(tableModels, driver.Value, name)
	if len(tables) == 0 {
		exitWithError(`no tables, you should build a table of your own business first.

see: http://www.go-admin.cn/en/docs/#/plugins/admin`)
	}

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "<enter> to select", "<space> to select", -1)

	chooseTables := selects(tables)
	if len(chooseTables) == 0 {
		exitWithError("no table is selected")
	}

	packageName := promptWithDefault("set package name", "main")
	connectionName := promptWithDefault("set connection name", "default")
	outputPath := promptWithDefault("set file output path", "./")

	fmt.Println(ansi.Color("✔", "green") + " generating: ")
	fmt.Println()

	fieldField := "Field"
	typeField := "Type"
	if driver.Value == "postgresql" {
		fieldField = "column_name"
		typeField = "udt_name"
	}

	bar := progressbar.New(len(chooseTables))
	for i := 0; i < len(chooseTables); i++ {
		_ = bar.Add(1)
		time.Sleep(10 * time.Millisecond)
		generateFile(chooseTables[i], conn, fieldField, typeField, packageName, connectionName, driver.Value, outputPath)
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

func getTablesFromSqlResult(models []map[string]interface{}, driver string, dbName string) []string {
	key := "Tables_in_" + strings.ToLower(dbName)
	if driver == "postgresql" {
		key = "tablename"
	}

	tables := make([]string, 0)

	for i := 0; i < len(models); i++ {
		if models[i][key].(string) != "goadmin_menu" && models[i][key].(string) != "goadmin_operation_log" &&
			models[i][key].(string) != "goadmin_permissions" && models[i][key].(string) != "goadmin_role_menu" &&
			models[i][key].(string) != "goadmin_roles" && models[i][key].(string) != "goadmin_session" &&
			models[i][key].(string) != "goadmin_users" && models[i][key].(string) != "goadmin_role_permissions" &&
			models[i][key].(string) != "goadmin_role_users" && models[i][key].(string) != "goadmin_user_permissions" {
			tables = append(tables, models[i][key].(string))
		}
	}

	return tables
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

func generateFile(table string, conn db.Connection, fieldField, typeField, packageName, connectionName, driver, outputPath string) {

	columnsModel, _ := db.WithDriver(conn.GetName()).Table(table).ShowColumns()

	var newTable = `table.NewDefaultTable(table.DefaultConfigWithDriver("` + driver + `"))`
	if connectionName != "default" {
		newTable = `table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("` + driver + `", "` + connectionName + `"))`
	}

	content := `package ` + packageName + `

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func Get` + strings.Title(table) + `Table() table.Table {

    ` + table + `Table := ` + newTable + `

	info := ` + table + `Table.GetInfo()
	
	`

	for _, model := range columnsModel {
		content += `info.AddField("` + strings.Title(model[fieldField].(string)) +
			`","` + model[fieldField].(string) +
			`", db.` + getType(model[typeField].(string)) + `)
	`
	}

	content += `
	info.SetTable("` + table + `").SetTitle("` + strings.Title(table) + `").SetDescription("` + strings.Title(table) + `")

	formList := ` + table + `Table.GetForm()
	
	`

	for _, model := range columnsModel {

		typeName := getType(model[typeField].(string))
		formType := form.GetFormTypeFromFieldType(db.DT(strings.ToUpper(typeName)), model[fieldField].(string))

		content += `formList.AddField("` + strings.Title(model[fieldField].(string)) + `","` +
			model[fieldField].(string) + `",db.` + typeName + `,` + formType + `)
	`
	}

	content += `
	formList.SetTable("` + table + `").SetTitle("` + strings.Title(table) + `").SetDescription("` + strings.Title(table) + `")

	return ` + table + `Table
}`

	err := ioutil.WriteFile(outputPath+"/"+table+".go", []byte(content), 0644)
	checkError(err)
}

func generateTables(outputPath string, tables []string, packageName string) {

	tableStr := ""
	commentStr := ""

	for i := 0; i < len(tables); i++ {
		tableStr += `
	"` + tables[i] + `": Get` + strings.Title(tables[i]) + `Table,`
		commentStr += `// "` + tables[i] + `" => http://localhost:9033/admin/info/` + tables[i] + `
`
	}

	content := `package ` + packageName + `

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
	err := ioutil.WriteFile(outputPath+"/tables.go", []byte(content), 0644)
	checkError(err)
}

func getType(typeName string) string {
	r, _ := regexp.Compile(`\(.*?\)`)
	typeName = r.ReplaceAllString(typeName, "")
	r2, _ := regexp.Compile(`unsigned(.*)`)
	return strings.TrimSpace(strings.Title(r2.ReplaceAllString(typeName, "")))
}
