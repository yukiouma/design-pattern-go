package biz

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Insert(t *Node, v int) *Node {
	if t == nil {
		return &Node{Value: v}
	}
	if t.Value > v {
		t.Left = Insert(t.Left, v)
	} else {
		t.Right = Insert(t.Right, v)
	}
	return t
}

func NewTree(data []int) *Node {
	var t *Node
	for _, v := range data {
		t = Insert(t, v)
	}
	return t
}
