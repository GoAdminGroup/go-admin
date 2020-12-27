// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package menu

import (
	"html/template"
	"regexp"
	"strconv"
	"strings"

	"github.com/GoAdminGroup/go-admin/modules/db/dialect"

	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin/models"
)

// Item is an menu item.
type Item struct {
	Name         string `json:"name"`
	ID           string `json:"id"`
	Url          string `json:"url"`
	IsLinkUrl    bool   `json:"isLinkUrl"`
	Icon         string `json:"icon"`
	Header       string `json:"header"`
	Active       string `json:"active"`
	ChildrenList []Item `json:"childrenList"`
}

// Menu contains list of menu items and other info.
type Menu struct {
	List        []Item              `json:"list"`
	Options     []map[string]string `json:"options"`
	MaxOrder    int64               `json:"maxOrder"`
	PluginName  string              `json:"pluginName"`
	ForceUpdate bool                `json:"forceUpdate"`
}

func (menu *Menu) GetUpdateJS(updateFlag bool) template.JS {
	if !updateFlag {
		return ""
	}

	forceUpdate := "false"
	if menu.ForceUpdate {
		forceUpdate = "true"
	}
	return template.JS(`$(function () {
	let curMenuPlug = $(".main-sidebar section.sidebar ul.sidebar-menu").attr("data-plug");
    if (curMenuPlug !== '` + menu.PluginName + `' || ` + forceUpdate + `) {
        $(".main-sidebar section.sidebar").html($("#sidebar-menu-tmpl").html())
    }
});`)
}

// SetMaxOrder set the max order of menu.
func (menu *Menu) SetMaxOrder(order int64) {
	menu.MaxOrder = order
}

// AddMaxOrder add the max order of menu.
func (menu *Menu) AddMaxOrder() {
	menu.MaxOrder++
}

// SetActiveClass set the active class of menu.
func (menu *Menu) SetActiveClass(path string) *Menu {

	reg, _ := regexp.Compile(`\?(.*)`)
	path = reg.ReplaceAllString(path, "")

	for i := 0; i < len(menu.List); i++ {
		menu.List[i].Active = ""
	}

	for i := 0; i < len(menu.List); i++ {
		if menu.List[i].Url == path && len(menu.List[i].ChildrenList) == 0 {
			menu.List[i].Active = "active"
			return menu
		}

		for j := 0; j < len(menu.List[i].ChildrenList); j++ {
			if menu.List[i].ChildrenList[j].Url == path {
				menu.List[i].Active = "active"
				menu.List[i].ChildrenList[j].Active = "active"
				return menu
			}

			menu.List[i].Active = ""
			menu.List[i].ChildrenList[j].Active = ""
		}
	}

	return menu
}

// FormatPath get template.HTML for front-end.
func (menu Menu) FormatPath() template.HTML {
	res := template.HTML(``)
	for i := 0; i < len(menu.List); i++ {
		if menu.List[i].Active != "" {
			if menu.List[i].Url != "#" && menu.List[i].Url != "" && len(menu.List[i].ChildrenList) > 0 {
				res += template.HTML(`<li><a href="` + menu.List[i].Url + `">` + menu.List[i].Name + `</a></li>`)
			} else {
				res += template.HTML(`<li>` + menu.List[i].Name + `</li>`)
				if len(menu.List[i].ChildrenList) == 0 {
					return res
				}
			}
			for j := 0; j < len(menu.List[i].ChildrenList); j++ {
				if menu.List[i].ChildrenList[j].Active != "" {
					return res + template.HTML(`<li>`+menu.List[i].ChildrenList[j].Name+`</li>`)
				}
			}
		}
	}
	return res
}

// GetEditMenuList return menu items list.
func (menu *Menu) GetEditMenuList() []Item {
	return menu.List
}

type NewMenuData struct {
	ParentId   int64  `json:"parent_id"`
	Type       int64  `json:"type"`
	Order      int64  `json:"order"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	PluginName string `json:"plugin_name"`
	Uri        string `json:"uri"`
	Header     string `json:"header"`
	Uuid       string `json:"uuid"`
}

func NewMenu(conn db.Connection, data NewMenuData) (int64, error) {
	maxOrder := data.Order
	checkOrder, _ := db.WithDriver(conn).Table("goadmin_menu").
		Where("plugin_name", "=", data.PluginName).
		OrderBy("order", "desc").
		First()

	if checkOrder != nil {
		maxOrder = checkOrder["order"].(int64)
	}

	id, err := db.WithDriver(conn).Table("goadmin_menu").
		Insert(dialect.H{
			"parent_id":   data.ParentId,
			"type":        data.Type,
			"order":       maxOrder,
			"title":       data.Title,
			"uuid":        data.Uuid,
			"icon":        data.Icon,
			"plugin_name": data.PluginName,
			"uri":         data.Uri,
			"header":      data.Header,
		})
	if !db.CheckError(err, db.INSERT) {
		return id, nil
	}
	return id, err
}

// GetGlobalMenu return Menu of given user model.
func GetGlobalMenu(user models.UserModel, conn db.Connection, lang string, pluginNames ...string) *Menu {

	var (
		menus      []map[string]interface{}
		menuOption = make([]map[string]string, 0)
		plugName   = ""
	)

	if len(pluginNames) > 0 {
		plugName = pluginNames[0]
	}

	user.WithRoles().WithMenus()

	if user.IsSuperAdmin() {
		menus, _ = db.WithDriver(conn).Table("goadmin_menu").
			Where("id", ">", 0).
			Where("plugin_name", "=", plugName).
			OrderBy("order", "asc").
			All()
	} else {

		var ids []interface{}
		for i := 0; i < len(user.MenuIds); i++ {
			ids = append(ids, user.MenuIds[i])
		}

		menus, _ = db.WithDriver(conn).Table("goadmin_menu").
			WhereIn("id", ids).
			Where("plugin_name", "=", plugName).
			OrderBy("order", "asc").
			All()
	}

	var title string
	for i := 0; i < len(menus); i++ {

		title = language.GetWithLang(menus[i]["title"].(string), lang)
		menuOption = append(menuOption, map[string]string{
			"id":    strconv.FormatInt(menus[i]["id"].(int64), 10),
			"title": title,
		})
	}

	menuList := constructMenuTree(menus, 0, lang)
	maxOrder := int64(0)
	if len(menus) > 0 {
		maxOrder = menus[len(menus)-1]["parent_id"].(int64)
	}

	return &Menu{
		List:       menuList,
		Options:    menuOption,
		MaxOrder:   maxOrder,
		PluginName: plugName,
	}
}

func constructMenuTree(menus []map[string]interface{}, parentID int64, lang string) []Item {

	branch := make([]Item, 0)

	var title string
	for j := 0; j < len(menus); j++ {
		if parentID == menus[j]["parent_id"].(int64) {
			if menus[j]["type"].(int64) == 1 {
				title = language.Get(menus[j]["title"].(string))
			} else {
				title = menus[j]["title"].(string)
			}

			header, _ := menus[j]["header"].(string)

			uri := menus[j]["uri"].(string)

			if lang != "" {
				if strings.Contains(uri, "?") {
					uri += "&__ga_lang=" + lang
				} else {
					uri += "?__ga_lang=" + lang
				}
			}

			child := Item{
				Name:         title,
				ID:           strconv.FormatInt(menus[j]["id"].(int64), 10),
				Url:          uri,
				Icon:         menus[j]["icon"].(string),
				Header:       header,
				Active:       "",
				ChildrenList: constructMenuTree(menus, menus[j]["id"].(int64), lang),
			}

			branch = append(branch, child)
		}
	}

	return branch
}
