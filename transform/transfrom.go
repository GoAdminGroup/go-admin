package transform

import (
	"goAdmin/config"
	"goAdmin/connections/mysql"
	"goAdmin/models"
	"strconv"
	"strings"
)

func TransfromData(page string, pageSize string, path string, prefix string) ([]string, []map[string]string, map[string]interface{}, string, string) {

	pageInt, _ := strconv.Atoi(page)

	title := config.GlobalTableList[prefix].Info.Title
	description := config.GlobalTableList[prefix].Info.Description

	thead, infoList, size := GetDataFromDatabase(prefix, pageSize, pageInt)

	paginator := GetPaginator(path, pageInt, page, pageSize, size, prefix)

	return thead, infoList, paginator, title, description
}

func TransfromFormData(id string, prefix string) ([]models.FormStruct, string, string) {

	title := config.GlobalTableList[prefix].Form.Title
	description := config.GlobalTableList[prefix].Form.Description

	formData := GetDataFromDatabaseWithId(prefix, id)

	return formData, title, description
}

func GetDataFromDatabaseWithId(prefix string, id string) []models.FormStruct {

	fields := ""

	for i := 0; i < len(config.GlobalTableList[prefix].Form.FormList); i++ {
		fields += config.GlobalTableList[prefix].Form.FormList[i].Field + ","
	}

	fields = SliceStr(fields)

	res, _ := mysql.Query("select "+fields+" from "+config.GlobalTableList[prefix].Form.Table+" where id = ?", id)

	for i := 0; i < len(config.GlobalTableList[prefix].Form.FormList); i++ {
		config.GlobalTableList[prefix].Form.FormList[i].Value = GetStringFromType(config.GlobalTableList[prefix].Form.FormList[i].TypeName, res[0][config.GlobalTableList[prefix].Form.FormList[i].Field])
	}

	return config.GlobalTableList[prefix].Form.FormList
}

func GetDataFromDatabase(prefix string, pageSize string, pageInt int) ([]string, []map[string]string, int) {

	thead := make([]string, 0)
	fields := ""

	for i := 0; i < len(config.GlobalTableList[prefix].Info.FieldList); i++ {
		if config.GlobalTableList[prefix].Info.FieldList[i].Field != "id" {
			fields += config.GlobalTableList[prefix].Info.FieldList[i].Field + ","
		}
		thead = append(thead, config.GlobalTableList[prefix].Info.FieldList[i].Head)
	}

	fields += "id"

	res, _ := mysql.Query("select "+fields+" from "+config.GlobalTableList[prefix].Info.Table+" where id > 0 order by created_at desc LIMIT ? OFFSET ?",
		pageSize, (pageInt-1)*10)

	infoList := make([]map[string]string, 0)

	for i := 0; i < len(res); i++ {

		// TODO: 加入对象池
		tempModelData := make(map[string]string, 0)

		for j := 0; j < len(config.GlobalTableList[prefix].Info.FieldList); j++ {
			tempModelData[config.GlobalTableList[prefix].Info.FieldList[j].Head] = config.GlobalTableList[prefix].Info.FieldList[j].ExcuFun(GetStringFromType(config.GlobalTableList[prefix].Info.FieldList[j].TypeName, res[i][config.GlobalTableList[prefix].Info.FieldList[j].Field]))
		}

		tempModelData["id"] = GetStringFromType("int", res[i]["id"])

		infoList = append(infoList, tempModelData)
	}

	total, _ := mysql.Query("select count(*) from "+config.GlobalTableList[prefix].Info.Table+" where id > ?", 0)
	size := int(total[0]["count(*)"].(int64))

	return thead, infoList, size
}

func UpdateDataFromDatabase(prefix string, dataList map[string][]string) {

	fields := ""
	valueList := make([]interface{}, 0)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_token" {
			fields += k + " = ?,"
			valueList = append(valueList, v[0])
		}
	}

	fields = SliceStr(fields)
	valueList = append(valueList, dataList["id"][0])

	mysql.Exec("update "+config.GlobalTableList[prefix].Form.Table+" set "+fields+" where id = ?", valueList...)
}

func InsertDataFromDatabase(prefix string, dataList map[string][]string) {

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

	fields = SliceStr(fields)
	queStr = SliceStr(queStr)
	valueList = append(valueList, dataList["id"][0])

	mysql.Exec("insert into "+config.GlobalTableList[prefix].Form.Table+"("+fields+") values ("+queStr+")", valueList...)
}

func DeleteDataFromDatabase(prefix string, id string) {
	mysql.Exec("delete from "+config.GlobalTableList[prefix].Form.Table+" where id = ?", id)
}

func SliceStr(s string) string {
	rs := []rune(s)
	length := len(rs)
	return string(rs[0 : length-1])
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
