package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
	form2 "github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DefaultTable struct {
	info             *types.InfoPanel
	form             *types.FormPanel
	detail           *types.InfoPanel
	connectionDriver string
	connection       string
	canAdd           bool
	editable         bool
	deletable        bool
	exportable       bool
	primaryKey       PrimaryKey
	sourceURL        string
	getDataFun       GetDataFun
}

type GetDataFun func(path string, params parameter.Parameters, isAll bool, ids []string) ([]map[string]interface{}, int)

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

func NewDefaultTable(cfg Config) Table {
	return DefaultTable{
		info:             types.NewInfoPanel(cfg.PrimaryKey.Name),
		form:             types.NewFormPanel(),
		detail:           types.NewInfoPanel(cfg.PrimaryKey.Name),
		connectionDriver: cfg.Driver,
		connection:       cfg.Connection,
		canAdd:           cfg.CanAdd,
		editable:         cfg.Editable,
		deletable:        cfg.Deletable,
		exportable:       cfg.Exportable,
		primaryKey:       cfg.PrimaryKey,
		sourceURL:        cfg.SourceURL,
		getDataFun:       cfg.GetDataFun,
	}
}

func (tb DefaultTable) Copy() Table {
	return DefaultTable{
		form: types.NewFormPanel().SetTable(tb.form.Table).
			SetDescription(tb.form.Description).
			SetTitle(tb.form.Title),
		info: types.NewInfoPanel(tb.primaryKey.Name).SetTable(tb.info.Table).
			SetDescription(tb.info.Description).
			SetTitle(tb.info.Title),
		connectionDriver: tb.connectionDriver,
		connection:       tb.connection,
		canAdd:           tb.canAdd,
		editable:         tb.editable,
		deletable:        tb.deletable,
		exportable:       tb.exportable,
		primaryKey:       tb.primaryKey,
		sourceURL:        tb.sourceURL,
		getDataFun:       tb.getDataFun,
	}
}

func (tb DefaultTable) GetInfo() *types.InfoPanel {
	return tb.info
}

func (tb DefaultTable) GetDetail() *types.InfoPanel {
	return tb.detail
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

func (tb DefaultTable) IsShowDetail() bool {
	return !tb.info.IsHideDetailButton
}

func (tb DefaultTable) GetExportable() bool {
	return tb.exportable && !tb.info.IsHideExportButton
}

// GetData query the data set.
func (tb DefaultTable) GetData(path string, params parameter.Parameters, isAll bool) (PanelInfo, error) {

	var (
		data      []map[string]interface{}
		size      int
		beginTime = time.Now()
	)

	if tb.getDataFun != nil {
		data, size = tb.getDataFun(path, params, isAll, []string{})
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(path, params, isAll, []string{})
	} else if tb.info.GetDataFn != nil {
		data, size = tb.info.GetDataFn(params.WithIsAll(isAll))
	} else if isAll {
		return tb.getAllDataFromDatabase(path, params)
	} else {
		return tb.getDataFromDatabase(path, params, []string{})
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

	endTime := time.Now()

	return PanelInfo{
		Thead:    thead,
		InfoList: infoList,
		Paginator: paginator.Get(path, params, size, tb.info.GetPageSizeList()).
			SetExtraInfo(template.HTML(fmt.Sprintf("<b>" + language.Get("query time") + ": </b>" +
				fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000)))),
		Title:       tb.info.Title,
		FormData:    filterForm,
		Description: tb.info.Description,
	}, nil
}

type GetDataFromURLRes struct {
	Data []map[string]interface{}
	Size int
}

func (tb DefaultTable) getDataFromURL(path string, params parameter.Parameters, isAll bool, ids []string) ([]map[string]interface{}, int) {

	u := ""
	if strings.Contains(tb.sourceURL, "?") {
		u = tb.sourceURL + "&" + params.Join()
	} else {
		u = tb.sourceURL + "?" + params.Join()
	}
	res, err := http.Get(u + "&pk=" + strings.Join(ids, ","))

	if err != nil {
		return []map[string]interface{}{}, 0
	}

	defer func() {
		_ = res.Body.Close()
	}()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return []map[string]interface{}{}, 0
	}

	var data GetDataFromURLRes

	err = json.Unmarshal(body, &data)

	if err != nil {
		return []map[string]interface{}{}, 0
	}

	return data.Data, data.Size
}

// GetDataWithIds query the data set.
func (tb DefaultTable) GetDataWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error) {

	var (
		data      []map[string]interface{}
		size      int
		beginTime = time.Now()
	)

	if tb.getDataFun != nil {
		data, size = tb.getDataFun(path, params, false, ids)
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(path, params, false, ids)
	} else if tb.info.GetDataFn != nil {
		data, size = tb.info.GetDataFn(params.WithPK(ids...))
	} else {
		return tb.getDataFromDatabase(path, params, ids)
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

	endTime := time.Now()

	return PanelInfo{
		Thead:    thead,
		InfoList: infoList,
		Paginator: paginator.Get(path, params, size, tb.info.GetPageSizeList()).
			SetExtraInfo(template.HTML(fmt.Sprintf("<b>" + language.Get("query time") + ": </b>" +
				fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000)))),
		Title:       tb.info.Title,
		FormData:    filterForm,
		Description: tb.info.Description,
	}, nil
}

func (tb DefaultTable) getTempModelData(res map[string]interface{}, params parameter.Parameters, columns Columns) map[string]template.HTML {

	tempModelData := make(map[string]template.HTML)
	headField := ""

	primaryKeyValue := db.GetValueFromDatabaseType(tb.primaryKey.Type, res[tb.primaryKey.Name], len(columns) == 0)

	for _, field := range tb.info.FieldList {

		headField = field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + "_goadmin_join_" + field.Field
		}

		if field.Hide {
			continue
		}
		if !modules.InArrayWithoutEmpty(params.Columns, headField) {
			continue
		}

		typeName := field.TypeName

		if field.Join.Valid() {
			typeName = db.Varchar
		}

		combineValue := db.GetValueFromDatabaseType(typeName, res[headField], len(columns) == 0).String()

		var value interface{}
		if len(columns) == 0 || modules.InArray(columns, headField) || field.Join.Valid() {
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
	return tempModelData
}

func (tb DefaultTable) getAllDataFromDatabase(path string, params parameter.Parameters) (PanelInfo, error) {
	var (
		connection     = tb.db()
		queryStatement = "select %s from %s %s %s order by " + modules.Delimiter(connection.GetDelimiter(), "%s") + " %s"
	)

	columns, _ := tb.getColumns(tb.info.Table)

	thead, fields, joins := tb.info.FieldList.GetThead(types.TableInfo{}, params, columns)

	fields += tb.info.Table + "." + modules.FilterField(tb.primaryKey.Name, connection.GetDelimiter())

	var (
		wheres    = ""
		whereArgs = make([]interface{}, 0)
		existKeys = make([]string, 0)
	)

	wheres, whereArgs, existKeys = params.Statement(wheres, connection.GetDelimiter(), whereArgs, columns, existKeys,
		tb.info.FieldList.GetFieldFilterProcessValue, tb.info.FieldList.GetFieldJoinTable)
	wheres, whereArgs = tb.info.Wheres.Statement(wheres, connection.GetDelimiter(), whereArgs, existKeys, columns)
	wheres, whereArgs = tb.info.WhereRaws.Statement(wheres, whereArgs)

	if wheres != "" {
		wheres = " where " + wheres
	}

	if !modules.InArray(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
	}

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, wheres, params.SortField, params.SortType)

	logger.LogSQL(queryCmd, []interface{}{})

	res, err := connection.QueryWithConnection(tb.connection, queryCmd, whereArgs...)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {
		infoList = append(infoList, tb.getTempModelData(res[i], params, columns))
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
		placeholder    = modules.Delimiter(connection.GetDelimiter(), "%s")
		queryStatement string
		countStatement string
	)

	beginTime := time.Now()

	if len(ids) > 0 {
		queryStatement = "select %s from %s %s where " + tb.primaryKey.Name + " in (%s) %s order by " + placeholder + " %s"
		countStatement = "select count(*) from " + placeholder + " %s where " + tb.primaryKey.Name + " in (%s)"
	} else {
		if connection.Name() == "mssql" {
			queryStatement = "SELECT * FROM (SELECT ROW_NUMBER() OVER (ORDER BY " + placeholder + " %s) as ROWNUMBER_, %s from " +
				placeholder + "%s %s %s  ) as TMP_ WHERE TMP_.ROWNUMBER_ > ? AND TMP_.ROWNUMBER_ <= ?"
			countStatement = "select count(*) as [size] from " + placeholder + " %s %s"
		} else {
			queryStatement = "select %s from " + placeholder + "%s %s %s order by " + placeholder + " %s LIMIT ? OFFSET ?"
			countStatement = "select count(*) from " + placeholder + " %s %s"
		}
	}

	columns, _ := tb.getColumns(tb.info.Table)

	thead, fields, joins, joinTables, filterForm := tb.getTheadAndFilterForm(params, columns)

	fields += tb.info.Table + "." + modules.FilterField(tb.primaryKey.Name, connection.GetDelimiter())

	if !modules.InArray(columns, params.SortField) {
		params.SortField = tb.primaryKey.Name
	}

	var (
		wheres    = ""
		whereArgs = make([]interface{}, 0)
		args      = make([]interface{}, 0)
		existKeys = make([]string, 0)
	)

	if len(ids) > 0 {
		for _, value := range ids {
			if value != "" {
				wheres += value + ","
			}
		}
		wheres = wheres[:len(wheres)-1]
	} else {

		wheres, whereArgs, existKeys = params.Statement(wheres, connection.GetDelimiter(), whereArgs, columns, existKeys,
			tb.info.FieldList.GetFieldFilterProcessValue, tb.info.FieldList.GetFieldJoinTable)
		wheres, whereArgs = tb.info.Wheres.Statement(wheres, connection.GetDelimiter(), whereArgs, existKeys, columns)
		wheres, whereArgs = tb.info.WhereRaws.Statement(wheres, whereArgs)

		if wheres != "" {
			wheres = " where " + wheres
		}

		if connection.Name() == "mssql" {
			args = append(whereArgs, (params.PageInt-1)*params.PageSizeInt, params.PageInt*params.PageSizeInt)
		} else {
			args = append(whereArgs, params.PageSize, (params.PageInt-1)*params.PageSizeInt)
		}
	}

	groupBy := ""
	if len(joinTables) > 0 {
		groupBy = " GROUP BY " + tb.info.Table + "." + modules.FilterField(tb.GetPrimaryKey().Name, connection.GetDelimiter())
	}

	queryCmd := ""
	if connection.Name() == "mssql" {
		queryCmd = fmt.Sprintf(queryStatement, params.SortField, params.SortType, fields, tb.info.Table, joins, wheres, groupBy)
	} else {
		queryCmd = fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, wheres, groupBy, params.SortField, params.SortType)
	}

	logger.LogSQL(queryCmd, args)

	res, err := connection.QueryWithConnection(tb.connection, queryCmd, args...)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]template.HTML, 0)

	for i := 0; i < len(res); i++ {
		infoList = append(infoList, tb.getTempModelData(res[i], params, columns))
	}

	// TODO: use the dialect

	if len(ids) > 0 {
		joins = ""
	}

	countCmd := fmt.Sprintf(countStatement, tb.info.Table, joins, wheres)

	total, err := connection.QueryWithConnection(tb.connection, countCmd, whereArgs...)

	if err != nil {
		return PanelInfo{}, err
	}

	logger.LogSQL(countCmd, nil)

	var size int
	if tb.connectionDriver == "postgresql" {
		size = int(total[0]["count"].(int64))
	} else if tb.connectionDriver == "mssql" {
		size = int(total[0]["size"].(int64))
	} else {
		size = int(total[0]["count(*)"].(int64))
	}

	endTime := time.Now()

	return PanelInfo{
		Thead:    thead,
		InfoList: infoList,
		Paginator: paginator.Get(path, params, size, tb.info.GetPageSizeList()).
			SetExtraInfo(template.HTML(fmt.Sprintf("<b>" + language.Get("query time") + ": </b>" +
				fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000)))),
		Title:       tb.info.Title,
		FormData:    filterForm,
		Description: tb.info.Description,
	}, nil
}

// GetDataWithId query the single row of data.
func (tb DefaultTable) GetDataWithId(id string) ([]types.FormField, [][]types.FormField, []string, string, string, error) {

	var (
		res     map[string]interface{}
		columns Columns
		fields  = make([]string, 0)
	)

	if tb.getDataFun != nil {
		list, _ := tb.getDataFun("", parameter.BaseParam(), false, []string{id})
		if len(list) > 0 {
			res = list[0]
		}
	} else if tb.sourceURL != "" {
		list, _ := tb.getDataFromURL("", parameter.BaseParam(), false, []string{id})
		if len(list) > 0 {
			res = list[0]
		}
	} else if tb.info.GetDataFn != nil {
		list, _ := tb.info.GetDataFn(parameter.BaseParam().WithPK(id))
		if len(list) > 0 {
			res = list[0]
		}
	} else {

		columns, _ = tb.getColumns(tb.form.Table)

		var err error

		res, err = tb.sql().
			Table(tb.form.Table).Select(fields...).
			Where(tb.primaryKey.Name, "=", id).
			First()

		if err != nil {
			return nil, nil, nil, "", "", err
		}
	}

	formList := tb.form.FieldList.Copy()

	for i := 0; i < len(tb.form.FieldList); i++ {
		if modules.InArray(columns, formList[i].Field) {
			fields = append(fields, formList[i].Field)
		}
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
						rowValue := modules.AorB(modules.InArray(columns, field.Field) || len(columns) == 0,
							db.GetValueFromDatabaseType(field.TypeName, res[field.Field], len(columns) == 0).String(), "")
						list[j] = field.UpdateValue(id, rowValue, res)
						if list[j].FormType == form2.File && list[j].Value != template.HTML("") {
							list[j].Value2 = "/" + config.Get().Store.Prefix + "/" + string(list[j].Value)
						}
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
		rowValue := modules.AorB(modules.InArray(columns, field.Field) || len(columns) == 0,
			db.GetValueFromDatabaseType(field.TypeName, res[field.Field], len(columns) == 0).String(), "")
		formList[key] = field.UpdateValue(id, rowValue, res)

		if formList[key].FormType == form2.File && formList[key].Value != template.HTML("") {
			formList[key].Value2 = "/" + config.Get().Store.Prefix + "/" + string(formList[key].Value)
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

	if tb.form.UpdateFn != nil {
		return tb.form.UpdateFn(dataList)
	}

	if tb.form.PreProcessFn != nil {
		dataList.Add("__go_admin_post_type", "0")
		dataList = tb.form.PreProcessFn(dataList)
	}

	_, err := tb.sql().Table(tb.form.Table).
		Where(tb.primaryKey.Name, "=", dataList.Get(tb.primaryKey.Name)).
		Update(tb.getInjectValueFromFormValue(dataList))

	// TODO: some errors should be ignored.
	if err != nil && !strings.Contains(err.Error(), "no affect") {
		if tb.connectionDriver != db.DriverPostgresql && tb.connectionDriver != db.DriverMssql {
			return err
		}
		if !strings.Contains(err.Error(), "LastInsertId is not supported") {
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

			dataList.Add("__go_admin_post_type", "0")

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

	if tb.form.PreProcessFn != nil {
		dataList.Add("__go_admin_post_type", "1")
		dataList = tb.form.PreProcessFn(dataList)
	}

	id, err := tb.sql().Table(tb.form.Table).Insert(tb.getInjectValueFromFormValue(dataList))

	// TODO: some errors should be ignored.
	if err != nil {
		if tb.connectionDriver != db.DriverPostgresql && tb.connectionDriver != db.DriverMssql {
			return err
		}
		if !strings.Contains(err.Error(), "LastInsertId is not supported") {
			return err
		}
	}

	dataList.Add(tb.GetPrimaryKey().Name, strconv.Itoa(int(id)))

	if tb.form.PostHook != nil {
		go func() {

			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			dataList.Add("__go_admin_post_type", "1")

			err := tb.form.PostHook(dataList)
			if err != nil {
				logger.Error(err)
			}
		}()
	}

	return nil
}

func (tb DefaultTable) getInjectValueFromFormValue(dataList form.Values) dialect.H {

	var (
		value        = make(dialect.H)
		exceptString = make([]string, 0)

		columns, auto = tb.getColumns(tb.form.Table)

		fun types.PostFieldFilterFn
	)

	if auto {
		exceptString = []string{tb.primaryKey.Name, "_previous_", "_method", "_t"}
	} else {
		exceptString = []string{"_previous_", "_method", "_t"}
	}

	if !dataList.IsSingleUpdatePost() {
		for _, field := range tb.form.FieldList {
			if field.FormType.IsMultiSelect() {
				if _, ok := dataList[field.Field+"[]"]; !ok {
					dataList[field.Field+"[]"] = []string{""}
				}
			}
		}
	}

	dataList = dataList.RemoveRemark()

	for k, v := range dataList {
		k = strings.Replace(k, "[]", "", -1)
		if !modules.InArray(exceptString, k) {
			if modules.InArray(columns, k) {
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

func GetNewFormList(groupHeaders []string,
	group [][]string,
	old []types.FormField) ([]types.FormField, [][]types.FormField, []string) {

	if len(group) == 0 {
		var newForm []types.FormField
		for _, v := range old {
			if !v.NotAllowAdd {
				v.Editable = true
				newForm = append(newForm, v.UpdateDefaultValue())
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
					if !v.NotAllowAdd {
						v.Editable = true
						list = append(list, v.UpdateDefaultValue())
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

func (tb DefaultTable) delete(table, key, id string) {
	_ = tb.sql().Table(table).
		Where(key, "=", id).
		Delete()
}

func (tb DefaultTable) getTheadAndFilterForm(params parameter.Parameters, columns Columns) ([]map[string]string,
	string, string, []string, []types.FormField) {
	return tb.info.FieldList.GetTheadAndFilterForm(types.TableInfo{
		Table:      tb.info.Table,
		Delimiter:  tb.db().GetDelimiter(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.primaryKey.Name,
	}, params, columns)
}

// db is a helper function return raw db connection.
func (tb DefaultTable) db() db.Connection {
	return db.GetConnectionFromService(services.Get(tb.connectionDriver))
}

// sql is a helper function return db sql.
func (tb DefaultTable) sql() *db.SQL {
	return db.WithDriverAndConnection(tb.connection, db.GetConnectionFromService(services.Get(tb.connectionDriver)))
}

type Columns []string

func (tb DefaultTable) getColumns(table string) (Columns, bool) {

	columnsModel, _ := tb.sql().Table(table).ShowColumns()

	columns := make(Columns, len(columnsModel))
	switch tb.connectionDriver {
	case "postgresql":
		auto := false
		for key, model := range columnsModel {
			columns[key] = model["column_name"].(string)
			if columns[key] == tb.primaryKey.Name {
				if v, ok := model["column_default"].(string); ok {
					if strings.Contains(v, "nextval") {
						auto = true
					}
				}
			}
		}
		return columns, auto
	case "mysql":
		auto := false
		for key, model := range columnsModel {
			columns[key] = model["Field"].(string)
			if columns[key] == tb.primaryKey.Name {
				if v, ok := model["Extra"].(string); ok {
					if v == "auto_increment" {
						auto = true
					}
				}
			}
		}
		return columns, auto
	case "sqlite":
		for key, model := range columnsModel {
			columns[key] = string(model["name"].(string))
		}

		num, _ := tb.sql().Table("sqlite_sequence").
			Where("name", "=", tb.GetForm().Table).Count()

		return columns, num > 0
	case "mssql":
		for key, model := range columnsModel {
			columns[key] = string(model["column_name"].(string))
		}
		return columns, true
	default:
		panic("wrong driver")
	}
}
