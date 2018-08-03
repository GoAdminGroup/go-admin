package models

import (
	"goAdmin/connections/mysql"
	"strconv"
	"strings"
	"goAdmin/modules"
)

// 表单列
type FormStruct struct {
	Field    string
	TypeName string
	Head     string
	Default  string
	Editable bool
	FormType string
	Value    string
	Options  []map[string]string
	ExcuFun  FieldValueFun
}

type RowModel struct {
	ID    int64
	Value string
}

// 数据过滤函数
type FieldValueFun func(value RowModel) string

// 展示列
type FieldStruct struct {
	ExcuFun  FieldValueFun
	Field    string
	TypeName string
	Head     string
}

func (field *FieldStruct) SetHead(head string) *FieldStruct {
	(*field).Head = head
	return field
}

func (field *FieldStruct) SetTypeName(typeName string) *FieldStruct {
	(*field).TypeName = typeName
	return field
}

func (field *FieldStruct) SetField(fieldName string) *FieldStruct {
	(*field).Field = fieldName
	return field
}

// 展示面板
type InfoPanel struct {
	FieldList   []FieldStruct
	Table       string
	Title       string
	Description string
}

// 表单面板
type FormPanel struct {
	FormList    []FormStruct
	Table       string
	Title       string
	Description string
}

// 一个管理数据模块的抽象表示
type GlobalTable struct {
	Info InfoPanel
	Form FormPanel
}

// 查数据
func (tableModel GlobalTable) GetDataFromDatabase(queryParam map[string]string) ([]string, []map[string]string, map[string]interface{}, string, string) {

	pageInt, _ := strconv.Atoi(queryParam["page"])

	title := tableModel.Info.Title
	description := tableModel.Info.Description

	thead := make([]string, 0)
	fields := ""

	columnsModel, _ := mysql.Query("show columns in " + tableModel.Info.Table)

	for i := 0; i < len(tableModel.Info.FieldList); i++ {
		if tableModel.Info.FieldList[i].Field != "id" && CheckInTable(columnsModel, tableModel.Info.FieldList[i].Field) {
			fields += tableModel.Info.FieldList[i].Field + ","
		}
		thead = append(thead, tableModel.Info.FieldList[i].Head)
	}

	fields += "id"

	res, _ := mysql.Query("select "+fields+" from "+tableModel.Info.Table+" where id > 0 order by created_at desc LIMIT ? OFFSET ?",
		queryParam["pageSize"], (pageInt-1)*10)

	infoList := make([]map[string]string, 0)

	for i := 0; i < len(res); i++ {

		// TODO: 加入对象池
		tempModelData := make(map[string]string, 0)

		for j := 0; j < len(tableModel.Info.FieldList); j++ {

			if CheckInTable(columnsModel, tableModel.Info.FieldList[j].Field) {
				tempModelData[tableModel.Info.FieldList[j].Head] = tableModel.Info.FieldList[j].ExcuFun(RowModel{
					res[i]["id"].(int64),
					GetStringFromType(tableModel.Info.FieldList[j].TypeName, res[i][tableModel.Info.FieldList[j].Field]),
				})
			} else {
				tempModelData[tableModel.Info.FieldList[j].Head] = tableModel.Info.FieldList[j].ExcuFun(RowModel{
					res[i]["id"].(int64),
					"",
				})
			}
		}

		tempModelData["id"] = GetStringFromType("int", res[i]["id"])

		infoList = append(infoList, tempModelData)
	}

	total, _ := mysql.Query("select count(*) from "+tableModel.Info.Table+" where id > ?", 0)
	size := int(total[0]["count(*)"].(int64))

	paginator := GetPaginator(queryParam["path"], pageInt, queryParam["page"], queryParam["pageSize"], size, queryParam["prefix"])

	return thead, infoList, paginator, title, description

}

// 查单个数据
func (tableModel GlobalTable) GetDataFromDatabaseWithId(prefix string, id string) ([]FormStruct, string, string) {

	fields := ""

	for i := 0; i < len(tableModel.Form.FormList); i++ {
		fields += tableModel.Form.FormList[i].Field + ","
	}

	fields = fields[0 : len(fields)-1]

	res, _ := mysql.Query("select "+fields+" from "+tableModel.Form.Table+" where id = ?", id)
	Idint64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tableModel.Form.FormList); i++ {
		tableModel.Form.FormList[i].Value = tableModel.Form.FormList[i].ExcuFun(RowModel{
			Idint64,
			GetStringFromType(tableModel.Form.FormList[i].TypeName, res[0][tableModel.Form.FormList[i].Field]),
		})
		if tableModel.Form.FormList[i].FormType == "select" {
			valueArr := strings.Split(tableModel.Form.FormList[i].Value, ",")
			for _, v := range tableModel.Form.FormList[i].Options {
				if modules.InArray(valueArr, v["value"]) {
					v["selected"] = "selected"
				}
			}
		}
	}

	return tableModel.Form.FormList, tableModel.Form.Title, tableModel.Form.Description
}

// 改数据
func (tableModel GlobalTable) UpdateDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	valueList := make([]interface{}, 0)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_token" {
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

	mysql.Exec("update "+tableModel.Form.Table+" set "+fields+" where id = ?", valueList...)
}

// 增数据
func (tableModel GlobalTable) InsertDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	queStr := ""
	valueList := make([]interface{}, 0)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_token" {
			fields += k + " = ?,"
			queStr += "?,"
			valueList = append(valueList, v[0])
		}
	}

	fields = fields[0 : len(fields)-1]
	queStr = queStr[0 : len(queStr)-1]
	valueList = append(valueList, dataList["id"][0])

	mysql.Exec("insert into "+tableModel.Form.Table+"("+fields+") values ("+queStr+")", valueList...)
}

// 删数据
func (tableModel GlobalTable) DeleteDataFromDatabase(prefix string, id string) {
	mysql.Exec("delete from "+tableModel.Form.Table+" where id = ?", id)
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
