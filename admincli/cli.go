// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/AlecAivazis/survey"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/system"
	"github.com/chenhg5/go-admin/template/types/form"
	"github.com/go-bindata/go-bindata"
	cli "github.com/jawher/mow.cli"
	"github.com/mgutz/ansi"
	"github.com/schollz/progressbar"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			if errs, ok := err.(error); ok {
				fmt.Println()
				if runtime.GOOS == "windows" && errs.Error() == "Incorrect function." {
					fmt.Println(ansi.Color("go-admin cli error: cli has not supported MINGW64 for now, please use cmd terminal instead.", "red"))
					fmt.Println("know more here: http://forum.go-admin.cn/threads/2")
				} else {
					fmt.Println(ansi.Color("go-admin cli error: "+errs.Error(), "red"))
				}
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

	app.Command("compile", "compile template file for template or component", func(cmd *cli.Cmd) {
		cmd.Command("tpl", "compile template file for template or component", func(cmd *cli.Cmd) {
			var (
				rootPath   = cmd.StringOpt("path", "./template/adminlte/resource/pages/", "compile root path")
				outputPath = cmd.StringOpt("out", "./template/adminlte/tmpl/template.go", "compile output path")
			)

			cmd.Action = func() {
				compileTmpl(*rootPath, *outputPath)
			}
		})

		cmd.Command("asset", "compile asset file for template or component", func(cmd *cli.Cmd) {
			var (
				rootPath   = cmd.StringOpt("path", "./template/adminlte/resource/assets/", "compile root path")
				outputPath = cmd.StringOpt("out", "./template/adminlte/resource/assets.go", "compile output path")
			)

			cmd.Action = func() {
				compileAsset(*rootPath, *outputPath)
			}
		})
	})

	app.Command("generate", "generate table model files", func(cmd *cli.Cmd) {
		cmd.Action = func() {
			generating()
		}
	})

	_ = app.Run(os.Args)
}

func generating() {

	clear(runtime.GOOS)
	fmt.Println("GoAdmin CLI " + system.Version + compareVersion(system.Version))
	fmt.Println()

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

	var result = make(map[string]interface{})

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
				Host:       host,
				Port:       port,
				User:       user,
				Pwd:        password,
				Name:       name,
				MaxIdleCon: 50,
				MaxOpenCon: 150,
				Driver:     driver,
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
				Driver: driver,
				File:   file,
			},
		}
	}

	// step 1. test connection
	conn.InitDB(cfg)

	// step 2. show tables
	tableModels, _ := db.WithDriver(conn.GetName()).ShowTables()

	tables := getTablesFromSqlResult(tableModels, driver, name)
	if len(tables) == 0 {
		exitWithError(`no tables, you should build a table of your own business first.

see: http://www.go-admin.cn/en/docs/#/plugins/admin`)
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

func exitWithError(msg string) {
	fmt.Println()
	fmt.Println(ansi.Color("go-admin cli error: "+msg, "red"))
	fmt.Println()
	os.Exit(-1)
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

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func compileAsset(rootPath, outputPath string) {
	cfg := bindata.NewConfig()
	cfg.Output = outputPath
	cfg.Input = make([]bindata.InputConfig, 0)
	cfg.Input = append(cfg.Input, parseInput(rootPath))
	checkError(bindata.Translate(cfg))
}

func parseInput(path string) bindata.InputConfig {
	if strings.HasSuffix(path, "/...") {
		return bindata.InputConfig{
			Path:      filepath.Clean(path[:len(path)-4]),
			Recursive: true,
		}
	} else {
		return bindata.InputConfig{
			Path:      filepath.Clean(path),
			Recursive: false,
		}
	}
}

func compileTmpl(rootPath, outputPath string) {
	content := `package tmpl

var List = map[string]string{`

	content = getContentFromDir(content, fixPath(rootPath), fixPath(rootPath))

	content += `}`

	_ = ioutil.WriteFile(outputPath, []byte(content), 0644)
}

func fixPath(p string) string {
	if p[len(p)-1] != '/' {
		return p + "/"
	}
	return p
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
	"github.com/chenhg5/go-admin/template/types/form"
)

func Get` + strings.Title(table) + `Table() table.Table {

    ` + table + `Table := table.NewDefaultTable(table.DefaultConfigWithDriver("` + driver + `"))
	` + table + `Table.GetInfo().FieldList = []types.Field{`

	for _, model := range columnsModel {
		content += `{
			Head:     "` + strings.Title(model[fieldField].(string)) + `",
			Field:    "` + model[fieldField].(string) + `",
			TypeName: db.` + getType(model[typeField].(string)) + `,
			Sortable: false,
			FilterFn: func(model types.RowModel) interface{} {
				return model.Value
			},
		},`
	}

	content += `}

	` + table + `Table.GetInfo().Table = "` + table + `"
	` + table + `Table.GetInfo().Title = "` + strings.Title(table) + `"
	` + table + `Table.GetInfo().Description = "` + strings.Title(table) + `"

	` + table + `Table.GetForm().FormList = []types.Form{`

	// TODO: identify the form type from filed type

	for _, model := range columnsModel {

		typeName := getType(model[typeField].(string))
		formType := form.GetFormTypeFromFieldType(db.DT(strings.ToUpper(typeName)), model[fieldField].(string))

		content += `{
			Head:     "` + strings.Title(model[fieldField].(string)) + `",
			Field:    "` + model[fieldField].(string) + `",
			TypeName: db.` + typeName + `,
			Default:  "",
			Editable: true,
			FormType: ` + formType + `,
			FilterFn: func(model types.RowModel) interface{} {
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
	commentStr := ""

	for i := 0; i < len(tables); i++ {
		tableStr += `
	"` + tables[i] + `": Get` + strings.Title(tables[i]) + `Table,`
		commentStr += `// "` + tables[i] + `" => http://localhost:9033/admin/info/` + tables[i] + `
`
	}

	content := `package ` + packageName + `

import "github.com/chenhg5/go-admin/plugins/admin/modules/table"

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
	r, _ := regexp.Compile(`\\(.*\\)`)
	typeName = r.ReplaceAllString(typeName, "")
	return strings.Title(strings.Replace(typeName, " unsigned", "", -1))
}

func getLatestVersion() string {
	res, err := http.Get("https://goproxy.cn/github.com/chenhg5/go-admin/@v/list")

	if err != nil || res.Body == nil {
		return ""
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil || body == nil {
		return ""
	}

	versionsArr := strings.Split(string(body), "\n")

	return versionsArr[len(versionsArr)-1]
}

func isRequireUpdate(srcVersion, toCompareVersion string) bool {
	if toCompareVersion == "" {
		return false
	}

	srcVersionArr := strings.Split(strings.Replace(srcVersion, "v", "", -1), ".")
	toCompareVersionArr := strings.Split(strings.Replace(toCompareVersion, "v", "", -1), ".")

	for i := len(srcVersionArr) - 1; i > -1; i-- {
		v, err := strconv.Atoi(srcVersionArr[i])
		if err != nil {
			return false
		}
		vv, err := strconv.Atoi(toCompareVersionArr[i])
		if err != nil {
			return false
		}
		if v < vv {
			return true
		} else if v > vv {
			return false
		} else {
			continue
		}
	}

	return false
}

func compareVersion(srcVersion string) string {
	toCompareVersion := getLatestVersion()
	if isRequireUpdate(srcVersion, toCompareVersion) {
		return ", the latest version is " + toCompareVersion + " now."
	} else {
		return ""
	}
}
