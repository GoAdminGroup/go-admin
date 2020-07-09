package menu

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestMenu_AddMaxOrder(t *testing.T) {
	menus := Menu{
		MaxOrder: 0,
	}
	menus.AddMaxOrder()
	assert.Equal(t, menus.MaxOrder, int64(1))
}

func TestMenu_SetMaxOrder(t *testing.T) {
	menus := Menu{
		MaxOrder: 0,
	}
	menus.SetMaxOrder(2)
	assert.Equal(t, menus.MaxOrder, int64(2))
}

func TestMenu_SetActiveClass(t *testing.T) {
	menus := Menu{
		List: []Item{
			{
				Name: "item1",
				ID:   "1",
				Url:  "/item1",
				Icon: "icon",
			}, {
				Name: "item2",
				ID:   "2",
				Url:  "/item2",
				Icon: "icon",
			}, {
				Name: "item3",
				ID:   "3",
				Url:  "/item3",
				Icon: "icon",
			}, {
				Name: "item4",
				ID:   "4",
				Url:  "/item4",
				Icon: "icon",
				ChildrenList: []Item{
					{
						Name: "item5",
						ID:   "5",
						Url:  "/item5",
						Icon: "icon",
					}, {
						Name: "item6",
						ID:   "6",
						Url:  "/item6",
						Icon: "icon",
					},
				},
			},
		},
		Options:  []map[string]string{},
		MaxOrder: 0,
	}

	menus.SetActiveClass("/item3")

	assert.Equal(t, menus.List[0].Active, "")
	assert.Equal(t, menus.List[1].Active, "")
	assert.Equal(t, menus.List[2].Active, "active")
	assert.Equal(t, menus.List[3].Active, "")

	menus.SetActiveClass("/item5")

	assert.Equal(t, menus.List[0].Active, "")
	assert.Equal(t, menus.List[1].Active, "")
	assert.Equal(t, menus.List[2].Active, "")
	assert.Equal(t, menus.List[3].Active, "active")
	assert.Equal(t, menus.List[3].ChildrenList[0].Active, "active")
	assert.Equal(t, menus.List[3].ChildrenList[1].Active, "")
}
