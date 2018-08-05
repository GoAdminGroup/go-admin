package hero

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const (
	TypeImport = iota
	TypeDefinition
	TypeExtend
	TypeInclude
	TypeBlock
	TypeCode
	TypeEscapedValue
	TypeRawValue
	TypeNote
	TypeHTML
	TypeRoot
)

const (
	Bool      = "b"
	Int       = "i"
	Uint      = "u"
	Float     = "f"
	String    = "s"
	Bytes     = "bs"
	Interface = "v"
)

const (
	OpenBrace   = '{'
	CloseBrace  = '}'
	LT          = '<'
	GT          = '>'
	Percent     = '%'
	Exclamation = '!'
	Colon       = ':'
	Tilde       = '~'
	Plus        = '+'
	Equal       = '='
	At          = '@'
	Pound       = '#'
	Space       = ' '
	BreakLine   = '\n'
)

var prefixTypeMap = map[byte]uint8{
	Exclamation: TypeImport,
	Colon:       TypeDefinition,
	Tilde:       TypeExtend,
	Plus:        TypeInclude,
	Equal:       TypeEscapedValue,
	At:          TypeBlock,
	Pound:       TypeNote,
}

var (
	openTag      = []byte{LT, Percent} // <%
	closeTag     = []byte{Percent, GT} // %>
	openBraceTag = []byte{OpenBrace}   // {
)

var parsedNodes map[string]*node
var dependencies *sort

func init() {
	cleanGlobal()
}

func cleanGlobal() {
	parsedNodes = make(map[string]*node)
	dependencies = newSort()
}

type node struct {
	t        uint8
	subtype  string
	children []*node
	chunk    *bytes.Buffer
}

func newNode(t uint8, chunk []byte) *node {
	n := &node{
		t:        t,
		children: make([]*node, 0),
		chunk:    new(bytes.Buffer),
	}

	if chunk != nil {
		n.chunk.Write(chunk)
	}

	return n
}

func splitByEndBlock(content []byte) ([]byte, []byte) {
	for i, open := 0, 0; i < len(content); i++ {
		switch content[i] {
		case OpenBrace:
			open++
		case CloseBrace:
			open--
		}

		if open == -1 {
			j := bytes.LastIndex(content[:i], openTag)
			k := bytes.Index(content[i+1:], closeTag)

			if j == -1 || k == -1 ||
				len(bytes.TrimSpace(content[j+2:i+1+k])) != 1 {
				goto Panic
			}
			return content[:j], content[i+k+3:]
		}
	}

Panic:
	panic("invalid endblock")
}

func (n *node) insert(dir, subpath string, content []byte) {
	path, err := filepath.Abs(filepath.Join(dir, subpath))
	if err != nil {
		log.Fatal(err)
	}

	for len(content) > 0 {
		i := bytes.Index(content, openTag)
		if i == -1 {
			i = len(content)
		}

		if i != 0 {
			c := bytes.TrimSpace(content[:i])
			if len(c) > 0 {
				n.children = append(
					n.children,
					newNode(TypeHTML, content[:i]),
				)
			}
			content = content[i:]
			continue
		}

		// starts with "<%"
		content = content[2:]

		i = bytes.Index(content, closeTag)
		if i == -1 {
			log.Fatalf("'<%%' not closed in file `%s`", path)
		}

		switch content[0] {
		case Exclamation, Colon, Pound:
			t, c := prefixTypeMap[content[0]], content[1:i]
			if len(bytes.TrimSpace(c)) > 0 {
				n.children = append(n.children, newNode(t, c))
			}
		case Equal:
			var (
				t       uint8
				subtype string
				c       []byte
			)

			if content[0] == Equal && content[1] == Equal {
				t, c = TypeRawValue, content[2:i]
			} else {
				t, c = TypeEscapedValue, content[1:i]
			}

			parts := bytes.Split(c, []byte{Space})
			if len(parts) > 0 {
				subtype = string(parts[0])
				if subtype == "" || subtype == String {
					subtype = String
				} else if subtype != Int && subtype != Uint &&
					subtype != Float && subtype != Bool &&
					subtype != Bytes && subtype != Interface {
					log.Fatalf("unknown value type %s", subtype)
				}

				c = bytes.TrimSpace(bytes.Join(parts[1:], []byte{Space}))
				if len(c) > 0 {
					child := newNode(t, bytes.TrimSpace(c))
					child.subtype = subtype

					n.children = append(n.children, child)
					goto ResetContent
				}
			}

			log.Fatal("lack of variable name")
		case Tilde, Plus:
			c := bytes.TrimSpace(content[1:i])

			parent := string(c[1 : len(c)-1])
			if !filepath.IsAbs(parent) {
				if parent, err = filepath.Abs(
					filepath.Join(dir, parent)); err != nil {
					log.Fatal(err)
				}
			}

			n.children = append(
				n.children,
				newNode(prefixTypeMap[content[0]], []byte(parent)),
			)

			dependencies.addVertices(parent, path)
			dependencies.addEdge(parent, path)
		case At:
			chunk := bytes.TrimSpace(content[1:i])
			if !bytes.HasSuffix(chunk, openBraceTag) {
				log.Fatalf("block not ended with `{` in file `%s`", path)
			}

			child := newNode(
				TypeBlock,
				bytes.TrimSpace(chunk[:len(chunk)-1]),
			)

			blockName := child.chunk.String()
			if b := n.findBlockByName(blockName); b != nil {
				log.Fatalf("duplicate block %s in file `%s`", blockName, path)
			}

			var childContent []byte

			childContent, content = splitByEndBlock(content[i+2:])
			child.insert(dir, subpath, childContent)
			n.children = append(n.children, child)
			continue
		default:
			n.children = append(
				n.children, newNode(TypeCode, content[:i]),
			)
		}

	ResetContent:
		content = content[i+2:]
	}
}

func (n *node) childrenByType(t uint8) []*node {
	var children []*node

	for _, child := range n.children {
		if child.t == t {
			children = append(children, child)
		}
	}

	return children
}

func (n *node) findBlockByName(name string) *node {
	for _, child := range n.children {
		if child.t == TypeBlock && child.chunk.String() == name {
			return child
		}
	}
	return nil
}

func (n *node) rebuild() {
	var pNode *node

	nodes := n.childrenByType(TypeExtend)
	if len(nodes) > 0 {
		pNode = parsedNodes[nodes[0].chunk.String()]
	}

	if pNode != nil {
		var children []*node

		for _, t := range []uint8{TypeImport, TypeDefinition} {
			children = append(children, n.childrenByType(t)...)
		}

		for _, child := range pNode.children {
			switch child.t {
			case TypeHTML, TypeCode, TypeEscapedValue,
				TypeRawValue, TypeInclude:
				children = append(children, child)
			case TypeBlock:
				block := n.findBlockByName(child.chunk.String())
				if block != nil {
					block.rebuild()
					children = append(children, block)
				}
			}
		}

		n.children = children
		return
	}

	newChildren := make([]*node, 0)
	for _, child := range n.children {
		switch child.t {
		case TypeInclude:
			key := child.chunk.String()
			includeNode, ok := parsedNodes[key]
			if !ok {
				var keys []string
				for k := range parsedNodes {
					keys = append(keys, k)
				}
				log.Fatalf("node \"%s\" not found. have: %v", key, keys)
			}

			for _, t := range []uint8{
				TypeExtend,
				TypeInclude,
				TypeDefinition} {
				if len(includeNode.childrenByType(t)) != 0 {
					log.Fatalf(
						"there shouldn't be any of statement of Extend, "+
							"Include or Definition in sub-template %s",
						key,
					)
				}
			}

			for _, childNode := range includeNode.children {
				if childNode.t == TypeImport {
					newChildren = append([]*node{childNode}, newChildren...)
				} else {
					newChildren = append(newChildren, childNode)
				}
			}
		case TypeBlock:
			child.rebuild()
			fallthrough
		default:
			newChildren = append(newChildren, child)
		}
	}

	n.children = newChildren
}

func CheckExtension(path string, extensions []string) bool {
	extension := filepath.Ext(path)
	for _, item := range extensions {
		if item == extension {
			return true
		}
	}
	return false
}

func parseFile(dir, subpath string) *node {
	path, err := filepath.Abs(filepath.Join(dir, subpath))
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// add dependency.
	dependencies.addVertices(path)

	root := newNode(TypeRoot, nil)
	root.insert(dir, subpath, content)

	for _, t := range []uint8{TypeExtend, TypeDefinition} {
		children := root.childrenByType(t)
		if len(children) > 1 {
			log.Fatalf(
				"there should be at most one Extend or Definition in file %s",
				path,
			)
		}
	}
	parsedNodes[path] = root

	return root
}

func parseDir(dir string, extensions []string) {
	err := filepath.Walk(dir, func(
		path string, info os.FileInfo, err error) error {

		stat, err := os.Stat(path)
		if err != nil {
			return err
		}

		if !stat.IsDir() && CheckExtension(path, extensions) {
			parseFile(dir, path[len(dir):])
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}

func rebuild() {
	queue := dependencies.sort()
	for _, path := range queue {
		if _, err := os.Stat(path); err == nil {
			continue
		} else if os.IsNotExist(err) {
			log.Fatal(path, " not found")
		} else {
			log.Fatal(err)
		}
	}

	for _, path := range queue {
		if node, ok := parsedNodes[path]; !ok {
			log.Fatalf("dependency %s not parsed", path)
		} else {
			node.rebuild()
		}
	}
}
