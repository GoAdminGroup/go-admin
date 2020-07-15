package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/mgutz/ansi"
)

func getPluginTemplate(modulePath, pluginName string) {
	data := PluginTmplData{
		ModulePath:  modulePath,
		PluginName:  pluginName,
		PluginTitle: strings.Title(pluginName),
	}
	checkError(ioutil.WriteFile("./"+pluginName+".go", parsePluginTmpl("main", data), 0644))
	checkError(ioutil.WriteFile("./router.go", parsePluginTmpl("router", data), 0644))
	checkError(os.Mkdir("controller", os.ModePerm))
	checkError(ioutil.WriteFile("./controller/common.go", parsePluginTmpl("controller", data), 0644))
	checkError(ioutil.WriteFile("./controller/example.go", parsePluginTmpl("controller_example", data), 0644))
	checkError(os.Mkdir("guard", os.ModePerm))
	checkError(ioutil.WriteFile("./guard/guard.go", parsePluginTmpl("guard", data), 0644))
	checkError(ioutil.WriteFile("./guard/example.go", parsePluginTmpl("guard_example", data), 0644))
	checkError(os.Mkdir("modules", os.ModePerm))
	checkError(os.Mkdir("./modules/language", os.ModePerm))
	checkError(ioutil.WriteFile("./modules/language/language.go", parsePluginTmpl("language", data), 0644))
	checkError(ioutil.WriteFile("./modules/language/cn.go", parsePluginTmpl("language_cn", data), 0644))
	checkError(ioutil.WriteFile("./modules/language/en.go", parsePluginTmpl("language_en", data), 0644))
	checkError(ioutil.WriteFile("./Makefile", []byte(pluginTemplate["makefile"]), 0644))
	checkError(os.Mkdir("example", os.ModePerm))
	checkError(os.Mkdir("tests", os.ModePerm))
	fmt.Println()
	fmt.Println(ansi.Color(getWord("Generate plugin template success~~🍺🍺"), "green"))
	fmt.Println()
}

type PluginTmplData struct {
	PluginName  string
	PluginTitle string
	ModulePath  string
}

func parsePluginTmpl(name string, data PluginTmplData) []byte {
	t, err := template.New("plugin").Parse(pluginTemplate[name])
	checkError(err)
	buf := new(bytes.Buffer)
	checkError(t.Execute(buf, data))
	c, err := format.Source(buf.Bytes())
	checkError(err)
	return c
}
