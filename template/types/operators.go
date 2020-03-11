package types

import "html/template"

type FilterOperator string

const (
	FilterOperatorLike           FilterOperator = "like"
	FilterOperatorGreater        FilterOperator = ">"
	FilterOperatorGreaterOrEqual FilterOperator = ">="
	FilterOperatorEqual          FilterOperator = "="
	FilterOperatorNotEqual       FilterOperator = "!="
	FilterOperatorLess           FilterOperator = "<"
	FilterOperatorLessOrEqual    FilterOperator = "<="
	FilterOperatorFree           FilterOperator = "free"
)

func GetOperatorFromValue(value string) FilterOperator {
	switch value {
	case "like":
		return FilterOperatorLike
	case "gr":
		return FilterOperatorGreater
	case "gq":
		return FilterOperatorGreaterOrEqual
	case "eq":
		return FilterOperatorEqual
	case "ne":
		return FilterOperatorNotEqual
	case "le":
		return FilterOperatorLess
	case "lq":
		return FilterOperatorLessOrEqual
	case "free":
		return FilterOperatorFree
	default:
		return FilterOperatorEqual
	}
}

func (o FilterOperator) Value() string {
	switch o {
	case FilterOperatorLike:
		return "like"
	case FilterOperatorGreater:
		return "gr"
	case FilterOperatorGreaterOrEqual:
		return "gq"
	case FilterOperatorEqual:
		return "eq"
	case FilterOperatorNotEqual:
		return "ne"
	case FilterOperatorLess:
		return "le"
	case FilterOperatorLessOrEqual:
		return "lq"
	case FilterOperatorFree:
		return "free"
	default:
		return "eq"
	}
}

func (o FilterOperator) String() string {
	return string(o)
}

func (o FilterOperator) Label() template.HTML {
	if o == FilterOperatorLike {
		return ""
	}
	return template.HTML(o)
}

func (o FilterOperator) AddOrNot() bool {
	return string(o) != "" && o != FilterOperatorFree
}

func (o FilterOperator) Valid() bool {
	switch o {
	case FilterOperatorLike, FilterOperatorGreater, FilterOperatorGreaterOrEqual,
		FilterOperatorLess, FilterOperatorLessOrEqual, FilterOperatorFree:
		return true
	default:
		return false
	}
}
