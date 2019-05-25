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

type Table interface {
	GetInfo() *types.InfoPanel
	GetForm() *types.FormPanel
	GetCanAdd() bool
	GetEditable() bool
	GetDeletable() bool
	GetFiltersMap() []map[string]string
	GetDataFromDatabase(path string, params *Parameters) PanelInfo
	GetDataFromDatabaseWithId(id string) ([]types.Form, string, string)
	UpdateDataFromDatabase(dataList map[string][]string)
	InsertDataFromDatabase(dataList map[string][]string)
	DeleteDataFromDatabase(id string)
}

type DefaultTable struct {
	info             *types.InfoPanel
	form             *types.FormPanel
	connectionDriver string
	canAdd           bool
	editable         bool
	deletable        bool
}

type PanelInfo struct {
	Thead       []map[string]string
	InfoList    []map[string]template.HTML
	Paginator   types.PaginatorAttribute
	Title       string
	Description string
	CanAdd      bool
	Editable    bool
	Deletable   bool
}

func NewDefaultTable(connectionDriver string, canAdd, editable, deletable bool) Table {
	tb := &DefaultTable{
		info: &types.InfoPanel{},
		form: &types.FormPanel{},
		connectionDriver: connectionDriver,
		canAdd: canAdd,
		editable: editable,
		deletable: deletable,
	}
	return tb
}

func (tb DefaultTable) GetInfo() *types.InfoPanel {
	return tb.info
}

func (tb DefaultTable) GetForm() *types.FormPanel {
	return tb.form
}

func (tb DefaultTable) GetCanAdd() bool {
	return tb.canAdd
}

func (tb DefaultTable) GetEditable() bool {
	return tb.editable
}

func (tb DefaultTable) GetDeletable() bool {
	return tb.deletable
}

func (tb DefaultTable) GetFiltersMap() []map[string]string {
	var filters = make([]map[string]string, 0)
	for _, value := range tb.info.FieldList {
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
func (tb DefaultTable) GetDataFromDatabase(path string, params *Parameters) PanelInfo {

	pageInt, _ := strconv.Atoi(params.Page)

	title := tb.info.Title
	description := tb.info.Description

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := db.WithDriver(tb.connectionDriver).Table(tb.info.Table).ShowColumns()

	columns := GetColumns(columnsModel, tb.connectionDriver)

	var sortable string
	for i := 0; i < len(tb.info.FieldList); i++ {
		if tb.info.FieldList[i].Field != "id" && CheckInTable(columns, tb.info.FieldList[i].Field) {
			fields += tb.info.FieldList[i].Field + ","
		}
		if tb.info.FieldList[i].Hide {
			continue
		}
		sortable = "0"
		if tb.info.FieldList[i].Sortable {
			sortable = "1"
		}
		thead = append(thead, map[string]string{
			"head":     tb.info.FieldList[i].Head,
			"sortable": sortable,
			"field":    tb.info.FieldList[i].Field,
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

	res, _ := tb.db().Query("select "+fields+" from "+tb.info.Table+wheres+" order by "+params.SortField+" "+
		params.SortType+" LIMIT ? OFFSET ?", args...)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: add object pool

		tempModelData := make(map[string]template.HTML, 0)
		row := res[i]

		for j := 0; j < len(tb.info.FieldList); j++ {
			if tb.info.FieldList[j].Hide {
				continue
			}
			if CheckInTable(columns, tb.info.FieldList[j].Field) {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].ExcuFun(types.RowModel{
					ID:    row["id"].(int64),
					Value: GetStringFromType(tb.info.FieldList[j].TypeName, row[tb.info.FieldList[j].Field]),
					Row:   row,
				}).(string))
			} else {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].ExcuFun(types.RowModel{
					ID:    row["id"].(int64),
					Value: "",
					Row:   row,
				}).(string))
			}
		}

		tempModelData["id"] = template.HTML(GetStringFromType("int", res[i]["id"]))

		infoList = append(infoList, tempModelData)
	}

	// TODO: use the dialect

	total, _ := tb.db().Query("select count(*) from "+tb.info.Table+wheres, whereArgs...)
	var size int
	if tb.connectionDriver == "sqlite" {
		size = int(total[0]["count(*)"].(int64))
	} else if tb.connectionDriver == "postgresql" {
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
		CanAdd:      tb.canAdd,
		Editable:    tb.editable,
		Deletable:   tb.deletable,
	}

}

// GetDataFromDatabaseWithId query the single row of data.
func (tb DefaultTable) GetDataFromDatabaseWithId(id string) ([]types.Form, string, string) {

	fields := make([]string, 0)

	columnsModel, _ := db.WithDriver(tb.connectionDriver).Table(tb.form.Table).ShowColumns()
	columns := GetColumns(columnsModel, tb.connectionDriver)

	for i := 0; i < len(tb.form.FormList); i++ {
		if CheckInTable(columns, tb.form.FormList[i].Field) {
			fields = append(fields, tb.form.FormList[i].Field)
		}
	}

	res, _ := db.WithDriver(tb.connectionDriver).
		Table(tb.form.Table).Select(fields...).
		Where("id", "=", id).
		First()

	idint64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tb.form.FormList); i++ {
		if CheckInTable(columns, tb.form.FormList[i].Field) {
			if tb.form.FormList[i].FormType == "select" || tb.form.FormList[i].FormType == "selectbox" || tb.form.FormList[i].FormType == "select_single" {
				valueArr := tb.form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]),
					Row:   res,
				}).([]string)
				for _, v := range tb.form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.form.FormList[i].Value = tb.form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]),
					Row:   res,
				}).(string)
			}
		} else {
			if tb.form.FormList[i].FormType == "select" || tb.form.FormList[i].FormType == "selectbox" {
				valueArr := tb.form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: GetStringFromType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]),
					Row:   res,
				}).([]string)
				for _, v := range tb.form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.form.FormList[i].Value = tb.form.FormList[i].ExcuFun(types.RowModel{
					ID:    idint64,
					Value: tb.form.FormList[i].Field,
					Row:   res,
				}).(string)
			}
		}
	}

	return tb.form.FormList, tb.form.Title, tb.form.Description
}

// UpdateDataFromDatabase update data.
func (tb DefaultTable) UpdateDataFromDatabase(dataList map[string][]string) {
	_, _ = db.WithDriver(tb.connectionDriver).
		Table(tb.form.Table).
		Where("id", "=", dataList["id"][0]).
		Update(tb.getValues(dataList))

}

// InsertDataFromDatabase insert data.
func (tb DefaultTable) InsertDataFromDatabase(dataList map[string][]string) {
	_, _ = db.WithDriver(tb.connectionDriver).
		Table(tb.form.Table).
		Insert(tb.getValues(dataList))
}

func (tb DefaultTable) getValues(dataList map[string][]string) dialect.H {
	value := make(dialect.H, 0)

	columnsModel, _ := db.WithDriver(tb.connectionDriver).Table(tb.form.Table).ShowColumns()

	var id = int64(0)
	if idArr, ok := dataList["id"]; ok {
		idInt, _ := strconv.Atoi(idArr[0])
		id = int64(idInt)
	}

	columns := GetColumns(columnsModel, tb.connectionDriver)
	var fun types.FieldValueFun
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && CheckInTable(columns, k) {
			for i := 0; i < len(tb.form.FormList); i++ {
				if k == tb.form.FormList[i].Field {
					fun = tb.form.FormList[i].PostFun
				}
			}
			if len(v) > 0 {
				if fun != nil {
					value[strings.Replace(k, "[]", "", -1)] = fun(types.RowModel{
						ID:    id,
						Value: strings.Join(modules.RemoveBlankFromArray(v), ","),
					})
				} else {
					value[strings.Replace(k, "[]", "", -1)] = strings.Join(modules.RemoveBlankFromArray(v), ",")
				}
			} else {
				if fun != nil {
					value[strings.Replace(k, "[]", "", -1)] = fun(types.RowModel{
						ID:    id,
						Value: v[0],
					})
				} else {
					value[strings.Replace(k, "[]", "", -1)] = v[0]
				}
			}
		}
	}
	return value
}

// DeleteDataFromDatabase delete data.
func (tb DefaultTable) DeleteDataFromDatabase(id string) {
	idArr := strings.Split(id, ",")
	for _, id := range idArr {
		tb.delete(tb.form.Table, "id", id)
	}
	if tb.form.Table == "goadmin_roles" {
		tb.delete("goadmin_role_users", "role_id", id)
		tb.delete("goadmin_role_permissions", "role_id", id)
		tb.delete("goadmin_role_menu", "role_id", id)
	}
	if tb.form.Table == "goadmin_users" {
		tb.delete("goadmin_role_users", "user_id", id)
		tb.delete("goadmin_user_permissions", "user_id", id)
	}
	if tb.form.Table == "goadmin_permissions" {
		tb.delete("goadmin_role_permissions", "permission_id", id)
		tb.delete("goadmin_user_permissions", "permission_id", id)
	}
	if tb.form.Table == "goadmin_menu" {
		tb.delete("goadmin_role_menu", "menu_id", id)
	}
}

func (tb DefaultTable) delete(table, key, id string) {
	_ = db.WithDriver(tb.connectionDriver).
		Table(table).
		Where(key, "=", id).
		Delete()
}

// db is a helper function return db connection.
func (tb DefaultTable) db() db.Connection {
	return db.GetConnectionByDriver(tb.connectionDriver)
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
			columns[key] = string(model["name"].([]uint8))
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
		return string(value.([]uint8))
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
