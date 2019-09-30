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
	Generators = make(GeneratorList)
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
	GetPrimaryKey() PrimaryKey
	GetFiltersMap() []map[string]string
	GetDataFromDatabase(path string, params parameter.Parameters) PanelInfo
	GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) PanelInfo
	GetDataFromDatabaseWithId(id string) ([]types.Form, string, string)
	UpdateDataFromDatabase(dataList form.Values)
	InsertDataFromDatabase(dataList form.Values)
	DeleteDataFromDatabase(id string)
}

type PrimaryKey struct {
	Type db.DatabaseType
	Name string
}

const (
	DefaultPrimaryKeyName = "id"
	DefaultConnectionName = "default"
)

type DefaultTable struct {
	info             *types.InfoPanel
	form             *types.FormPanel
	connectionDriver string
	connection       string
	canAdd           bool
	editable         bool
	deletable        bool
	exportable       bool
	primaryKey       PrimaryKey
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
	PrimaryKey PrimaryKey
}

var DefaultConfig = Config{
	Driver:     db.DriverMysql,
	CanAdd:     true,
	Editable:   true,
	Deletable:  true,
	Exportable: false,
	Connection: DefaultConnectionName,
	PrimaryKey: PrimaryKey{
		Type: db.Int,
		Name: DefaultPrimaryKeyName,
	},
}

func (config Config) SetPrimaryKeyType(typ string) Config {
	config.PrimaryKey.Type = db.GetDTAndCheck(typ)
	return config
}

func (config Config) SetCanAdd(canAdd bool) Config {
	config.CanAdd = canAdd
	return config
}

func (config Config) SetEditable(editable bool) Config {
	config.Editable = editable
	return config
}

func (config Config) SetDeletable(deletable bool) Config {
	config.Deletable = deletable
	return config
}

func (config Config) SetExportable(exportable bool) Config {
	config.Exportable = exportable
	return config
}

func (config Config) SetConnection(connection string) Config {
	config.Connection = connection
	return config
}

func DefaultConfigWithDriver(driver string) Config {
	return Config{
		Driver:     driver,
		Connection: DefaultConnectionName,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: false,
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	}
}

func DefaultConfigWithDriverAndConnection(driver, conn string) Config {
	return Config{
		Driver:     driver,
		Connection: conn,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: false,
		PrimaryKey: PrimaryKey{
			Type: db.Int,
			Name: DefaultPrimaryKeyName,
		},
	}
}

func NewDefaultTable(cfg Config) Table {
	return DefaultTable{
		info:             &types.InfoPanel{},
		form:             &types.FormPanel{},
		connectionDriver: cfg.Driver,
		connection:       cfg.Connection,
		canAdd:           cfg.CanAdd,
		editable:         cfg.Editable,
		deletable:        cfg.Deletable,
		exportable:       cfg.Exportable,
		primaryKey:       cfg.PrimaryKey,
	}
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

func (tb DefaultTable) GetPrimaryKey() PrimaryKey {
	return tb.primaryKey
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
			"name":  tb.primaryKey.Name,
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
		if tb.info.FieldList[i].Field != tb.primaryKey.Name && checkInTable(columns, tb.info.FieldList[i].Field) {
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

	fields += tb.primaryKey.Name

	if !checkInTable(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
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

		tempModelData := make(map[string]template.HTML)
		row := res[i]

		primaryKeyValue := db.GetValueFromDatabaseType(tb.primaryKey.Type, res[i][tb.primaryKey.Name])

		for j := 0; j < len(tb.info.FieldList); j++ {
			if tb.info.FieldList[j].Hide {
				continue
			}
			if checkInTable(columns, tb.info.FieldList[j].Field) {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    primaryKeyValue.String(),
					Value: db.GetValueFromDatabaseType(tb.info.FieldList[j].TypeName, row[tb.info.FieldList[j].Field]).String(),
					Row:   row,
				}).(string))
			} else {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    primaryKeyValue.String(),
					Value: "",
					Row:   row,
				}).(string))
			}
		}

		tempModelData[tb.primaryKey.Name] = template.HTML(primaryKeyValue.String())

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
		if tb.info.FieldList[i].Field != tb.primaryKey.Name && checkInTable(columns, tb.info.FieldList[i].Field) {
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

	fields += tb.primaryKey.Name

	if !checkInTable(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
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

		tempModelData := make(map[string]template.HTML)
		row := res[i]

		primaryKeyValue := db.GetValueFromDatabaseType(tb.primaryKey.Type, res[i][tb.primaryKey.Name])

		for j := 0; j < len(tb.info.FieldList); j++ {
			if tb.info.FieldList[j].Hide {
				continue
			}
			if checkInTable(columns, tb.info.FieldList[j].Field) {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    primaryKeyValue.String(),
					Value: db.GetValueFromDatabaseType(tb.info.FieldList[j].TypeName, row[tb.info.FieldList[j].Field]).String(),
					Row:   row,
				}).(string))
			} else {
				tempModelData[tb.info.FieldList[j].Head] = template.HTML(tb.info.FieldList[j].FilterFn(types.RowModel{
					ID:    primaryKeyValue.String(),
					Value: "",
					Row:   row,
				}).(string))
			}
		}

		tempModelData[tb.primaryKey.Name] = template.HTML(primaryKeyValue.String())

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
		Where(tb.primaryKey.Name, "=", id).
		First()

	for i := 0; i < len(tb.form.FormList); i++ {
		if checkInTable(columns, tb.form.FormList[i].Field) {
			if tb.form.FormList[i].FormType == "select" || tb.form.FormList[i].FormType == "selectbox" || tb.form.FormList[i].FormType == "select_single" {
				valueArr := tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    id,
					Value: db.GetValueFromDatabaseType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]).String(),
					Row:   res,
				}).([]string)
				for _, v := range tb.form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.form.FormList[i].Value = tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    id,
					Value: db.GetValueFromDatabaseType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]).String(),
					Row:   res,
				}).(string)
			}
		} else {
			if tb.form.FormList[i].FormType == "select" || tb.form.FormList[i].FormType == "selectbox" {
				valueArr := tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    id,
					Value: db.GetValueFromDatabaseType(tb.form.FormList[i].TypeName, res[tb.form.FormList[i].Field]).String(),
					Row:   res,
				}).([]string)
				for _, v := range tb.form.FormList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				tb.form.FormList[i].Value = tb.form.FormList[i].FilterFn(types.RowModel{
					ID:    id,
					Value: tb.form.FormList[i].Field,
					Row:   res,
				}).(string)
			}
		}
	}

	return tb.form.FormList, tb.form.Title, tb.form.Description
}

// UpdateDataFromDatabase update data.
func (tb DefaultTable) UpdateDataFromDatabase(dataList form.Values) {
	_, _ = tb.sql().Table(tb.form.Table).
		Where(tb.primaryKey.Name, "=", dataList.Get(tb.primaryKey.Name)).
		Update(tb.getValues(dataList))

}

// InsertDataFromDatabase insert data.
func (tb DefaultTable) InsertDataFromDatabase(dataList form.Values) {
	_, _ = tb.sql().Table(tb.form.Table).Insert(tb.getValues(dataList))
}

func (tb DefaultTable) getValues(dataList form.Values) dialect.H {
	value := make(dialect.H)

	columnsModel, _ := tb.sql().Table(tb.form.Table).ShowColumns()

	var id = int64(0)
	if idArr, ok := dataList[tb.primaryKey.Name]; ok {
		idInt, _ := strconv.Atoi(idArr[0])
		id = int64(idInt)
	}

	columns := getColumns(columnsModel, tb.connectionDriver)
	var fun types.PostFieldFilterFn
	for k, v := range dataList {
		k = strings.Replace(k, "[]", "", -1)
		if k != tb.primaryKey.Name && k != "_previous_" && k != "_method" && k != "_t" && checkInTable(columns, k) {
			delimiter := ","
			for i := 0; i < len(tb.form.FormList); i++ {
				if k == tb.form.FormList[i].Field {
					fun = tb.form.FormList[i].PostFilterFn
					delimiter = modules.SetDefault(tb.form.FormList[i].DefaultOptionDelimiter, ",")
				}
			}
			vv := modules.RemoveBlankFromArray(v)
			if fun != nil {
				value[k] = fun(types.PostRowModel{
					ID:    id,
					Value: vv,
				})
			} else {
				if len(vv) > 1 {
					value[k] = strings.Join(vv, delimiter)
				} else {
					value[k] = vv[0]
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
		tb.delete(tb.form.Table, tb.primaryKey.Name, id)
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

func GetNewFormList(old []types.Form, primaryKey string) []types.Form {
	var newForm []types.Form
	for _, v := range old {
		v.Value = ""
		if v.Field != primaryKey && v.Field != "created_at" && v.Field != "updated_at" {
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
