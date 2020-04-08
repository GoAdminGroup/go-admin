package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	errs "github.com/GoAdminGroup/go-admin/modules/errors"
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

type GetDataFun func(params parameter.Parameters) ([]map[string]interface{}, int)

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
		data, size = tb.getDataFun(params)
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(params)
	} else if tb.Info.GetDataFn != nil {
		data, size = tb.Info.GetDataFn(params)
	} else if params.IsAll() {
		return tb.getAllDataFromDatabase(params)
	} else {
		return tb.getDataFromDatabase(params)
	}

	infoList := make(types.InfoList, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

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

func (tb DefaultTable) getDataFromURL(params parameter.Parameters) ([]map[string]interface{}, int) {

	u := ""
	if strings.Contains(tb.sourceURL, "?") {
		u = tb.sourceURL + "&" + params.Join()
	} else {
		u = tb.sourceURL + "?" + params.Join()
	}
	res, err := http.Get(u + "&pk=" + strings.Join(params.PKs(), ","))

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
		data, size = tb.getDataFun(params)
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(params)
	} else if tb.Info.GetDataFn != nil {
		data, size = tb.Info.GetDataFn(params)
	} else {
		return tb.getDataFromDatabase(params)
	}

	infoList := make([]map[string]types.InfoItem, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

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

		// TODO: ToDisplay some same logic execute repeatedly, it can be improved.
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

func (tb DefaultTable) getAllDataFromDatabase(params parameter.Parameters) (PanelInfo, error) {
	var (
		connection     = tb.db()
		queryStatement = "select %s from %s %s %s %s order by " + modules.Delimiter(connection.GetDelimiter(), "%s") + " %s"
	)

	columns, _ := tb.getColumns(tb.Info.Table)

	thead, fields, joins := tb.Info.FieldList.GetThead(types.TableInfo{
		Table:      tb.Info.Table,
		Delimiter:  tb.db().GetDelimiter(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.PrimaryKey.Name,
	}, params, columns)

	fields += tb.Info.Table + "." + modules.FilterField(tb.PrimaryKey.Name, connection.GetDelimiter())

	groupBy := ""
	if joins != "" {
		groupBy = " GROUP BY " + tb.Info.Table + "." + modules.Delimiter(connection.GetDelimiter(), tb.PrimaryKey.Name)
	}

	var (
		wheres    = ""
		whereArgs = make([]interface{}, 0)
		existKeys = make([]string, 0)
	)

	wheres, whereArgs, existKeys = params.Statement(wheres, tb.Info.Table, connection.GetDelimiter(), whereArgs, columns, existKeys,
		tb.Info.FieldList.GetFieldFilterProcessValue, tb.Info.FieldList.GetFieldJoinTable)
	wheres, whereArgs = tb.Info.Wheres.Statement(wheres, connection.GetDelimiter(), whereArgs, existKeys, columns)
	wheres, whereArgs = tb.Info.WhereRaws.Statement(wheres, whereArgs)

	if wheres != "" {
		wheres = " where " + wheres
	}

	if !modules.InArray(columns, params.SortField) {
		params.SortField = tb.PrimaryKey.Name
	}

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.Info.Table, joins, wheres, groupBy, params.SortField, params.SortType)

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

// TODO: refactor
func (tb DefaultTable) getDataFromDatabase(params parameter.Parameters) (PanelInfo, error) {

	var (
		connection     = tb.db()
		placeholder    = modules.Delimiter(connection.GetDelimiter(), "%s")
		queryStatement string
		countStatement string
		ids            = params.PKs()
		pk             = tb.Info.Table + "." + modules.Delimiter(connection.GetDelimiter(), tb.PrimaryKey.Name)
	)

	if connection.Name() == db.DriverPostgresql {
		placeholder = "%s"
	}

	beginTime := time.Now()

	if len(ids) > 0 {
		countExtra := ""
		if connection.Name() == db.DriverMssql {
			countExtra = "as [size]"
		}
		// %s means: fields, table, join table, pk values, group by, order by field,  order by type
		queryStatement = "select %s from " + placeholder + " %s where " + pk + " in (%s) %s ORDER BY %s." + placeholder + " %s"
		// %s means: table, join table, pk values
		countStatement = "select count(*) " + countExtra + " from " + placeholder + " %s where " + pk + " in (%s)"
	} else {
		if connection.Name() == db.DriverMssql {
			// %s means: order by field, order by type, fields, table, join table, wheres, group by
			queryStatement = "SELECT * FROM (SELECT ROW_NUMBER() OVER (ORDER BY %s." + placeholder + " %s) as ROWNUMBER_, %s from " +
				placeholder + "%s %s %s ) as TMP_ WHERE TMP_.ROWNUMBER_ > ? AND TMP_.ROWNUMBER_ <= ?"
			// %s means: table, join table, wheres
			countStatement = "select count(*) as [size] from " + placeholder + " %s %s"
		} else {
			// %s means: fields, table, join table, wheres, group by, order by field, order by type
			queryStatement = "select %s from " + placeholder + "%s %s %s order by %s." + placeholder + " %s LIMIT ? OFFSET ?"
			// %s means: table, join table, wheres
			countStatement = "select count(*) from " + placeholder + " %s %s"
		}
	}

	columns, _ := tb.getColumns(tb.Info.Table)

	thead, fields, joinFields, joins, joinTables, filterForm := tb.getTheadAndFilterForm(params, columns)

	fields += pk

	allFields := fields
	groupFields := fields

	if joinFields != "" {
		allFields += "," + joinFields[:len(joinFields)-1]
		if connection.Name() == db.DriverMssql {
			for _, field := range tb.Info.FieldList {
				if field.TypeName == db.Text || field.TypeName == db.Longtext {
					f := modules.Delimiter(connection.GetDelimiter(), field.Field)
					headField := tb.Info.Table + "." + f
					allFields = strings.Replace(allFields, headField, "CAST("+headField+" AS NVARCHAR(MAX)) as "+f, -1)
					groupFields = strings.Replace(groupFields, headField, "CAST("+headField+" AS NVARCHAR(MAX))", -1)
				}
			}
		}
	}

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
				wheres += "?,"
				args = append(args, value)
			}
		}
		wheres = wheres[:len(wheres)-1]
	} else {

		// parameter
		wheres, whereArgs, existKeys = params.Statement(wheres, tb.Info.Table, connection.GetDelimiter(), whereArgs, columns, existKeys,
			tb.Info.FieldList.GetFieldFilterProcessValue, tb.Info.FieldList.GetFieldJoinTable)
		// pre query
		wheres, whereArgs = tb.Info.Wheres.Statement(wheres, connection.GetDelimiter(), whereArgs, existKeys, columns)
		wheres, whereArgs = tb.Info.WhereRaws.Statement(wheres, whereArgs)

		if wheres != "" {
			wheres = " where " + wheres
		}

		if connection.Name() == db.DriverMssql {
			args = append(whereArgs, (params.PageInt-1)*params.PageSizeInt, params.PageInt*params.PageSizeInt)
		} else {
			args = append(whereArgs, params.PageSize, (params.PageInt-1)*params.PageSizeInt)
		}
	}

	groupBy := ""
	if len(joinTables) > 0 {
		if connection.Name() == db.DriverMssql {
			groupBy = " GROUP BY " + groupFields
		} else {
			groupBy = " GROUP BY " + pk
		}
	}

	queryCmd := ""
	if connection.Name() == db.DriverMssql && len(ids) == 0 {
		queryCmd = fmt.Sprintf(queryStatement, tb.Info.Table, params.SortField, params.SortType,
			allFields, tb.Info.Table, joins, wheres, groupBy)
	} else {
		queryCmd = fmt.Sprintf(queryStatement, allFields, tb.Info.Table, joins, wheres, groupBy,
			tb.Info.Table, params.SortField, params.SortType)
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
	var size int

	if len(ids) == 0 {
		countCmd := fmt.Sprintf(countStatement, tb.Info.Table, joins, wheres)

		total, err := connection.QueryWithConnection(tb.connection, countCmd, whereArgs...)

		if err != nil {
			return PanelInfo{}, err
		}

		logger.LogSQL(countCmd, nil)

		if tb.connectionDriver == "postgresql" {
			size = int(total[0]["count"].(int64))
		} else if tb.connectionDriver == db.DriverMssql {
			size = int(total[0]["size"].(int64))
		} else {
			size = int(total[0]["count(*)"].(int64))
		}
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

func getDataRes(list []map[string]interface{}, _ int) map[string]interface{} {
	if len(list) > 0 {
		return list[0]
	}
	return nil
}

// GetDataWithId query the single row of data.
func (tb DefaultTable) GetDataWithId(param parameter.Parameters) (FormInfo, error) {

	var (
		res     map[string]interface{}
		columns Columns
		custom  = tb.getDataFun != nil || tb.sourceURL != "" || tb.Info.GetDataFn != nil
		id      = param.PK()
	)

	if tb.getDataFun != nil {
		res = getDataRes(tb.getDataFun(param))
	} else if tb.sourceURL != "" {
		res = getDataRes(tb.getDataFromURL(param))
	} else if tb.Detail.GetDataFn != nil {
		res = getDataRes(tb.Detail.GetDataFn(param))
	} else if tb.Info.GetDataFn != nil {
		res = getDataRes(tb.Info.GetDataFn(param))
	} else {

		columns, _ = tb.getColumns(tb.Form.Table)

		var (
			fields, joinFields, joins, groupBy = "", "", "", ""

			err            error
			joinTables     = make([]string, 0)
			args           = []interface{}{id}
			connection     = tb.db()
			delimiter      = connection.GetDelimiter()
			tableName      = tb.GetForm().Table
			pk             = tableName + "." + modules.Delimiter(delimiter, tb.PrimaryKey.Name)
			queryStatement = "select %s from " + modules.Delimiter(delimiter, "%s") + " %s where " + pk + " = ? %s "
		)

		if connection.Name() == db.DriverPostgresql {
			queryStatement = "select %s from %s %s where " + pk + " = ? %s "
		}

		for _, field := range tb.Form.FieldList {

			if field.Field != pk && modules.InArray(columns, field.Field) &&
				!field.Join.Valid() {
				fields += tableName + "." + modules.FilterField(field.Field, delimiter) + ","
			}

			headField := field.Field

			if field.Join.Valid() {
				headField = field.Join.Table + parameter.FilterParamJoinInfix + field.Field
				joinFields += db.GetAggregationExpression(connection.Name(), field.Join.Table+"."+
					modules.FilterField(field.Field, delimiter), headField, types.JoinFieldValueDelimiter) + ","
				if !modules.InArray(joinTables, field.Join.Table) {
					joinTables = append(joinTables, field.Join.Table)
					joins += " left join " + modules.FilterField(field.Join.Table, delimiter) + " on " +
						field.Join.Table + "." + modules.FilterField(field.Join.JoinField, delimiter) + " = " +
						tableName + "." + modules.FilterField(field.Join.Field, delimiter)
				}
			}
		}

		fields += pk
		groupFields := fields

		if joinFields != "" {
			fields += "," + joinFields[:len(joinFields)-1]
			if connection.Name() == db.DriverMssql {
				for _, field := range tb.Form.FieldList {
					if field.TypeName == db.Text || field.TypeName == db.Longtext {
						f := modules.Delimiter(connection.GetDelimiter(), field.Field)
						headField := tb.Info.Table + "." + f
						fields = strings.Replace(fields, headField, "CAST("+headField+" AS NVARCHAR(MAX)) as "+f, -1)
						groupFields = strings.Replace(groupFields, headField, "CAST("+headField+" AS NVARCHAR(MAX))", -1)
					}
				}
			}
		}

		if len(joinTables) > 0 {
			if connection.Name() == db.DriverMssql {
				groupBy = " GROUP BY " + groupFields
			} else {
				groupBy = " GROUP BY " + pk
			}
		}

		queryCmd := fmt.Sprintf(queryStatement, fields, tableName, joins, groupBy)

		logger.LogSQL(queryCmd, args)

		result, err := connection.QueryWithConnection(tb.connection, queryCmd, args...)

		if err != nil {
			return FormInfo{Title: tb.Form.Title, Description: tb.Form.Description}, err
		}

		if len(result) == 0 {
			return FormInfo{Title: tb.Form.Title, Description: tb.Form.Description}, errors.New(errs.WrongID)
		}

		res = result[0]
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

	var (
		errMsg = ""
		err    error
	)

	if tb.Form.PostHook != nil {
		defer func() {
			dataList.Add(form.PostTypeKey, "0")
			dataList.Add(form.PostResultKey, errMsg)
			go func() {
				defer func() {
					if err := recover(); err != nil {
						logger.Error(err)
					}
				}()

				err := tb.Form.PostHook(dataList)
				if err != nil {
					logger.Error(err)
				}
			}()
		}()
	}

	if tb.Form.Validator != nil {
		if err := tb.Form.Validator(dataList); err != nil {
			errMsg = "post error: " + err.Error()
			return err
		}
	}

	if tb.Form.PreProcessFn != nil {
		dataList = tb.Form.PreProcessFn(dataList)
	}

	if tb.Form.UpdateFn != nil {
		dataList.Delete(form.PostTypeKey)
		err = tb.Form.UpdateFn(dataList)
		if err != nil {
			errMsg = "post error: " + err.Error()
		}
		return err
	}

	_, err = tb.sql().Table(tb.Form.Table).
		Where(tb.PrimaryKey.Name, "=", dataList.Get(tb.PrimaryKey.Name)).
		Update(tb.getInjectValueFromFormValue(dataList))

	// NOTE: some errors should be ignored.
	if db.CheckError(err, db.UPDATE) {
		if err != nil {
			errMsg = "post error: " + err.Error()
		}
		return err
	}

	return nil
}

// InsertData insert data.
func (tb DefaultTable) InsertData(dataList form.Values) error {

	dataList.Add(form.PostTypeKey, "1")

	var (
		id     = int64(0)
		err    error
		errMsg = ""
	)

	if tb.Form.PostHook != nil {
		defer func() {
			dataList.Add(form.PostTypeKey, "1")
			dataList.Add(tb.GetPrimaryKey().Name, strconv.Itoa(int(id)))
			dataList.Add(form.PostResultKey, errMsg)

			go func() {
				defer func() {
					if err := recover(); err != nil {
						logger.Error(err)
					}
				}()

				err := tb.Form.PostHook(dataList)
				if err != nil {
					logger.Error(err)
				}
			}()
		}()
	}

	if tb.Form.Validator != nil {
		if err := tb.Form.Validator(dataList); err != nil {
			errMsg = "post error: " + err.Error()
			return err
		}
	}

	if tb.Form.PreProcessFn != nil {
		dataList = tb.Form.PreProcessFn(dataList)
	}

	if tb.Form.InsertFn != nil {
		dataList.Delete(form.PostTypeKey)
		err = tb.Form.InsertFn(dataList)
		if err != nil {
			errMsg = "post error: " + err.Error()
		}
		return err
	}

	id, err = tb.sql().Table(tb.Form.Table).Insert(tb.getInjectValueFromFormValue(dataList))

	// NOTE: some errors should be ignored.
	if db.CheckError(err, db.INSERT) {
		errMsg = "post error: " + err.Error()
		return err
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

	var (
		idArr = strings.Split(id, ",")
		err   error
	)

	if tb.Info.DeleteHook != nil {
		defer func() {
			go func() {
				defer func() {
					if recoverErr := recover(); recoverErr != nil {
						logger.Error(recoverErr)
					}
				}()

				if hookErr := tb.Info.DeleteHook(idArr); hookErr != nil {
					logger.Error(hookErr)
				}
			}()
		}()
	}

	if tb.Info.DeleteHookWithRes != nil {
		defer func() {
			go func() {
				defer func() {
					if recoverErr := recover(); recoverErr != nil {
						logger.Error(recoverErr)
					}
				}()

				if hookErr := tb.Info.DeleteHookWithRes(idArr, err); hookErr != nil {
					logger.Error(hookErr)
				}
			}()
		}()
	}

	if tb.Info.PreDeleteFn != nil {
		if err = tb.Info.PreDeleteFn(idArr); err != nil {
			return err
		}
	}

	if tb.Info.DeleteFn != nil {
		err = tb.Info.DeleteFn(idArr)
		return err
	}

	if len(idArr) == 0 || tb.Info.Table == "" {
		err = errors.New("delete error: wrong parameter")
		return err
	}

	err = tb.delete(tb.Info.Table, tb.PrimaryKey.Name, idArr)
	return err
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

func (tb DefaultTable) delete(table, key string, values []string) error {

	var vals = make([]interface{}, len(values))
	for i, v := range values {
		vals[i] = v
	}

	return tb.sql().Table(table).
		WhereIn(key, vals).
		Delete()
}

func (tb DefaultTable) getTheadAndFilterForm(params parameter.Parameters, columns Columns) (types.Thead,
	string, string, string, []string, []types.FormField) {

	return tb.Info.FieldList.GetTheadAndFilterForm(types.TableInfo{
		Table:      tb.Info.Table,
		Delimiter:  tb.delimiter(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.PrimaryKey.Name,
	}, params, columns, func() *db.SQL {
		return tb.sql()
	})
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
	case db.DriverPostgresql:
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
	case db.DriverMysql:
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
	case db.DriverSqlite:
		for key, model := range columnsModel {
			columns[key] = string(model["name"].(string))
		}

		num, _ := tb.sql().Table("sqlite_sequence").
			Where("name", "=", tb.GetForm().Table).Count()

		return columns, num > 0
	case db.DriverMssql:
		for key, model := range columnsModel {
			columns[key] = string(model["column_name"].(string))
		}
		return columns, true
	default:
		panic("wrong driver")
	}
}
