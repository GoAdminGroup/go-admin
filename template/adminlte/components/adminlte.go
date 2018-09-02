package components

import (
	"bytes"
	"html/template"
	"fmt"
	"github.com/chenhg5/go-admin/template/adminlte"
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
		text += adminlte.Adminlte["components/" + v]
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

func GetTemplate(isPjax bool) (tmpl *template.Template, name string) {
	var (
		err error
	)

	if !isPjax {
		name = "layout"
		tmpl, err = template.New("layout").Parse(adminlte.Adminlte["layout"] +
			adminlte.Adminlte["head"] + adminlte.Adminlte["header"] + adminlte.Adminlte["sidebar"] +
			adminlte.Adminlte["footer"] + adminlte.Adminlte["js"] + adminlte.Adminlte["menu"] +
			adminlte.Adminlte["admin_panel"] + adminlte.Adminlte["content"])
	} else {
		name = "content"
		tmpl, err = template.New("content").Parse(adminlte.Adminlte["admin_panel"] + adminlte.Adminlte["content"])
	}

	if err != nil {
		fmt.Println(err)
	}

	return
}