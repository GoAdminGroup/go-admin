package models

import (
	"github.com/chenhg5/go-admin/modules/connections"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"strconv"
	"strings"
	"github.com/chenhg5/go-admin/template/adminlte/components"
	"html/template"
)

type ErrStruct struct {
	Class   string
	Message string
}

// 一个管理数据模块的抽象表示
type GlobalTable struct {
	Info components.InfoPanel
	Form components.FormPanel
}

// 查数据
func (tableModel GlobalTable) GetDataFromDatabase(queryParam map[string]string) ([]map[string]string, []map[string]template.HTML, *components.PaninatorAttribute, string, string) {

	pageInt, _ := strconv.Atoi(queryParam["page"])

	title := tableModel.Info.Title
	description := tableModel.Info.Description

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := connections.GetConnection().Query("show columns in " + tableModel.Info.Table)

	var sortable string
	for i := 0; i < len(tableModel.Info.FieldList); i++ {
		if tableModel.Info.FieldList[i].Field != "id" && CheckInTable(columnsModel, tableModel.Info.FieldList[i].Field) {
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
	if !CheckInTable(columnsModel, queryParam["sortField"]) {
		queryParam["sortField"] = "id"
	}

	res, _ := connections.GetConnection().Query("select " + fields + " from " + tableModel.Info.Table + " where id > 0 order by " + queryParam["sortField"] + " "+
		queryParam["sortType"]+ " LIMIT ? OFFSET ?", queryParam["pageSize"], (pageInt-1)*10)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: 加入对象池
		tempModelData := make(map[string]template.HTML, 0)

		for j := 0; j < len(tableModel.Info.FieldList); j++ {
			if CheckInTable(columnsModel, tableModel.Info.FieldList[j].Field) {
				tempModelData[tableModel.Info.FieldList[j].Head] = template.HTML(tableModel.Info.FieldList[j].ExcuFun(components.RowModel{
					ID:    res[i]["id"].(int64),
					Value: GetStringFromType(tableModel.Info.FieldList[j].TypeName, res[i][tableModel.Info.FieldList[j].Field]),
				}).(string))
			} else {
				tempModelData[tableModel.Info.FieldList[j].Head] = template.HTML(tableModel.Info.FieldList[j].ExcuFun(components.RowModel{
					ID:    res[i]["id"].(int64),
					Value: "",
				}).(string))
			}
		}

		tempModelData["id"] = template.HTML(GetStringFromType("int", res[i]["id"]))

		infoList = append(infoList, tempModelData)
	}

	total, _ := connections.GetConnection().Query("select count(*) from "+tableModel.Info.Table+" where id > ?", 0)
	size := int(total[0]["count(*)"].(int64))

	paginator := GetPaginator(queryParam["path"], pageInt, queryParam["page"], queryParam["pageSize"], queryParam["sortField"], queryParam["sortType"], size)

	return thead, infoList, paginator, title, description

}

// 查单个数据
func (tableModel GlobalTable) GetDataFromDatabaseWithId(prefix string, id string) ([]components.FormStruct, string, string) {

	fields := ""

	columnsModel, _ := connections.GetConnection().Query("show columns in " + tableModel.Form.Table)

	for i := 0; i < len(tableModel.Form.FormList); i++ {
		if CheckInTable(columnsModel, tableModel.Form.FormList[i].Field) {
			fields += tableModel.Form.FormList[i].Field + ","
		}
	}

	fields = fields[0 : len(fields)-1]

	res, _ := connections.GetConnection().Query("select "+fields+" from "+tableModel.Form.Table+" where id = ?", id)
	idint64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tableModel.Form.FormList); i++ {
		if CheckInTable(columnsModel, tableModel.Form.FormList[i].Field) {
			if tableModel.Form.FormList[i].FormType == "select" || tableModel.Form.FormList[i].FormType == "selectbox" {
				valueArr := tableModel.Form.FormList[i].ExcuFun(components.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tableModel.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tableModel.Form.FormList[i].Value = tableModel.Form.FormList[i].ExcuFun(components.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
				}).(string)
			}
		} else {
			if tableModel.Form.FormList[i].FormType == "select" || tableModel.Form.FormList[i].FormType == "selectbox" {
				valueArr := tableModel.Form.FormList[i].ExcuFun(components.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tableModel.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tableModel.Form.FormList[i].Value = tableModel.Form.FormList[i].ExcuFun(components.RowModel{
					ID:    idint64,
					Value: tableModel.Form.FormList[i].Field,
				}).(string)
			}
		}
	}

	return tableModel.Form.FormList, tableModel.Form.Title, tableModel.Form.Description
}

// 改数据
func (tableModel GlobalTable) UpdateDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	valueList := make([]interface{}, 0)
	columnsModel, _ := connections.GetConnection().Query("show columns in " + tableModel.Form.Table)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columnsModel, k) {
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

	connections.GetConnection().Exec("update "+tableModel.Form.Table+" set "+fields+" where id = ?", valueList...)
}

// 增数据
func (tableModel GlobalTable) InsertDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	queStr := ""
	var valueList []interface{}
	columnsModel, _ := connections.GetConnection().Query("show columns in " + tableModel.Form.Table)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columnsModel, k) {
			fields += k + ","
			queStr += "?,"
			valueList = append(valueList, v[0])
		}
	}

	fields = fields[0 : len(fields)-1]
	queStr = queStr[0 : len(queStr)-1]

	// TODO: 过滤
	connections.GetConnection().Exec("insert into "+tableModel.Form.Table+"("+fields+") values ("+queStr+")", valueList...)
}

// 删数据
func (tableModel GlobalTable) DeleteDataFromDatabase(prefix string, id string) {
	idArr := strings.Split(id, ",")
	for _, id := range idArr {
		connections.GetConnection().Exec("delete from "+tableModel.Form.Table+" where id = ?", id)
	}
}

func GetNewFormList(old []components.FormStruct) []components.FormStruct {
	var newForm []components.FormStruct
	for _, v := range old {
		if v.Field != "id" && v.Field != "created_at" && v.Field != "updated_at" {
			newForm = append(newForm, v)
		}
	}
	return newForm
}

// 检查字段是否在数据表中
func CheckInTable(colums []map[string]interface{}, find string) bool {
	for i := 0; i < len(colums); i++ {
		if colums[i]["Field"].(string) == find {
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
