package config

func GetGlobalMenu() []map[string]string {
	return []map[string]string{
		{
			"url":          "/user/info",
			"title":        "用户表",
			"active_class": "active",
		},
		{
			"url":          "/menu",
			"title":        "菜单管理",
			"active_class": "",
		},
	}
}