package models

import "strconv"

func GetPaginator(path string, pageInt int, page string, pageSize string, size int, prefix string) map[string]interface{} {
	paginator := make(map[string]interface{}, 0)

	pageSizeInt, _ := strconv.Atoi(pageSize)

	if page == "1" {
		paginator["previou_class"] = "disabled"
		paginator["previou_url"] = path
	} else {
		paginator["previou_class"] = ""
		paginator["previou_url"] = path + "?page=" + strconv.Itoa(pageInt-1) + "&pageSize=" + pageSize
	}

	paginator["delete_url"] = "/delete/" + prefix
	paginator["new_url"] = "/info/" + prefix + "/new?page=" + page + "&pageSize=" + pageSize
	paginator["edit_url"] = "/info/" + prefix + "/edit?page=" + page + "&pageSize=" + pageSize
	paginator["pageSize"] = pageSize
	paginator["next_class"] = ""
	paginator["next_url"] = path + "?page=" + strconv.Itoa(pageInt+1) + "&pageSize=" + pageSize
	paginator["url"] = path + "?page=" + page
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

	return paginator
}
