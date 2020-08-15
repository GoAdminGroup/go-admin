package components

import (
	"bytes"
	"html/template"

	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/utils"
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

	defineName := utils.ReplaceAll(templateName[0], "table/", "", "form/", "")

	err = tmpl.ExecuteTemplate(buffer, defineName, compo)
	if err != nil {
		logger.Error(tmplName+" ComposeHtml Error:", err)
	}
	return template.HTML(buffer.String())
}
