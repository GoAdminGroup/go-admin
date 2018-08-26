package adminlte

import (
	"bytes"
	"html/template"
	"fmt"
	tmp "goAdmin/template"
	"strings"
)

type AdminlteStruct struct {
	Name       string
	Components AdminlteComponents
}

type AdminlteComponents struct {

}

var Adminlte = AdminlteStruct{
	Name: "adminlte",
	Components: AdminlteComponents{},
}

func ComposeHtml(compo interface{}, templateName... string) template.HTML {
	var text = ""
	for _, v := range templateName {
		text += tmp.Adminlte["components/" + v]
	}

	tmpla, err := template.New("comp").Parse(text)
	if err != nil {
		fmt.Println("ComposeHtml Error:", err)
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

func GetTemplate(isPjax bool) *template.Template {
	var (
		tmpl *template.Template
		err error
	)

	if !isPjax {

		tmpl, err = template.New("layout").Parse(tmp.Adminlte["layout"] +
			tmp.Adminlte["head"] + tmp.Adminlte["header"] + tmp.Adminlte["sidebar"] +
			tmp.Adminlte["footer"] + tmp.Adminlte["js"] + tmp.Adminlte["menu"] +
			tmp.Adminlte["admin_panel"] + tmp.Adminlte["content"])

	} else {
		tmpl, err = template.New("content").Parse(tmp.Adminlte["admin_panel"] + tmp.Adminlte["content"])
	}

	if err != nil {
		fmt.Println(err)
	}

	return tmpl
}