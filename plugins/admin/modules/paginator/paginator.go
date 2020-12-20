package paginator

import (
	"fmt"
	"html/template"
	"math"
	"strconv"

	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/parameter"
	template2 "github.com/GoAdminGroup/go-admin/template"
	"github.com/GoAdminGroup/go-admin/template/components"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Config struct {
	Size         int
	Param        parameter.Parameters
	PageSizeList []string
}

func Get(cfg Config) types.PaginatorAttribute {

	paginator := template2.Default().Paginator().(*components.PaginatorAttribute)

	totalPage := int(math.Ceil(float64(cfg.Size) / float64(cfg.Param.PageSizeInt)))

	if cfg.Param.PageInt == 1 {
		paginator.PreviousClass = "disabled"
		paginator.PreviousUrl = cfg.Param.URLPath
	} else {
		paginator.PreviousClass = ""
		paginator.PreviousUrl = cfg.Param.URLPath + cfg.Param.GetLastPageRouteParamStr(true)
	}

	if cfg.Param.PageInt == totalPage {
		paginator.NextClass = "disabled"
		paginator.NextUrl = cfg.Param.URLPath
	} else {
		paginator.NextClass = ""
		paginator.NextUrl = cfg.Param.URLPath + cfg.Param.GetNextPageRouteParamStr(true)
	}
	paginator.Url = cfg.Param.URLPath + cfg.Param.GetRouteParamStrWithoutPageSize("1") + "&" + form.NoAnimationKey + "=true"
	paginator.CurPageEndIndex = strconv.Itoa((cfg.Param.PageInt) * cfg.Param.PageSizeInt)
	paginator.CurPageStartIndex = strconv.Itoa((cfg.Param.PageInt-1)*cfg.Param.PageSizeInt + 1)
	paginator.Total = strconv.Itoa(cfg.Size)

	if len(cfg.PageSizeList) == 0 {
		cfg.PageSizeList = []string{"10", "20", "50", "100"}
	}

	paginator.Option = make(map[string]template.HTML, len(cfg.PageSizeList))
	for i := 0; i < len(cfg.PageSizeList); i++ {
		paginator.Option[cfg.PageSizeList[i]] = template.HTML("")
	}

	paginator.Option[cfg.Param.PageSize] = template.HTML("selected")

	paginator.Pages = []map[string]string{}

	var pagesArr []map[string]string

	if totalPage < 10 {
		for i := 1; i < totalPage+1; i++ {
			if i == cfg.Param.PageInt {
				pagesArr = addPageLink(pagesArr, cfg.Param, cfg.Param.PageInt, "active")
			} else {
				pagesArr = addPageLink(pagesArr, cfg.Param, i, "")
			}
		}
	} else {
		if cfg.Param.PageInt < 6 {
			for i := 1; i < totalPage+1; i++ {

				if i == cfg.Param.PageInt {
					pagesArr = addPageLink(pagesArr, cfg.Param, cfg.Param.PageInt, "active")
				} else {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "")
				}

				if i == 6 {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "split")
					i = totalPage - 2
				}
			}
		} else if cfg.Param.PageInt < totalPage-4 {
			for i := 1; i < totalPage+1; i++ {

				if i == cfg.Param.PageInt {
					pagesArr = addPageLink(pagesArr, cfg.Param, cfg.Param.PageInt, "active")
				} else {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "")
				}

				if i == 2 {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "split")
					i = cfg.Param.PageInt - 3
				}

				if i == cfg.Param.PageInt+2 {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "split")
					i = totalPage - 1
				}
			}
		} else {
			for i := 1; i < totalPage+1; i++ {

				if i == cfg.Param.PageInt {
					pagesArr = addPageLink(pagesArr, cfg.Param, cfg.Param.PageInt, "active")
				} else {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "")
				}

				if i == 2 {
					pagesArr = addPageLink(pagesArr, cfg.Param, i, "split")
					i = totalPage - 6
				}
			}
		}
	}

	paginator.Pages = pagesArr

	endNum := paginator.CurPageEndIndex
	if cfg.Size < cfg.Param.PageSizeInt {
		endNum = paginator.Total
	}

	paginator.SetEntriesInfo(template.HTML(fmt.Sprintf(language.Get("showing <b>%s</b> to <b>%s</b> of <b>%s</b> entries"),
		paginator.CurPageStartIndex, endNum, paginator.Total)))

	return paginator.SetPageSizeList(cfg.PageSizeList)
}

func addPageLink(arr []map[string]string, params parameter.Parameters, page int, active string) []map[string]string {

	pageStr := strconv.Itoa(page)
	isSplit := "0"

	if active == "split" {
		isSplit = "1"
		active = ""
		pageStr = ""
	}

	return append(arr, map[string]string{
		"page":    pageStr,
		"active":  active,
		"isSplit": isSplit,
		"url":     params.URLNoAnimation(pageStr),
	})
}
