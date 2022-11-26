package structuralfacade

import (
	"structural-facade/internal/facade"
	"testing"
)

func TestFacadeFilePrinter(t *testing.T) {
	var dir string
	printer := facade.NewIoFilePrinter()
	dir = "/root/playground/golang/design-pattern/structural-facade/go.mod"
	if err := printer.PrintFile(dir); err != nil {
		t.Fatalf("failed to print file because: %v", err)
	}
	dir = "/root/playground/golang/design-pattern/structural-facade/go.mod1"
	if err := printer.PrintFile(dir); err == nil {
		t.Fatalf("can not print unexisted file")
	}
}
