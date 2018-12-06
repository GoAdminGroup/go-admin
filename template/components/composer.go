package components

import (
	"bytes"
	"fmt"
	"github.com/chenhg5/go-admin/modules/language"
	"html/template"
	"strings"
)

func ComposeHtml(temList map[string]string, compo interface{}, templateName ...string) template.HTML {
	var text = ""
	for _, v := range templateName {
		text += temList["components/"+v]
	}

	tmpla, err := template.New("comp").Funcs(template.FuncMap{
		"lang":     language.Get,
		"langHtml": language.GetFromHtml,
	}).Parse(text)
	if err != nil {
		panic("ComposeHtml Error:" + err.Error())
	}
	buffer := new(bytes.Buffer)

	defineName := strings.Replace(templateName[0], "table/", "", -1)
	defineName = strings.Replace(defineName, "form/", "", -1)

	err = tmpla.ExecuteTemplate(buffer, defineName, compo)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
