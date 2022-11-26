package biz

type BFSIterator struct {
	element []int
}

func NewBFSIterator(tree *Node) Iterator {
	iter := &BFSIterator{
		element: make([]int, 0),
	}
	if tree == nil {
		return iter
	}

	// bfs
	queue := make([]*Node, 0)
	queue = append(queue, tree)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		iter.element = append(iter.element, node.Value)
		if l := node.Left; l != nil {
			queue = append(queue, l)
		}
		if r := node.Right; r != nil {
			queue = append(queue, r)
		}
	}

	return iter
}

func (t *BFSIterator) Next() int {
	list := t.element
	result := list[0]
	list = list[1:]
	t.element = list
	return result
}
func (t *BFSIterator) HasNext() bool {
	return len(t.element) != 0
}
