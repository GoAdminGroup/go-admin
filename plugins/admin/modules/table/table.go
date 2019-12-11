package table

import (
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"html/template"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Generator func() Table

type GeneratorList map[string]Generator

func (g GeneratorList) Add(key string, gen Generator) {
	g[key] = gen
}

var (
	generators = make(GeneratorList)
	tableList  = map[string]Table{}
)

func Get(key string) Table {
	return tableList[key]
}

func InitTableList() {
	for prefix, generator := range generators {
		tableList[prefix] = generator()
	}
}

// RefreshTableList refresh the table list when the table relationship changed.
func RefreshTableList() {
	for k, v := range generators {
		tableList[k] = v()
	}
}

// SetGenerators update generators.
func SetGenerators(gens map[string]Generator) {
	for key, gen := range gens {
		generators[key] = gen
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
	GetDataFromDatabase(path string, params parameter.Parameters, isAll bool) (PanelInfo, error)
	GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error)
	GetDataFromDatabaseWithId(id string) ([]types.FormField, [][]types.FormField, []string, string, string, error)
	UpdateDataFromDatabase(dataList form.Values) error
	InsertDataFromDatabase(dataList form.Values) error
	DeleteDataFromDatabase(id string) error
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
	FormData    []types.FormField
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

func DefaultConfig() Config {
	return Config{
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
	return tb.canAdd && !tb.info.IsHideNewButton
}

func (tb DefaultTable) GetPrimaryKey() PrimaryKey {
	return tb.primaryKey
}

func (tb DefaultTable) GetEditable() bool {
	return tb.editable && !tb.info.IsHideEditButton
}

func (tb DefaultTable) GetDeletable() bool {
	return tb.deletable && !tb.info.IsHideDeleteButton
}

func (tb DefaultTable) GetExportable() bool {
	return tb.exportable && !tb.info.IsHideExportButton
}

// GetDataFromDatabase query the data set.
func (tb DefaultTable) GetDataFromDatabase(path string, params parameter.Parameters, isAll bool) (PanelInfo, error) {
	if isAll {
		return tb.getAllDataFromDatabase(path, params)
	}
	return tb.getDataFromDatabase(path, params, []string{})
}

// GetDataFromDatabaseWithIds query the data set.
func (tb DefaultTable) GetDataFromDatabaseWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error) {
	return tb.getDataFromDatabase(path, params, ids)
}

func (tb DefaultTable) getTempModelData(res map[string]interface{}, resNext map[string]interface{}, params parameter.Parameters,
	columns Columns, isLast bool) (map[string]template.HTML, bool) {

	tempModelData := make(map[string]template.HTML)
	headField := ""

	primaryKeyValue := db.GetValueFromDatabaseType(tb.primaryKey.Type, res[tb.primaryKey.Name])

	// check if is join query result, which
	isCombine := !isLast && res[tb.primaryKey.Name] == resNext[tb.primaryKey.Name]

	for _, field := range tb.info.FieldList {

		headField = field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + "_" + field.Field
		}

		if field.Hide {
			continue
		}
		if !modules.InArrayWithoutEmpty(params.Columns, headField) {
			continue
		}

		var combineValue = ""

		if isCombine {
			var (
				resValue     = db.GetValueFromDatabaseType(field.TypeName, res[headField]).String()
				resNextValue = db.GetValueFromDatabaseType(field.TypeName, resNext[headField]).String()
			)
			if resValue == resNextValue {
				combineValue = resValue
			} else {
				combineValue = resValue + "|" + resNextValue
			}
		} else {
			combineValue = db.GetValueFromDatabaseType(field.TypeName, res[headField]).String()
		}

		var value interface{}
		if inArray(columns, headField) || field.Join.Valid() {
			value = field.ToDisplay(types.FieldModel{
				ID:    primaryKeyValue.String(),
				Value: combineValue,
				Row:   res,
			})
		} else {
			value = field.ToDisplay(types.FieldModel{
				ID:    primaryKeyValue.String(),
				Value: "",
				Row:   res,
			})
		}
		if valueStr, ok := value.(string); ok {
			tempModelData[headField] = template.HTML(valueStr)
		} else {
			tempModelData[headField] = value.(template.HTML)
		}
	}

	tempModelData[tb.primaryKey.Name] = template.HTML(primaryKeyValue.String())
	return tempModelData, isCombine
}

func (tb DefaultTable) getAllDataFromDatabase(path string, params parameter.Parameters) (PanelInfo, error) {
	var (
		connection     = tb.db()
		placeholder    = delimiter(connection.GetDelimiter(), "%s")
		queryStatement = "select %s from %s %s order by " + placeholder + " %s"
	)

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)

	var (
		fields     string
		joins      string
		headField  string
		joinTables = make([]string, 0)
		thead      = make([]map[string]string, 0)
	)
	for _, field := range tb.info.FieldList {
		if field.Field != tb.primaryKey.Name && inArray(columns, field.Field) &&
			!field.Join.Valid() {
			fields += tb.info.Table + "." + filterFiled(field.Field, connection.GetDelimiter()) + ","
		}

		headField = field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + "_" + field.Field
			fields += field.Join.Table + "." + filterFiled(field.Field, connection.GetDelimiter()) + " as " + headField + ","
			if !modules.InArray(joinTables, field.Join.Table) {
				joinTables = append(joinTables, field.Join.Table)
				joins += " left join " + filterFiled(field.Join.Table, connection.GetDelimiter()) + " on " +
					field.Join.Table + "." + filterFiled(field.Join.JoinField, connection.GetDelimiter()) + " = " +
					tb.info.Table + "." + filterFiled(field.Join.Field, connection.GetDelimiter())
			}
		}

		if field.Hide {
			continue
		}

		thead = append(thead, map[string]string{
			"head":       field.Head,
			"field":      headField,
			"edittype":   field.EditType.String(),
			"editoption": field.GetEditOptions(),
			"width":      strconv.Itoa(field.Width),
		})
	}

	fields += tb.info.Table + "." + filterFiled(tb.primaryKey.Name, connection.GetDelimiter())

	if !inArray(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
	}

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, params.SortField, params.SortType)

	logger.LogSQL(queryCmd, []interface{}{})

	res, err := connection.QueryWithConnection(tb.connection, queryCmd)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		rexNext := make(map[string]interface{})
		if i != len(res)-1 {
			rexNext = res[i+1]
		}

		tempModelData, isCombine := tb.getTempModelData(res[i], rexNext, params, columns, i == len(res)-1)

		if isCombine {
			if i != len(res)-2 {
				res = append(res[:i+1], res[i+2:]...)
			} else {
				res = res[:i+1]
			}
		}

		infoList = append(infoList, tempModelData)
	}

	return PanelInfo{
		InfoList:    infoList,
		Thead:       thead,
		Title:       tb.info.Title,
		Description: tb.info.Description,
	}, nil
}

func (tb DefaultTable) getDataFromDatabase(path string, params parameter.Parameters, ids []string) (PanelInfo, error) {

	var (
		connection     = tb.db()
		placeholder    = delimiter(connection.GetDelimiter(), "%s")
		queryStatement string
		countStatement string
	)

	if len(ids) > 0 {
		queryStatement = "select %s from %s %s where " + tb.primaryKey.Name + " in (%s) order by " + placeholder + " %s"
		countStatement = "select count(*) from " + placeholder + " where " + tb.primaryKey.Name + " in (%s)"
	} else {
		queryStatement = "select %s from " + placeholder + "%s %s order by " + placeholder + " %s LIMIT ? OFFSET ?"
		countStatement = "select count(*) from " + placeholder + "%s"
	}

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)

	var (
		sortable   string
		editable   string
		hide       string
		joins      string
		headField  string
		joinTables = make([]string, 0)
		filterForm = make([]types.FormField, 0)
	)
	for _, field := range tb.info.FieldList {
		if field.Field != tb.primaryKey.Name && inArray(columns, field.Field) &&
			!field.Join.Valid() {
			fields += tb.info.Table + "." + filterFiled(field.Field, connection.GetDelimiter()) + ","
		}

		headField = field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + "_" + field.Field
			fields += field.Join.Table + "." + filterFiled(field.Field, connection.GetDelimiter()) + " as " + headField + ","
			if !modules.InArray(joinTables, field.Join.Table) {
				joinTables = append(joinTables, field.Join.Table)
				joins += " left join " + filterFiled(field.Join.Table, connection.GetDelimiter()) + " on " +
					field.Join.Table + "." + filterFiled(field.Join.JoinField, connection.GetDelimiter()) + " = " +
					tb.info.Table + "." + filterFiled(field.Join.Field, connection.GetDelimiter())
			}
		}

		if field.Filterable {

			var value, value2 string

			if field.FilterType.IsRange() {
				value = params.GetFieldValue(field.Field + "_start__goadmin")
				value2 = params.GetFieldValue(field.Field + "_end__goadmin")
			} else {
				if field.FilterOperator == types.FilterOperatorFree {
					value2 = params.GetFieldOperator(field.Field).String()
				}
				value = params.GetFieldValue(field.Field)
			}

			filterForm = append(filterForm, types.FormField{
				Field:     field.Field,
				Head:      modules.AorB(field.FilterHead == "", field.Head, field.FilterHead),
				TypeName:  field.TypeName,
				HelpMsg:   field.FilterHelpMsg,
				FormType:  field.FilterType,
				Editable:  true,
				Value:     value,
				Value2:    value2,
				Options:   field.FilterOptions.SetSelected(params.GetFieldValue(field.Field), field.FilterType.SelectedLabel()),
				OptionExt: field.FilterOptionExt,
				Label:     field.FilterOperator.Label(),
			})

			if field.FilterOperator.AddOrNot() {
				filterForm = append(filterForm, types.FormField{
					Field:    field.Field + "__operator__",
					Head:     field.Head,
					TypeName: field.TypeName,
					Value:    field.FilterOperator.Value(),
					FormType: field.FilterType,
					Hide:     true,
				})
			}
		}

		if field.Hide {
			continue
		}
		sortable = modules.AorB(field.Sortable, "1", "0")
		editable = modules.AorB(field.EditAble, "true", "false")
		hide = modules.AorB(modules.InArrayWithoutEmpty(params.Columns, headField), "0", "1")
		thead = append(thead, map[string]string{
			"head":       field.Head,
			"sortable":   sortable,
			"field":      headField,
			"hide":       hide,
			"editable":   editable,
			"edittype":   field.EditType.String(),
			"editoption": field.GetEditOptions(),
			"width":      strconv.Itoa(field.Width),
		})
	}

	fields += tb.info.Table + "." + filterFiled(tb.primaryKey.Name, connection.GetDelimiter())

	if !inArray(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
	}

	var (
		wheres    = ""
		whereArgs = make([]interface{}, 0)
		args      = make([]interface{}, 0)
	)

	if len(ids) > 0 {
		for _, value := range ids {
			if value != "" {
				wheres += value + ","
			}
		}
		wheres = wheres[:len(wheres)-1]
	} else {
		wheres = " where "
		if len(params.Fields) == 0 {
			wheres = ""
		} else {
			for key, value := range params.Fields {

				var op types.FilterOperator
				if strings.Contains(key, "_end__goadmin") {
					key = strings.Replace(key, "_end__goadmin", "", -1)
					op = "<="
				} else if strings.Contains(key, "_start__goadmin") {
					key = strings.Replace(key, "_start__goadmin", "", -1)
					op = ">="
				} else if !strings.Contains(key, "__operator__") {
					op = params.GetFieldOperator(key)
				}

				if inArray(columns, key) {
					wheres += filterFiled(key, connection.GetDelimiter()) + " " + op.String() + " ? and "
					if op == types.FilterOperatorLike && !strings.Contains(value, "%") {
						whereArgs = append(whereArgs, "%"+value+"%")
					} else {
						whereArgs = append(whereArgs, value)
					}
				}

				if field := tb.info.FieldList.GetFieldByFieldName(key); field.Exist() && field.Join.Table != "" {
					wheres += field.Join.Table + "." + filterFiled(key, connection.GetDelimiter()) + " " + op.String() + " ? and "
					if op == types.FilterOperatorLike && !strings.Contains(value, "%") {
						whereArgs = append(whereArgs, "%"+value+"%")
					} else {
						whereArgs = append(whereArgs, value)
					}
				}
			}
			if wheres != " where " {
				wheres = wheres[:len(wheres)-4]
			}
		}
		args = append(whereArgs, params.PageSize, (modules.GetPage(params.Page)-1)*10)
	}

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, wheres, params.SortField, params.SortType)

	logger.LogSQL(queryCmd, args)

	res, err := connection.QueryWithConnection(tb.connection, queryCmd, args...)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {

		rexNext := make(map[string]interface{})
		if i != len(res)-1 {
			rexNext = res[i+1]
		}

		tempModelData, isCombine := tb.getTempModelData(res[i], rexNext, params, columns, i == len(res)-1)

		if isCombine {
			if i != len(res)-2 {
				res = append(res[:i+1], res[i+2:]...)
			} else {
				res = res[:i+1]
			}
		}

		infoList = append(infoList, tempModelData)
	}

	// TODO: use the dialect

	countCmd := fmt.Sprintf(countStatement, tb.info.Table, wheres)

	total, err := connection.QueryWithConnection(tb.connection, countCmd, whereArgs...)

	if err != nil {
		return PanelInfo{}, err
	}

	logger.LogSQL(countCmd, nil)

	var size int
	if tb.connectionDriver == "postgresql" {
		size = int(total[0]["count"].(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	return PanelInfo{
		Thead:       thead,
		InfoList:    infoList,
		Paginator:   paginator.Get(path, params, size, tb.info.GetPageSizeList()),
		Title:       tb.info.Title,
		FormData:    filterForm,
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
		if inArray(columns, formList[i].Field) {
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
				for _, field := range tb.form.FieldList {
					if value[j] == field.Field {
						rowValue := modules.AorB(inArray(columns, field.Field),
							db.GetValueFromDatabaseType(field.TypeName, res[field.Field]).String(), "")
						list[j] = field.UpdateValue(id, rowValue, res)
						break
					}
				}
			}

			groupFormList = append(groupFormList, list)
			groupHeaders = append(groupHeaders, tb.form.TabHeaders[key])
		}
		return tb.form.FieldList, groupFormList, groupHeaders, tb.form.Title, tb.form.Description, nil
	}

	for key, field := range formList {
		rowValue := modules.AorB(inArray(columns, field.Field),
			db.GetValueFromDatabaseType(field.TypeName, res[field.Field]).String(), "")
		formList[key] = field.UpdateValue(id, rowValue, res)
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

	if tb.form.UpdateFn != nil {
		return tb.form.UpdateFn(dataList)
	}

	_, err := tb.sql().Table(tb.form.Table).
		Where(tb.primaryKey.Name, "=", dataList.Get(tb.primaryKey.Name)).
		Update(tb.getInjectValueFromFormValue(dataList))

	// TODO: some errors should be ignored.
	if err != nil {
		if tb.connectionDriver != db.DriverPostgresql {
			return err
		}
		if !strings.Contains(err.Error(), "LastInsertId is not supported by this driver") {
			return err
		}
	}

	// NOTE: Database Transaction may be considered here.

	if tb.form.PostHook != nil {
		go func() {

			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			err := tb.form.PostHook(dataList)
			if err != nil {
				logger.Error(err)
			}
		}()
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

	if tb.form.InsertFn != nil {
		return tb.form.InsertFn(dataList)
	}

	id, err := tb.sql().Table(tb.form.Table).Insert(tb.getInjectValueFromFormValue(dataList))

	// TODO: some errors should be ignored.
	if err != nil {
		if tb.connectionDriver != db.DriverPostgresql {
			return err
		}
		if !strings.Contains(err.Error(), "LastInsertId is not supported by this driver") {
			return err
		}
	}

	dataList.Add(tb.GetPrimaryKey().Name, strconv.Itoa(int(id)))

	// NOTE: Database Transaction may be considered here.

	if tb.form.PostHook != nil {
		go func() {

			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			err := tb.form.PostHook(dataList)
			if err != nil {
				logger.Error(err)
			}
		}()
	}

	return nil
}

func (tb DefaultTable) getInjectValueFromFormValue(dataList form.Values) dialect.H {
	value := make(dialect.H)

	columnsModel, _ := tb.sql().Table(tb.form.Table).ShowColumns()

	columns := getColumns(columnsModel, tb.connectionDriver)
	var (
		fun          types.PostFieldFilterFn
		exceptString = []string{tb.primaryKey.Name, "_previous_", "_method", "_t"}
	)
	for k, v := range dataList {
		k = strings.Replace(k, "[]", "", -1)
		if !modules.InArray(exceptString, k) {
			if inArray(columns, k) {
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
					} else if len(vv) > 0 {
						value[k] = vv[0]
					} else {
						value[k] = ""
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
func (tb DefaultTable) DeleteDataFromDatabase(id string) error {
	idArr := strings.Split(id, ",")

	if tb.info.DeleteFn != nil {

		if len(idArr) == 0 {
			return errors.New("wrong parameter")
		}

		return tb.info.DeleteFn(idArr)
	}

	if tb.info.PreDeleteFn != nil && len(idArr) > 0 {
		if err := tb.info.PreDeleteFn(idArr); err != nil {
			return err
		}
	}

	for _, id := range idArr {
		tb.delete(tb.form.Table, tb.primaryKey.Name, id)
	}

	if tb.info.DeleteHook != nil && len(idArr) > 0 {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			if err := tb.info.DeleteHook(idArr); err != nil {
				logger.Error(err)
			}
		}()
	}

	return nil
}

func (tb DefaultTable) delete(table, key, id string) {
	_ = tb.sql().Table(table).
		Where(key, "=", id).
		Delete()
}

// db is a helper function return raw db connection.
func (tb DefaultTable) db() db.Connection {
	return db.GetConnectionFromService(services.Get(tb.connectionDriver))
}

var services service.List

func SetServices(srv service.List) {
	services = srv
}

// sql is a helper function return db sql.
func (tb DefaultTable) sql() *db.SQL {
	return db.WithDriverAndConnection(tb.connection, db.GetConnectionFromService(services.Get(tb.connectionDriver)))
}

func GetNewFormList(groupHeaders []string,
	group [][]string,
	old []types.FormField) ([]types.FormField, [][]types.FormField, []string) {

	if len(group) == 0 {
		var newForm []types.FormField
		for _, v := range old {
			v.Value = v.Default
			if !v.NotAllowAdd {
				v.Editable = true
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
					if !v.NotAllowAdd {
						v.Editable = true
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

func delimiter(del, s string) string {
	return del + s + del
}

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

// inArray checks the find string is in the columns or not.
func inArray(columns []string, find string) bool {
	for i := 0; i < len(columns); i++ {
		if columns[i] == find {
			return true
		}
	}
	return false
}
