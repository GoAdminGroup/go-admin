// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"fmt"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	_ "github.com/chenhg5/go-admin/modules/db/mysql"
	_ "github.com/chenhg5/go-admin/modules/db/postgresql"
	cli "github.com/jawher/mow.cli"
	"github.com/manifoldco/promptui"
	"github.com/mgutz/ansi"
	"github.com/schollz/progressbar"
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
	promptSelect := promptui.Select{
		Label:     "choose a driver",
		Items:     []string{"mysql", "mssql", "postgresql", "sqlite"},
		Templates: selectTemplateWithTitle("choose a driver"),
	}

	_, driver, err := promptSelect.Run()

	checkError(err)

	var (
		cfg  map[string]config.Database
		name string
		conn = db.GetConnectionByDriver(driver)
	)

	if driver != "sqlite" {
		host := prompt("sql address")
		port := prompt("sql port")
		user := prompt("sql username")
		password := promptPassword("sql password")

		name = prompt("sql database name")

		if conn == nil {
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
		file := promptPassword("sql file")

		name = prompt("sql database name")

		if conn == nil {
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

	fmt.Println(ansi.Color("✔", "green") + " choose tables: ")

	tables := getTablesFromSqlResult(tableModels, driver, name)
	if len(tables) == 0 {
		panic("no tables")
	}
	chooseTables := make([]string, 0)
	var value string
	for {
		value = selects(tables)
		if value == "[finish]" {
			break
		}
		chooseTables = append(chooseTables, value)
		tables = removeItem(tables, value)
		if tables[0] == "[finish]" {
			break
		}
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
	generateTables(outputPath, chooseTables)

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color("generate success~~🍺🍺", "green"))
	fmt.Println()
	fmt.Println("see the docs: " + ansi.Color("http://doc.go-admin.cn/#/introduce/plugins/plugins", "blue"))
	fmt.Println()
	fmt.Println()
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

	return append(tables, "[finish]")
}

func selectTemplateWithTitle(title string) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}: ",
		Active:   ansi.Color("❯", "cyan") + " {{ . | cyan }}",
		Inactive: "  {{ . }}",
		Selected: ansi.Color("✔", "green") + " " + title + ": {{ . | cyan }}",
	}
}

func tableSelectTemplateWithTitle() *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   ansi.Color("❯", "cyan") + " {{ . | cyan }}",
		Inactive: "  {{ . }}",
		Selected: "    " + ansi.Color("✔", "green") + " {{ . | cyan }}",
		Help:     "Use the arrow keys to navigate: ↓ ↑ → ←, choose [finish] to exit",
	}
}

func promptTemplateWithTitle(title string) *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Success: ansi.Color("✔", "green") + " " + title + ": ",
	}
}

func prompt(label string) string {

	validate := func(input string) error {
		if input == "" {
			return errors.New(label + " is empty")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: promptTemplateWithTitle(label),
		Validate:  validate,
	}

	result, err := prompt.Run()

	checkError(err)

	return result
}

func promptWithDefault(label string, defaultValue string) string {

	validate := func(input string) error {
		if input == "" {
			return errors.New(label + " is empty")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:     label,
		Templates: promptTemplateWithTitle(label),
		Default:   defaultValue,
		Validate:  validate,
	}

	result, err := prompt.Run()

	checkError(err)

	return result
}

func promptPassword(label string) string {
	prompt := promptui.Prompt{
		Label:     label,
		Mask:      '*',
		Templates: promptTemplateWithTitle(label),
	}

	result, err := prompt.Run()

	checkError(err)

	return result
}

func selects(tables []string) string {
	promptSelect := promptui.Select{
		Label:     "choose table to generate",
		Items:     tables,
		Templates: tableSelectTemplateWithTitle(),
	}

	_, result, err := promptSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return ""
	}

	return result
}

func removeItem(tables []string, table string) []string {
	index := 0
	for i := 0; i < len(tables); i++ {
		if tables[i] == table {
			index = i
		}
	}
	return append(tables[:index], tables[index+1:]...)
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
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func Get` + strings.Title(table) + `Table() (` + table + `Table models.Table) {

	` + table + `Table.Info.FieldList = []types.Field{`

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

	` + table + `Table.Info.Table = "` + table + `"
	` + table + `Table.Info.Title = "` + strings.Title(table) + `"
	` + table + `Table.Info.Description = "` + strings.Title(table) + `"

	` + table + `Table.Form.FormList = []types.Form{`

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

	` + table + `Table.Form.Table = "` + table + `"
	` + table + `Table.Form.Title = "` + strings.Title(table) + `"
	` + table + `Table.Form.Description = "` + strings.Title(table) + `"

	` + table + `Table.Editable = true"
	` + table + `Table.ConnectionDriver = "` + driver + `"

	return
}`

	err := ioutil.WriteFile(outputPath+"/"+table+".go", []byte(content), 0644)
	checkError(err)
}

func generateTables(outputPath string, tables []string) {

	tableStr := ""

	for i := 0; i < len(tables); i++ {
		tableStr += `
	"` + tables[i] + `": Get` + strings.Title(tables[i]) + `Table,`
	}

	content := `package datamodel

import "github.com/chenhg5/go-admin/plugins/admin/models"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
var Generators = map[string]models.TableGenerator{` + tableStr + `
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
