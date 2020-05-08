package main

var langs = map[string]map[string]string{
	"cn": {
		"user login name": "用户登录名",
		"user nickname":   "用户昵称",
		"user password":   "用户密码",

		"choose a driver":          "选择数据库驱动",
		"sql address":              "连接地址",
		"sql port":                 "端口",
		"sql username":             "用户名",
		"sql schema":               "Schema",
		"sql database name":        "数据库名",
		"sql file":                 "文件地址",
		"sql password":             "密码",
		"choose table to generate": "选择要生成的表格",

		"wrong config file path": "错误的配置文件路径",
		"user record exists":     "用户记录已存在",
		"empty tables":           "表格不能为空",

		"tables to generate, use comma to split": "要生成权限的表格，用逗号分隔",

		"no tables, you should build a table of your own business first.": "表格不能为空，请先创建您的业务表",
		"no table is selected": "没有选择表格",

		"set package name":     "设置包名",
		"set connection name":  "设置连接",
		"set file output path": "设置文件输出路径",

		"generate permission records for tables, Y on behalf of yes": "是否生成表格权限，Y 代表是",

		"Query":                 "查询",
		"Show Edit Form Page":   "编辑页显示",
		"Show Create Form Page": "新建记录页显示",
		"Edit":                  "编辑",
		"Create":                "新建",
		"Delete":                "删除",
		"Export":                "导出",

		"Use arrows to move, type to filter, enter to select": "使用方向键去移动，空格键选择，输入进行筛选",
		"select all": "选择全部",
		"Use arrows to move, space to select, type to filter": "使用方向键去移动，空格键选择，输入进行筛选",
		"Add admin user success~~🍺🍺":                          "增加用户成功~~🍺🍺",
		"Add table permissions success~~🍺🍺":                   "增加表格权限成功~~🍺🍺",
		"Generate data table models success~~🍺🍺":              "生成数据模型文件成功~~🍺🍺",
		"see the docs: ": "查看文档：",
		"generating: ":   "生成中：",
	},
	"en": {},
}

var defaultLang = "en"

func setDefaultLangSet(set string) {
	if set != "" && (set == "cn" || set == "en") {
		defaultLang = set
	}
}

func getWord(msg string) string {
	if word, ok := langs[defaultLang][msg]; ok {
		return word
	}
	return msg
}
