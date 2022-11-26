package biz

type Folder struct {
	child []Node
}

var _ Node = (*Folder)(nil)

func NewFolder() Folder {
	return Folder{child: make([]Node, 0)}
}

func (f *Folder) Count() int {
	result := 0
	for _, v := range f.child {
		result += v.Count()
	}
	return result
}

func (f *Folder) Add(node ...Node) {
	f.child = append(f.child, node...)
}
