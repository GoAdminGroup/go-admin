package main

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/tools"
	"github.com/mgutz/ansi"
	"github.com/schollz/progressbar"
	"gopkg.in/ini.v1"
)

var systemGoAdminTables = []string{
	"goadmin_menu",
	"goadmin_operation_log",
	"goadmin_permissions",
	"goadmin_role_menu",
	"goadmin_site",
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
		info = new(dbInfo)

		connection, packageName, outputPath, generatePermissionFlag string

		chooseTables = make([]string, 0)

		cfgModel *ini.File
		err      error
	)

	if cfgFile != "" {
		cfgModel, err = ini.Load(cfgFile)

		if err != nil {
			panic(errors.New("wrong config file path"))
		}

		languageCfg, err := cfgModel.GetSection("language")

		if err == nil {
			setDefaultLangSet(languageCfg.Key("language").Value())
		}

		modelCfgModel, exist2 := cfgModel.GetSection("model")

		if exist2 == nil {
			connection = modelCfgModel.Key("connection").Value()
			packageName = modelCfgModel.Key("package").Value()
			outputPath = modelCfgModel.Key("output").Value()
			generatePermissionFlag = modelCfgModel.Key("generate_permission_flag").Value()
		}

		info = getDBInfoFromINIConfig(cfgModel, connection)
	}

	// step 1. get connection
	conn := askForDBConnection(info)

	// step 2. show tables
	if len(chooseTables) == 0 {
		tables, err := db.WithDriver(conn).Table(info.Database).ShowTables()

		if err != nil {
			panic(err)
		}

		tables = filterTables(tables)

		if len(tables) == 0 {
			panic(newError(`no tables, you should build a table of your own business first.`))
		}
		tables = append([]string{"[" + getWord("select all") + "]"}, tables...)

		survey.SelectQuestionTemplate = strings.Replace(survey.SelectQuestionTemplate, "<enter> to select", "<space> to select", -1)

		chooseTables = selects(tables)
		if len(chooseTables) == 0 {
			panic(newError("no table is selected"))
		}
		if modules.InArray(chooseTables, "["+getWord("select all")+"]") {
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

	if generatePermissionFlag == "" {
		generatePermissionFlag = promptWithDefault("generate permission records for tables, Y on behalf of yes",
			"Y")
	}

	if generatePermissionFlag == "Y" {
		if connection == "default" {
			for _, table := range chooseTables {
				insertPermissionOfTable(conn, table)
			}
		} else {
			var defInfo = new(dbInfo)
			if cfgFile != "" {
				defInfo = getDBInfoFromINIConfig(cfgModel, "")
			}
			defConn := askForDBConnection(defInfo)
			for _, table := range chooseTables {
				insertPermissionOfTable(defConn, table)
			}
		}
	}

	fmt.Println()
	fmt.Println(ansi.Color("✔", "green") + " " + getWord("generating: "))
	fmt.Println()

	bar := progressbar.New(len(chooseTables))
	for i := 0; i < len(chooseTables); i++ {
		_ = bar.Add(1)
		time.Sleep(10 * time.Millisecond)
		checkError(tools.Generate(tools.NewParam(tools.Config{
			Connection:     connection,
			Driver:         info.DriverName,
			Package:        packageName,
			HideFilterArea: true,
			Table:          chooseTables[i],
			Schema:         info.Schema,
			Output:         outputPath,
			Conn:           conn,
		})))
	}
	generateTables(outputPath, chooseTables, packageName)

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color(getWord("Generate data table models success~~🍺🍺"), "green"))
	fmt.Println()
	fmt.Println(getWord("see the docs: ") + ansi.Color("http://doc.go-admin.cn/en/#/introduce/plugins/admin",
		"blue"))
	fmt.Println(getWord("visit forum: ") + ansi.Color("http://discuss.go-admin.com",
		"blue"))
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
			Prompt:   &survey.Input{Message: getWord(label)},
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
			Prompt:   &survey.Input{Message: getWord(label), Default: defaultValue},
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
		Message: getWord("sql password"),
	}
	err := survey.AskOne(prompt, &password, nil)

	checkError(err)

	return password
}

func selects(tables []string) []string {

	chooseTables := make([]string, 0)
	prompt := &survey.MultiSelect{
		Message:  getWord("choose table to generate"),
		Options:  tables,
		PageSize: 10,
	}
	err := survey.AskOne(prompt, &chooseTables, nil)

	checkError(err)

	return chooseTables
}

func singleSelect(msg string, options []string, def string) string {
	var qs = []*survey.Question{
		{
			Name: "question",
			Prompt: &survey.Select{
				Message: msg,
				Options: options,
				Default: def,
			},
		},
	}
	var result = make(map[string]interface{})
	err := survey.Ask(qs, &result)
	checkError(err)

	return result["question"].(core.OptionAnswer).Value

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