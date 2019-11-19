package context

import (
	"github.com/magiconair/properties/assert"
	"testing"
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
	tree.addPath(stringToArr("/adm"), "GET", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admi"), "GET", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin"), "GET", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin/info/menu"), "GET", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin/info/me"), "GET", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin/info/mefr"), "GET", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin/info/user"), "POST", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin/menu/new"), "POST", []Handler{func(ctx *Context) {}})
	tree.addPath(stringToArr("/admin/menu/new"), "GET", []Handler{func(ctx *Context) {}})
	assert.Equal(t, tree.findPath(stringToArr("/admin/menu/new"), "GET") != nil, true)
	assert.Equal(t, tree.findPath(stringToArr("/admin/menu/new"), "POST") != nil, true)
	assert.Equal(t, tree.findPath(stringToArr("/admin/me/new"), "POST") == nil, true)
	tree.printChildren()
}
