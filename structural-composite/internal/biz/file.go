package biz

type File struct{}

var _ Node = (*File)(nil)

func NewFile() File {
	return File{}
}

func (f File) Count() int {
	return 1
}
