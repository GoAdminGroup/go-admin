package tools

import (
	"bytes"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"go/format"
	"io/ioutil"
	"regexp"
	"strings"
	"text/template"
)

type Param struct {
	Connection string `json:"connection"`
	Driver     string `json:"driver"`
	Package    string `json:"package"`

	Table      string `json:"table"`
	RowTable   string `json:"row_table"`
	TableTitle string `json:"table_title"`
	TableName  string `json:"table_name"`

	HideFilterArea bool `json:"hide_filter_area"`

	Fields Fields

	Output string `json:"output"`
}

type Config struct {
	Connection     string        `json:"connection"`
	Driver         string        `json:"driver"`
	Package        string        `json:"package"`
	Table          string        `json:"table"`
	Schema         string        `json:"schema"`
	Output         string        `json:"output"`
	Conn           db.Connection `json:"conn"`
	HideFilterArea bool          `json:"hide_filter_area"`
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

	return Param{
		Connection:     cfg.Connection,
		Driver:         cfg.Driver,
		Package:        cfg.Package,
		Table:          fixedTable(ta),
		TableTitle:     strings.Title(ta),
		TableName:      dbTable,
		HideFilterArea: cfg.HideFilterArea,
		RowTable:       cfg.Table,
		Fields:         getFieldsFromConn(cfg.Conn, dbTable, cfg.Driver),
		Output:         cfg.Output,
	}
}

func NewParamWithFields(cfg Config, fields Fields) Param {
	ta := camelcase(cfg.Table)
	dbTable := cfg.Table
	if cfg.Schema != "" {
		dbTable = cfg.Schema + "." + cfg.Table
	}

	return Param{
		Connection:     cfg.Connection,
		Driver:         cfg.Driver,
		Package:        cfg.Package,
		Table:          ta,
		TableTitle:     strings.Title(cfg.Table),
		TableName:      dbTable,
		RowTable:       cfg.Table,
		HideFilterArea: cfg.HideFilterArea,
		Fields:         fields,
		Output:         cfg.Output,
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
	return ioutil.WriteFile(param.Output+"/"+param.RowTable+".go", c, 0644)
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
