package context

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSlash(t *testing.T) {
	assert.Equal(t, "/abc", slash("/abc"))
	assert.Equal(t, "/", slash(""))
	assert.Equal(t, "/abc", slash("abc/"))
	assert.Equal(t, "/", slash("/"))
	assert.Equal(t, "/abc", slash("/abc/"))
	assert.Equal(t, "/", slash("//"))
}

func TestJoin(t *testing.T) {
	assert.Equal(t, "/abc/abc", join(slash("/abc"), slash("/abc")))
	assert.Equal(t, "/", join(slash("/"), slash("/")))
	assert.Equal(t, "/abc", join(slash("/"), slash("/abc")))
	assert.Equal(t, "/abc", join(slash("abc/"), slash("/")))
	assert.Equal(t, "/abc", join(slash("/abc/"), slash("/")))
}

func TestTree(t *testing.T) {
	tree := tree()
	tree.addPath(stringToArr("/adm"), "GET", []Handler{func(ctx *Context) { fmt.Println(1) }})
	tree.addPath(stringToArr("/admi"), "GET", []Handler{func(ctx *Context) { fmt.Println(1) }})
	tree.addPath(stringToArr("/admin"), "GET", []Handler{func(ctx *Context) { fmt.Println(1) }})
	tree.addPath(stringToArr("/admin/menu/new"), "POST", []Handler{func(ctx *Context) { fmt.Println(1) }})
	tree.addPath(stringToArr("/admin/menu/new"), "GET", []Handler{func(ctx *Context) { fmt.Println(1) }})
	tree.addPath(stringToArr("/admin/info/:__prefix"), "GET", []Handler{
		func(ctx *Context) { fmt.Println("auth") },
		func(ctx *Context) { fmt.Println("init") },
		func(ctx *Context) { fmt.Println("info") },
	})
	tree.addPath(stringToArr("/admin/info/:__prefix/detail"), "GET", []Handler{
		func(ctx *Context) { fmt.Println("auth") },
		func(ctx *Context) { fmt.Println("detail") },
	})

	fmt.Println("/admin/menu/new", "GET")
	h := tree.findPath(stringToArr("/admin/menu/new"), "GET")
	assert.Equal(t, h != nil, true)
	printHandler(h)
	fmt.Println("/admin/menu/new", "POST")
	h = tree.findPath(stringToArr("/admin/menu/new"), "POST")
	assert.Equal(t, h != nil, true)
	printHandler(h)
	fmt.Println("/admin/me/new", "POST")
	h = tree.findPath(stringToArr("/admin/me/new"), "POST")
	assert.Equal(t, h == nil, true)
	printHandler(h)
	fmt.Println("/admin/info/user", "GET")
	h = tree.findPath(stringToArr("/admin/info/user"), "GET")
	assert.Equal(t, h != nil, true)
	printHandler(h)
	fmt.Println("/admin/info/user/detail", "GET")
	h = tree.findPath(stringToArr("/admin/info/user/detail"), "GET")
	assert.Equal(t, h != nil, true)
	printHandler(h)
	fmt.Println("=========== printChildren ===========")
	tree.printChildren()
}

func printHandler(h []Handler) {
	for _, value := range h {
		value(&Context{})
	}
}
