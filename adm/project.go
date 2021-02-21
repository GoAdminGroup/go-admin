package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	"github.com/GoAdminGroup/go-admin/modules/db"

	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/mgutz/ansi"
	"gopkg.in/ini.v1"
)

type Project struct {
	Port         string
	Theme        string
	Prefix       string
	Language     string
	Driver       string
	DriverModule string
	Framework    string
	Module       string
	Orm          string
}

func buildProject(cfgFile string) {

	clear(runtime.GOOS)
	cliInfo()

	var (
		p        Project
		cfgModel *ini.File
		err      error
		info     = new(dbInfo)
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

		projectCfgModel, err := cfgModel.GetSection("project")

		if err == nil {
			p.Theme = projectCfgModel.Key("theme").Value()
			p.Framework = projectCfgModel.Key("framework").Value()
			p.Language = projectCfgModel.Key("language").Value()
			p.Port = projectCfgModel.Key("port").Value()
			p.Driver = projectCfgModel.Key("driver").Value()
			p.Prefix = projectCfgModel.Key("prefix").Value()
			p.Module = projectCfgModel.Key("module").Value()
			p.Orm = projectCfgModel.Key("orm").Value()
		}

		info = getDBInfoFromINIConfig(cfgModel, "default")
	}

	// generate main.go

	initSurvey()

	if p.Module == "" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic(err)
		}
		p.Module = promptWithDefault(getWord("module path"), filepath.Base(dir))
	}
	if p.Framework == "" {
		p.Framework = singleSelect(getWord("choose framework"),
			[]string{"gin", "beego", "buffalo", "fasthttp", "echo", "chi", "gf", "gorilla", "iris"}, "gin")
	}
	if p.Theme == "" {
		p.Theme = singleSelect(getWord("choose a theme"), template2.DefaultThemeNames, "sword")
	}
	if p.Language == "" {
		p.Language = singleSelect(getWord("choose language"),
			[]string{getWord("cn"), getWord("en"), getWord("jp"), getWord("tc")},
			getWord("cn"))
		switch p.Language {
		case getWord("cn"):
			p.Language = language.CN
		case getWord("en"):
			p.Language = language.EN
		case getWord("jp"):
			p.Language = language.JP
		case getWord("tc"):
			p.Language = language.TC
		}
	}
	if p.Port == "" {
		p.Port = promptWithDefault(getWord("port"), "80")
	}
	if p.Prefix == "" {
		p.Prefix = promptWithDefault(getWord("url prefix"), "admin")
	}
	if p.Driver == "" {
		p.Driver = singleSelect(getWord("choose a driver"),
			[]string{"mysql", "postgresql", "sqlite", "mssql"}, "mysql")
	}
	p.DriverModule = p.Driver
	if p.Driver == db.DriverPostgresql {
		p.DriverModule = "postgres"
	}

	rootPath, err := os.Getwd()

	if err != nil {
		rootPath = "."
	} else {
		rootPath = filepath.ToSlash(rootPath)
	}

	var cfg = config.SetDefault(&config.Config{
		Debug: true,
		Env:   config.EnvLocal,
		Theme: p.Theme,
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:          p.Language,
		UrlPrefix:         p.Prefix,
		IndexUrl:          "/",
		AccessLogPath:     rootPath + "/logs/access.log",
		ErrorLogPath:      rootPath + "/logs/error.log",
		InfoLogPath:       rootPath + "/logs/info.log",
		BootstrapFilePath: rootPath + "/bootstrap.go",
		GoModFilePath:     rootPath + "/go.mod",
	})

	if info.DriverName == "" && p.Driver != "" {
		info.DriverName = p.Driver
	}

	cfg.Databases = askForDBConfig(info)

	if p.Orm == "" {
		p.Orm = singleSelect(getWord("choose a orm"),
			[]string{getWord("none"), "gorm"}, getWord("none"))
		if p.Orm == getWord("none") {
			p.Orm = ""
		}
	}

	installProjectTmpl(p, cfg, cfgFile, info)

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color(getWord("Generate project template success~~🍺🍺"), "green"))
	fmt.Println()
	fmt.Println(getWord("1 Import and initialize database:"))
	fmt.Println()
	if defaultLang == "cn" || p.Language == language.CN || p.Language == "cn" {
		fmt.Println("- sqlite: " + ansi.Color("https://gitee.com/go-admin/go-admin/raw/master/data/admin.db", "blue"))
		fmt.Println("- mssql: " + ansi.Color("https://gitee.com/go-admin/go-admin/raw/master/data/admin.mssql", "blue"))
		fmt.Println("- postgresql: " + ansi.Color("https://gitee.com/go-admin/go-admin/raw/master/data/admin.pgsql", "blue"))
		fmt.Println("- mysql: " + ansi.Color("https://gitee.com/go-admin/go-admin/raw/master/data/admin.sql", "blue"))
	} else {
		fmt.Println("- sqlite: " + ansi.Color("https://github.com/GoAdminGroup/go-admin/raw/master/data/admin.db", "blue"))
		fmt.Println("- mssql: " + ansi.Color("https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.mssql", "blue"))
		fmt.Println("- postgresql: " + ansi.Color("https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.pgsql", "blue"))
		fmt.Println("- mysql: " + ansi.Color("https://raw.githubusercontent.com/GoAdminGroup/go-admin/master/data/admin.sql", "blue"))
	}
	fmt.Println()
	fmt.Println(getWord("2 Execute the following command to run:"))
	fmt.Println()
	if runtime.GOOS == "windows" {
		fmt.Println("> GO111MODULE=on go mod init " + p.Module)
		if defaultLang == "cn" || p.Language == language.CN || p.Language == "cn" {
			fmt.Println("> GORPOXY=https://goproxy.io GO111MODULE=on go mod tidy")
		} else {
			fmt.Println("> GO111MODULE=on go mod tidy")
		}
		fmt.Println("> GO111MODULE=on go run .")
	} else {
		fmt.Println("> make init module=" + p.Module)
		if defaultLang == "cn" || p.Language == language.CN || p.Language == "cn" {
			fmt.Println("> GORPOXY=https://goproxy.io make install")
		} else {
			fmt.Println("> make install")
		}
		fmt.Println("> make serve")
	}
	fmt.Println()
	fmt.Println(getWord("3 Visit and login:"))
	fmt.Println()
	if p.Port != "80" {
		fmt.Println("-  " + getWord("Login: ") + ansi.Color("http://127.0.0.1:"+p.Port+"/"+p.Prefix+"/login", "blue"))
	} else {
		fmt.Println("-  " + getWord("Login: ") + ansi.Color("http://127.0.0.1/"+p.Prefix+"/login", "blue"))
	}
	fmt.Println(getWord("account: admin  password: admin"))
	fmt.Println()
	fmt.Println("-  " + getWord("Generate CRUD models: ") + ansi.Color("http://127.0.0.1:"+p.Port+"/"+p.Prefix+"/info/generate/new", "blue"))
	fmt.Println()
	fmt.Println(getWord("4 See more in README.md"))
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

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(dir, "\\", "/")
}

func installProjectTmpl(p Project, cfg *config.Config, cfgFile string, info *dbInfo) {

	t, err := template.New("project").Funcs(map[string]interface{}{
		"title": strings.Title,
	}).Parse(projectTemplate[p.Framework])
	checkError(err)
	buf := new(bytes.Buffer)
	checkError(t.Execute(buf, p))
	c, err := format.Source(buf.Bytes())
	checkError(err)
	checkError(ioutil.WriteFile("./main.go", c, 0644))

	checkError(os.Mkdir("pages", os.ModePerm))
	checkError(os.Mkdir("tables", os.ModePerm))
	checkError(os.Mkdir("logs", os.ModePerm))
	checkError(os.Mkdir("uploads", os.ModePerm))
	checkError(os.Mkdir("html", os.ModePerm))
	checkError(os.Mkdir("build", os.ModePerm))
	if p.Orm == "gorm" {
		checkError(os.Mkdir("models", os.ModePerm))
		checkError(ioutil.WriteFile("./models/base.go", []byte(`package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	orm, err = gorm.Open("`+p.DriverModule+`", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
}
`), os.ModePerm))
	}

	checkError(ioutil.WriteFile("./logs/access.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./logs/info.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./logs/error.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./bootstrap.go", []byte(`package main`), os.ModePerm))

	if p.Theme == "sword" {
		checkError(ioutil.WriteFile("./pages/index.go", swordIndexPage, 0644))
	} else {
		checkError(ioutil.WriteFile("./pages/index.go", adminlteIndexPage, 0644))
	}

	if defaultLang == "cn" || p.Language == language.CN || p.Language == "cn" {
		checkError(ioutil.WriteFile("./main_test.go", mainTestCN, 0644))
		checkError(ioutil.WriteFile("./README.md", []byte(fmt.Sprintf(readmeCN, p.Port+"/"+p.Prefix)), 0644))
	} else {
		checkError(ioutil.WriteFile("./main_test.go", mainTest, 0644))
		checkError(ioutil.WriteFile("./README.md", []byte(fmt.Sprintf(readme, p.Port+"/"+p.Prefix)), 0644))
	}

	makefileContent := makefile
	if p.Driver == db.DriverSqlite {
		makefileContent = bytes.ReplaceAll(makefileContent, []byte("CGO_ENABLED=0"), []byte("CGO_ENABLED=1"))
	}

	checkError(ioutil.WriteFile("./Makefile", makefileContent, 0644))

	checkError(ioutil.WriteFile("./html/hello.tmpl", []byte(`<div class="hello">
    <h1>{{index . "msg"}}</h1>
</div>

<style>
    .hello {
        padding: 50px;
        width: 100%;
        text-align: center;
    }
</style>
`), 0644))

	checkError(ioutil.WriteFile("./tables/tables.go", []byte(`// This file is generated by GoAdmin CLI adm.
package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// Generators is a map of table models.
//
// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// example end
//
var Generators = map[string]table.Generator{

	// generators end
}
`), 0644))

	// generate config json

	if cfgFile == "" {
		t, err := template.New("ini").Parse(admINI)
		checkError(err)
		buf := new(bytes.Buffer)
		checkError(t.Execute(buf, info))
		checkError(ioutil.WriteFile("./adm.ini", buf.Bytes(), 0644))
	}

	configByte, err := json.MarshalIndent(cfg, "", "	")
	configByte = bytes.ReplaceAll(configByte, []byte(`
	"logger": {
		"encoder": {},
		"rotate": {}
	},`), []byte{})
	configByte = bytes.ReplaceAll(configByte, []byte(`,
	"animation": {}`), []byte{})
	checkError(err)
	checkError(ioutil.WriteFile("./config.json", configByte, 0644))

}
