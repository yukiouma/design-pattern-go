package biz

type Node interface {
	SetName(string)
	Name() string
	Clone() Node
	Print()
}
