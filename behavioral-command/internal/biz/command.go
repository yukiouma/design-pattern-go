package biz

type Command interface {
	Execute()
	Undo()
}
