package biz

import "fmt"

type Folder struct {
	name     string
	children []Node
}

func NewFolder(s string, node ...Node) Node {
	return &Folder{name: s, children: node}
}

var _ Node = (*Folder)(nil)

func (e *Folder) Name() string {
	return e.name
}

func (e *Folder) Clone() Node {
	copy := *e
	copyChildren := make([]Node, 0, len(e.children))
	for _, v := range e.children {
		copyChildren = append(copyChildren, v.Clone())
	}
	copy.children = copyChildren
	return &copy
}

func (e *Folder) SetName(s string) {
	e.name = s
}

func (e *Folder) Print() {
	fmt.Println(e.Name())
	for _, v := range e.children {
		v.Print()
	}
}
