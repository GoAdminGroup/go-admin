// Copyright 2019 GoAdmin Core Team. All rights reserved.
// Use of this source code is governed by a Apache-2.0 style
// license that can be found in the LICENSE file.

package context

import "fmt"

type node struct {
	children []*node
	value    string
	method   []string
	handle   [][]Handler
}

func tree() *node {
	return &node{
		children: make([]*node, 0),
		value:    "/",
		handle:   nil,
	}
}

func (n *node) hasMethod(method string) int {
	for k, m := range n.method {
		if m == method {
			return k
		}
	}
	return -1
}

func (n *node) addMethodAndHandler(method string, handler []Handler) {
	n.method = append(n.method, method)
	n.handle = append(n.handle, handler)
}

func (n *node) addChild(child *node) {
	n.children = append(n.children, child)
}

func (n *node) addContent(value string) *node {
	var child = n.search(value)
	if child == nil {
		child = &node{
			children: make([]*node, 0),
			value:    value,
		}
		n.addChild(child)
	}
	return child
}

func (n *node) search(value string) *node {
	for _, child := range n.children {
		if child.value == value || child.value == "*" {
			return child
		}
	}
	return nil
}

func (n *node) addPath(paths []string, method string, handler []Handler) {
	child := n
	for i := 0; i < len(paths); i++ {
		child = child.addContent(paths[i])
	}
	child.addMethodAndHandler(method, handler)
}

func (n *node) findPath(paths []string, method string) []Handler {
	child := n
	for i := 0; i < len(paths); i++ {
		child = child.search(paths[i])
		if child == nil {
			return nil
		}
	}

	methodIndex := child.hasMethod(method)
	if methodIndex == -1 {
		return nil
	}

	return child.handle[methodIndex]
}

func (n *node) print() {
	fmt.Println(n)
}

func (n *node) printChildren() {
	n.print()
	for _, child := range n.children {
		child.printChildren()
	}
}

func stringToArr(path string) []string {
	var (
		paths      = make([]string, 0)
		start      = 0
		end        int
		isWildcard = false
	)
	for i := 0; i < len(path); i++ {
		if i == 0 && path[0] == '/' {
			start = 1
			continue
		}
		if path[i] == ':' {
			isWildcard = true
		}
		if i == len(path)-1 {
			end = i + 1
			if isWildcard {
				paths = append(paths, "*")
			} else {
				paths = append(paths, path[start:end])
			}
		}
		if path[i] == '/' {
			end = i
			if isWildcard {
				paths = append(paths, "*")
			} else {
				paths = append(paths, path[start:end])
			}
			start = i + 1
			isWildcard = false
		}
	}
	return paths
}
