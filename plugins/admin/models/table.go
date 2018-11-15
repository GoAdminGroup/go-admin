package models

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
	"strconv"
	"strings"
)

type TableGenerator func() Table

var (
	Generators = map[string]TableGenerator{}
	TableList  = map[string]Table{}
)

func InitTableList() {
	TableList = make(map[string]Table, len(Generators))
	for prefix, generator := range Generators {
		TableList[prefix] = generator()
	}
}

// RefreshTableList refresh the table list when the table
// relationship changed.
func RefreshTableList() {
	for k, v := range Generators {
		TableList[k] = v()
	}
}

func SetGenerators(generators map[string]TableGenerator) {
	for key, generator := range generators {
		Generators[key] = generator
	}
}

type Table struct {
	Info             types.InfoPanel
	Form             types.FormPanel
	ConnectionDriver string
}

type Columns []string

func GetColumns(columnsModel []map[string]interface{}, driver string) Columns {
	columns := make(Columns, len(columnsModel))
	switch driver {
	case "mysql":
		for key, model := range columnsModel {
			columns[key] = model["Field"].(string)
		}
		return columns
	case "sqlite":
		for key, model := range columnsModel {
			columns[key] = string((*(model["name"].(*interface{}))).([]uint8))
		}
		return columns
	default:
		panic("wrong driver")
	}
}

// GetDataFromDatabase query the data set.
// TODO: set return param as a struct
func (tableModel Table) GetDataFromDatabase(queryParam map[string]string) ([]map[string]string, []map[string]template.HTML, types.PaginatorAttribute, string, string) {

	pageInt, _ := strconv.Atoi(queryParam["page"])

	title := tableModel.Info.Title
	description := tableModel.Info.Description

	thead := make([]map[string]string, 0)
	fields := ""

	showColumns := "show columns in " + tableModel.Info.Table
	if tableModel.ConnectionDriver == "sqlite" {
		showColumns = "PRAGMA table_info(" + tableModel.Info.Table + ");"
	}

	columnsModel, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query(showColumns)
	columns := GetColumns(columnsModel, tableModel.ConnectionDriver)

	var sortable string
	for i := 0; i < len(tableModel.Info.FieldList); i++ {
		if tableModel.Info.FieldList[i].Field != "id" && CheckInTable(columns, tableModel.Info.FieldList[i].Field) {
			fields += tableModel.Info.FieldList[i].Field + ","
		}
		sortable = "0"
		if tableModel.Info.FieldList[i].Sortable {
			sortable = "1"
		}
		thead = append(thead, map[string]string{
			"head":     tableModel.Info.FieldList[i].Head,
			"sortable": sortable,
			"field":    tableModel.Info.FieldList[i].Field,
		})
	}

	fields += "id"

	if queryParam["sortType"] != "desc" && queryParam["sortType"] != "asc" {
		queryParam["sortType"] = "desc"
	}
	if !CheckInTable(columns, queryParam["sortField"]) {
		queryParam["sortField"] = "id"
	}

	res, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query("select "+fields+" from "+tableModel.Info.Table+" where id > 0 order by "+queryParam["sortField"]+" "+
		queryParam["sortType"]+" LIMIT ? OFFSET ?", queryParam["pageSize"], (pageInt-1)*10)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: 加入对象池
		tempModelData := make(map[string]template.HTML, 0)

		for j := 0; j < len(tableModel.Info.FieldList); j++ {
			if CheckInTable(columns, tableModel.Info.FieldList[j].Field) {
				tempModelData[tableModel.Info.FieldList[j].Head] = template.HTML(tableModel.Info.FieldList[j].ExcuFun(types.RowModel{
					ID:    res[i]["id"].(int64),
					Value: GetStringFromType(tableModel.Info.FieldList[j].TypeName, res[i][tableModel.Info.FieldList[j].Field]),
				}).(string))
			} else {
				tempModelData[tableModel.Info.FieldList[j].Head] = template.HTML(tableModel.Info.FieldList[j].ExcuFun(types.RowModel{
					ID:    res[i]["id"].(int64),
					Value: "",
				}).(string))
			}
		}

		tempModelData["id"] = template.HTML(GetStringFromType("int", res[i]["id"]))

		infoList = append(infoList, tempModelData)
	}

	total, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query("select count(*) from "+tableModel.Info.Table+" where id > ?", 0)
	var size int
	if tableModel.ConnectionDriver == "sqlite" {
		size = int((*(total[0]["count(*)"].(*interface{}))).(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	paginator := GetPaginator(queryParam["path"], pageInt, queryParam["page"], queryParam["pageSize"], queryParam["sortField"], queryParam["sortType"], size)

	return thead, infoList, paginator, title, description

}

// GetDataFromDatabaseWithId query the single row of data.
func (tableModel Table) GetDataFromDatabaseWithId(prefix string, id string) ([]types.FormStruct, string, string) {

	fields := ""

	columnsModel, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query("show columns in " + tableModel.Form.Table)
	columns := GetColumns(columnsModel, tableModel.ConnectionDriver)

	for i := 0; i < len(tableModel.Form.FormList); i++ {
		if CheckInTable(columns, tableModel.Form.FormList[i].Field) {
			fields += tableModel.Form.FormList[i].Field + ","
		}
	}

	fields = fields[0 : len(fields)-1]

	res, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query("select "+fields+" from "+tableModel.Form.Table+" where id = ?", id)
	idint64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tableModel.Form.FormList); i++ {
		if CheckInTable(columns, tableModel.Form.FormList[i].Field) {
			if tableModel.Form.FormList[i].FormType == "select" || tableModel.Form.FormList[i].FormType == "selectbox" || tableModel.Form.FormList[i].FormType == "select_single" {
				valueArr := tableModel.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tableModel.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tableModel.Form.FormList[i].Value = tableModel.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
				}).(string)
			}
		} else {
			if tableModel.Form.FormList[i].FormType == "select" || tableModel.Form.FormList[i].FormType == "selectbox" {
				valueArr := tableModel.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tableModel.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tableModel.Form.FormList[i].Value = tableModel.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: tableModel.Form.FormList[i].Field,
				}).(string)
			}
		}
	}

	return tableModel.Form.FormList, tableModel.Form.Title, tableModel.Form.Description
}

// UpdateDataFromDatabase update data.
func (tableModel Table) UpdateDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	valueList := make([]interface{}, 0)
	columnsModel, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query("show columns in " + tableModel.Form.Table)
	columns := GetColumns(columnsModel, tableModel.ConnectionDriver)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columns, k) {
			fields += strings.Replace(k, "[]", "", -1) + " = ?,"
			if len(v) > 0 {
				valueList = append(valueList, strings.Join(modules.RemoveBlackFromArray(v), ","))
			} else {
				valueList = append(valueList, v[0])
			}
		}
	}

	fields = fields[0 : len(fields)-1]
	valueList = append(valueList, dataList["id"][0])

	db.GetConnectionByDriver(tableModel.ConnectionDriver).Exec("update "+tableModel.Form.Table+" set "+fields+" where id = ?", valueList...)
}

// InsertDataFromDatabase insert data.
func (tableModel Table) InsertDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	queStr := ""
	var valueList []interface{}
	columnsModel, _ := db.GetConnectionByDriver(tableModel.ConnectionDriver).Query("show columns in " + tableModel.Form.Table)
	columns := GetColumns(columnsModel, tableModel.ConnectionDriver)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columns, k) {
			fields += k + ","
			queStr += "?,"
			valueList = append(valueList, v[0])
		}
	}

	fields = fields[0 : len(fields)-1]
	queStr = queStr[0 : len(queStr)-1]

	// TODO: 过滤
	db.GetConnectionByDriver(tableModel.ConnectionDriver).Exec("insert into "+tableModel.Form.Table+"("+fields+") values ("+queStr+")", valueList...)
}

// DeleteDataFromDatabase delete data.
func (tableModel Table) DeleteDataFromDatabase(prefix string, id string) {
	idArr := strings.Split(id, ",")
	for _, id := range idArr {
		db.GetConnectionByDriver(tableModel.ConnectionDriver).Exec("delete from "+tableModel.Form.Table+" where id = ?", id)
	}
}

func GetNewFormList(old []types.FormStruct) []types.FormStruct {
	var newForm []types.FormStruct
	for _, v := range old {
		v.Value = ""
		if v.Field != "id" && v.Field != "created_at" && v.Field != "updated_at" {
			newForm = append(newForm, v)
		}
	}
	fmt.Println("form", newForm)
	return newForm
}

// CheckInTable checks the find string is in the columns or not.
func CheckInTable(columns []string, find string) bool {
	for i := 0; i < len(columns); i++ {
		if columns[i] == find {
			return true
		}
	}
	return false
}

func GetStringFromType(typeName string, value interface{}) string {
	typeName = strings.ToUpper(typeName)
	if value == nil {
		return ""
	}
	switch typeName {
	case "INT":
		return strconv.FormatInt(value.(int64), 10)
	case "TINYINT":
		return strconv.FormatInt(value.(int64), 10)
	case "MEDIUMINT":
		return strconv.FormatInt(value.(int64), 10)
	case "SMALLINT":
		return strconv.FormatInt(value.(int64), 10)
	case "BIGINT":
		return strconv.FormatInt(value.(int64), 10)
	case "FLOAT":
		return strconv.FormatFloat(value.(float64), 'g', 5, 32)
	case "DOUBLE":
		return strconv.FormatFloat(value.(float64), 'g', 5, 32)
	case "DECIMAL":
		return string(value.(uint8))
	case "DATE":
		return value.(string)
	case "TIME":
		return value.(string)
	case "YEAR":
		return value.(string)
	case "DATETIME":
		return value.(string)
	case "TIMESTAMP":
		return value.(string)
	case "VARCHAR":
		return value.(string)
	case "MEDIUMTEXT":
		return value.(string)
	case "LONGTEXT":
		return value.(string)
	case "TINYTEXT":
		return value.(string)
	case "TEXT":
		return value.(string)
	default:
		return ""
	}
}
