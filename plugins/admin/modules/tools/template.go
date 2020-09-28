package tools

const tableModelTmpl = `{{define "table_model"}}
package {{.Package}}

import (
	{{.ExtraImport}}
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

	info := {{.Table}}.GetInfo(){{if .HideFilterArea}}.HideFilterArea(){{end}}

	{{if .HideNewButton}}info.HideNewButton(){{end}}
	{{if .HideExportButton}}info.HideExportButton(){{end}}
	{{if .HideEditButton}}info.HideEditButton(){{end}}
	{{if .HideDeleteButton}}info.HideDeleteButton(){{end}}
	{{if .HideDetailButton}}info.HideDetailButton(){{end}}
	{{if .HideFilterButton}}info.HideFilterButton(){{end}}
	{{if .HideRowSelector}}info.HideRowSelector(){{end}}
	{{if .HidePagination}}info.HidePagination(){{end}}
	{{if .HideQueryInfo}}info.HideQueryInfo(){{end}}
	{{if not .FilterFormLayout.Default}}info.SetFilterFormLayout(form.{{.FilterFormLayout.String}}){{end}}

	{{- range $key, $field := .Fields}}
	info.AddField("{{$field.Head}}", "{{$field.Name}}", db.{{$field.DBType}}){{if $field.Filterable}}.
		FieldFilterable(){{end -}}{{if $field.Sortable}}.
		FieldSortable(){{end -}}{{if $field.InfoEditable}}.
		FieldEditAble(){{end -}}{{if $field.Hide}}.
		FieldHide(){{end -}}
	{{- end}}

	info.SetTable("{{.TableName}}").SetTitle("{{.TablePageTitle}}").SetDescription("{{.TableDescription}}")

	formList := {{.Table}}.GetForm()

	{{- range $key, $field := .FormFields}}
	formList.AddField("{{$field.Head}}", "{{$field.Name}}", db.{{$field.DBType}}, form.{{$field.FormType}}){{if ne $field.Default ""}}.
		FieldDefault({{$field.Default}}){{end -}}{{if not $field.CanAdd}}.
		FieldDisableWhenCreate(){{end -}}{{if not $field.Editable}}.
		FieldDisableWhenUpdate(){{end -}}{{if $field.FormHide}}.
		FieldHide(){{end -}}{{if $field.EditHide}}.
		FieldHideWhenUpdate(){{end -}}{{if $field.CreateHide}}.
		FieldHideWhenCreate(){{end -}}{{$field.ExtraFun}}
	{{- end}}

	{{if .HideContinueEditCheckBox}}formList.HideContinueEditCheckBox(){{end}} 
	{{if .HideContinueNewCheckBox}}formList.HideContinueNewCheckBox(){{end}}  
	{{if .HideResetButton}}formList.HideResetButton(){{end}}          
	{{if .HideBackButton}}formList.HideBackButton(){{end}}           

	formList.SetTable("{{.TableName}}").SetTitle("{{.FormTitle}}").SetDescription("{{.FormDescription}}")


	{{if eq .DetailDisplay 1}}
		detail := {{.Table}}.GetDetailFromInfo()
	{{else if eq .DetailDisplay 2}}
		detail := {{.Table}}.GetDetail()
	{{end}}

	{{- range $key, $field := .DetailFields}}
	detail.AddField("{{$field.Head}}", "{{$field.Name}}", db.{{$field.DBType}})
	{{- end}}

	{{if ne .DetailDisplay 0}}
	detail.SetTable("{{.TableName}}").SetTitle("{{.TablePageTitle}}").SetDescription("{{.TableDescription}}")
	{{end}}

	{{.ExtraCode}}

	return {{.Table}}
}
{{end}}`
