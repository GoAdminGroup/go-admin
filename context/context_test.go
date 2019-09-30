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
	tree := Tree()
	tree.addPath([]string{}, "GET", []Handler{})
	tree.printChildren()
}
