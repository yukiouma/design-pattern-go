package structuralbridge

import (
	"structural-bridge/internal/biz"
	"testing"
)

const (
	path = "."
)

func TestBridge(t *testing.T) {
	windows := biz.WrapWindows(biz.NewWindows())
	winapp := biz.NewApp(windows)
	winapp.LoadPicture(path)
	winapp.EditPicture()
	winapp.SavePicture(path)

	linux := biz.WrapLinux(biz.NewLinux())
	linuxapp := biz.NewApp(linux)
	linuxapp.LoadPicture(path)
	linuxapp.EditPicture()
	linuxapp.SavePicture(path)
}
