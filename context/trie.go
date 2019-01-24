// Copyright 2018 cg33.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package context

import "fmt"

type node struct {
	children []*node
	value    string
	method   string
	handle   Handler
}

func Tree() *node {
	return &node{
		children: make([]*node, 0),
		value:    "/",
		handle:   nil,
	}
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

func (n *node) addPath(paths []string, method string, handler Handler) {
	if len(paths) > 0 {
		child := n.addContent(paths[0])
		if len(paths) > 1 {
			child.addPath(paths[1:], method, handler)
		} else {
			child.method = method
			child.handle = handler
		}
	}
}

func (n *node) findPath(paths []string, method string) Handler {
	if len(paths) > 0 {
		child := n.search(paths[0])
		if child == nil {
			return nil
		} else {
			if len(paths) > 1 {
				return child.findPath(paths[1:], method)
			} else {
				if child.method != method {
					return nil
				} else {
					return child.handle
				}
			}
		}
	}
	return nil
}

func (n *node) print() {
	fmt.Println(n.value)
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
		end        = 0
		iswildcard = false
	)
	for i := 0; i < len(path); i++ {
		if i == 0 && path[0] == '/' {
			start = 1
			continue
		}
		if path[i] == ':' {
			iswildcard = true
		}
		if i == len(path)-1 {
			end = i + 1
			if iswildcard {
				paths = append(paths, "*")
			} else {
				paths = append(paths, path[start:end])
			}
		}
		if path[i] == '/' {
			end = i
			if iswildcard {
				paths = append(paths, "*")
			} else {
				paths = append(paths, path[start:end])
			}
			start = i + 1
			iswildcard = false
		}
	}
	return paths
}
