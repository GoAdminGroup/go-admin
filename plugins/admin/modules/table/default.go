package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DefaultTable struct {
	*BaseTable
	connectionDriver string
	connection       string
	sourceURL        string
	getDataFun       GetDataFun
}

type GetDataFun func(path string, params parameter.Parameters, isAll bool, ids []string) ([]map[string]interface{}, int)

func NewDefaultTable(cfgs ...Config) Table {

	var cfg Config

	if len(cfgs) > 0 && cfgs[0].PrimaryKey.Name != "" {
		cfg = cfgs[0]
	} else {
		cfg = DefaultConfig()
	}

	return DefaultTable{
		BaseTable: &BaseTable{
			Info:       types.NewInfoPanel(cfg.PrimaryKey.Name),
			Form:       types.NewFormPanel(),
			Detail:     types.NewInfoPanel(cfg.PrimaryKey.Name),
			CanAdd:     cfg.CanAdd,
			Editable:   cfg.Editable,
			Deletable:  cfg.Deletable,
			Exportable: cfg.Exportable,
			PrimaryKey: cfg.PrimaryKey,
		},
		connectionDriver: cfg.Driver,
		connection:       cfg.Connection,
		sourceURL:        cfg.SourceURL,
		getDataFun:       cfg.GetDataFun,
	}
}

func (tb DefaultTable) Copy() Table {
	return DefaultTable{
		BaseTable: &BaseTable{
			Form: types.NewFormPanel().SetTable(tb.Form.Table).
				SetDescription(tb.Form.Description).
				SetTitle(tb.Form.Title),
			Info: types.NewInfoPanel(tb.PrimaryKey.Name).SetTable(tb.Info.Table).
				SetDescription(tb.Info.Description).
				SetTitle(tb.Info.Title).
				SetGetDataFn(tb.Info.GetDataFn),
			Detail: types.NewInfoPanel(tb.PrimaryKey.Name).SetTable(tb.Detail.Table).
				SetDescription(tb.Detail.Description).
				SetTitle(tb.Detail.Title).
				SetGetDataFn(tb.Detail.GetDataFn),
			CanAdd:     tb.CanAdd,
			Editable:   tb.Editable,
			Deletable:  tb.Deletable,
			Exportable: tb.Exportable,
			PrimaryKey: tb.PrimaryKey,
		},
		connectionDriver: tb.connectionDriver,
		connection:       tb.connection,
		sourceURL:        tb.sourceURL,
		getDataFun:       tb.getDataFun,
	}
}

// GetData query the data set.
func (tb DefaultTable) GetData(params parameter.Parameters) (PanelInfo, error) {

	var (
		data      []map[string]interface{}
		size      int
		beginTime = time.Now()
	)

	if tb.getDataFun != nil {
		data, size = tb.getDataFun(params.URLPath, params, params.IsAll(), []string{})
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(params.URLPath, params, params.IsAll(), []string{})
	} else if tb.Info.GetDataFn != nil {
		data, size = tb.Info.GetDataFn(params)
	} else if params.IsAll() {
		return tb.getAllDataFromDatabase(params.URLPath, params)
	} else {
		return tb.getDataFromDatabase(params.URLPath, params, []string{})
	}

	infoList := make(types.InfoList, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

	endTime := time.Now()

	return PanelInfo{
		Thead:    thead,
		InfoList: infoList,
		Paginator: paginator.Get(paginator.Config{
			Size:         size,
			Param:        params,
			PageSizeList: tb.Info.GetPageSizeList(),
		}).SetExtraInfo(template.HTML(fmt.Sprintf("<b>" + language.Get("query time") + ": </b>" +
			fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000)))),
		Title:          tb.Info.Title,
		FilterFormData: filterForm,
		Description:    tb.Info.Description,
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
func (tb DefaultTable) GetDataWithIds(params parameter.Parameters) (PanelInfo, error) {

	var (
		data      []map[string]interface{}
		size      int
		beginTime = time.Now()
	)

	if tb.getDataFun != nil {
		data, size = tb.getDataFun(params.URLPath, params, false, params.PKs())
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(params.URLPath, params, false, params.PKs())
	} else if tb.Info.GetDataFn != nil {
		data, size = tb.Info.GetDataFn(params)
	} else {
		return tb.getDataFromDatabase(params.URLPath, params, params.PKs())
	}

	infoList := make([]map[string]types.InfoItem, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

	endTime := time.Now()

	return PanelInfo{
		Thead:    thead,
		InfoList: infoList,
		Paginator: paginator.Get(paginator.Config{
			Size:         size,
			Param:        params,
			PageSizeList: tb.Info.GetPageSizeList(),
		}).
			SetExtraInfo(template.HTML(fmt.Sprintf("<b>" + language.Get("query time") + ": </b>" +
				fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000)))),
		Title:          tb.Info.Title,
		FilterFormData: filterForm,
		Description:    tb.Info.Description,
	}, nil
}

func (tb DefaultTable) getTempModelData(res map[string]interface{}, params parameter.Parameters, columns Columns) map[string]types.InfoItem {

	var tempModelData = make(map[string]types.InfoItem)
	headField := ""

	primaryKeyValue := db.GetValueFromDatabaseType(tb.PrimaryKey.Type, res[tb.PrimaryKey.Name], len(columns) == 0)

	for _, field := range tb.Info.FieldList {

		headField = field.Field

		if field.Join.Valid() {
			headField = field.Join.Table + parameter.FilterParamJoinInfix + field.Field
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
			tempModelData[headField] = types.InfoItem{
				Content: template.HTML(valueStr),
				Value:   combineValue,
			}
		} else {
			tempModelData[headField] = types.InfoItem{
				Content: value.(template.HTML),
				Value:   combineValue,
			}
		}
	}

	tempModelData[tb.PrimaryKey.Name] = types.InfoItem{
		Content: template.HTML(primaryKeyValue.String()),
		Value:   primaryKeyValue.String(),
	}
	return tempModelData
}

func (tb DefaultTable) getAllDataFromDatabase(path string, params parameter.Parameters) (PanelInfo, error) {
	var (
		connection     = tb.db()
		queryStatement = "select %s from %s %s %s order by " + modules.Delimiter(connection.GetDelimiter(), "%s") + " %s"
	)

	columns, _ := tb.getColumns(tb.Info.Table)

	thead, fields, joins := tb.Info.FieldList.GetThead(types.TableInfo{
		Table:      tb.Info.Table,
		Delimiter:  tb.db().GetDelimiter(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.PrimaryKey.Name,
	}, params, columns)

	fields += tb.Info.Table + "." + modules.FilterField(tb.PrimaryKey.Name, connection.GetDelimiter())

	var (
		wheres    = ""
		whereArgs = make([]interface{}, 0)
		existKeys = make([]string, 0)
	)

	wheres, whereArgs, existKeys = params.Statement(wheres, connection.GetDelimiter(), whereArgs, columns, existKeys,
		tb.Info.FieldList.GetFieldFilterProcessValue, tb.Info.FieldList.GetFieldJoinTable)
	wheres, whereArgs = tb.Info.Wheres.Statement(wheres, connection.GetDelimiter(), whereArgs, existKeys, columns)
	wheres, whereArgs = tb.Info.WhereRaws.Statement(wheres, whereArgs)

	if wheres != "" {
		wheres = " where " + wheres
	}

	if !modules.InArray(columns, params.SortField) {
		params.SortField = tb.PrimaryKey.Name
	}

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.Info.Table, joins, wheres, params.SortField, params.SortType)

	logger.LogSQL(queryCmd, []interface{}{})

	res, err := connection.QueryWithConnection(tb.connection, queryCmd, whereArgs...)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]types.InfoItem, 0)

	for i := 0; i < len(res); i++ {
		infoList = append(infoList, tb.getTempModelData(res[i], params, columns))
	}

	return PanelInfo{
		InfoList:    infoList,
		Thead:       thead,
		Title:       tb.Info.Title,
		Description: tb.Info.Description,
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
		queryStatement = "select %s from %s %s where " + tb.PrimaryKey.Name + " in (%s) %s order by " + placeholder + " %s"
		countStatement = "select count(*) from " + placeholder + " %s where " + tb.PrimaryKey.Name + " in (%s)"
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

	columns, _ := tb.getColumns(tb.Info.Table)

	thead, fields, joins, joinTables, filterForm := tb.getTheadAndFilterForm(params, columns)

	fields += tb.Info.Table + "." + modules.FilterField(tb.PrimaryKey.Name, connection.GetDelimiter())

	if !modules.InArray(columns, params.SortField) {
		params.SortField = tb.PrimaryKey.Name
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

		// parameter
		wheres, whereArgs, existKeys = params.Statement(wheres, connection.GetDelimiter(), whereArgs, columns, existKeys,
			tb.Info.FieldList.GetFieldFilterProcessValue, tb.Info.FieldList.GetFieldJoinTable)
		// pre query
		wheres, whereArgs = tb.Info.Wheres.Statement(wheres, connection.GetDelimiter(), whereArgs, existKeys, columns)
		wheres, whereArgs = tb.Info.WhereRaws.Statement(wheres, whereArgs)

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
		groupBy = " GROUP BY " + tb.Info.Table + "." + modules.FilterField(tb.GetPrimaryKey().Name, connection.GetDelimiter())
	}

	queryCmd := ""
	if connection.Name() == "mssql" {
		queryCmd = fmt.Sprintf(queryStatement, params.SortField, params.SortType, fields, tb.Info.Table, joins, wheres, groupBy)
	} else {
		queryCmd = fmt.Sprintf(queryStatement, fields, tb.Info.Table, joins, wheres, groupBy, params.SortField, params.SortType)
	}

	logger.LogSQL(queryCmd, args)

	res, err := connection.QueryWithConnection(tb.connection, queryCmd, args...)

	if err != nil {
		return PanelInfo{}, err
	}

	infoList := make([]map[string]types.InfoItem, 0)

	for i := 0; i < len(res); i++ {
		infoList = append(infoList, tb.getTempModelData(res[i], params, columns))
	}

	// TODO: use the dialect

	if len(ids) > 0 {
		joins = ""
	}

	countCmd := fmt.Sprintf(countStatement, tb.Info.Table, joins, wheres)

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
		Paginator: tb.GetPaginator(size, params,
			template.HTML(fmt.Sprintf("<b>"+language.Get("query time")+": </b>"+
				fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000)))),
		Title:          tb.Info.Title,
		FilterFormData: filterForm,
		Description:    tb.Info.Description,
	}, nil
}

// GetDataWithId query the single row of data.
func (tb DefaultTable) GetDataWithId(id string) (FormInfo, error) {

	var (
		res     map[string]interface{}
		columns Columns
		custom  = false
	)

	if tb.getDataFun != nil {
		list, _ := tb.getDataFun("", parameter.BaseParam(), false, []string{id})
		if len(list) > 0 {
			res = list[0]
		}
		custom = true
	} else if tb.sourceURL != "" {
		list, _ := tb.getDataFromURL("", parameter.BaseParam(), false, []string{id})
		if len(list) > 0 {
			res = list[0]
		}
		custom = true
	} else if tb.Detail.GetDataFn != nil {
		list, _ := tb.Detail.GetDataFn(parameter.BaseParam().WithPKs(id))
		if len(list) > 0 {
			res = list[0]
		}
		custom = true
	} else if tb.Info.GetDataFn != nil {
		list, _ := tb.Info.GetDataFn(parameter.BaseParam().WithPKs(id))
		if len(list) > 0 {
			res = list[0]
		}
		custom = true
	} else {

		columns, _ = tb.getColumns(tb.Form.Table)

		var (
			err    error
			fields = make([]string, 0)
		)

		for i := 0; i < len(tb.Form.FieldList); i++ {
			if modules.InArray(columns, tb.Form.FieldList[i].Field) {
				fields = append(fields, tb.Form.FieldList[i].Field)
			}
		}

		res, err = tb.sql().
			Table(tb.Form.Table).Select(fields...).
			Where(tb.PrimaryKey.Name, "=", id).
			First()

		if err != nil {
			return FormInfo{Title: tb.Form.Title, Description: tb.Form.Description}, err
		}
	}

	var (
		groupFormList = make([]types.FormFields, 0)
		groupHeaders  = make([]string, 0)
	)

	if len(tb.Form.TabGroups) > 0 {
		if custom {
			groupFormList, groupHeaders = tb.Form.GroupFieldWithValue(id, columns, res)
		} else {
			groupFormList, groupHeaders = tb.Form.GroupFieldWithValue(id, columns, res, tb.sql)
		}
		return FormInfo{
			FieldList:         tb.Form.FieldList,
			GroupFieldList:    groupFormList,
			GroupFieldHeaders: groupHeaders,
			Title:             tb.Form.Title,
			Description:       tb.Form.Description,
		}, nil
	}

	var fieldList types.FormFields
	if custom {
		fieldList = tb.Form.FieldsWithValue(id, columns, res)
	} else {
		fieldList = tb.Form.FieldsWithValue(id, columns, res, tb.sql)
	}

	return FormInfo{
		FieldList:         fieldList.FillCustomContent(),
		GroupFieldList:    groupFormList,
		GroupFieldHeaders: groupHeaders,
		Title:             tb.Form.Title,
		Description:       tb.Form.Description,
	}, nil
}

// UpdateData update data.
func (tb DefaultTable) UpdateData(dataList form.Values) error {

	dataList.Add(form.PostTypeKey, "0")

	if tb.Form.Validator != nil {
		if err := tb.Form.Validator(dataList); err != nil {
			return err
		}
	}

	if tb.Form.UpdateFn != nil {
		dataList.Delete(form.PostTypeKey)
		return tb.Form.UpdateFn(dataList)
	}

	if tb.Form.PreProcessFn != nil {
		dataList = tb.Form.PreProcessFn(dataList)
	}

	_, err := tb.sql().Table(tb.Form.Table).
		Where(tb.PrimaryKey.Name, "=", dataList.Get(tb.PrimaryKey.Name)).
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

	if tb.Form.PostHook != nil {
		go func() {

			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			dataList.Add(form.PostTypeKey, "0")

			err := tb.Form.PostHook(dataList)
			if err != nil {
				logger.Error(err)
			}
		}()
	}

	return nil
}

// InsertData insert data.
func (tb DefaultTable) InsertData(dataList form.Values) error {

	dataList.Add(form.PostTypeKey, "1")

	if tb.Form.Validator != nil {
		if err := tb.Form.Validator(dataList); err != nil {
			return err
		}
	}

	if tb.Form.InsertFn != nil {
		dataList.Delete(form.PostTypeKey)
		return tb.Form.InsertFn(dataList)
	}

	if tb.Form.PreProcessFn != nil {
		dataList = tb.Form.PreProcessFn(dataList)
	}

	id, err := tb.sql().Table(tb.Form.Table).Insert(tb.getInjectValueFromFormValue(dataList))

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

	if tb.Form.PostHook != nil {
		go func() {

			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			dataList.Add(form.PostTypeKey, "1")

			err := tb.Form.PostHook(dataList)
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

		columns, auto = tb.getColumns(tb.Form.Table)

		fun types.PostFieldFilterFn
	)

	if auto {
		exceptString = []string{tb.PrimaryKey.Name, form.PreviousKey, form.MethodKey, form.TokenKey}
	} else {
		exceptString = []string{form.PreviousKey, form.MethodKey, form.TokenKey}
	}

	if !dataList.IsSingleUpdatePost() {
		for _, field := range tb.Form.FieldList {
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
				for i := 0; i < len(tb.Form.FieldList); i++ {
					if k == tb.Form.FieldList[i].Field {
						fun = tb.Form.FieldList[i].PostFilterFn
						delimiter = modules.SetDefault(tb.Form.FieldList[i].DefaultOptionDelimiter, ",")
					}
				}
				vv := modules.RemoveBlankFromArray(v)
				if fun != nil {
					value[k] = fun(types.PostFieldModel{
						ID:    dataList.Get(tb.PrimaryKey.Name),
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
				fun := tb.Form.FieldList.FindByFieldName(k).PostFilterFn
				if fun != nil {
					fun(types.PostFieldModel{
						ID:    dataList.Get(tb.PrimaryKey.Name),
						Value: modules.RemoveBlankFromArray(v),
					})
				}
			}
		}
	}
	return value
}

// DeleteData delete data.
func (tb DefaultTable) DeleteData(id string) error {
	idArr := strings.Split(id, ",")

	if tb.Info.DeleteFn != nil {

		if len(idArr) == 0 {
			return errors.New("wrong parameter")
		}

		return tb.Info.DeleteFn(idArr)
	}

	if tb.Info.PreDeleteFn != nil && len(idArr) > 0 {
		if err := tb.Info.PreDeleteFn(idArr); err != nil {
			return err
		}
	}

	// TODO: use where in
	for _, id := range idArr {
		tb.delete(tb.Form.Table, tb.PrimaryKey.Name, id)
	}

	if tb.Info.DeleteHook != nil && len(idArr) > 0 {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(err)
				}
			}()

			if err := tb.Info.DeleteHook(idArr); err != nil {
				logger.Error(err)
			}
		}()
	}

	return nil
}

func (tb DefaultTable) GetNewForm() FormInfo {

	if len(tb.Form.TabGroups) == 0 {
		return FormInfo{FieldList: tb.Form.FieldsWithDefaultValue(tb.sql).FillCustomContent()}
	}

	newForm, headers := tb.Form.GroupField(tb.sql)

	return FormInfo{GroupFieldList: newForm, GroupFieldHeaders: headers}
}

// ***************************************
// helper function for database operation
// ***************************************

func (tb DefaultTable) delete(table, key, id string) {
	_ = tb.sql().Table(table).
		Where(key, "=", id).
		Delete()
}

func (tb DefaultTable) getTheadAndFilterForm(params parameter.Parameters, columns Columns) (types.Thead,
	string, string, []string, []types.FormField) {
	return tb.Info.FieldList.GetTheadAndFilterForm(types.TableInfo{
		Table:      tb.Info.Table,
		Delimiter:  tb.delimiter(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.PrimaryKey.Name,
	}, params, columns)
}

// db is a helper function return raw db connection.
func (tb DefaultTable) db() db.Connection {
	if tb.connectionDriver != "" && tb.getDataFromDB() {
		return db.GetConnectionFromService(services.Get(tb.connectionDriver))
	}
	return nil
}

func (tb DefaultTable) delimiter() string {
	if tb.getDataFromDB() {
		return tb.db().GetDelimiter()
	}
	return ""
}

func (tb DefaultTable) getDataFromDB() bool {
	return tb.sourceURL == "" && tb.getDataFun == nil && tb.Info.GetDataFn == nil && tb.Detail.GetDataFn == nil
}

// sql is a helper function return db sql.
func (tb DefaultTable) sql() *db.SQL {
	if tb.connectionDriver != "" && tb.getDataFromDB() {
		return db.WithDriverAndConnection(tb.connection, db.GetConnectionFromService(services.Get(tb.connectionDriver)))
	}
	return nil
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
			if columns[key] == tb.PrimaryKey.Name {
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
			if columns[key] == tb.PrimaryKey.Name {
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
