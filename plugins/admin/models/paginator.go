package models

import (
	template2 "github.com/chenhg5/go-admin/template"
	"github.com/chenhg5/go-admin/template/adminlte/components"
	"github.com/chenhg5/go-admin/template/types"
	"html/template"
	"math"
	"strconv"
)

func GetPaginator(path string, pageInt int, page, pageSize, sortField, sortType string, size int) types.PaginatorAttribute {

	paginator := template2.Get("adminlte").Paginator().(*components.PaginatorAttribute)

	pageSizeInt, _ := strconv.Atoi(pageSize)
	totalPage := int(math.Ceil(float64(size) / float64(pageSizeInt)))

	if page == "1" {
		paginator.PreviousClass = "disabled"
		paginator.PreviousUrl = path
	} else {
		paginator.PreviousClass = ""
		paginator.PreviousUrl = path + "?page=" + strconv.Itoa(pageInt-1) + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	}

	if pageInt == totalPage {
		paginator.NextClass = "disabled"
		paginator.NextUrl = path
	} else {
		paginator.NextClass = ""
		paginator.NextUrl = path + "?page=" + strconv.Itoa(pageInt+1) + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	}
	paginator.Url = path + "?page=" + page + "&sort=" + sortField + "&sort_type=" + sortType
	paginator.CurPageEndIndex = strconv.Itoa((pageInt) * pageSizeInt)
	paginator.CurPageStartIndex = strconv.Itoa((pageInt - 1) * pageSizeInt)
	paginator.Total = strconv.Itoa(size)
	paginator.Option = map[string]template.HTML{
		"10":  template.HTML(""),
		"20":  template.HTML(""),
		"30":  template.HTML(""),
		"50":  template.HTML(""),
		"100": template.HTML(""),
	}
	paginator.Option[pageSize] = template.HTML("selected")

	paginator.Pages = []map[string]string{}

	if totalPage < 10 {
		var pagesArr []map[string]string
		for i := 1; i < totalPage+1; i++ {
			if i == pageInt {
				pagesArr = append(pagesArr, map[string]string{
					"page":    strconv.Itoa(i),
					"active":  "active",
					"isSplit": "0",
					"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType,
				})
			} else {
				pagesArr = append(pagesArr, map[string]string{
					"page":    strconv.Itoa(i),
					"active":  "",
					"isSplit": "0",
					"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType,
				})
			}
		}
		paginator.Pages = pagesArr
	} else {
		var pagesArr []map[string]string
		if pageInt < 6 {
			for i := 1; i < totalPage+1; i++ {

				if i == pageInt {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "active",
						"isSplit": "0",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
				} else {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "",
						"isSplit": "0",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
				}

				if i == 6 {
					pagesArr = append(pagesArr, map[string]string{
						"page":    "",
						"active":  "",
						"isSplit": "1",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
					i = totalPage - 1
				}
			}
		} else if pageInt < totalPage-4 {
			for i := 1; i < totalPage+1; i++ {

				if i == pageInt {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "active",
						"isSplit": "0",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
				} else {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "",
						"isSplit": "0",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
				}

				if i == 2 {
					pagesArr = append(pagesArr, map[string]string{
						"page":    "",
						"active":  "",
						"isSplit": "1",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
					if pageInt < 7 {
						i = 5
					} else {
						i = pageInt - 2
					}
				}

				if pageInt < 7 {
					if i == pageInt+5 {
						pagesArr = append(pagesArr, map[string]string{
							"page":    "",
							"active":  "",
							"isSplit": "1",
							"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
						})
						i = totalPage - 1
					}
				} else {
					if i == pageInt+3 {
						pagesArr = append(pagesArr, map[string]string{
							"page":    "",
							"active":  "",
							"isSplit": "1",
							"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
						})
						i = totalPage - 1
					}
				}
			}
		} else {
			for i := 1; i < totalPage+1; i++ {

				if i == pageInt {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "active",
						"isSplit": "0",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
				} else {
					pagesArr = append(pagesArr, map[string]string{
						"page":    strconv.Itoa(i),
						"active":  "",
						"isSplit": "0",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
				}

				if i == 2 {
					pagesArr = append(pagesArr, map[string]string{
						"page":    "",
						"active":  "",
						"isSplit": "1",
						"url":     path + "?page=" + strconv.Itoa(i) + "&pageSize=" + pageSize,
					})
					i = totalPage - 4
				}
			}
		}
		paginator.Pages = pagesArr
	}

	return paginator
}
