package biz

import "fmt"

type File struct {
	name string
}

func NewFile(s string) Node {
	return &File{name: s}
}

var _ Node = (*File)(nil)

func (e *File) Name() string {
	return e.name
}

func (e *File) Clone() Node {
	copy := *e
	return &copy
}

func (e *File) SetName(s string) {
	e.name = s
}

func (e *File) Print() {
	fmt.Println(e.Name())
}
