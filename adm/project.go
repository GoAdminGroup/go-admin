package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/mgutz/ansi"
	"go/format"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"text/template"
)

type Project struct {
	Port      string
	Theme     string
	Prefix    string
	Language  string
	Driver    string
	Framework string
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
		}

		info = getDBInfoFromINIConfig(cfgModel, "default")
	}

	// generate main.go

	initSurvey()

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

	t, err := template.New("project").Funcs(map[string]interface{}{
		"title": func(s string) string {
			return strings.Title(s)
		},
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

	checkError(ioutil.WriteFile("./logs/access.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./logs/info.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./logs/error.log", []byte{}, os.ModePerm))

	if p.Theme == "sword" {
		checkError(ioutil.WriteFile("./index.go", swordIndexPage, 0644))
	} else {
		checkError(ioutil.WriteFile("./index.go", adminlteIndexPage, 0644))
	}

	if defaultLang == "cn" || p.Language == language.CN || p.Language == "cn" {
		checkError(ioutil.WriteFile("./main_test.go", mainTestCN, 0644))
		checkError(ioutil.WriteFile("./README.md", readmeCN, 0644))
	} else {
		checkError(ioutil.WriteFile("./main_test.go", mainTest, 0644))
		checkError(ioutil.WriteFile("./README.md", readme, 0644))
	}

	checkError(ioutil.WriteFile("./Makefile", makefile, 0644))

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

	// generate config json

	var cfg = config.SetDefault(config.Config{
		Debug: true,
		Env:   config.EnvLocal,
		Theme: p.Theme,
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:      p.Language,
		UrlPrefix:     p.Prefix,
		IndexUrl:      "/",
		AccessLogPath: "./logs/access.log",
		ErrorLogPath:  "./logs/error.log",
		InfoLogPath:   "./logs/info.log",
	})

	if info.DriverName == "" && p.Driver != "" {
		info.DriverName = p.Driver
	}

	cfg.Databases = askForDBConfig(info)

	if cfgFile == "" {
		t, err := template.New("ini").Parse(admINI)
		checkError(err)
		buf := new(bytes.Buffer)
		checkError(t.Execute(buf, info))
		checkError(ioutil.WriteFile("./adm.ini", buf.Bytes(), 0644))
	}

	configByte, err := json.MarshalIndent(cfg, "", "	")
	configByte = bytes.Replace(configByte, []byte(`
	"logger": {
		"encoder": {},
		"rotate": {}
	},`), []byte{}, -1)
	configByte = bytes.Replace(configByte, []byte(`,
	"animation": {}`), []byte{}, -1)
	checkError(err)
	checkError(ioutil.WriteFile("./config.json", configByte, 0644))

	fmt.Println()
	fmt.Println()
	fmt.Println(ansi.Color(getWord("Generate project template success~~🍺🍺"), "green"))
	fmt.Println()
	fmt.Println(getWord("1 Initialize database:"))
	fmt.Println()
	fmt.Println("- sqlite: " + ansi.Color("https://github.com/GoAdminGroup/go-admin/blob/master/data/admin.db", "blue"))
	fmt.Println("- mssql: " + ansi.Color("https://github.com/GoAdminGroup/go-admin/blob/master/data/admin.mssql", "blue"))
	fmt.Println("- postgresql: " + ansi.Color("https://github.com/GoAdminGroup/go-admin/blob/master/data/admin.pgsql", "blue"))
	fmt.Println("- mysql: " + ansi.Color("https://github.com/GoAdminGroup/go-admin/blob/master/data/admin.sql", "blue"))
	fmt.Println()
	fmt.Println(getWord("2 Execute the following command to run:"))
	fmt.Println()
	if runtime.GOOS == "windows" {
		fmt.Println("> go mod init xxxx.com/xxxxx/xxxx")
		fmt.Println("> go mod tidy")
		fmt.Println("> GO111MODULE=on go run .")
	} else {
		fmt.Println("> make init module=xxxx.com/xxxxx/xxxx")
		fmt.Println("> make install")
		fmt.Println("> make serve")
	}
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
