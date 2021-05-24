package tools

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"

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

	TablePageTitle    string `json:"table_title"`
	TableDescription  string `json:"table_description"`
	FormTitle         string `json:"form_title"`
	FormDescription   string `json:"form_description"`
	DetailTitle       string `json:"detail_title"`
	DetailDescription string `json:"detail_description"`

	ExtraImport string `json:"extra_import"`
	ExtraCode   string `json:"extra_code"`

	Fields       Fields `json:"fields"`
	FormFields   Fields `json:"form_fields"`
	DetailFields Fields `json:"detail_fields"`

	DetailDisplay uint8 `json:"detail_display"`

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
	ExtraImport      string        `json:"extra_import"`

	TableTitle        string `json:"table_title"`
	TableDescription  string `json:"table_description"`
	FormTitle         string `json:"form_title"`
	FormDescription   string `json:"form_description"`
	DetailTitle       string `json:"detail_title"`
	DetailDescription string `json:"detail_description"`

	HideContinueEditCheckBox bool `json:"hide_continue_edit_check_box"`
	HideContinueNewCheckBox  bool `json:"hide_continue_new_check_box"`
	HideResetButton          bool `json:"hide_reset_button"`
	HideBackButton           bool `json:"hide_back_button"`

	DetailDisplay uint8 `json:"detail_display"`

	ExtraCode string `json:"extra_code"`
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

func NewParam(cfg Config) *Param {
	ta := camelcase(cfg.Table)
	dbTable := cfg.Table
	if cfg.Schema != "" {
		dbTable = cfg.Schema + "." + cfg.Table
	}

	fields := getFieldsFromConn(cfg.Conn, dbTable, cfg.Driver)
	tt := strings.Title(ta)

	return &Param{
		Connection:               cfg.Connection,
		Driver:                   cfg.Driver,
		Package:                  cfg.Package,
		Table:                    fixedTable(ta),
		TableTitle:               tt,
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
		DetailDisplay:            cfg.DetailDisplay,
		Output:                   cfg.Output,
		ExtraImport:              cfg.ExtraImport,
		ExtraCode:                cfg.ExtraCode,
		TablePageTitle:           utils.SetDefault(cfg.TableTitle, "", tt),
		TableDescription:         utils.SetDefault(cfg.TableDescription, "", tt),
		FormTitle:                utils.SetDefault(cfg.FormTitle, "", tt),
		FormDescription:          utils.SetDefault(cfg.FormDescription, "", tt),
	}
}

func NewParamWithFields(cfg Config, fields ...Fields) *Param {
	ta := camelcase(cfg.Table)
	dbTable := cfg.Table
	if cfg.Schema != "" {
		dbTable = cfg.Schema + "." + cfg.Table
	}

	if len(cfg.Output) > 0 && cfg.Output[len(cfg.Output)-1] == '/' {
		cfg.Output = cfg.Output[:len(cfg.Output)-1]
	}

	tt := strings.Title(ta)

	detailFields := make(Fields, 0)
	if len(fields) > 2 {
		detailFields = fields[2]
	}

	return &Param{
		Connection:               cfg.Connection,
		Driver:                   cfg.Driver,
		Package:                  cfg.Package,
		Table:                    ta,
		TableTitle:               tt,
		TableName:                dbTable,
		RowTable:                 cfg.Table,
		Fields:                   fields[0],
		FormFields:               fields[1],
		DetailFields:             detailFields,
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
		DetailDisplay:            cfg.DetailDisplay,
		Output:                   cfg.Output,
		ExtraImport:              cfg.ExtraImport,
		ExtraCode:                cfg.ExtraCode,
		TablePageTitle:           utils.SetDefault(cfg.TableTitle, "", tt),
		TableDescription:         utils.SetDefault(cfg.TableDescription, "", tt),
		FormTitle:                utils.SetDefault(cfg.FormTitle, "", tt),
		FormDescription:          utils.SetDefault(cfg.FormDescription, "", tt),
		DetailTitle:              utils.SetDefault(cfg.DetailTitle, "", tt),
		DetailDescription:        utils.SetDefault(cfg.DetailDescription, "", tt),
	}
}

type Fields []Field

type Field struct {
	Head         string `json:"head"`
	Name         string `json:"name"`
	DBType       string `json:"db_type"`
	FormType     string `json:"form_type"`
	Filterable   bool   `json:"filterable"`
	Sortable     bool   `json:"sortable"`
	InfoEditable bool   `json:"info_editable"`
	Editable     bool   `json:"editable"`
	Hide         bool   `json:"hide"`
	FormHide     bool   `json:"form_hide"`
	EditHide     bool   `json:"edit_hide"`
	CreateHide   bool   `json:"create_hide"`
	Default      string `json:"default"`
	CanAdd       bool   `json:"can_add"`
	ExtraFun     string `json:"extra_fun"`
}

func Generate(param *Param) error {
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

const (
	commentStrEnd = "example end"
	tablesEnd     = "generators end"
)

func GenerateTables(outputPath, packageName string, tables []string, isNew bool) error {

	if len(outputPath) > 0 && outputPath[len(outputPath)-1] == '/' {
		outputPath = outputPath[:len(outputPath)-1]
	}

	outputPath = filepath.FromSlash(outputPath)
	fileExist := utils.FileExist(outputPath + "/tables.go")

	if !isNew && !fileExist {
		return nil
	}

	var (
		tableStr          = ""
		commentStr        = ""
		tablesContentByte []byte
		tablesContent     string
		err               error
	)
	if fileExist {
		tablesContentByte, err = ioutil.ReadFile(outputPath + "/tables.go")
		if err != nil {
			return err
		}
		tablesContent = string(tablesContentByte)
	}

	for i := 0; i < len(tables); i++ {
		lowerTable := strings.ToLower(tables[i])
		if !strings.Contains(tablesContent, `"`+lowerTable+`"`) {
			tableStr += fmt.Sprintf(`
	"%s": Get%sTable, `, lowerTable, strings.Title(camelcase(tables[i])))

			if commentStr != "" {
				commentStr += `
`
			}
			commentStr += fmt.Sprintf(`// "%s" => http://localhost:9033/admin/info/%s`, lowerTable, lowerTable)
		}
	}

	commentStr += `
//
// ` + commentStrEnd
	tableStr += `

	// ` + tablesEnd

	content := ""

	if tablesContent != "" && strings.Contains(tablesContent, "/") {
		replacer := strings.NewReplacer(`// `+commentStrEnd, commentStr, `// `+tablesEnd, tableStr)
		tablesContent = replacer.Replace(tablesContent)
		keep := `// example:
//`
		keep2 := `,

	// ` + tablesEnd
		replacer2 := strings.NewReplacer(keep, keep, keep2, keep2,
			`//
//
`, "//", `//

//`, "//", `//
// "`, `// "`,

			`,

`, ",")
		content = replacer2.Replace(tablesContent)
	} else {
		content = fmt.Sprintf(`// This file is generated by GoAdmin CLI adm.
package %s

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
%s
var Generators = map[string]table.Generator{
%s
}
`, packageName, commentStr, tableStr)
	}

	c, err := format.Source([]byte(content))

	if err != nil {
		return err
	}
	return ioutil.WriteFile(outputPath+"/tables.go", c, 0644)
}

func InsertPermissionOfTable(conn db.Connection, table string) {
	table = strings.ToLower(table)
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("query", "generator"), table+"_query", "GET", "/info/"+table)
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("show edit form page", "generator"), table+"_show_edit", "GET",
		"/info/"+table+"/edit")
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("show create form page", "generator"), table+"_show_create", "GET",
		"/info/"+table+"/new")
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("edit", "generator"), table+"_edit", "POST",
		"/edit/"+table)
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("create", "generator"), table+"_create", "POST",
		"/new/"+table)
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("delete", "generator"), table+"_delete", "POST",
		"/delete/"+table)
	InsertPermissionInfoDB(conn, table+" "+language.GetWithScope("export", "generator"), table+"_export", "POST",
		"/export/"+table)
}

func InsertPermissionInfoDB(conn db.Connection, name, slug, httpMethod, httpPath string) {
	checkExist, err := db.WithDriver(conn).Table("goadmin_permissions").
		Where("slug", "=", slug).
		First()

	if db.CheckError(err, db.QUERY) {
		panic(err)
	}

	if checkExist != nil {
		return
	}

	_, err = db.WithDriver(conn).Table("goadmin_permissions").
		Insert(dialect.H{
			"name":        name,
			"slug":        slug,
			"http_method": httpMethod,
			"http_path":   httpPath,
		})

	if db.CheckError(err, db.INSERT) {
		panic(err)
	}
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
			CanAdd:   true,
			Editable: true,
			FormType: form.GetFormTypeFromFieldType(db.DT(strings.ToUpper(typeName)), model[fieldField].(string)),
		}
		if model[fieldField].(string) == "id" {
			fields[i].Filterable = true
		}
	}

	return fields
}
