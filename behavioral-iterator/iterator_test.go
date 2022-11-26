package behavioraliterator

import (
	"behavioral-iterator/internal/biz"
	"testing"
)

func TestIterator(t *testing.T) {
	sli := []int{5, 3, 7, 2, 4, 6, 9, 1, 8}
	tree := biz.NewTree(sli)
	bfsIterator := biz.NewBFSIterator(tree)
	dfsIterator := biz.NewDFSIterator(tree)

	bfsResult := make([]int, 0, len(sli))
	for bfsIterator.HasNext() {
		bfsResult = append(bfsResult, bfsIterator.Next())
	}

	bfsResultExpect := []int{5, 3, 7, 2, 4, 6, 9, 1, 8}
	for i, v := range bfsResultExpect {
		if v != bfsResult[i] {
			t.Error("error in bfs")
			return
		}
	}

	dfsResult := make([]int, 0, len(sli))
	for dfsIterator.HasNext() {
		dfsResult = append(dfsResult, dfsIterator.Next())
	}

	dfsResultExpect := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range dfsResultExpect {
		if v != dfsResult[i] {
			t.Error("error in dfs")
			return
		}
	}

}
