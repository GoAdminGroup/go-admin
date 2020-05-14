package tools

const tableModelTmpl = `{{define "table_model"}}
package {{.Package}}

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func Get{{.TableTitle}}Table(ctx *context.Context) table.Table {
	
	{{if eq .Connection "default"}}
	{{.Table}} := table.NewDefaultTable(table.DefaultConfigWithDriver("{{.Driver}}"))
	{{else}}
	{{.Table}} := table.NewDefaultTable(table.DefaultConfigWithDriverAndConnection("{{.Driver}}", "{{.Connection}}"))
	{{end}}

	info := {{.Table}}.GetInfo()

	{{- range $key, $field := .Fields}}
	info.AddField("{{$field.Head}}", "{{$field.Name}}", db.{{$field.DBType}}){{if $field.Filterable}}.FieldFilterable(){{end -}}
	{{- end}}

	info.SetTable("{{.TableName}}").SetTitle("{{.TableTitle}}").SetDescription("{{.TableTitle}}")

	formList := {{.Table}}.GetForm()

	{{- range $key, $field := .Fields}}
	formList.AddField("{{$field.Head}}", "{{$field.Name}}", db.{{$field.DBType}}, form.{{$field.FormType}}){{if $field.NotAllowAdd}}.FieldNotAllowAdd(){{end -}}
	{{- end}}

	formList.SetTable("{{.TableName}}").SetTitle("{{.TableTitle}}").SetDescription("{{.TableTitle}}")

	return {{.Table}}
}
{{end}}`
