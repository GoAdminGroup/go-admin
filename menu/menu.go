package menu

import (
	"goAdmin/connections/mysql"
	"strconv"
)

type Menu struct {
	GlobalMenuList   []MenuItem
	GlobalMenuOption []map[string]string
	MaxOrder         int64
}

var GlobalMenu = &Menu{
	[]MenuItem{},
	[]map[string]string{},
	0,
}

func InitMenu() {
	SetGlobalMenu()
}

func SetGlobalMenu() {
	menus, _ := mysql.Query("select * from goadmin_menu where id > 0 order by `order` asc")

	menuOption := make([]map[string]string, 0)

	for i := 0; i < len(menus); i++ {
		menuOption = append(menuOption, map[string]string{
			"id":    strconv.FormatInt(menus[i]["id"].(int64), 10),
			"title": menus[i]["title"].(string),
		})
	}

	menulist := ConstructMenuTree(menus, 0)

	(*GlobalMenu).GlobalMenuOption = menuOption
	(*GlobalMenu).GlobalMenuList = menulist
	(*GlobalMenu).MaxOrder = menus[len(menus)-1]["parent_id"].(int64)
}

func (menu *Menu) SexMaxOrder(order int64) {
	menu.MaxOrder = order
}

func ConstructMenuTree(menus []map[string]interface{}, parentId int64) []MenuItem {

	branch := make([]MenuItem, 0)

	for j := 0; j < len(menus); j++ {
		if parentId == menus[j]["parent_id"].(int64) {

			childList := ConstructMenuTree(menus, menus[j]["id"].(int64))

			child := MenuItem{
				menus[j]["title"].(string),
				strconv.FormatInt(menus[j]["id"].(int64), 10),
				menus[j]["uri"].(string),
				menus[j]["icon"].(string),
				"",
				childList,
			}

			branch = append(branch, child)
		}
	}

	return branch
}

func GetMenuItemById(id string) MenuItem {
	menu, _ := mysql.Query("select * from goadmin_menu where id = ?", id)

	return MenuItem{
		menu[0]["title"].(string),
		strconv.FormatInt(menu[0]["id"].(int64), 10),
		menu[0]["uri"].(string),
		menu[0]["icon"].(string),
		"",
		[]MenuItem{},
	}
}

func (menu *Menu) SetActiveClass(path string) {
	for i := 0; i < len((*menu).GlobalMenuList); i++ {
		if (*menu).GlobalMenuList[i].Url == path {
			(*menu).GlobalMenuList[i].Active = "active"
		} else {
			(*menu).GlobalMenuList[i].Active = ""
		}
	}
}

type MenuItem struct {
	Name         string
	ID           string
	Url          string
	Icon         string
	Active       string
	ChildrenList []MenuItem
}

func (menu *Menu) GetEditMenuList() []MenuItem {
	return (*menu).GlobalMenuList
}
