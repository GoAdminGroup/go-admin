// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	_ "github.com/chenhg5/go-admin/modules/db/mysql"
	_ "github.com/chenhg5/go-admin/modules/db/postgresql"
	cli "github.com/jawher/mow.cli"
	"github.com/mgutz/ansi"
	"github.com/schollz/progressbar"
	"gopkg.in/AlecAivazis/survey.v1"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			if errs, ok := err.(error); ok {
				fmt.Println()
				fmt.Println(ansi.Color("go-admin cli error: "+errs.Error(), "red"))
				fmt.Println()
			}
		}
	}()

	app := cli.App("go-admin cli tool", "cli tool for developing and generating")

	app.Spec = "[-v]"

	var verbose = app.BoolOpt("v verbose", false, "out debug info")

	app.Before = func() {
		if *verbose {
			fmt.Println("debug mode is on")
		}
	}

	app.Command("compile", "compile template file for developing", func(cmd *cli.Cmd) {
		var (
			rootPath   = cmd.StringOpt("path", "./template/adminlte/resource/pages/", "compile root path")
			outputPath = cmd.StringOpt("out", "./template/adminlte/tmpl/template.go", "compile output path")
		)

		cmd.Action = func() {
			compileTmpl(*rootPath, *outputPath)
		}
	})

	app.Command("generate", "generate table model files", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			generating()
		}
	})

	_ = app.Run(os.Args)
}

func generating() {

	print("\033[H\033[2J")
	fmt.Println("GoAdmin CLI v1.0.0")

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "space to select", "<enter> to select", -1)

	var qs = []*survey.Question{
		{
			Name: "driver",
			Prompt: &survey.Select{
				Message: "choose a driver",
				Options: []string{"mysql", "mssql", "postgresql", "sqlite"},
				Default: "mysql",
			},
		},
	}

	var result = make(map[string]interface{}, 0)

	err := survey.Ask(qs, &result)
	checkError(err)
	driver := result["driver"].(string)

	var (
		cfg  map[string]config.Database
		name string
		conn = db.GetConnectionByDriver(driver)
	)

	if driver != "sqlite" {
		host := promptWithDefault("sql address", "127.0.0.1")
		port := promptWithDefault("sql port", "3306")
		user := promptWithDefault("sql username", "root")
		password := promptPassword()

		name = prompt("sql database name")

		if conn == nil {
			exitWithError("invalid db connection")
			panic("invalid db connection")
		}
		cfg = map[string]config.Database{
			"default": {
				HOST:         host,
				PORT:         port,
				USER:         user,
				PWD:          password,
				NAME:         name,
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       driver,
				FILE:         "",
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
				DRIVER: driver,
				FILE:   file,
			},
		}
	}

	// step 1. test connection
	conn.InitDB(cfg)

	// step 2. show tables
	tableModels, _ := db.WithDriver(conn.GetName()).ShowTables()

	tables := getTablesFromSqlResult(tableModels, driver, name)
	if len(tables) == 0 {
		exitWithError("no tables")
	}

	survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "<enter> to select", "<space> to select", -1)

	chooseTables := selects(tables)
	if len(chooseTables) == 0 {
		exitWithError("no choosing tables")
	}

	packageName := promptWithDefault("set package name", "main")
	outputPath := promptWithDefault("set file output path", "./")

	fmt.Println(ansi.Color("✔", "green") + " generating: ")
	fmt.Println()

	fieldField := "Field"
	typeField := "Type"
	if driver == "postgresql" {
		fieldField = "column_name"
		typeField = "udt_name"
	}

	bar := progressbar.New(len(chooseTables))
	for i := 0; i < len(chooseTables); i++ {
		_ = bar.Add(1)
		time.Sleep(10 * time.Millisecond)
		generateFile(chooseTables[i], conn, fieldField, typeField, packageName, driver, outputPath)
	}
	generateTables(outputPath, chooseTables, packageName)

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color("generate success~~🍺🍺", "green"))
	fmt.Println()
	fmt.Println("see the docs: " + ansi.Color("http://doc.go-admin.cn/#/introduce/plugins/admin", "blue"))
	fmt.Println()
	fmt.Println()
}

func exitWithError(msg string) {
	fmt.Println()
	fmt.Println(ansi.Color("go-admin cli error: "+msg, "red"))
	fmt.Println()
	os.Exit(-1)
}

func getTablesFromSqlResult(models []map[string]interface{}, driver string, dbName string) []string {
	key := "Tables_in_" + dbName
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

	var result = make(map[string]interface{}, 0)

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

	var result = make(map[string]interface{}, 0)

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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func compileTmpl(rootPath, outputPath string) {
	content := `package tmpl

var List = map[string]string{`

	content = getContentFromDir(content, rootPath, rootPath)

	content += `}`

	_ = ioutil.WriteFile(outputPath, []byte(content), 0644)
}

func getContentFromDir(content, dirPath, rootPath string) string {
	files, _ := ioutil.ReadDir(dirPath)

	for _, f := range files {

		if f.IsDir() {
			content = getContentFromDir(content, dirPath+f.Name()+"/", rootPath)
			continue
		}

		b, err := ioutil.ReadFile(dirPath + f.Name())
		if err != nil {
			fmt.Print(err)
		}
		str := string(b)

		suffix := path.Ext(f.Name())
		onlyName := strings.TrimSuffix(f.Name(), suffix)

		if suffix == ".tmpl" {
			fmt.Println(dirPath + f.Name())
			content += `"` + strings.Replace(dirPath, rootPath, "", -1) + onlyName + `":` + "`" + str + "`,"
		}
	}

	return content
}

func generateFile(table string, conn db.Connection, fieldField, typeField, packageName, driver, outputPath string) {

	columnsModel, _ := db.WithDriver(conn.GetName()).Table(table).ShowColumns()

	content := `package ` + packageName + `

import (
	"github.com/chenhg5/go-admin/plugins/admin/modules/table"
	"github.com/chenhg5/go-admin/template/types"	
)

func Get` + strings.Title(table) + `Table() table.Table {

    ` + table + `Table := table.NewDefaultTable(table.DefaultTableConfigWithDriver("` + driver + `"))
	` + table + `Table.GetInfo().FieldList = []types.Field{`

	for _, model := range columnsModel {
		content += `{
			Head:     "` + strings.Title(model[fieldField].(string)) + `",
			Field:    "` + model[fieldField].(string) + `",
			TypeName: "` + GetType(model[typeField].(string)) + `",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},`
	}

	content += `}

	` + table + `Table.GetInfo().Table = "` + table + `"
	` + table + `Table.GetInfo().Title = "` + strings.Title(table) + `"
	` + table + `Table.GetInfo().Description = "` + strings.Title(table) + `"

	` + table + `Table.GetForm().FormList = []types.Form{`

	for _, model := range columnsModel {

		formType := "text"
		if model[fieldField].(string) == "id" {
			formType = "default"
		}

		content += `{
			Head:     "` + strings.Title(model[fieldField].(string)) + `",
			Field:    "` + model[fieldField].(string) + `",
			TypeName: "` + GetType(model[typeField].(string)) + `",
			Default:  "",
			Editable: false,
			FormType: "` + formType + `",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},`
	}

	content += `	}

	` + table + `Table.GetForm().Table = "` + table + `"
	` + table + `Table.GetForm().Title = "` + strings.Title(table) + `"
	` + table + `Table.GetForm().Description = "` + strings.Title(table) + `"


	return ` + table + `Table
}`

	err := ioutil.WriteFile(outputPath+"/"+table+".go", []byte(content), 0644)
	checkError(err)
}

func generateTables(outputPath string, tables []string, packageName string) {

	tableStr := ""

	for i := 0; i < len(tables); i++ {
		tableStr += `
	"` + tables[i] + `": Get` + strings.Title(tables[i]) + `Table,`
	}

	content := `package ` + packageName + `

import "github.com/chenhg5/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
var Generators = map[string]table.Generator{` + tableStr + `
}
`
	err := ioutil.WriteFile(outputPath+"/tables.go", []byte(content), 0644)
	checkError(err)
}

func GetType(typeName string) string {
	r, _ := regexp.Compile("\\(.*\\)")
	typeName = r.ReplaceAllString(typeName, "")
	return strings.ToLower(strings.Replace(typeName, " unsigned", "", -1))
}
