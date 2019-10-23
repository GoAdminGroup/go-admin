package table

import (
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
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
	GetDataFromDatabase(path string, params parameter.Parameters) (PanelInfo, error)
	GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error)
	GetDataFromDatabaseWithId(id string) ([]types.FormField, [][]types.FormField, []string, string, string, error)
	UpdateDataFromDatabase(dataList form.Values) error
	InsertDataFromDatabase(dataList form.Values) error
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
	Thead       Thead
	InfoList    InfoList
	Paginator   types.PaginatorAttribute
	Title       string
	Description string
}

type Thead []map[string]string

func (t Thead) GroupBy(group [][]string) []Thead {
	var res = make([]Thead, len(group))

	for key, value := range group {
		var newThead = make(Thead, len(t))

		for index, info := range t {
			if modules.InArray(value, info["field"]) {
				newThead[index] = info
			}
		}

		res[key] = newThead
	}

	return res
}

type InfoList []map[string]template.HTML

func (i InfoList) GroupBy(groups types.TabGroups) []InfoList {

	var res = make([]InfoList, len(groups))

	for key, value := range groups {
		var newInfoList = make(InfoList, len(i))

		for index, info := range i {
			var newRow = make(map[string]template.HTML)
			for mk, m := range info {
				if modules.InArray(value, mk) {
					newRow[mk] = m
				}
			}
			newInfoList[index] = newRow
		}

		res[key] = newInfoList
	}

	return res
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
		info:             types.NewInfoPanel(),
		form:             types.NewFormPanel(),
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
		if value.Filterable {
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
func (tb DefaultTable) GetDataFromDatabase(path string, params parameter.Parameters) (PanelInfo, error) {

	connection := tb.db()

	var (
		queryStatement = "select %s from " + connection.GetDelimiter() + "%s" + connection.GetDelimiter() +
			"%s %s order by " + connection.GetDelimiter() + "%s" + connection.GetDelimiter() + " %s LIMIT ? OFFSET ?"
		countStatement = "select count(*) from " + connection.GetDelimiter() + "%s" + connection.GetDelimiter() + "%s"
	)

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)

	var (
		sortable   string
		hide       string
		joins      string
		headField  string
		joinTables = make([]string, 0)
	)
	for i := 0; i < len(tb.info.FieldList); i++ {
		if tb.info.FieldList[i].Field != tb.primaryKey.Name && checkInTable(columns, tb.info.FieldList[i].Field) &&
			!tb.info.FieldList[i].Join.Valid() {
			fields += tb.info.Table + "." + filterFiled(tb.info.FieldList[i].Field, connection.GetDelimiter()) + ","
		}

		headField = tb.info.FieldList[i].Field

		if tb.info.FieldList[i].Join.Valid() {
			headField = tb.info.FieldList[i].Join.Table + "_" + tb.info.FieldList[i].Field
			fields += tb.info.FieldList[i].Join.Table + "." + filterFiled(tb.info.FieldList[i].Field, connection.GetDelimiter()) + " as " + headField + ","
			if !modules.InArray(joinTables, tb.info.FieldList[i].Join.Table) {
				joinTables = append(joinTables, tb.info.FieldList[i].Join.Table)
				joins += " left join " + filterFiled(tb.info.FieldList[i].Join.Table, connection.GetDelimiter()) + " on " +
					tb.info.FieldList[i].Join.Table + "." + filterFiled(tb.info.FieldList[i].Join.JoinField, connection.GetDelimiter()) + " = " +
					tb.info.Table + "." + filterFiled(tb.info.FieldList[i].Join.Field, connection.GetDelimiter())
			}
		}

		if tb.info.FieldList[i].Hide {
			continue
		}
		sortable = modules.AorB(tb.info.FieldList[i].Sortable, "1", "0")
		hide = modules.AorB(modules.InArrayWithoutEmpty(params.Columns, headField), "0", "1")
		thead = append(thead, map[string]string{
			"head":     tb.info.FieldList[i].Head,
			"sortable": sortable,
			"field":    headField,
			"hide":     hide,
			"width":    strconv.Itoa(tb.info.FieldList[i].Width),
		})
	}

	fields += tb.info.Table + "." + filterFiled(tb.primaryKey.Name, connection.GetDelimiter())

	if !checkInTable(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
	}

	wheres := " where "
	whereArgs := make([]interface{}, 0)
	if len(params.Fields) == 0 {
		wheres = ""
	} else {
		for key, value := range params.Fields {
			if checkInTable(columns, key) {
				wheres += filterFiled(key, connection.GetDelimiter()) + " = ? and "
				whereArgs = append(whereArgs, value)
			}
		}
		wheres = wheres[:len(wheres)-4]
	}
	args := append(whereArgs, params.PageSize, (modules.GetPage(params.Page)-1)*10)

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, wheres, params.SortField, params.SortType)

	logger.LogSql(queryCmd, args)

	res, err := connection.QueryWithConnection(tb.connection, queryCmd, args...)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: add object pool

		tempModelData := make(map[string]template.HTML)
		row := res[i]

		primaryKeyValue := db.GetValueFromDatabaseType(tb.primaryKey.Type, res[i][tb.primaryKey.Name])

		for j := 0; j < len(tb.info.FieldList); j++ {

			headField = tb.info.FieldList[j].Field

			if tb.info.FieldList[j].Join.Valid() {
				headField = tb.info.FieldList[j].Join.Table + "_" + tb.info.FieldList[j].Field
			}

			if tb.info.FieldList[j].Hide {
				continue
			}
			if !modules.InArrayWithoutEmpty(params.Columns, headField) {
				continue
			}
			var value interface{}
			if checkInTable(columns, headField) || tb.info.FieldList[j].Join.Valid() {
				value = tb.info.FieldList[j].ToDisplay(types.FieldModel{
					ID:    primaryKeyValue.String(),
					Value: db.GetValueFromDatabaseType(tb.info.FieldList[j].TypeName, row[headField]).String(),
					Row:   row,
				})
			} else {
				value = tb.info.FieldList[j].ToDisplay(types.FieldModel{
					ID:    primaryKeyValue.String(),
					Value: "",
					Row:   row,
				})
			}
			if valueStr, ok := value.(string); ok {
				tempModelData[headField] = template.HTML(valueStr)
			} else {
				tempModelData[headField] = value.(template.HTML)
			}
		}

		tempModelData[tb.primaryKey.Name] = template.HTML(primaryKeyValue.String())

		infoList = append(infoList, tempModelData)
	}

	// TODO: use the dialect

	countCmd := fmt.Sprintf(countStatement, tb.info.Table, wheres)

	total, err := connection.QueryWithConnection(tb.connection, countCmd, whereArgs...)

	if err != nil {
		return PanelInfo{}, err
	}

	logger.LogSql(countCmd, whereArgs)

	var size int
	if tb.connectionDriver == "postgresql" {
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
	}, nil
}

// GetDataFromDatabaseWithIds query the data set.
func (tb DefaultTable) GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error) {

	connection := tb.db()

	var (
		queryStatement = "select %s from %s %s where " + tb.primaryKey.Name + " in (%s) order by " + connection.GetDelimiter() +
			"%s" + connection.GetDelimiter() + " %s"
		countStatement = "select count(*) from " + connection.GetDelimiter() + "%s" + connection.GetDelimiter() +
			" where " + tb.primaryKey.Name + " in (%s)"
	)

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)

	var (
		sortable   string
		hide       string
		joins      string
		headField  string
		joinTables = make([]string, 0)
	)
	for i := 0; i < len(tb.info.FieldList); i++ {
		if tb.info.FieldList[i].Field != tb.primaryKey.Name && checkInTable(columns, tb.info.FieldList[i].Field) &&
			!tb.info.FieldList[i].Join.Valid() {
			fields += tb.info.Table + "." + filterFiled(tb.info.FieldList[i].Field, connection.GetDelimiter()) + ","
		}

		headField = tb.info.FieldList[i].Field

		if tb.info.FieldList[i].Join.Valid() {
			headField = tb.info.FieldList[i].Join.Table + "_" + tb.info.FieldList[i].Field
			fields += tb.info.FieldList[i].Join.Table + "." + filterFiled(tb.info.FieldList[i].Field, connection.GetDelimiter()) + " as " + headField + ","
			if !modules.InArray(joinTables, tb.info.FieldList[i].Join.Table) {
				joinTables = append(joinTables, tb.info.FieldList[i].Join.Table)
				joins += " left join " + filterFiled(tb.info.FieldList[i].Join.Table, connection.GetDelimiter()) + " on " +
					tb.info.FieldList[i].Join.Table + "." + filterFiled(tb.info.FieldList[i].Join.JoinField, connection.GetDelimiter()) + " = " +
					tb.info.Table + "." + filterFiled(tb.info.FieldList[i].Join.Field, connection.GetDelimiter())
			}
		}

		if tb.info.FieldList[i].Hide {
			continue
		}
		sortable = modules.AorB(tb.info.FieldList[i].Sortable, "1", "0")
		hide = modules.AorB(modules.InArrayWithoutEmpty(params.Columns, headField), "0", "1")
		thead = append(thead, map[string]string{
			"head":     tb.info.FieldList[i].Head,
			"sortable": sortable,
			"field":    headField,
			"hide":     hide,
			"width":    strconv.Itoa(tb.info.FieldList[i].Width),
		})
	}

	fields += tb.info.Table + "." + filterFiled(tb.primaryKey.Name, connection.GetDelimiter())

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

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, whereIds, params.SortField, params.SortType)

	res, err := connection.QueryWithConnection(tb.connection, queryCmd)

	if err != nil {
		return PanelInfo{}, err
	}

	logger.LogSql(queryCmd, nil)

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		// TODO: add object pool

		tempModelData := make(map[string]template.HTML)
		row := res[i]

		primaryKeyValue := db.GetValueFromDatabaseType(tb.primaryKey.Type, res[i][tb.primaryKey.Name])

		for j := 0; j < len(tb.info.FieldList); j++ {

			headField = tb.info.FieldList[i].Field

			if tb.info.FieldList[j].Join.Valid() {
				headField = tb.info.FieldList[j].Join.Table + "_" + tb.info.FieldList[j].Field
			}

			if tb.info.FieldList[j].Hide {
				continue
			}
			if !modules.InArrayWithoutEmpty(params.Columns, headField) {
				continue
			}
			var value interface{}
			if checkInTable(columns, headField) {
				value = tb.info.FieldList[j].ToDisplay(types.FieldModel{
					ID:    primaryKeyValue.String(),
					Value: db.GetValueFromDatabaseType(tb.info.FieldList[j].TypeName, row[headField]).String(),
					Row:   row,
				})
			} else {
				value = tb.info.FieldList[j].ToDisplay(types.FieldModel{
					ID:    primaryKeyValue.String(),
					Value: "",
					Row:   row,
				})
			}

			if valueStr, ok := value.(string); ok {
				tempModelData[headField] = template.HTML(valueStr)
			} else {
				tempModelData[headField] = value.(template.HTML)
			}
		}

		tempModelData[tb.primaryKey.Name] = template.HTML(primaryKeyValue.String())

		infoList = append(infoList, tempModelData)
	}

	// TODO: use the dialect

	countCmd := fmt.Sprintf(countStatement, tb.info.Table, whereIds)

	total, err := connection.QueryWithConnection(tb.connection, countCmd)

	if err != nil {
		return PanelInfo{}, err
	}

	logger.LogSql(countCmd, nil)

	var size int
	if tb.connectionDriver == "postgresql" {
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
	}, nil
}

// GetDataFromDatabaseWithId query the single row of data.
func (tb DefaultTable) GetDataFromDatabaseWithId(id string) ([]types.FormField, [][]types.FormField, []string, string, string, error) {

	fields := make([]string, 0)

	columnsModel, err := tb.sql().Table(tb.form.Table).ShowColumns()

	if err != nil {
		return nil, nil, nil, "", "", err
	}

	columns := getColumns(columnsModel, tb.connectionDriver)

	formList := tb.form.FieldList.Copy()

	for i := 0; i < len(tb.form.FieldList); i++ {
		if checkInTable(columns, formList[i].Field) {
			fields = append(fields, formList[i].Field)
		}
	}

	res, err := tb.sql().
		Table(tb.form.Table).Select(fields...).
		Where(tb.primaryKey.Name, "=", id).
		First()

	if err != nil {
		return nil, nil, nil, "", "", err
	}

	var (
		groupFormList = make([][]types.FormField, 0)
		groupHeaders  = make([]string, 0)
	)

	if len(tb.form.TabGroups) > 0 {
		for key, value := range tb.form.TabGroups {
			list := make([]types.FormField, len(value))
			for j := 0; j < len(value); j++ {
				for i := 0; i < len(tb.form.FieldList); i++ {
					if value[j] == formList[i].Field {
						if checkInTable(columns, formList[i].Field) {
							if formList[i].FormType.IsSelect() {
								valueArr := formList[i].ToDisplay(types.FieldModel{
									ID:    id,
									Value: db.GetValueFromDatabaseType(formList[i].TypeName, res[formList[i].Field]).String(),
									Row:   res,
								}).([]string)
								for _, v := range formList[i].Options {
									if modules.InArray(valueArr, v["value"]) {
										v["selected"] = "selected"
									} else {
										v["selected"] = ""
									}
								}
							} else if formList[i].FormType.IsRadio() {
								value := formList[i].ToDisplay(types.FieldModel{
									ID:    id,
									Value: db.GetValueFromDatabaseType(formList[i].TypeName, res[formList[i].Field]).String(),
									Row:   res,
								}).(string)
								for _, v := range formList[i].Options {
									if value == v["value"] {
										v["selected"] = "checked"
									} else {
										v["selected"] = ""
									}
								}
							} else {
								formList[i].Value = formList[i].ToDisplay(types.FieldModel{
									ID:    id,
									Value: db.GetValueFromDatabaseType(formList[i].TypeName, res[formList[i].Field]).String(),
									Row:   res,
								}).(string)
							}
						} else {
							if formList[i].FormType.IsSelect() {
								valueArr := formList[i].ToDisplay(types.FieldModel{
									ID:    id,
									Value: "",
									Row:   res,
								}).([]string)
								for _, v := range formList[i].Options {
									if modules.InArray(valueArr, v["value"]) {
										v["selected"] = "selected"
									} else {
										v["selected"] = ""
									}
								}
							} else if formList[i].FormType.IsRadio() {
								value := formList[i].ToDisplay(types.FieldModel{
									ID:    id,
									Value: "",
									Row:   res,
								}).(string)
								for _, v := range formList[i].Options {
									if value == v["value"] {
										v["selected"] = "checked"
									} else {
										v["selected"] = ""
									}
								}
							} else {
								formList[i].Value = formList[i].ToDisplay(types.FieldModel{
									ID:    id,
									Value: "",
									Row:   res,
								}).(string)
							}
						}
						list[j] = formList[i]
						break
					}
				}
			}

			groupFormList = append(groupFormList, list)
			groupHeaders = append(groupHeaders, tb.form.TabHeaders[key])
		}
		return tb.form.FieldList, groupFormList, groupHeaders, tb.form.Title, tb.form.Description, nil
	}

	for i := 0; i < len(tb.form.FieldList); i++ {
		if checkInTable(columns, formList[i].Field) {
			if formList[i].FormType.IsSelect() {
				valueArr := formList[i].ToDisplay(types.FieldModel{
					ID:    id,
					Value: db.GetValueFromDatabaseType(formList[i].TypeName, res[formList[i].Field]).String(),
					Row:   res,
				}).([]string)
				for _, v := range formList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				formList[i].Value = formList[i].ToDisplay(types.FieldModel{
					ID:    id,
					Value: db.GetValueFromDatabaseType(formList[i].TypeName, res[formList[i].Field]).String(),
					Row:   res,
				}).(string)
			}
		} else {
			if formList[i].FormType.IsSelect() {
				valueArr := formList[i].ToDisplay(types.FieldModel{
					ID:    id,
					Value: "",
					Row:   res,
				}).([]string)
				for _, v := range formList[i].Options {
					if modules.InArray(valueArr, v["value"]) {
						v["selected"] = "selected"
					}
				}
			} else {
				formList[i].Value = formList[i].ToDisplay(types.FieldModel{
					ID:    id,
					Value: "",
					Row:   res,
				}).(string)
			}
		}
	}

	return formList, groupFormList, groupHeaders, tb.form.Title, tb.form.Description, nil
}

// UpdateDataFromDatabase update data.
func (tb DefaultTable) UpdateDataFromDatabase(dataList form.Values) error {

	if tb.form.Validator != nil {
		if err := tb.form.Validator(dataList); err != nil {
			return err
		}
	}

	_, _ = tb.sql().Table(tb.form.Table).
		Where(tb.primaryKey.Name, "=", dataList.Get(tb.primaryKey.Name)).
		Update(tb.getValues(dataList))

	//if err != nil {
	//	return err
	//}

	if tb.form.PostHook != nil {
		go tb.form.PostHook(dataList)
	}

	return nil
}

// InsertDataFromDatabase insert data.
func (tb DefaultTable) InsertDataFromDatabase(dataList form.Values) error {

	if tb.form.Validator != nil {
		if err := tb.form.Validator(dataList); err != nil {
			return err
		}
	}

	id, _ := tb.sql().Table(tb.form.Table).Insert(tb.getValues(dataList))

	//if err != nil {
	//	return err
	//}

	dataList.Add(tb.GetPrimaryKey().Name, strconv.Itoa(int(id)))

	if tb.form.PostHook != nil {
		go tb.form.PostHook(dataList)
	}

	return nil
}

func (tb DefaultTable) getValues(dataList form.Values) dialect.H {
	value := make(dialect.H)

	columnsModel, _ := tb.sql().Table(tb.form.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)
	var fun types.PostFieldFilterFn
	for k, v := range dataList {
		k = strings.Replace(k, "[]", "", -1)
		if k != tb.primaryKey.Name && k != "_previous_" && k != "_method" && k != "_t" {
			if checkInTable(columns, k) {
				delimiter := ","
				for i := 0; i < len(tb.form.FieldList); i++ {
					if k == tb.form.FieldList[i].Field {
						fun = tb.form.FieldList[i].PostFilterFn
						delimiter = modules.SetDefault(tb.form.FieldList[i].DefaultOptionDelimiter, ",")
					}
				}
				vv := modules.RemoveBlankFromArray(v)
				if fun != nil {
					value[k] = fun(types.PostFieldModel{
						ID:    dataList.Get(tb.primaryKey.Name),
						Value: vv,
					})
				} else {
					if len(vv) > 1 {
						value[k] = strings.Join(vv, delimiter)
					} else {
						value[k] = vv[0]
					}
				}
			} else {
				fun := tb.form.FieldList.FindByFieldName(k).PostFilterFn
				if fun != nil {
					fun(types.PostFieldModel{
						ID:    dataList.Get(tb.primaryKey.Name),
						Value: modules.RemoveBlankFromArray(v),
					})
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

func GetNewFormList(groupHeaders []string,
	group [][]string,
	old []types.FormField,
	primaryKey string) ([]types.FormField, [][]types.FormField, []string) {

	if len(group) == 0 {
		var newForm []types.FormField
		for _, v := range old {
			v.Value = v.Default
			if v.Field != primaryKey && !v.NotAllowAdd {
				newForm = append(newForm, v)
			}
		}
		return newForm, [][]types.FormField{}, []string{}
	}

	var (
		newForm = make([][]types.FormField, 0)
		headers = make([]string, 0)
	)

	for key, value := range group {
		list := make([]types.FormField, 0)

		for i := 0; i < len(value); i++ {
			for _, v := range old {
				if v.Field == value[i] {
					v.Value = v.Default
					if v.Field != primaryKey && !v.NotAllowAdd {
						list = append(list, v)
						break
					}
				}
			}
		}

		newForm = append(newForm, list)
		headers = append(headers, groupHeaders[key])
	}

	return []types.FormField{}, newForm, headers
}

// ***************************************
// helper function for database operation
// ***************************************

func filterFiled(filed, delimiter string) string {
	return delimiter + filed + delimiter
}

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
