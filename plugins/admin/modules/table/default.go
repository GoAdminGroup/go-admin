package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	errs "github.com/GoAdminGroup/go-admin/modules/errors"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
)

// DefaultTable is an implementation of table.Table
type DefaultTable struct {
	*BaseTable
	connectionDriver     string
	connectionDriverMode string
	connection           string
	sourceURL            string
	getDataFun           GetDataFun

	dbObj db.Connection
}

type GetDataFun func(params parameter.Parameters) ([]map[string]interface{}, int)

func NewDefaultTable(ctx *context.Context, cfgs ...Config) Table {

	var cfg Config

	if len(cfgs) > 0 && cfgs[0].PrimaryKey.Name != "" {
		cfg = cfgs[0]
	} else {
		cfg = DefaultConfig()
	}

	return &DefaultTable{
		BaseTable: &BaseTable{
			Info:           types.NewInfoPanel(ctx, cfg.PrimaryKey.Name),
			Form:           types.NewFormPanel(),
			NewForm:        types.NewFormPanel(),
			Detail:         types.NewInfoPanel(ctx, cfg.PrimaryKey.Name),
			CanAdd:         cfg.CanAdd,
			Editable:       cfg.Editable,
			Deletable:      cfg.Deletable,
			Exportable:     cfg.Exportable,
			PrimaryKey:     cfg.PrimaryKey,
			OnlyNewForm:    cfg.OnlyNewForm,
			OnlyUpdateForm: cfg.OnlyUpdateForm,
			OnlyDetail:     cfg.OnlyDetail,
			OnlyInfo:       cfg.OnlyInfo,
		},
		connectionDriver:     cfg.Driver,
		connectionDriverMode: cfg.DriverMode,
		connection:           cfg.Connection,
		sourceURL:            cfg.SourceURL,
		getDataFun:           cfg.GetDataFun,
	}
}

// Copy copy a new table.Table from origin DefaultTable
func (tb *DefaultTable) Copy() Table {
	return &DefaultTable{
		BaseTable: &BaseTable{
			Form: types.NewFormPanel().SetTable(tb.Form.Table).
				SetDescription(tb.Form.Description).
				SetTitle(tb.Form.Title),
			NewForm: types.NewFormPanel().SetTable(tb.Form.Table).
				SetDescription(tb.Form.Description).
				SetTitle(tb.Form.Title),
			Info: types.NewInfoPanel(tb.Info.Ctx, tb.PrimaryKey.Name).SetTable(tb.Info.Table).
				SetDescription(tb.Info.Description).
				SetTitle(tb.Info.Title).
				SetGetDataFn(tb.Info.GetDataFn),
			Detail: types.NewInfoPanel(tb.Info.Ctx, tb.PrimaryKey.Name).SetTable(tb.Detail.Table).
				SetDescription(tb.Detail.Description).
				SetTitle(tb.Detail.Title).
				SetGetDataFn(tb.Detail.GetDataFn),
			CanAdd:     tb.CanAdd,
			Editable:   tb.Editable,
			Deletable:  tb.Deletable,
			Exportable: tb.Exportable,
			PrimaryKey: tb.PrimaryKey,
		},
		connectionDriver:     tb.connectionDriver,
		connectionDriverMode: tb.connectionDriverMode,
		connection:           tb.connection,
		sourceURL:            tb.sourceURL,
		getDataFun:           tb.getDataFun,
	}
}

// GetData query the data set.
func (tb *DefaultTable) GetData(ctx *context.Context, params parameter.Parameters) (PanelInfo, error) {

	var (
		data      []map[string]interface{}
		size      int
		beginTime = time.Now()
	)

	if tb.Info.UpdateParametersFns != nil {
		for _, fn := range tb.Info.UpdateParametersFns {
			fn(&params)
		}
	}

	if tb.Info.QueryFilterFn != nil {
		var ids []string
		var stopQuery bool

		if tb.getDataFun == nil && tb.Info.GetDataFn == nil {
			ids, stopQuery = tb.Info.QueryFilterFn(params, tb.db())
		} else {
			ids, stopQuery = tb.Info.QueryFilterFn(params, nil)
		}

		if stopQuery {
			return tb.GetDataWithIds(ctx, params.WithPKs(ids...))
		}
	}

	if tb.getDataFun != nil {
		data, size = tb.getDataFun(params)
	} else if tb.sourceURL != "" {
		data, size = tb.getDataFromURL(params)
	} else if tb.Info.GetDataFn != nil {
		data, size = tb.Info.GetDataFn(params)
	} else if params.IsAll() {
		return tb.getAllDataFromDatabase(params)
	} else {
		return tb.getDataFromDatabase(ctx, params)
	}

	infoList := make(types.InfoList, 0)

	for i := 0; i < len(data); i++ {
		infoList = append(infoList, tb.getTempModelData(data[i], params, []string{}))
	}

	thead, _, _, _, _, filterForm := tb.getTheadAndFilterForm(params, []string{})

	endTime := time.Now()

	extraInfo := ""

	if !tb.Info.IsHideQueryInfo {
		extraInfo = fmt.Sprintf("<b>" + language.Get("query time") + ": </b>" +
			fmt.Sprintf("%.3fms", endTime.Sub(beginTime).Seconds()*1000))
	}

	return PanelInfo{
		Thead:    thead,
		InfoList: infoList,
		Paginator: paginator.Get(ctx, paginator.Config{
			Size:         size,
			Param:        params,
			PageSizeList: tb.Info.GetPageSizeList(),
		}).SetExtraInfo(template.HTML(extraInfo)),
		Title:          tb.Info.Title,
		FilterFormData: filterForm,
		Description:    tb.Info.Description,
	}, nil
}

type GetDataFromURLRes struct {
	Data []map[string]interface{}
	Size int
}

func (tb *DefaultTable) getDataFromURL(params parameter.Parameters) ([]map[string]interface{}, int) {

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
func (tb *DefaultTable) GetDataWithIds(ctx *context.Context, params parameter.Parameters) (PanelInfo, error) {

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
		return tb.getDataFromDatabase(ctx, params)
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
		Paginator: paginator.Get(ctx, paginator.Config{
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

func (tb *DefaultTable) getTempModelData(res map[string]interface{}, params parameter.Parameters, columns Columns) map[string]types.InfoItem {

	var tempModelData = map[string]types.InfoItem{
		"__goadmin_edit_params":   {},
		"__goadmin_delete_params": {},
		"__goadmin_detail_params": {},
	}
	headField := ""
	editParams := ""
	deleteParams := ""
	detailParams := ""

	primaryKeyValue := db.GetValueFromDatabaseType(tb.PrimaryKey.Type, res[tb.PrimaryKey.Name], len(columns) == 0)

	for _, field := range tb.Info.FieldList {

		headField = field.Field

		if field.Joins.Valid() {
			headField = field.Joins.Last().GetTableName() + parameter.FilterParamJoinInfix + field.Field
		}

		if field.Hide {
			continue
		}
		if !modules.InArrayWithoutEmpty(params.Columns, headField) {
			continue
		}

		typeName := field.TypeName

		if field.Joins.Valid() {
			typeName = db.Varchar
		}

		combineValue := db.GetValueFromDatabaseType(typeName, res[headField], len(columns) == 0).String()

		// TODO: ToDisplay some same logic execute repeatedly, it can be improved.
		var value interface{}
		if len(columns) == 0 || modules.InArray(columns, headField) || field.Joins.Valid() {
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
		var valueStr string
		var ok bool
		if valueStr, ok = value.(string); ok {
			tempModelData[headField] = types.InfoItem{
				Content: template.HTML(valueStr),
				Value:   combineValue,
			}
		} else {
			valueStr = string(value.(template.HTML))
			tempModelData[headField] = types.InfoItem{
				Content: value.(template.HTML),
				Value:   combineValue,
			}
		}

		if field.IsEditParam {
			editParams += "__goadmin_edit_" + field.Field + "=" + valueStr + "&"
		}
		if field.IsDeleteParam {
			deleteParams += "__goadmin_delete_" + field.Field + "=" + valueStr + "&"
		}
		if field.IsDetailParam {
			detailParams += "__goadmin_detail_" + field.Field + "=" + valueStr + "&"
		}
	}

	if editParams != "" {
		tempModelData["__goadmin_edit_params"] = types.InfoItem{Content: template.HTML("&" + editParams[:len(editParams)-1])}
	}
	if deleteParams != "" {
		tempModelData["__goadmin_delete_params"] = types.InfoItem{Content: template.HTML("&" + deleteParams[:len(deleteParams)-1])}
	}
	if detailParams != "" {
		tempModelData["__goadmin_detail_params"] = types.InfoItem{Content: template.HTML("&" + detailParams[:len(detailParams)-1])}
	}

	primaryKeyField := tb.Info.FieldList.GetFieldByFieldName(tb.PrimaryKey.Name)
	value := primaryKeyField.ToDisplay(types.FieldModel{
		ID:    primaryKeyValue.String(),
		Value: primaryKeyValue.String(),
		Row:   res,
	})
	if valueStr, ok := value.(string); ok {
		tempModelData[tb.PrimaryKey.Name] = types.InfoItem{
			Content: template.HTML(valueStr),
			Value:   primaryKeyValue.String(),
		}
	} else {
		tempModelData[tb.PrimaryKey.Name] = types.InfoItem{
			Content: value.(template.HTML),
			Value:   primaryKeyValue.String(),
		}
	}

	return tempModelData
}

func (tb *DefaultTable) getAllDataFromDatabase(params parameter.Parameters) (PanelInfo, error) {
	var (
		connection     = tb.db()
		queryStatement = "select %s from %s %s %s %s order by " + modules.Delimiter(connection.GetDelimiter(), connection.GetDelimiter2(), "%s") + " %s"
	)

	columns, _ := tb.getColumns(tb.Info.Table)

	thead, fields, joins := tb.Info.FieldList.GetThead(types.TableInfo{
		Table:      tb.Info.Table,
		Delimiter:  connection.GetDelimiter(),
		Delimiter2: connection.GetDelimiter2(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.PrimaryKey.Name,
	}, params, columns)

	fields += tb.Info.Table + "." + modules.FilterField(tb.PrimaryKey.Name, connection.GetDelimiter(), connection.GetDelimiter2())

	groupBy := ""
	if joins != "" {
		groupBy = " GROUP BY " + tb.Info.Table + "." + modules.Delimiter(connection.GetDelimiter(), connection.GetDelimiter2(), tb.PrimaryKey.Name)
	}

	var (
		wheres    = ""
		whereArgs = make([]interface{}, 0)
		existKeys = make([]string, 0)
	)

	wheres, whereArgs, existKeys = params.Statement(wheres, tb.Info.Table, connection.GetDelimiter(), connection.GetDelimiter2(), whereArgs, columns, existKeys,
		tb.Info.FieldList.GetFieldFilterProcessValue)
	wheres, whereArgs = tb.Info.Wheres.Statement(wheres, connection.GetDelimiter(), connection.GetDelimiter2(), whereArgs, existKeys, columns)
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
func (tb *DefaultTable) getDataFromDatabase(ctx *context.Context, params parameter.Parameters) (PanelInfo, error) {

	var (
		connection     = tb.db()
		delimiter      = connection.GetDelimiter()
		delimiter2     = connection.GetDelimiter2()
		placeholder    = modules.Delimiter(delimiter, delimiter2, "%s")
		queryStatement string
		countStatement string
		ids            = params.PKs()
		table          = modules.Delimiter(delimiter, delimiter2, tb.Info.Table)
		pk             = table + "." + modules.Delimiter(delimiter, delimiter2, tb.PrimaryKey.Name)
	)

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
			countStatement = "select count(*) as [size] from (select 1 as [size] from " + placeholder + " %s %s %s) src"
		} else {
			// %s means: fields, table, join table, wheres, group by, order by field, order by type
			queryStatement = "select %s from " + placeholder + "%s %s %s order by " + placeholder + "." + placeholder + " %s LIMIT ? OFFSET ?"
			// %s means: table, join table, wheres
			countStatement = "select count(*) from (select " + pk + " from " + placeholder + " %s %s %s) src"
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
					f := modules.Delimiter(connection.GetDelimiter(), connection.GetDelimiter2(), field.Field)
					headField := table + "." + f
					allFields = strings.ReplaceAll(allFields, headField, "CAST("+headField+" AS NVARCHAR(MAX)) as "+f)
					groupFields = strings.ReplaceAll(groupFields, headField, "CAST("+headField+" AS NVARCHAR(MAX))")
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
		wheres, whereArgs, existKeys = params.Statement(wheres, tb.Info.Table, connection.GetDelimiter(), connection.GetDelimiter2(), whereArgs, columns, existKeys,
			tb.Info.FieldList.GetFieldFilterProcessValue)
		// pre query
		wheres, whereArgs = tb.Info.Wheres.Statement(wheres, connection.GetDelimiter(), connection.GetDelimiter2(), whereArgs, existKeys, columns)
		wheres, whereArgs = tb.Info.WhereRaws.Statement(wheres, whereArgs)

		if wheres != "" {
			wheres = " where " + wheres
		}

		if connection.Name() == db.DriverMssql {
			args = append(whereArgs, (params.PageInt-1)*params.PageSizeInt, params.PageInt*params.PageSizeInt)
		} else {
			args = append(whereArgs, params.PageSizeInt, (params.PageInt-1)*params.PageSizeInt)
		}
	}

	groupBy := ""
	if len(joinTables) > 0 {
		if connection.Name() == db.DriverMssql || connection.Name() == db.DriverPostgresql {
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
		countCmd := fmt.Sprintf(countStatement, tb.Info.Table, joins, wheres, groupBy)

		total, err := connection.QueryWithConnection(tb.connection, countCmd, whereArgs...)

		if err != nil {
			return PanelInfo{}, err
		}

		logger.LogSQL(countCmd, nil)

		if tb.connectionDriver == "postgresql" {
			if tb.connectionDriverMode == "h2" {
				size = int(total[0]["count(*)"].(int64))
			} else if config.GetDatabases().GetDefault().DriverMode == "h2" {
				size = int(total[0]["count(*)"].(int64))
			} else {
				size = int(total[0]["count"].(int64))
			}
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
		Paginator: tb.GetPaginator(ctx, size, params,
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
func (tb *DefaultTable) GetDataWithId(param parameter.Parameters) (FormInfo, error) {

	var (
		res     map[string]interface{}
		columns Columns
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
			delimiter2     = connection.GetDelimiter2()
			tableName      = modules.Delimiter(delimiter, delimiter2, tb.GetForm().Table)
			pk             = tableName + "." + modules.Delimiter(delimiter, delimiter2, tb.PrimaryKey.Name)
			queryStatement = "select %s from %s %s where " + pk + " = ? %s "
		)

		for i := 0; i < len(tb.Form.FieldList); i++ {

			if tb.Form.FieldList[i].Field != pk && modules.InArray(columns, tb.Form.FieldList[i].Field) &&
				!tb.Form.FieldList[i].Joins.Valid() {
				fields += tableName + "." + modules.FilterField(tb.Form.FieldList[i].Field, delimiter, delimiter2) + ","
			}

			if tb.Form.FieldList[i].Joins.Valid() {
				headField := tb.Form.FieldList[i].Joins.Last().GetTableName() + parameter.FilterParamJoinInfix + tb.Form.FieldList[i].Field
				joinFields += db.GetAggregationExpression(connection.Name(), tb.Form.FieldList[i].Joins.Last().GetTableName(delimiter, delimiter2)+"."+
					modules.FilterField(tb.Form.FieldList[i].Field, delimiter, delimiter2), headField, types.JoinFieldValueDelimiter) + ","
				for _, join := range tb.Form.FieldList[i].Joins {
					if !modules.InArray(joinTables, join.GetTableName(delimiter, delimiter2)) {
						joinTables = append(joinTables, join.GetTableName(delimiter, delimiter2))
						if join.BaseTable == "" {
							join.BaseTable = tableName
						}
						joins += " left join " + modules.FilterField(join.Table, delimiter, delimiter2) + " " + join.TableAlias + " on " +
							join.GetTableName(delimiter, delimiter2) + "." + modules.FilterField(join.JoinField, delimiter, delimiter2) + " = " +
							join.BaseTable + "." + modules.FilterField(join.Field, delimiter, delimiter2)
					}
				}
			}
		}

		fields += pk
		groupFields := fields

		if joinFields != "" {
			fields += "," + joinFields[:len(joinFields)-1]
			if connection.Name() == db.DriverMssql {
				for i := 0; i < len(tb.Form.FieldList); i++ {
					if tb.Form.FieldList[i].TypeName == db.Text || tb.Form.FieldList[i].TypeName == db.Longtext {
						f := modules.Delimiter(connection.GetDelimiter(), connection.GetDelimiter2(), tb.Form.FieldList[i].Field)
						headField := tb.Info.Table + "." + f
						fields = strings.ReplaceAll(fields, headField, "CAST("+headField+" AS NVARCHAR(MAX)) as "+f)
						groupFields = strings.ReplaceAll(groupFields, headField, "CAST("+headField+" AS NVARCHAR(MAX))")
					}
				}
			}
		}

		if len(joinTables) > 0 {
			if connection.Name() == db.DriverMssql || connection.Name() == db.DriverPostgresql {
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
		groupFormList, groupHeaders = tb.Form.GroupFieldWithValue(tb.PrimaryKey.Name, id, columns, res, tb.sqlObjOrNil)
		return FormInfo{
			FieldList:         tb.Form.FieldList,
			GroupFieldList:    groupFormList,
			GroupFieldHeaders: groupHeaders,
			Title:             tb.Form.Title,
			Description:       tb.Form.Description,
		}, nil
	}

	var fieldList = tb.Form.FieldsWithValue(tb.PrimaryKey.Name, id, columns, res, tb.sqlObjOrNil)

	return FormInfo{
		FieldList:         fieldList,
		GroupFieldList:    groupFormList,
		GroupFieldHeaders: groupHeaders,
		Title:             tb.Form.Title,
		Description:       tb.Form.Description,
	}, nil
}

// UpdateData update data.
func (tb *DefaultTable) UpdateData(ctx *context.Context, dataList form.Values) error {

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
						logger.ErrorCtx(ctx, "UpdateData error %+v", err)
					}
				}()

				err := tb.Form.PostHook(dataList)
				if err != nil {
					logger.ErrorCtx(ctx, "UpdateData PostHook error %+v", err)
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
		err = tb.Form.UpdateFn(tb.PreProcessValue(dataList, types.PostTypeUpdate))
		if err != nil {
			errMsg = "post error: " + err.Error()
		}
		return err
	}

	if len(dataList) == 0 {
		return nil
	}

	_, err = tb.sql().Table(tb.Form.Table).
		Where(tb.PrimaryKey.Name, "=", dataList.Get(tb.PrimaryKey.Name)).
		Update(tb.getInjectValueFromFormValue(dataList, types.PostTypeUpdate))

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
func (tb *DefaultTable) InsertData(ctx *context.Context, dataList form.Values) error {

	dataList.Add(form.PostTypeKey, "1")

	var (
		id     = int64(0)
		err    error
		errMsg = ""
		f      = tb.GetActualNewForm()
	)

	if f.PostHook != nil {
		defer func() {
			dataList.Add(form.PostTypeKey, "1")
			dataList.Add(tb.GetPrimaryKey().Name, strconv.Itoa(int(id)))
			dataList.Add(form.PostResultKey, errMsg)

			go func() {
				defer func() {
					if err := recover(); err != nil {
						logger.ErrorCtx(ctx, "InsertData error %+v", err)
					}
				}()

				err := f.PostHook(dataList)
				if err != nil {
					logger.ErrorCtx(ctx, "InsertData PostHook error %+v", err)
				}
			}()
		}()
	}

	if f.Validator != nil {
		if err := f.Validator(dataList); err != nil {
			errMsg = "post error: " + err.Error()
			return err
		}
	}

	if f.PreProcessFn != nil {
		dataList = f.PreProcessFn(dataList)
	}

	if f.InsertFn != nil {
		dataList.Delete(form.PostTypeKey)
		err = f.InsertFn(tb.PreProcessValue(dataList, types.PostTypeCreate))
		if err != nil {
			errMsg = "post error: " + err.Error()
		}
		return err
	}

	if len(dataList) == 0 {
		return nil
	}

	id, err = tb.sql().Table(f.Table).Insert(tb.getInjectValueFromFormValue(dataList, types.PostTypeCreate))

	// NOTE: some errors should be ignored.
	if db.CheckError(err, db.INSERT) {
		errMsg = "post error: " + err.Error()
		return err
	}

	return nil
}

func (tb *DefaultTable) getInjectValueFromFormValue(dataList form.Values, typ types.PostType) dialect.H {

	var (
		value         = make(dialect.H)
		exceptString  = make([]string, 0)
		columns, auto = tb.getColumns(tb.Form.Table)

		fun types.PostFieldFilterFn
	)

	// If a key is a auto increment primary key, it can`t be insert or update.
	if auto {
		exceptString = []string{tb.PrimaryKey.Name, form.PreviousKey, form.MethodKey, form.TokenKey,
			constant.IframeKey, constant.IframeIDKey}
	} else {
		exceptString = []string{form.PreviousKey, form.MethodKey, form.TokenKey,
			constant.IframeKey, constant.IframeIDKey}
	}

	if !dataList.IsSingleUpdatePost() {
		for i := 0; i < len(tb.Form.FieldList); i++ {
			if tb.Form.FieldList[i].FormType.IsMultiSelect() {
				if _, ok := dataList[tb.Form.FieldList[i].Field+"[]"]; !ok {
					dataList[tb.Form.FieldList[i].Field+"[]"] = []string{""}
				}
			}
		}
	}

	dataList = dataList.RemoveRemark()

	for k, v := range dataList {
		k = strings.ReplaceAll(k, "[]", "")
		if !modules.InArray(exceptString, k) {
			if modules.InArray(columns, k) {
				field := tb.Form.FieldList.FindByFieldName(k)
				delimiter := ","
				if field != nil {
					fun = field.PostFilterFn
					delimiter = modules.SetDefault(field.DefaultOptionDelimiter, ",")
				}
				vv := modules.RemoveBlankFromArray(v)
				if fun != nil {
					value[k] = fun(types.PostFieldModel{
						ID:       dataList.Get(tb.PrimaryKey.Name),
						Value:    vv,
						Row:      dataList.ToMap(),
						PostType: typ,
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
				field := tb.Form.FieldList.FindByFieldName(k)
				if field != nil && field.PostFilterFn != nil {
					field.PostFilterFn(types.PostFieldModel{
						ID:       dataList.Get(tb.PrimaryKey.Name),
						Value:    modules.RemoveBlankFromArray(v),
						Row:      dataList.ToMap(),
						PostType: typ,
					})
				}
			}
		}
	}
	return value
}

func (tb *DefaultTable) PreProcessValue(dataList form.Values, typ types.PostType) form.Values {

	exceptString := []string{form.PreviousKey, form.MethodKey, form.TokenKey,
		constant.IframeKey, constant.IframeIDKey}
	dataList = dataList.RemoveRemark()
	var fun types.PostFieldFilterFn

	for k, v := range dataList {
		k = strings.ReplaceAll(k, "[]", "")
		if !modules.InArray(exceptString, k) {
			field := tb.Form.FieldList.FindByFieldName(k)
			if field != nil {
				fun = field.PostFilterFn
			}
			vv := modules.RemoveBlankFromArray(v)
			if fun != nil {
				dataList.Add(k, fmt.Sprintf("%s", fun(types.PostFieldModel{
					ID:       dataList.Get(tb.PrimaryKey.Name),
					Value:    vv,
					Row:      dataList.ToMap(),
					PostType: typ,
				})))
			}
		}
	}
	return dataList
}

// DeleteData delete data.
func (tb *DefaultTable) DeleteData(id string) error {

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

func (tb *DefaultTable) GetNewFormInfo() FormInfo {

	f := tb.GetActualNewForm()

	if len(f.TabGroups) == 0 {
		return FormInfo{FieldList: f.FieldsWithDefaultValue(tb.sqlObjOrNil)}
	}

	newForm, headers := f.GroupField(tb.sqlObjOrNil)

	return FormInfo{GroupFieldList: newForm, GroupFieldHeaders: headers}
}

// ***************************************
// helper function for database operation
// ***************************************

func (tb *DefaultTable) delete(table, key string, values []string) error {

	var vals = make([]interface{}, len(values))
	for i, v := range values {
		vals[i] = v
	}

	return tb.sql().Table(table).
		WhereIn(key, vals).
		Delete()
}

func (tb *DefaultTable) getTheadAndFilterForm(params parameter.Parameters, columns Columns) (types.Thead,
	string, string, string, []string, []types.FormField) {

	return tb.Info.FieldList.GetTheadAndFilterForm(types.TableInfo{
		Table:      tb.Info.Table,
		Delimiter:  tb.delimiter(),
		Delimiter2: tb.delimiter2(),
		Driver:     tb.connectionDriver,
		PrimaryKey: tb.PrimaryKey.Name,
	}, params, columns, tb.sqlObjOrNil)
}

// db is a helper function return raw db connection.
func (tb *DefaultTable) db() db.Connection {
	if tb.dbObj == nil {
		tb.dbObj = db.GetConnectionFromService(services.Get(tb.connectionDriver))
	}
	return tb.dbObj
}

func (tb *DefaultTable) delimiter() string {
	if tb.getDataFromDB() {
		return tb.db().GetDelimiter()
	}
	return ""
}

func (tb *DefaultTable) delimiter2() string {
	if tb.getDataFromDB() {
		return tb.db().GetDelimiter2()
	}
	return ""
}

func (tb *DefaultTable) getDataFromDB() bool {
	return tb.sourceURL == "" && tb.getDataFun == nil && tb.Info.GetDataFn == nil && tb.Detail.GetDataFn == nil
}

// sql is a helper function return db sql.
func (tb *DefaultTable) sql() *db.SQL {
	return db.WithDriverAndConnection(tb.connection, tb.db())
}

// sqlObjOrNil is a helper function return db sql obj or nil.
func (tb *DefaultTable) sqlObjOrNil() *db.SQL {
	if tb.connectionDriver != "" && tb.getDataFromDB() {
		return db.WithDriverAndConnection(tb.connection, tb.db())
	}
	return nil
}

type Columns []string

func (tb *DefaultTable) getColumns(table string) (Columns, bool) {

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
