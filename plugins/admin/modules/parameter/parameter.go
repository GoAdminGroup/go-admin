package parameter

import (
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/constant"
	"net/url"
	"strconv"
	"strings"
)

type Parameters struct {
	Page        string
	PageInt     int
	PageSize    string
	PageSizeInt int
	SortField   string
	Columns     []string
	SortType    string
	Fields      map[string]string
}

const (
	Page     = "__page"
	PageSize = "__pageSize"
	Sort     = "__sort"
	SortType = "__sort_type"
	Columns  = "__columns"
	Prefix   = "__prefix"
	Pjax     = "_pjax"

	operatorSuffix = "__operator__"

	sortTypeDesc = "desc"
	sortTypeAsc  = "asc"

	IsAll      = "is_all"
	PrimaryKey = "pk"

	true  = "true"
	false = "false"

	FilterRangeParamStartSuffix = "_start__goadmin"
	FilterRangeParamEndSuffix   = "_end__goadmin"
	FilterParamJoinInfix        = "_goadmin_join_"
	FilterParamOperatorSuffix   = "__operator__"
)

var operators = map[string]string{
	"like": "like",
	"gr":   ">",
	"gq":   ">=",
	"eq":   "=",
	"ne":   "!=",
	"le":   "<",
	"lq":   "<=",
	"free": "free",
}

var keys = []string{Page, PageSize, Sort, Columns, Prefix, Pjax}

func BaseParam() Parameters {
	return Parameters{Page: "1", PageSize: "1", Fields: make(map[string]string)}
}

func (param Parameters) WithPK(id ...string) Parameters {
	param.Fields["pk"] = strings.Join(id, ",")
	return param
}

func (param Parameters) PK() []string {
	return strings.Split(param.Fields[PrimaryKey], ",")
}

func (param Parameters) IsAll() bool {
	return param.Fields[IsAll] == true
}

func (param Parameters) WithIsAll(isAll bool) Parameters {
	if isAll {
		param.Fields[IsAll] = true
	} else {
		param.Fields[IsAll] = false
	}
	return param
}

func GetParam(values url.Values, defaultPageSize int, primaryKey, defaultSort string) Parameters {
	page := getDefault(values, Page, "1")
	pageSize := getDefault(values, PageSize, strconv.Itoa(defaultPageSize))
	sortField := getDefault(values, Sort, primaryKey)
	sortType := getDefault(values, SortType, defaultSort)
	columns := getDefault(values, Columns, "")

	fields := make(map[string]string)

	for key, value := range values {
		if !modules.InArray(keys, key) && value[0] != "" {
			if key == SortType {
				if value[0] != sortTypeDesc && value[0] != sortTypeAsc {
					fields[key] = sortTypeDesc
				}
			} else {
				if strings.Contains(key, operatorSuffix) &&
					values.Get(strings.Replace(key, operatorSuffix, "", -1)) == "" {
					continue
				}
				fields[key] = value[0]
			}
		}
	}

	columnsArr := make([]string, 0)
	if columns != "" {
		columnsArr = strings.Split(columns, ",")
	}

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	return Parameters{
		Page:        page,
		PageSize:    pageSize,
		PageSizeInt: pageSizeInt,
		PageInt:     pageInt,
		SortField:   sortField,
		SortType:    sortType,
		Fields:      fields,
		Columns:     columnsArr,
	}
}

func (param Parameters) GetFilterFieldValueStart(field string) string {
	return param.Fields[field] + FilterRangeParamStartSuffix
}

func (param Parameters) GetFilterFieldValueEnd(field string) string {
	return param.Fields[field] + FilterRangeParamEndSuffix
}

func (param Parameters) GetFieldValue(field string) string {
	return param.Fields[field]
}

func (param Parameters) GetFieldOperator(field string) string {
	if param.Fields[field+operatorSuffix] == "" {
		return "eq"
	}
	return param.Fields[field+operatorSuffix]
}

func GetParamFromUrl(value string, fromList bool, defaultPageSize int, primaryKey, defaultSort string) Parameters {

	if !fromList {
		return BaseParam()
	}

	prevUrlArr := strings.Split(value, "?")
	paramArr := strings.Split(prevUrlArr[1], "&")

	var (
		page      = "1"
		pageSize  = strconv.Itoa(defaultPageSize)
		sortField = primaryKey
		sortType  = defaultSort
		columns   = make([]string, 0)
	)

	for i := 0; i < len(paramArr); i++ {
		arr := strings.Split(paramArr[i], "=")
		switch arr[0] {
		case PageSize:
			pageSize = arr[1]
		case Page:
			page = arr[1]
		case Sort:
			sortField = arr[1]
		case SortType:
			sortType = arr[1]
		case Columns:
			columns = strings.Split(arr[1], ",")
		}
	}

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	return Parameters{
		Page:        page,
		PageSize:    pageSize,
		PageSizeInt: pageSizeInt,
		PageInt:     pageInt,
		SortField:   sortField,
		SortType:    sortType,
		Columns:     columns,
		Fields:      make(map[string]string),
	}
}

func (param Parameters) Join() string {
	p := param.GetFixedParamStr()
	p.Add(Page, param.Page)
	return p.Encode()
}

func (param Parameters) SetPage(page string) Parameters {
	param.Page = page
	return param
}

func (param Parameters) GetRouteParamStr() string {
	p := param.GetFixedParamStr()
	p.Add(Page, param.Page)
	return "?" + p.Encode()
}

func (param Parameters) GetRouteParamStrWithoutPageSize() string {
	p := url.Values{}
	p.Add(Sort, param.SortField)
	p.Add(Page, param.Page)
	p.Add(SortType, param.SortType)
	if len(param.Columns) > 0 {
		p.Add(Columns, strings.Join(param.Columns, ","))
	}
	for key, value := range param.Fields {
		p.Add(key, value)
	}
	return "?" + p.Encode()
}

func (param Parameters) GetLastPageRouteParamStr() string {
	p := param.GetFixedParamStr()
	p.Add(Page, strconv.Itoa(param.PageInt-1))
	return "?" + p.Encode()
}

func (param Parameters) GetNextPageRouteParamStr() string {
	p := param.GetFixedParamStr()
	p.Add(Page, strconv.Itoa(param.PageInt+1))
	return "?" + p.Encode()
}

func (param Parameters) GetFixedParamStr() url.Values {
	p := url.Values{}
	p.Add(Sort, param.SortField)
	p.Add(PageSize, param.PageSize)
	p.Add(SortType, param.SortType)
	if len(param.Columns) > 0 {
		p.Add(Columns, strings.Join(param.Columns, ","))
	}
	for key, value := range param.Fields {
		if key != constant.EditPKKey && key != constant.DetailPKKey {
			p.Add(key, value)
		}
	}
	return p
}

func (param Parameters) Statement(wheres, delimiter string, whereArgs []interface{}, columns, existKeys []string,
	filterProcess func(string, string) string, getJoinTable func(string) string) (string, []interface{}, []string) {
	for key, value := range param.Fields {

		if modules.InArray(existKeys, key) {
			continue
		}

		var op string
		if strings.Contains(key, FilterRangeParamEndSuffix) {
			key = strings.Replace(key, FilterRangeParamEndSuffix, "", -1)
			op = "<="
		} else if strings.Contains(key, FilterRangeParamStartSuffix) {
			key = strings.Replace(key, FilterRangeParamStartSuffix, "", -1)
			op = ">="
		} else if !strings.Contains(key, FilterParamOperatorSuffix) {
			op = operators[param.GetFieldOperator(key)]
		}

		if modules.InArray(columns, key) {
			wheres += modules.FilterField(key, delimiter) + " " + op + " ? and "
			if op == "like" && !strings.Contains(value, "%") {
				whereArgs = append(whereArgs, "%"+filterProcess(key, value)+"%")
			} else {
				whereArgs = append(whereArgs, value)
			}
		} else {
			keys := strings.Split(key, FilterParamJoinInfix)
			if len(keys) > 1 {
				if joinTable := getJoinTable(key); joinTable != "" {
					value := filterProcess(key, value)
					wheres += joinTable + "." + modules.FilterField(keys[1], delimiter) + " " + op + " ? and "
					if op == "like" && !strings.Contains(value, "%") {
						whereArgs = append(whereArgs, "%"+value+"%")
					} else {
						whereArgs = append(whereArgs, value)
					}
				}
			}
		}

		existKeys = append(existKeys, key)
	}

	if len(wheres) > 3 {
		wheres = wheres[:len(wheres)-4]
	}

	return wheres, whereArgs, existKeys
}

func getDefault(values url.Values, key, def string) string {
	value := values.Get(key)
	if value == "" {
		return def
	}
	return value
}
