package biz

type DFSIterator struct {
	element []int
}

func NewDFSIterator(tree *Node) Iterator {
	element := make([]int, 0)

	// dfs
	var dfs func(*Node)
	dfs = func(t *Node) {
		if t == nil {
			return
		}
		dfs(t.Left)
		element = append(element, t.Value)
		dfs(t.Right)
	}
	dfs(tree)

	return &DFSIterator{
		element: element,
	}
}

func (t *DFSIterator) Next() int {
	list := t.element
	result := list[0]
	list = list[1:]
	t.element = list
	return result
}
func (t *DFSIterator) HasNext() bool {
	return len(t.element) != 0
}
