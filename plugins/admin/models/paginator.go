package models

import (
	"goAdmin/modules/auth"
	"strconv"
)

func GetPaginator(path string, pageInt int, page, pageSize, sortField, sortType string, size int, prefix string) map[string]interface{} {
	paginator := make(map[string]interface{}, 0)

	pageSizeInt, _ := strconv.Atoi(pageSize)

	if page == "1" {
		paginator["previou_class"] = "disabled"
		paginator["previou_url"] = path
	} else {
		paginator["previou_class"] = ""
		paginator["previou_url"] = path + "?page=" + strconv.Itoa(pageInt-1) + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	}

	paginator["delete_url"] = "/delete/" + prefix
	paginator["new_url"] = "/info/" + prefix + "/new?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	paginator["edit_url"] = "/info/" + prefix + "/edit?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	paginator["pageSize"] = pageSize
	paginator["next_class"] = ""
	paginator["next_url"] = path + "?page=" + strconv.Itoa(pageInt+1) + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	paginator["url"] = path + "?page=" + page + "&pageSize=" + pageSize + "&sort=" + sortField + "&sort_type=" + sortType
	paginator["curPageEndIndex"] = strconv.Itoa((pageInt) * pageSizeInt)
	paginator["curPageStartIndex"] = strconv.Itoa((pageInt - 1) * pageSizeInt)
	paginator["total"] = strconv.Itoa(size)
	paginator["curpage"] = page
	paginator["option"] = map[string]string{
		"10":  "",
		"20":  "",
		"30":  "",
		"50":  "",
		"100": "",
	}
	paginator["option"].(map[string]string)[pageSize] = "selected=''"
	paginator["success"] = false
	paginator["token"] = auth.TokenHelper.AddToken()

	paginator["pages"] = []map[string]string{}
	totalPage := size/pageSizeInt + 1
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
		paginator["pages"] = pagesArr
	} else {
		var pagesArr []map[string]string
		if pageInt < 5 {
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
		paginator["pages"] = pagesArr
	}

	return paginator
}
