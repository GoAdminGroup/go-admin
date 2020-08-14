package main

import (
	"errors"
	"fmt"
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

func generating(cfgFile, connName string) {

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

		if connection == "" {
			connection = connName
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

		survey.SelectQuestionTemplate = strings.ReplaceAll(survey.SelectQuestionTemplate, "<enter> to select",
			"<space> to select")

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
		generatePermissionFlag = singleSelect(getWord("generate permission records for tables"),
			[]string{getWord("yes"), getWord("no")}, getWord("yes"))
	}

	if generatePermissionFlag == getWord("yes") {
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

	if err := tools.GenerateTables(outputPath, packageName, chooseTables, true); err != nil {
		panic(err)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color(getWord("Generate data table models success~~🍺🍺"), "green"))
	fmt.Println()
	if defaultLang == "cn" {
		fmt.Println(getWord("see the docs: ") + ansi.Color("http://doc.go-admin.cn",
			"blue"))
	} else {
		fmt.Println(getWord("see the docs: ") + ansi.Color("https://book.go-admin.com",
			"blue"))
	}
	fmt.Println(getWord("visit forum: ") + ansi.Color("http://discuss.go-admin.com",
		"blue"))
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
