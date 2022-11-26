package creationalprototype

import (
	"creational-prototype/internal/biz"
	"testing"
)

func TestPrototype(t *testing.T) {
	// check deep copy of file
	file1 := biz.NewFile("demo1")
	file2 := file1.Clone()
	file2.SetName("demo2")
	if file1.Name() == file2.Name() {
		t.Error("file clone does not pass")
		return
	}

	dir1 := biz.NewFolder("dir1", file1, file2)
	dir2 := dir1.Clone()
	dir1.Print()
	dir2.Print()
}
