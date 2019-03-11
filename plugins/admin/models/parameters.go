package models

import (
	"net/url"
	"strconv"
	"strings"
)

type Parameters struct {
	Page      string
	PageSize  string
	SortField string
	SortType  string
	Fields    map[string]string
}

func GetParam(values url.Values) *Parameters {
	page := GetDefault(values, "page", "1")
	pageSize := GetDefault(values, "pageSize", "10")
	sortField := GetDefault(values, "sort", "id")
	sortType := GetDefault(values, "sort_type", "desc")

	fields := make(map[string]string, 0)

	for key, value := range values {
		if key != "page" &&
			key != "pageSize" &&
			key != "sort" &&
			key != "sort_type" &&
			key != "prefix" &&
			key != "_pjax" &&
			value[0] != "" {
			if key == "sort_type" {
				if value[0] != "desc" && value[0] != "asc" {
					fields[key] = "desc"
				}
			} else {
				fields[key] = value[0]
			}
		}
	}

	return &Parameters{
		Page:      page,
		PageSize:  pageSize,
		SortField: sortField,
		SortType:  sortType,
		Fields:    fields,
	}
}

func GetParamFromUrl(value string) *Parameters {
	prevUrlArr := strings.Split(value, "?")
	paramArr := strings.Split(prevUrlArr[1], "&")
	page := "1"
	pageSize := "10"
	sortField := "id"
	sortType := "desc"

	//fields := make(map[string]string, 0)

	for i := 0; i < len(paramArr); i++ {
		arr := strings.Split(paramArr[i], "=")
		switch arr[0] {
		case "pageSize":
			pageSize = arr[1]
		case "page":
			page = arr[1]
		case "sort":
			sortField = arr[1]
		case "sort_type":
			sortType = arr[1]
			//default:
			//	fields[arr[0]] = arr[1]
		}
	}

	return &Parameters{
		Page:      page,
		PageSize:  pageSize,
		SortField: sortField,
		SortType:  sortType,
		//Fields:    fields,
	}
}

func (param *Parameters) SetPage(page string) *Parameters {
	param.Page = page
	return param
}

func (param *Parameters) GetRouteParamStr() string {
	return "?page=" + param.Page + param.GetFixedParamStr()
}

func (param *Parameters) GetRouteParamStrWithoutPageSize() string {
	return "?page=" + param.Page + param.GetFixedParamStrWithoutPageSize()
}

func (param *Parameters) GetLastPageRouteParamStr() string {
	pageInt, _ := strconv.Atoi(param.Page)
	return "?page=" + strconv.Itoa(pageInt-1) + param.GetFixedParamStr()
}

func (param *Parameters) GetNextPageRouteParamStr() string {
	pageInt, _ := strconv.Atoi(param.Page)
	return "?page=" + strconv.Itoa(pageInt+1) + param.GetFixedParamStr()
}

func (param *Parameters) GetFixedParamStrWithoutPageSize() string {
	str := "&"
	for key, value := range param.Fields {
		str += key + "=" + value + "&"
	}
	return "&sort=" + param.SortField + "&sort_type=" + param.SortType + str[:len(str)-1]
}

func (param *Parameters) GetFixedParamStr() string {
	str := "&"
	for key, value := range param.Fields {
		str += key + "=" + value + "&"
	}
	return "&pageSize=" + param.PageSize + "&sort=" + param.SortField + "&sort_type=" + param.SortType + str[:len(str)-1]
}

func GetDefault(values url.Values, key, def string) string {
	value := values.Get(key)
	if value == "" {
		return def
	}
	return value
}
