package facade

import (
	"fmt"
	"io"
	"os"
)

type ioPrinter struct{}

var _ FilePrinter = (*ioPrinter)(nil)

func NewIoFilePrinter() FilePrinter {
	return &ioPrinter{}
}

func (p *ioPrinter) PrintFile(dir string) error {
	f, err := os.Open(dir)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	fmt.Printf("reading file: %s\n", dir)
	fmt.Println(string(data))
	return nil
}
