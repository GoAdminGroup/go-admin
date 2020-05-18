package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/mgutz/ansi"
	"go/format"
	"gopkg.in/ini.v1"
	"io/ioutil"
	"os"
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

	if p.Theme == "" {
		p.Theme = singleSelect(getWord("choose a theme"), []string{"sword", "adminlte"}, "sword")
	}
	if p.Framework == "" {
		p.Framework = singleSelect(getWord("choose framework"),
			[]string{"gin", "beego", "buffalo", "fasthttp", "echo", "chi", "gf", "gorilla", "iris"}, "gin")
	}
	if p.Language == "" {
		p.Language = singleSelect(getWord("choose language"),
			[]string{getWord("cn"), getWord("en"), getWord("jp"), getWord("tc")}, "cn")
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
	}).Parse(ProjectTemplate[p.Framework])
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

	checkError(ioutil.WriteFile("./logs/access.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./logs/info.log", []byte{}, os.ModePerm))
	checkError(ioutil.WriteFile("./logs/error.log", []byte{}, os.ModePerm))

	if p.Theme == "sword" {
		checkError(ioutil.WriteFile("./index.go", swordIndexPage, os.ModePerm))
	} else {
		checkError(ioutil.WriteFile("./index.go", adminlteIndexPage, os.ModePerm))
	}

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
	fmt.Println(getWord("see the docs: ") + ansi.Color("http://doc.go-admin.cn",
		"blue"))
	fmt.Println(getWord("visit forum: ") + ansi.Color("http://discuss.go-admin.com",
		"blue"))
	fmt.Println()
	fmt.Println()
}
