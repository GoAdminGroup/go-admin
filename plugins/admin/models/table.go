package models

import (
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

type PanelInfo struct {
	Thead       []map[string]string
	InfoList    []map[string]template.HTML
	Paginator   types.PaginatorAttribute
	Title       string
	Description string
}

func (tb Table) GetFiltersMap() []map[string]string {
	var filters = make([]map[string]string, 0)
	for _, value := range tb.Info.FieldList {
		if value.Filter {
			filters = append(filters, map[string]string{
				"title": value.Head,
				"name": value.Field,
			})
		}
	}
	if len(filters) == 0 {
		filters = append(filters, map[string]string{
			"title": "ID",
			"name": "id",
		})
	}
	return filters
}

// GetDataFromDatabase query the data set.
func (tb Table) GetDataFromDatabase(path string, params *Parameters) PanelInfo {

	pageInt, _ := strconv.Atoi(params.Page)

	title := tb.Info.Title
	description := tb.Info.Description

	thead := make([]map[string]string, 0)
	fields := ""

	showColumns := "show columns in " + tb.Info.Table
	if tb.ConnectionDriver == "sqlite" {
		showColumns = "PRAGMA table_info(" + tb.Info.Table + ");"
	}

	columnsModel, _ := tb.db().Query(showColumns)
	columns := GetColumns(columnsModel, tb.ConnectionDriver)

	var sortable string
	for i := 0; i < len(tb.Info.FieldList); i++ {
		if tb.Info.FieldList[i].Field != "id" && CheckInTable(columns, tb.Info.FieldList[i].Field) {
			fields += tb.Info.FieldList[i].Field + ","
		}
		sortable = "0"
		if tb.Info.FieldList[i].Sortable {
			sortable = "1"
		}
		thead = append(thead, map[string]string{
			"head":     tb.Info.FieldList[i].Head,
			"sortable": sortable,
			"field":    tb.Info.FieldList[i].Field,
		})
	}

	fields += "id"

	if !CheckInTable(columns, params.SortField) {
		params.SortField = "id"
	}

	wheres := " where "
	whereArgs := make([]interface{}, 0)
	if len(params.Fields) == 0 {
		wheres += "id > 0"
	} else {
		for key, value := range params.Fields {
			wheres += key + " = ? and"
			whereArgs = append(whereArgs, value)
		}
		wheres = wheres[:len(wheres)-4]
	}
	args := append(whereArgs, params.PageSize, (pageInt-1)*10)

	// TODO: add left join table relations

	res, _ := tb.db().Query("select " + fields + " from " + tb.Info.Table + wheres + " order by " + params.SortField + " "+
		params.SortType+ " LIMIT ? OFFSET ?", args...)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: 加入对象池
		tempModelData := make(map[string]template.HTML, 0)

		for j := 0; j < len(tb.Info.FieldList); j++ {
			if CheckInTable(columns, tb.Info.FieldList[j].Field) {
				tempModelData[tb.Info.FieldList[j].Head] = template.HTML(tb.Info.FieldList[j].ExcuFun(types.RowModel{
					ID:    res[i]["id"].(int64),
					Value: GetStringFromType(tb.Info.FieldList[j].TypeName, res[i][tb.Info.FieldList[j].Field]),
				}).(string))
			} else {
				tempModelData[tb.Info.FieldList[j].Head] = template.HTML(tb.Info.FieldList[j].ExcuFun(types.RowModel{
					ID:    res[i]["id"].(int64),
					Value: "",
				}).(string))
			}
		}

		tempModelData["id"] = template.HTML(GetStringFromType("int", res[i]["id"]))

		infoList = append(infoList, tempModelData)
	}

	total, _ := tb.db().Query("select count(*) from "+tb.Info.Table+wheres, whereArgs...)
	var size int
	if tb.ConnectionDriver == "sqlite" {
		size = int((*(total[0]["count(*)"].(*interface{}))).(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	paginator := GetPaginator(path, params, size)

	return PanelInfo{
		thead, infoList, paginator, title, description,
	}

}

// GetDataFromDatabaseWithId query the single row of data.
func (tb Table) GetDataFromDatabaseWithId(id string) ([]types.FormStruct, string, string) {

	fields := ""

	columnsModel, _ := tb.db().Query("show columns in " + tb.Form.Table)
	columns := GetColumns(columnsModel, tb.ConnectionDriver)

	for i := 0; i < len(tb.Form.FormList); i++ {
		if CheckInTable(columns, tb.Form.FormList[i].Field) {
			fields += tb.Form.FormList[i].Field + ","
		}
	}

	fields = fields[0 : len(fields)-1]

	res, _ := tb.db().Query("select "+fields+" from "+tb.Form.Table+" where id = ?", id)
	idint64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tb.Form.FormList); i++ {
		if CheckInTable(columns, tb.Form.FormList[i].Field) {
			if tb.Form.FormList[i].FormType == "select" || tb.Form.FormList[i].FormType == "selectbox" || tb.Form.FormList[i].FormType == "select_single" {
				valueArr := tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.Form.FormList[i].TypeName, res[0][tb.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tb.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.Form.FormList[i].Value = tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.Form.FormList[i].TypeName, res[0][tb.Form.FormList[i].Field]),
				}).(string)
			}
		} else {
			if tb.Form.FormList[i].FormType == "select" || tb.Form.FormList[i].FormType == "selectbox" {
				valueArr := tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.Form.FormList[i].TypeName, res[0][tb.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tb.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.Form.FormList[i].Value = tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: tb.Form.FormList[i].Field,
				}).(string)
			}
		}
	}

	return tb.Form.FormList, tb.Form.Title, tb.Form.Description
}

// UpdateDataFromDatabase update data.
func (tb Table) UpdateDataFromDatabase(dataList map[string][]string) {

	fields := ""
	valueList := make([]interface{}, 0)
	columnsModel, _ := tb.db().Query("show columns in " + tb.Form.Table)
	columns := GetColumns(columnsModel, tb.ConnectionDriver)
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

	tb.db().Exec("update "+tb.Form.Table+" set "+fields+" where id = ?", valueList...)
}

// InsertDataFromDatabase insert data.
func (tb Table) InsertDataFromDatabase(dataList map[string][]string) {

	fields := ""
	queStr := ""
	var valueList []interface{}
	columnsModel, _ := tb.db().Query("show columns in " + tb.Form.Table)
	columns := GetColumns(columnsModel, tb.ConnectionDriver)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columns, k) {
			fields += k + ","
			queStr += "?,"
			valueList = append(valueList, v[0])
		}
	}

	fields = fields[0 : len(fields)-1]
	queStr = queStr[0 : len(queStr)-1]

	tb.db().Exec("insert into "+tb.Form.Table+"("+fields+") values ("+queStr+")", valueList...)
}

// DeleteDataFromDatabase delete data.
func (tb Table) DeleteDataFromDatabase(id string) {
	idArr := strings.Split(id, ",")
	for _, id := range idArr {
		tb.db().Exec("delete from "+tb.Form.Table+" where id = ?", id)
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
	return newForm
}

func (tb Table) db() db.Connection {
	return db.GetConnectionByDriver(tb.ConnectionDriver)
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
