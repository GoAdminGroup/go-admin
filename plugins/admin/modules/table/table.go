package table

import (
	"fmt"
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/db/dialect"
	"github.com/chenhg5/go-admin/plugins/admin/modules"
	"github.com/chenhg5/go-admin/plugins/admin/modules/form"
	"github.com/chenhg5/go-admin/plugins/admin/modules/paginator"
	"github.com/chenhg5/go-admin/plugins/admin/modules/parameter"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
	"strconv"
	"strings"
)

type Generator func() Table

type GeneratorList map[string]Generator

func (g GeneratorList) Add(key string, gen Generator) {
	g[key] = gen
}

var (
	Generators = make(GeneratorList, 0)
	List       = map[string]Table{}
)

func InitTableList() {
	List = make(map[string]Table, len(Generators))
	for prefix, generator := range Generators {
		List[prefix] = generator()
	}
}

// RefreshTableList refresh the table list when the table relationship changed.
func RefreshTableList() {
	for k, v := range Generators {
		List[k] = v()
	}
}

// SetGenerators update Generators.
func SetGenerators(generators map[string]Generator) {
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
	GetExportable() bool
	GetFiltersMap() []map[string]string
	GetDataFromDatabase(path string, params parameter.Parameters) PanelInfo
	GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) PanelInfo
	GetDataFromDatabaseWithId(id string) ([]types.Form, string, string)
	UpdateDataFromDatabase(dataList form.FormValue)
	InsertDataFromDatabase(dataList form.FormValue)
	DeleteDataFromDatabase(id string)
}

type DefaultTable struct {
	info             *types.InfoPanel
	form             *types.FormPanel
	connectionDriver string
	connection       string
	canAdd           bool
	editable         bool
	deletable        bool
	exportable       bool
	prefix           string
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

type Config struct {
	Driver     string
	Connection string
	CanAdd     bool
	Editable   bool
	Deletable  bool
	Exportable bool
}

var DefaultConfig = &Config{
	Driver:     "mysql",
	CanAdd:     true,
	Editable:   true,
	Deletable:  true,
	Exportable: false,
	Connection: "default",
}

func DefaultConfigWithDriver(driver string) *Config {
	return &Config{
		Driver:     driver,
		Connection: "default",
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: false,
	}
}

func DefaultConfigWithDriverAndConnection(driver, conn string) *Config {
	return &Config{
		Driver:     driver,
		Connection: conn,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: false,
	}
}

func NewDefaultTable(cfg *Config) Table {
	tb := &DefaultTable{
		info:             &types.InfoPanel{},
		form:             &types.FormPanel{},
		connectionDriver: cfg.Driver,
		connection:       cfg.Connection,
		canAdd:           cfg.CanAdd,
		editable:         cfg.Editable,
		deletable:        cfg.Deletable,
		exportable:       cfg.Exportable,
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

func (tb DefaultTable) GetExportable() bool {
	return tb.exportable
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
func (tb DefaultTable) GetDataFromDatabase(path string, params parameter.Parameters) PanelInfo {

	const (
		queryStatement = "select %s from %s%s order by %s %s LIMIT ? OFFSET ?"
		countStatement = "select count(*) from %s%s"
	)

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)

	var sortable string
	for i := 0; i < len(tb.info.FieldList); i++ {
		if tb.info.FieldList[i].Field != "id" && checkInTable(columns, tb.info.FieldList[i].Field) {
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

	if !checkInTable(columns, params.SortField) {
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
	args := append(whereArgs, params.PageSize, (modules.GetPage(params.Page)-1)*10)

	// TODO: add left join table relations

	res, _ := tb.db().QueryWithConnection(tb.connection,
		fmt.Sprintf(queryStatement, fields, tb.info.Table, wheres, params.SortField, params.SortType),
		args...)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: add object pool

		tempModelData := make(map[string]template.HTML, 0)
		row := res[i]

		for j := 0; j < len(tb.info.FieldList); j++ {
			if tb.info.FieldList[j].Hide {
				continue
			}
			if checkInTable(columns, tb.info.FieldList[j].Field) {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    row["id"].(int64),
					Value: getStringFromType(tb.info.FieldList[j].TypeName, row[tb.info.FieldList[j].Field]),
					Row:   row,
				}).(string))
			} else {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    row["id"].(int64),
					Value: "",
					Row:   row,
				}).(string))
			}
		}

		tempModelData["id"] = template.HTML(getStringFromType("int", res[i]["id"]))

		infoList = append(infoList, tempModelData)
	}

	// TODO: use the dialect

	total, _ := tb.db().QueryWithConnection(tb.connection, fmt.Sprintf(countStatement, tb.info.Table, wheres), whereArgs...)
	var size int
	if tb.connectionDriver == "sqlite" {
		size = int(total[0]["count(*)"].(int64))
	} else if tb.connectionDriver == "postgresql" {
		size = int(total[0]["count"].(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	return PanelInfo{
		Thead:       thead,
		InfoList:    infoList,
		Paginator:   paginator.Get(path, params, size),
		Title:       tb.info.Title,
		Description: tb.info.Description,
		CanAdd:      tb.canAdd,
		Editable:    tb.editable,
		Deletable:   tb.deletable,
	}
}

// GetDataFromDatabase query the data set.
func (tb DefaultTable) GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) PanelInfo {

	const (
		queryStatement = "select %s from %s where id in (%s) order by %s %s"
		countStatement = "select count(*) from %s where id in (%s)"
	)

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)

	var sortable string
	for i := 0; i < len(tb.info.FieldList); i++ {
		if tb.info.FieldList[i].Field != "id" && checkInTable(columns, tb.info.FieldList[i].Field) {
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

	if !checkInTable(columns, params.SortField) {
		params.SortField = "id"
	}

	whereIds := ""

	for _, value := range ids {
		if value != "" {
			whereIds += value + ","
		}			
	}
	whereIds = whereIds[:len(whereIds)-1]

	// TODO: add left join table relations

	res, _ := tb.db().QueryWithConnection(tb.connection,
		fmt.Sprintf(queryStatement, fields, tb.info.Table, whereIds, params.SortField, params.SortType))

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: add object pool

		tempModelData := make(map[string]template.HTML, 0)
		row := res[i]

		for j := 0; j < len(tb.info.FieldList); j++ {
			if tb.info.FieldList[j].Hide {
				continue
			}
			if checkInTable(columns, tb.info.FieldList[j].Field) {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    row["id"].(int64),
					Value: getStringFromType(tb.info.FieldList[j].TypeName, row[tb.info.FieldList[j].Field]),
					Row:   row,
				}).(string))
			} else {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    row["id"].(int64),
					Value: "",
					Row:   row,
				}).(string))
			}
		}

		tempModelData["id"] = template.HTML(getStringFromType("int", res[i]["id"]))

		infoList = append(infoList, tempModelData)
	}

	// TODO: use the dialect

	total, _ := tb.db().QueryWithConnection(tb.connection, fmt.Sprintf(countStatement, tb.info.Table, whereIds))
	var size int
	if tb.connectionDriver == "sqlite" {
		size = int(total[0]["count(*)"].(int64))
	} else if tb.connectionDriver == "postgresql" {
		size = int(total[0]["count"].(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	return PanelInfo{
		Thead:       thead,
		InfoList:    infoList,
		Paginator:   paginator.Get(path, params, size),
		Title:       tb.info.Title,
		Description: tb.info.Description,
		CanAdd:      tb.canAdd,
		Editable:    tb.editable,
		Deletable:   tb.deletable,
	}
}

// GetDataFromDatabaseWithId query the single row of data.
func (tb DefaultTable) GetDataFromDatabaseWithId(id string) ([]types.Form, string, string) {

	fields := make([]string, 0)

	columnsModel, _ := tb.sql().Table(tb.form.Table).ShowColumns()
	columns := getColumns(columnsModel, tb.connectionDriver)

	for i := 0; i < len(tb.form.FormList); i++ {
		if checkInTable(columns, tb.form.FormList[i].Field) {
			fields = append(fields, tb.form.FormList[i].Field)
		}
	}

	res, _ := tb.sql().
		Table(tb.form.Table).Select(fields...).
		Where("id", "=", id).
		First()

	idInt64, _ := strconv.ParseInt(id, 10, 64)

	for i := 0; i < len(tb.form.FormList); i++ {
		if checkInTable(columns, tb.form.FormList[i].Field) {
			if tb.form.FormList[i].FormType == "select" || tb.form.FormList[i].FormType == "selectbox" || tb.form.FormList[i].FormType == "select_single" {
				valueArr := tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    idInt64,
					Value: getStringFromType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]),
					Row:   res,
				}).([]string)
				for _, v := range tb.form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.form.FormList[i].Value = tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    idInt64,
					Value: getStringFromType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]),
					Row:   res,
				}).(string)
			}
		} else {
			if tb.form.FormList[i].FormType == "select" || tb.form.FormList[i].FormType == "selectbox" {
				valueArr := tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    idInt64,
					Value: getStringFromType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]),
					Row:   res,
				}).([]string)
				for _, v := range tb.form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.form.FormList[i].Value = tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    idInt64,
					Value: tb.form.FormList[i].Field,
					Row:   res,
				}).(string)
			}
		}
	}

	return tb.form.FormList, tb.form.Title, tb.form.Description
}

// UpdateDataFromDatabase update data.
func (tb DefaultTable) UpdateDataFromDatabase(dataList form.FormValue) {
	_, _ = tb.sql().Table(tb.form.Table).
		Where("id", "=", dataList.Get("id")).
		Update(tb.getValues(dataList))

}

// InsertDataFromDatabase insert data.
func (tb DefaultTable) InsertDataFromDatabase(dataList form.FormValue) {
	_, _ = tb.sql().Table(tb.form.Table).Insert(tb.getValues(dataList))
}

func (tb DefaultTable) getValues(dataList form.FormValue) dialect.H {
	value := make(dialect.H, 0)

	columnsModel, _ := tb.sql().Table(tb.form.Table).ShowColumns()

	var id = int64(0)
	if idArr, ok := dataList["id"]; ok {
		idInt, _ := strconv.Atoi(idArr[0])
		id = int64(idInt)
	}

	columns := getColumns(columnsModel, tb.connectionDriver)
	var fun types.FieldFilterFn
	for k, v := range dataList {
		if k != "id" && k != "_previous_" && k != "_method" && k != "_t" && checkInTable(columns, k) {
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
	_ = tb.sql().Table(table).
		Where(key, "=", id).
		Delete()
}

// db is a helper function return raw db connection.
func (tb DefaultTable) db() db.Connection {
	return db.GetConnectionByDriver(tb.connectionDriver)
}

// sql is a helper function return db sql.
func (tb DefaultTable) sql() *db.Sql {
	return db.WithDriverAndConnection(tb.connection, tb.connectionDriver)
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

func getColumns(columnsModel []map[string]interface{}, driver string) Columns {
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
			columns[key] = string(model["name"].(string))
		}
		return columns
	default:
		panic("wrong driver")
	}
}

// checkInTable checks the find string is in the columns or not.
func checkInTable(columns []string, find string) bool {
	for i := 0; i < len(columns); i++ {
		if columns[i] == find {
			return true
		}
	}
	return false
}

// getStringFromType get the value from sql raw type.
func getStringFromType(typeName string, value interface{}) string {
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
