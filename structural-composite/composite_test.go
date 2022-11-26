package structuralcomposite

import (
	"structural-composite/internal/biz"
	"testing"
)

func TestComposite(t *testing.T) {
	root := biz.NewFolder()
	sub1, sub2 := biz.NewFolder(), biz.NewFolder()
	files := make([]biz.File, 3)
	for i := range files {
		files[i] = biz.NewFile()
	}
	root.Add(files[0])
	sub1.Add(files[1])
	sub2.Add(files[2])
	root.Add(&sub1, &sub2)
	if root.Count() != 3 {
		t.Error("composite did not pass")
	}
}
