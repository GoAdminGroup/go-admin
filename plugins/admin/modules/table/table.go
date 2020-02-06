package table

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/db/dialect"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/paginator"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	"github.com/GoAdminGroup/go-admin/template/types"
	form2 "github.com/GoAdminGroup/go-admin/template/types/form"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Generator func(ctx *context.Context) Table

type GeneratorList map[string]Generator

func (g GeneratorList) InjectRoutes(app *context.App) {
	for _, gen := range g {
		table := gen(context.NewContext(&http.Request{
			URL: &url.URL{},
		}))
		for _, cb := range table.GetInfo().Callbacks {
			app.AppendReqAndResp(cb.Path, cb.Method, cb.Handlers)
		}
		for _, cb := range table.GetForm().Callbacks {
			app.AppendReqAndResp(cb.Path, cb.Method, cb.Handlers)
		}
	}
}

func (g GeneratorList) Add(key string, gen Generator) {
	g[key] = gen
}

func (g GeneratorList) Combine(gg GeneratorList) {
	for key, gen := range gg {
		if _, ok := g[key]; !ok {
			g[key] = gen
		}
	}
}

func (g GeneratorList) CombineAll(ggg []GeneratorList) {
	for _, gg := range ggg {
		for key, gen := range gg {
			if _, ok := g[key]; !ok {
				g[key] = gen
			}
		}
	}
}

var (
	generators = make(GeneratorList)
	tableList  = map[string]Table{}
)

func Get(key string) Table {
	return tableList[key]
}

func InitTableList(ctx *context.Context) {
	for prefix, generator := range generators {
		tableList[prefix] = generator(ctx)
	}
}

// RefreshTableList refresh the table list when the table relationship changed.
func RefreshTableList(ctx *context.Context) {
	for k, v := range generators {
		tableList[k] = v(ctx)
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
	GetDetail() *types.InfoPanel
	GetForm() *types.FormPanel

	GetCanAdd() bool
	GetEditable() bool
	GetDeletable() bool
	GetExportable() bool
	IsShowDetail() bool

	GetPrimaryKey() PrimaryKey

	GetData(path string, params parameter.Parameters, isAll bool) (PanelInfo, error)
	GetDataWithIds(path string, params parameter.Parameters, ids []string) (PanelInfo, error)
	GetDataWithId(id string) ([]types.FormField, [][]types.FormField, []string, string, string, error)
	UpdateDataFromDatabase(dataList form.Values) error
	InsertDataFromDatabase(dataList form.Values) error
	DeleteDataFromDatabase(id string) error

	Copy() Table
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

type Config struct {
	Driver     string
	Connection string
	CanAdd     bool
	Editable   bool
	Deletable  bool
	Exportable bool
	PrimaryKey PrimaryKey
	SourceURL  string
	GetDataFun GetDataFun
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

func (config Config) SetSourceURL(url string) Config {
	config.SourceURL = url
	return config
}

func (config Config) SetGetDataFun(fun GetDataFun) Config {
	config.GetDataFun = fun
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

	thead, filterForm := tb.getTheadAndFilterForm(params)

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

func (tb DefaultTable) getTheadAndFilterForm(params parameter.Parameters) (Thead, []types.FormField) {

	var (
		sortable   string
		editable   string
		hide       string
		headField  string
		filterForm = make([]types.FormField, 0)
		thead      = make([]map[string]string, 0)
	)
	for _, field := range tb.info.FieldList {

		headField = field.Field

		if field.Filterable {

			var value, value2 string

			if field.FilterType.IsRange() {
				value = params.GetFieldValue(headField + "_start__goadmin")
				value2 = params.GetFieldValue(headField + "_end__goadmin")
			} else {
				if field.FilterOperator == types.FilterOperatorFree {
					value2 = types.GetOperatorFromValue(params.GetFieldOperator(headField)).String()
				}
				value = params.GetFieldValue(headField)
			}

			filterForm = append(filterForm, types.FormField{
				Field:     headField,
				Head:      modules.AorB(field.FilterHead == "", field.Head, field.FilterHead),
				TypeName:  field.TypeName,
				HelpMsg:   field.FilterHelpMsg,
				FormType:  field.FilterType,
				Editable:  true,
				Value:     template.HTML(value),
				Value2:    value2,
				Options:   field.FilterOptions.SetSelected(params.GetFieldValue(field.Field), field.FilterType.SelectedLabel()),
				OptionExt: field.FilterOptionExt,
				Label:     field.FilterOperator.Label(),
			})

			if field.FilterOperator.AddOrNot() {
				filterForm = append(filterForm, types.FormField{
					Field:    headField + "__operator__",
					Head:     field.Head,
					TypeName: field.TypeName,
					Value:    template.HTML(field.FilterOperator.Value()),
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
	return thead, filterForm
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

	thead, filterForm := tb.getTheadAndFilterForm(params)

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
		if len(columns) == 0 || inArray(columns, headField) || field.Join.Valid() {
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
		placeholder    = delimiter(connection.GetDelimiter(), "%s")
		queryStatement = "select %s from %s %s order by " + placeholder + " %s"
	)

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns, _ := tb.getColumns(columnsModel)

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
		placeholder    = delimiter(connection.GetDelimiter(), "%s")
		queryStatement string
		countStatement string
	)

	beginTime := time.Now()

	if len(ids) > 0 {
		queryStatement = "select %s from %s %s where " + tb.primaryKey.Name + " in (%s) %s order by " + placeholder + " %s"
		countStatement = "select count(*) from " + placeholder + " %s where " + tb.primaryKey.Name + " in (%s)"
	} else {
		queryStatement = "select %s from " + placeholder + "%s %s %s order by " + placeholder + " %s LIMIT ? OFFSET ?"
		countStatement = "select count(*) from " + placeholder + " %s %s"
		if connection.Name() == "mssql" {
			queryStatement = "SELECT * FROM (SELECT ROW_NUMBER() OVER (ORDER BY " + placeholder + " %s) as ROWNUMBER_, %s from " +
				placeholder + "%s %s %s  ) as TMP_ WHERE TMP_.ROWNUMBER_ > ? AND TMP_.ROWNUMBER_ <= ?"
			countStatement = "select count(*) as [size] from " + placeholder + " %s %s"
		}
	}

	thead := make([]map[string]string, 0)
	fields := ""

	columnsModel, _ := tb.sql().Table(tb.info.Table).ShowColumns()

	columns, _ := tb.getColumns(columnsModel)

	var (
		sortable   string
		editable   string
		hide       string
		joins      string
		headField  string
		joinTables = make([]string, 0)
		filterForm = make([]types.FormField, 0)
		hasJoin    = false
	)
	for _, field := range tb.info.FieldList {
		if field.Field != tb.primaryKey.Name && inArray(columns, field.Field) &&
			!field.Join.Valid() {
			fields += tb.info.Table + "." + filterFiled(field.Field, connection.GetDelimiter()) + ","
		}

		headField = field.Field

		if field.Join.Valid() {
			hasJoin = true
			headField = field.Join.Table + "_goadmin_join_" + field.Field
			fields += getAggregationExpression(tb.connectionDriver, field.Join.Table+"."+
				filterFiled(field.Field, connection.GetDelimiter()), headField, types.JoinFieldValueDelimiter) + ","
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
				value = params.GetFieldValue(headField + "_start__goadmin")
				value2 = params.GetFieldValue(headField + "_end__goadmin")
			} else {
				if field.FilterOperator == types.FilterOperatorFree {
					value2 = types.GetOperatorFromValue(params.GetFieldOperator(headField)).String()
				}
				value = params.GetFieldValue(headField)
			}

			filterForm = append(filterForm, types.FormField{
				Field:     headField,
				Head:      modules.AorB(field.FilterHead == "", field.Head, field.FilterHead),
				TypeName:  field.TypeName,
				HelpMsg:   field.FilterHelpMsg,
				FormType:  field.FilterType,
				Editable:  true,
				Value:     template.HTML(value),
				Value2:    value2,
				Options:   field.FilterOptions.SetSelected(params.GetFieldValue(field.Field), field.FilterType.SelectedLabel()),
				OptionExt: field.FilterOptionExt,
				Label:     field.FilterOperator.Label(),
			})

			if field.FilterOperator.AddOrNot() {
				filterForm = append(filterForm, types.FormField{
					Field:    headField + "__operator__",
					Head:     field.Head,
					TypeName: field.TypeName,
					Value:    template.HTML(field.FilterOperator.Value()),
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

		if len(params.Fields) == 0 && len(tb.info.Wheres) == 0 && tb.info.WhereRaws.Raw == "" {
			wheres = ""
		} else {

			wheres = " where "

			for key, value := range params.Fields {

				if modules.InArray(existKeys, key) {
					continue
				}

				var op types.FilterOperator
				if strings.Contains(key, "_end__goadmin") {
					key = strings.Replace(key, "_end__goadmin", "", -1)
					op = "<="
				} else if strings.Contains(key, "_start__goadmin") {
					key = strings.Replace(key, "_start__goadmin", "", -1)
					op = ">="
				} else if !strings.Contains(key, "__operator__") {
					op = types.GetOperatorFromValue(params.GetFieldOperator(key))
				}

				if inArray(columns, key) {
					wheres += filterFiled(key, connection.GetDelimiter()) + " " + op.String() + " ? and "
					field := tb.info.FieldList.GetFieldByFieldName(key)
					if field.FilterProcess != nil {
						value = field.FilterProcess(value)
					}
					if op == types.FilterOperatorLike && !strings.Contains(value, "%") {
						whereArgs = append(whereArgs, "%"+value+"%")
					} else {
						whereArgs = append(whereArgs, value)
					}
				} else {
					keys := strings.Split(key, "_goadmin_join_")
					if len(keys) > 1 {
						if field := tb.info.FieldList.GetFieldByFieldName(keys[1]); field.Exist() && field.Join.Table != "" {
							if field.FilterProcess != nil {
								value = field.FilterProcess(value)
							}
							wheres += field.Join.Table + "." + filterFiled(keys[1], connection.GetDelimiter()) + " " + op.String() + " ? and "
							if op == types.FilterOperatorLike && !strings.Contains(value, "%") {
								whereArgs = append(whereArgs, "%"+value+"%")
							} else {
								whereArgs = append(whereArgs, value)
							}
						}
					}
				}

				existKeys = append(existKeys, key)
			}

			for k, wh := range tb.info.Wheres {

				whFieldArr := strings.Split(wh.Field, ".")
				whField := ""
				whTable := ""
				if len(whFieldArr) > 1 {
					whField = whFieldArr[1]
					whTable = whFieldArr[0]
				} else {
					whField = whFieldArr[0]
				}

				if modules.InArray(existKeys, whField) {
					continue
				}

				// TODO: support like operation and join table
				if inArray(columns, whField) {

					joinMark := "and"
					if k != len(tb.info.Wheres)-1 {
						joinMark = tb.info.Wheres[k+1].Join
					}

					if whTable != "" {
						wheres += whTable + "." + filterFiled(whField, connection.GetDelimiter()) + " " + wh.Operator + " ? " + joinMark + " "
					} else {
						wheres += filterFiled(whField, connection.GetDelimiter()) + " " + wh.Operator + " ? " + joinMark + " "
					}
					whereArgs = append(whereArgs, wh.Arg)
				}
			}

			if wheres != " where " {
				wheres = wheres[:len(wheres)-4]
				if tb.info.WhereRaws.Raw != "" {
					checkGrammar := false
					for i := 0; i < len(tb.info.WhereRaws.Raw); i++ {
						if tb.info.WhereRaws.Raw[i] == ' ' {
							continue
						} else {
							if tb.info.WhereRaws.Raw[i] == 'a' {
								if len(tb.info.WhereRaws.Raw) < i+3 {
									break
								} else {
									if tb.info.WhereRaws.Raw[i+1] == 'n' && tb.info.WhereRaws.Raw[i+2] == 'd' {
										checkGrammar = true
									}
								}
							} else if tb.info.WhereRaws.Raw[i] == 'o' {
								if len(tb.info.WhereRaws.Raw) < i+2 {
									break
								} else {
									if tb.info.WhereRaws.Raw[i+1] == 'r' {
										checkGrammar = true
									}
								}
							} else {
								break
							}
						}
					}

					if checkGrammar {
						wheres += tb.info.WhereRaws.Raw + " "
					} else {
						wheres += " and " + tb.info.WhereRaws.Raw + " "
					}

					whereArgs = append(whereArgs, tb.info.WhereRaws.Args...)
				}
			} else {
				if tb.info.WhereRaws.Raw != "" {
					index := 0
					for i := 0; i < len(tb.info.WhereRaws.Raw); i++ {
						if tb.info.WhereRaws.Raw[i] == ' ' {
							continue
						} else {
							if tb.info.WhereRaws.Raw[i] == 'a' {
								if len(tb.info.WhereRaws.Raw) < i+3 {
									break
								} else {
									if tb.info.WhereRaws.Raw[i+1] == 'n' && tb.info.WhereRaws.Raw[i+2] == 'd' {
										index = i + 3
									}
								}
							} else if tb.info.WhereRaws.Raw[i] == 'o' {
								if len(tb.info.WhereRaws.Raw) < i+2 {
									break
								} else {
									if tb.info.WhereRaws.Raw[i+1] == 'r' {
										index = i + 2
									}
								}
							} else {
								break
							}
						}
					}
					wheres += tb.info.WhereRaws.Raw[index:] + " "
					whereArgs = append(whereArgs, tb.info.WhereRaws.Args...)
				} else {
					wheres = ""
				}
			}

		}
		pageSize, _ := strconv.Atoi(params.PageSize)
		if connection.Name() == "mssql" {
			args = append(whereArgs, (modules.GetPage(params.Page)-1)*pageSize, modules.GetPage(params.Page)*pageSize)
		} else {
			args = append(whereArgs, params.PageSize, (modules.GetPage(params.Page)-1)*pageSize)
		}
	}

	groupBy := ""
	if hasJoin {
		groupBy = " GROUP BY " + tb.info.Table + "." + filterFiled(tb.GetPrimaryKey().Name, connection.GetDelimiter())
	}

	queryCmd := fmt.Sprintf(queryStatement, fields, tb.info.Table, joins, wheres, groupBy, params.SortField, params.SortType)
	if connection.Name() == "mssql" {
		queryCmd = fmt.Sprintf(queryStatement, params.SortField, params.SortType, fields, tb.info.Table, joins, wheres, groupBy)
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
		columnsModel, err := tb.sql().Table(tb.form.Table).ShowColumns()

		if err != nil {
			return nil, nil, nil, "", "", err
		}

		columns, _ = tb.getColumns(columnsModel)

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
		if inArray(columns, formList[i].Field) {
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
						rowValue := modules.AorB(inArray(columns, field.Field) || len(columns) == 0,
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
		rowValue := modules.AorB(inArray(columns, field.Field) || len(columns) == 0,
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
	value := make(dialect.H)

	columnsModel, _ := tb.sql().Table(tb.form.Table).ShowColumns()

	columns, auto := tb.getColumns(columnsModel)
	var (
		fun          types.PostFieldFilterFn
		exceptString = make([]string, 0)
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
	if del == "[" {
		return "[" + s + "]"
	}
	return del + s + del
}

func filterFiled(filed, delimiter string) string {
	if delimiter == "[" {
		return filed
	}
	return delimiter + filed + delimiter
}

type Columns []string

func (tb DefaultTable) getColumns(columnsModel []map[string]interface{}) (Columns, bool) {
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

func getAggregationExpression(driver, field, headField, delimiter string) string {
	switch driver {
	case "postgresql":
		return fmt.Sprintf("string_agg(%s::character varying, '%s') as %s", field, delimiter, headField)
	case "mysql":
		return fmt.Sprintf("group_concat(%s separator '%s') as %s", field, delimiter, headField)
	case "sqlite":
		return fmt.Sprintf("group_concat(%s, '%s') as %s", field, delimiter, headField)
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
