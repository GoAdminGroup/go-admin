package tools

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

type Param struct {
	Connection string `json:"connection"`
	Driver     string `json:"driver"`
	Package    string `json:"package"`

	Table      string `json:"table"`
	RowTable   string `json:"row_table"`
	TableTitle string `json:"table_title"`
	TableName  string `json:"table_name"`

	HideFilterArea   bool        `json:"hide_filter_area"`
	HideNewButton    bool        `json:"hide_new_button"`
	HideExportButton bool        `json:"hide_export_button"`
	HideEditButton   bool        `json:"hide_edit_button"`
	HideDeleteButton bool        `json:"hide_delete_button"`
	HideDetailButton bool        `json:"hide_detail_button"`
	HideFilterButton bool        `json:"hide_filter_button"`
	HideRowSelector  bool        `json:"hide_row_selector"`
	HidePagination   bool        `json:"hide_pagination"`
	HideQueryInfo    bool        `json:"hide_query_info"`
	FilterFormLayout form.Layout `json:"filter_form_layout"`

	HideContinueEditCheckBox bool `json:"hide_continue_edit_check_box"`
	HideContinueNewCheckBox  bool `json:"hide_continue_new_check_box"`
	HideResetButton          bool `json:"hide_reset_button"`
	HideBackButton           bool `json:"hide_back_button"`

	Fields Fields

	FormFields Fields

	Output string `json:"output"`
}

type Config struct {
	Connection       string        `json:"connection"`
	Driver           string        `json:"driver"`
	Package          string        `json:"package"`
	Table            string        `json:"table"`
	Schema           string        `json:"schema"`
	Output           string        `json:"output"`
	Conn             db.Connection `json:"conn"`
	HideFilterArea   bool          `json:"hide_filter_area"`
	HideNewButton    bool          `json:"hide_new_button"`
	HideExportButton bool          `json:"hide_export_button"`
	HideEditButton   bool          `json:"hide_edit_button"`
	HideDeleteButton bool          `json:"hide_delete_button"`
	HideDetailButton bool          `json:"hide_detail_button"`
	HideFilterButton bool          `json:"hide_filter_button"`
	HideRowSelector  bool          `json:"hide_row_selector"`
	HidePagination   bool          `json:"hide_pagination"`
	HideQueryInfo    bool          `json:"hide_query_info"`
	FilterFormLayout form.Layout   `json:"filter_form_layout"`

	HideContinueEditCheckBox bool `json:"hide_continue_edit_check_box"`
	HideContinueNewCheckBox  bool `json:"hide_continue_new_check_box"`
	HideResetButton          bool `json:"hide_reset_button"`
	HideBackButton           bool `json:"hide_back_button"`
}

func fixedTable(table string) string {
	if utils.InArray(keyWords, table) {
		return table + "_"
	}
	return table
}

var keyWords = []string{"import", "package", "chan", "const", "func", "interface", "map", "struct", "type",
	"var", "break", "case", "continue", "default", "defer", "else", "fallthrough", "for", "go", "goto", "if",
	"range", "return", "select", "switch"}

func NewParam(cfg Config) Param {
	ta := camelcase(cfg.Table)
	dbTable := cfg.Table
	if cfg.Schema != "" {
		dbTable = cfg.Schema + "." + cfg.Table
	}

	fields := getFieldsFromConn(cfg.Conn, dbTable, cfg.Driver)

	return Param{
		Connection:               cfg.Connection,
		Driver:                   cfg.Driver,
		Package:                  cfg.Package,
		Table:                    fixedTable(ta),
		TableTitle:               strings.Title(ta),
		TableName:                dbTable,
		HideFilterArea:           cfg.HideFilterArea,
		HideNewButton:            cfg.HideNewButton,
		HideExportButton:         cfg.HideExportButton,
		HideEditButton:           cfg.HideEditButton,
		HideDeleteButton:         cfg.HideDeleteButton,
		HideDetailButton:         cfg.HideDetailButton,
		HideFilterButton:         cfg.HideFilterButton,
		HideRowSelector:          cfg.HideRowSelector,
		HidePagination:           cfg.HidePagination,
		HideQueryInfo:            cfg.HideQueryInfo,
		FilterFormLayout:         cfg.FilterFormLayout,
		HideContinueEditCheckBox: cfg.HideContinueEditCheckBox,
		HideContinueNewCheckBox:  cfg.HideContinueNewCheckBox,
		HideResetButton:          cfg.HideResetButton,
		HideBackButton:           cfg.HideBackButton,
		RowTable:                 cfg.Table,
		Fields:                   fields,
		FormFields:               fields,
		Output:                   cfg.Output,
	}
}

func NewParamWithFields(cfg Config, fields ...Fields) Param {
	ta := camelcase(cfg.Table)
	dbTable := cfg.Table
	if cfg.Schema != "" {
		dbTable = cfg.Schema + "." + cfg.Table
	}

	if len(cfg.Output) > 0 && cfg.Output[len(cfg.Output)-1] == '/' {
		cfg.Output = cfg.Output[:len(cfg.Output)-1]
	}

	return Param{
		Connection:               cfg.Connection,
		Driver:                   cfg.Driver,
		Package:                  cfg.Package,
		Table:                    ta,
		TableTitle:               strings.Title(ta),
		TableName:                dbTable,
		RowTable:                 cfg.Table,
		Fields:                   fields[0],
		FormFields:               fields[1],
		HideFilterArea:           cfg.HideFilterArea,
		HideNewButton:            cfg.HideNewButton,
		HideExportButton:         cfg.HideExportButton,
		HideEditButton:           cfg.HideEditButton,
		HideDeleteButton:         cfg.HideDeleteButton,
		HideDetailButton:         cfg.HideDetailButton,
		HideFilterButton:         cfg.HideFilterButton,
		HideRowSelector:          cfg.HideRowSelector,
		HidePagination:           cfg.HidePagination,
		HideQueryInfo:            cfg.HideQueryInfo,
		FilterFormLayout:         cfg.FilterFormLayout,
		HideContinueEditCheckBox: cfg.HideContinueEditCheckBox,
		HideContinueNewCheckBox:  cfg.HideContinueNewCheckBox,
		HideResetButton:          cfg.HideResetButton,
		HideBackButton:           cfg.HideBackButton,
		Output:                   cfg.Output,
	}
}

type Fields []Field

type Field struct {
	Head        string `json:"head"`
	Name        string `json:"name"`
	DBType      string `json:"db_type"`
	FormType    string `json:"form_type"`
	NotAllowAdd bool   `json:"not_allow_add"`
	Filterable  bool   `json:"filterable"`
	Sortable    bool   `json:"sortable"`
	Editable    bool   `json:"editable"`
	CanAdd      bool   `json:"can_add"`
}

func Generate(param Param) error {
	t, err := template.New("table_model").Parse(tableModelTmpl)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = t.Execute(buf, param)
	if err != nil {
		return err
	}
	c, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath.FromSlash(param.Output)+"/"+param.RowTable+".go", c, 0644)
}

func GenerateTables(outputPath, packageName string, tables []string, new bool) error {

	if len(outputPath) > 0 && outputPath[len(outputPath)-1] == '/' {
		outputPath = outputPath[:len(outputPath)-1]
	}

	outputPath = filepath.FromSlash(outputPath)

	if !new && !utils.FileExist(outputPath+"/tables.go") {
		return nil
	}

	tableStr := ""
	commentStr := ""
	const (
		commentStrEnd = "example end"
		tablesEnd     = "generators end"
	)

	for i := 0; i < len(tables); i++ {
		tableStr += `
	"` + tables[i] + `": Get` + strings.Title(camelcase(tables[i])) + `Table,`
		commentStr += `// "` + tables[i] + `" => http://localhost:9033/admin/info/` + tables[i] + `
`
	}
	commentStr += `//
// ` + commentStrEnd + `
`
	tableStr += `

	// ` + tablesEnd

	tablesContentByte, err := ioutil.ReadFile(outputPath + "/tables.go")
	tablesContent := string(tablesContentByte)

	content := ""

	if err == nil && tablesContent != "" && strings.Index(tablesContent, "/") != -1 {
		tablesContent = strings.Replace(tablesContent, commentStrEnd+`
//`, commentStr[3:]+"//", -1)
		content = strings.Replace(tablesContent, "// "+tablesEnd, tableStr, -1)
	} else {
		content = `package ` + packageName + `

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
` + commentStr + `//
var Generators = map[string]table.Generator{` + tableStr + `
}
`
	}

	c, err := format.Source([]byte(content))

	if err != nil {
		return err
	}
	return ioutil.WriteFile(outputPath+"/tables.go", c, 0644)
}

func camelcase(s string) string {
	arr := strings.Split(s, "_")
	var res = ""
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			res += arr[i]
		} else {
			res += strings.Title(arr[i])
		}
	}
	return res
}

func getType(typeName string) string {
	r, _ := regexp.Compile(`\(.*?\)`)
	typeName = r.ReplaceAllString(typeName, "")
	r2, _ := regexp.Compile(`unsigned(.*)`)
	return strings.TrimSpace(strings.Title(strings.ToLower(r2.ReplaceAllString(typeName, ""))))
}

func getFieldsFromConn(conn db.Connection, table, driver string) Fields {
	columnsModel, _ := db.WithDriver(conn).Table(table).ShowColumns()

	fields := make(Fields, len(columnsModel))

	fieldField := "Field"
	typeField := "Type"
	if driver == "postgresql" {
		fieldField = "column_name"
		typeField = "udt_name"
	}
	if driver == "sqlite" {
		fieldField = "name"
		typeField = "type"
	}
	if driver == "mssql" {
		fieldField = "column_name"
		typeField = "data_type"
	}

	for i, model := range columnsModel {
		typeName := getType(model[typeField].(string))
		fields[i] = Field{
			Head:     strings.Title(model[fieldField].(string)),
			Name:     model[fieldField].(string),
			DBType:   typeName,
			FormType: form.GetFormTypeFromFieldType(db.DT(strings.ToUpper(typeName)), model[fieldField].(string)),
		}
		if model[fieldField].(string) == "id" {
			fields[i].Filterable = true
		}
	}

	return fields
}
