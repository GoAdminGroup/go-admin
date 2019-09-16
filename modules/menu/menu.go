// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package menu

import (
	"github.com/chenhg5/go-admin/modules/db"
	"github.com/chenhg5/go-admin/modules/language"
	"github.com/chenhg5/go-admin/plugins/admin/models"
	"strconv"
)

type Item struct {
	Name         string
	ID           string
	Url          string
	Icon         string
	Header       string
	Active       string
	ChildrenList []Item
}

type Menu struct {
	List     []Item
	Options  []map[string]string
	MaxOrder int64
}

func (menu *Menu) SexMaxOrder(order int64) {
	menu.MaxOrder = order
}

func (menu *Menu) AddMaxOrder() {
	menu.MaxOrder += 1
}

func (menu *Menu) SetActiveClass(path string) *Menu {

	for i := 0; i < len((*menu).List); i++ {
		(*menu).List[i].Active = ""
	}

	for i := 0; i < len((*menu).List); i++ {
		if (*menu).List[i].Url == path {
			(*menu).List[i].Active = "active"
			return menu
		} else {
			for j := 0; j < len((*menu).List[i].ChildrenList); j++ {
				if (*menu).List[i].ChildrenList[j].Url == path {
					(*menu).List[i].Active = "active"
					return menu
				} else {
					(*menu).List[i].Active = ""
				}
			}
		}
	}

	return menu
}

func (menu *Menu) GetEditMenuList() []Item {
	return (*menu).List
}

func GetGlobalMenu(user models.UserModel) *Menu {

	var (
		menus      []map[string]interface{}
		menuOption = make([]map[string]string, 0)
	)

	user.WithRoles().WithMenus()

	if user.IsSuperAdmin() {
		menus, _ = db.Table("goadmin_menu").
			Where("id", ">", 0).
			OrderBy("order", "asc").
			All()
	} else {

		var ids []interface{}
		for i := 0; i < len(user.MenuIds); i++ {
			ids = append(ids, user.MenuIds[i])
		}

		menus, _ = db.Table("goadmin_menu").
			WhereIn("id", ids).
			OrderBy("order", "asc").
			All()
	}

	var title string
	for i := 0; i < len(menus); i++ {

		if menus[i]["type"].(int64) == 1 {
			title = language.Get(menus[i]["title"].(string))
		} else {
			title = menus[i]["title"].(string)
		}

		menuOption = append(menuOption, map[string]string{
			"id":    strconv.FormatInt(menus[i]["id"].(int64), 10),
			"title": title,
		})
	}

	menuList := constructMenuTree(menus, 0)

	return &Menu{
		List:     menuList,
		Options:  menuOption,
		MaxOrder: menus[len(menus)-1]["parent_id"].(int64),
	}
}

func constructMenuTree(menus []map[string]interface{}, parentId int64) []Item {

	branch := make([]Item, 0)

	var title string
	for j := 0; j < len(menus); j++ {
		if parentId == menus[j]["parent_id"].(int64) {

			childList := constructMenuTree(menus, menus[j]["id"].(int64))

			if menus[j]["type"].(int64) == 1 {
				title = language.Get(menus[j]["title"].(string))
			} else {
				title = menus[j]["title"].(string)
			}

			header, _ := menus[j]["header"].(string)

			child := Item{
				Name:         title,
				ID:           strconv.FormatInt(menus[j]["id"].(int64), 10),
				Url:          menus[j]["uri"].(string),
				Icon:         menus[j]["icon"].(string),
				Header:       header,
				Active:       "",
				ChildrenList: childList,
			}

			branch = append(branch, child)
		}
	}

	return branch
}
