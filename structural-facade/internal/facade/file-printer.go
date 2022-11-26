package facade

type FilePrinter interface {
	PrintFile(dir string) error
}
