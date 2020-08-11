package components

import (
	"bytes"
	"html/template"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/logger"

	template2 "github.com/GoAdminGroup/go-admin/template"
)

func ComposeHtml(temList map[string]string, compo interface{}, templateName ...string) template.HTML {
	var text = ""
	for _, v := range templateName {
		text += temList["components/"+v]
	}

	tmplName := ""
	if len(templateName) > 0 {
		tmplName = templateName[0] + " "
	}

	tmpl, err := template.New("comp").Funcs(template2.DefaultFuncMap).Parse(text)
	if err != nil {
		logger.Panic(tmplName + "ComposeHtml Error:" + err.Error())
		return ""
	}
	buffer := new(bytes.Buffer)

	defineName := strings.Replace(templateName[0], "table/", "", -1)
	defineName = strings.Replace(defineName, "form/", "", -1)

	err = tmpl.ExecuteTemplate(buffer, defineName, compo)
	if err != nil {
		logger.Error(tmplName+"ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
