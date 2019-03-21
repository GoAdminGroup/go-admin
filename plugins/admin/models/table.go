package models

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
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
				"name":  value.Field,
			})
		}
	}
	if len(filters) == 0 {
		filters = append(filters, map[string]string{
			"title": "ID",
			"name":  "id",
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

	columnsModel, _ := db.WithDriver(tb.ConnectionDriver).Table(tb.Info.Table).ShowColumns()

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
			wheres += key + " = ? and "
			whereArgs = append(whereArgs, value)
		}
		wheres = wheres[:len(wheres)-4]
	}
	args := append(whereArgs, params.PageSize, (pageInt-1)*10)

	// TODO: add left join table relations

	res, _ := tb.db().Query("select "+fields+" from "+tb.Info.Table+wheres+" order by "+params.SortField+" "+
		params.SortType+" LIMIT ? OFFSET ?", args...)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: add object pool

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
	} else if tb.ConnectionDriver == "postgresql" {
		size = int(total[0]["count"].(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	paginator := GetPaginator(path, params, size)

	return PanelInfo{
		Thead:       thead,
		InfoList:    infoList,
		Paginator:   paginator,
		Title:       title,
		Description: description,
	}

}

// GetDataFromDatabaseWithId query the single row of data.
func (tb Table) GetDataFromDatabaseWithId(id string) ([]types.Form, string, string) {

	fields := make([]string, 0)

	columnsModel, _ := db.WithDriver(tb.ConnectionDriver).Table(tb.Form.Table).ShowColumns()
	columns := GetColumns(columnsModel, tb.ConnectionDriver)

	for i := 0; i < len(tb.Form.FormList); i++ {
		if CheckInTable(columns, tb.Form.FormList[i].Field) {
			fields = append(fields, tb.Form.FormList[i].Field)
		}
	}

	res, _ := db.WithDriver(tb.ConnectionDriver).
		Table(tb.Form.Table).Select(fields...).
		Where("id", "=", id).
		First()

	idint64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tb.Form.FormList); i++ {
		if CheckInTable(columns, tb.Form.FormList[i].Field) {
			if tb.Form.FormList[i].FormType == "select" || tb.Form.FormList[i].FormType == "selectbox" || tb.Form.FormList[i].FormType == "select_single" {
				valueArr := tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.Form.FormList[i].TypeName, res[tb.Form.FormList[i].Field]),
				}).([]string)
				for _, v := range tb.Form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.Form.FormList[i].Value = tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.Form.FormList[i].TypeName, res[tb.Form.FormList[i].Field]),
				}).(string)
			}
		} else {
			if tb.Form.FormList[i].FormType == "select" || tb.Form.FormList[i].FormType == "selectbox" {
				valueArr := tb.Form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.Form.FormList[i].TypeName, res[tb.Form.FormList[i].Field]),
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

	value := make(dialect.H, 0)

	columnsModel, _ := db.WithDriver(tb.ConnectionDriver).Table(tb.Form.Table).ShowColumns()

	columns := GetColumns(columnsModel, tb.ConnectionDriver)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columns, k) {
			if len(v) > 0 {
				value[strings.Replace(k, "[]", "", -1)] = strings.Join(modules.RemoveBlankFromArray(v), ",")
			} else {
				value[strings.Replace(k, "[]", "", -1)] = v[0]
			}
		}
	}

	db.WithDriver(tb.ConnectionDriver).
		Table(tb.Form.Table).
		Where("id", "=", dataList["id"][0]).
		Update(value)

}

// InsertDataFromDatabase insert data.
func (tb Table) InsertDataFromDatabase(dataList map[string][]string) {

	value := make(dialect.H, 0)

	columnsModel, _ := db.WithDriver(tb.ConnectionDriver).Table(tb.Form.Table).ShowColumns()

	columns := GetColumns(columnsModel, tb.ConnectionDriver)
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columns, k) {
			if len(v) > 0 {
				value[strings.Replace(k, "[]", "", -1)] = strings.Join(modules.RemoveBlankFromArray(v), ",")
			} else {
				value[strings.Replace(k, "[]", "", -1)] = v[0]
			}
		}
	}

	db.WithDriver(tb.ConnectionDriver).
		Table(tb.Form.Table).
		Insert(value)
}

// DeleteDataFromDatabase delete data.
func (tb Table) DeleteDataFromDatabase(id string) {
	idArr := strings.Split(id, ",")
	for _, id := range idArr {
		db.WithDriver(tb.ConnectionDriver).
			Table(tb.Form.Table).
			Where("id", "=", id).
			Delete()
	}
	if tb.Form.Table == "goadmin_roles" {
		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_role_users").
			Where("role_id", "=", id).
			Delete()

		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_role_permissions").
			Where("role_id", "=", id).
			Delete()

		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_role_menu").
			Where("role_id", "=", id).
			Delete()
	}
	if tb.Form.Table == "goadmin_users" {
		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_role_users").
			Where("user_id", "=", id).
			Delete()

		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_user_permissions").
			Where("user_id", "=", id).
			Delete()
	}
	if tb.Form.Table == "goadmin_permissions" {
		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_role_permissions").
			Where("permission_id", "=", id).
			Delete()

		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_user_permissions").
			Where("permission_id", "=", id).
			Delete()
	}
	if tb.Form.Table == "goadmin_menu" {
		db.WithDriver(tb.ConnectionDriver).
			Table("goadmin_role_menu").
			Where("menu_id", "=", id).
			Delete()
	}
}

// db is a helper function return db connection.
func (tb Table) db() db.Connection {
	return db.GetConnectionByDriver(tb.ConnectionDriver)
}

func GetNewFormList(old []types.Form) []types.Form {
	var newForm []types.Form
	for _, v := range old {
		v.Value = ""
		if v.Field != "id" && v.Field != "created_at" && v.Field != "updated_at" {
			newForm = append(newForm, v)
		}
	}
	return newForm
}

// ***************************************
// helper function for database operation
// ***************************************

type Columns []string

func GetColumns(columnsModel []map[string]interface{}, driver string) Columns {
	columns := make(Columns, len(columnsModel))
	switch driver {
	case "postgresql":
		for key, model := range columnsModel {
			columns[key] = model["column_name"].(string)
		}
		return columns
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
