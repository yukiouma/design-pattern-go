package behavioralcor

import (
	"behavioral-cor/internal/cor"
	"testing"
)

const (
	a = "a"
	b = "b"
	c = "c"
	d = "d"
)

func TestReaderReadFile(t *testing.T) {
	app := cor.NewApp()
	app.Register(a, cor.Reader)
	app.Register(b, cor.Appender)
	app.Register(c, cor.Updater)
	app.Register(d, cor.Deleter)

	guard := NewGuard(app)

	if !guard.CanReadFile(a) {
		t.Fatal("a can read file")
	}

	if guard.CanAppendFile(a) {
		t.Fatal("a can not append file")
	}

	if guard.CanUpdateFile(b) {
		t.Fatal("a can not update file")
	}

	if !guard.CanUpdateFile(c) {
		t.Fatal("c can update file")
	}

	if guard.CanDeleteFile(c) {
		t.Fatal("c can not delete file")
	}

	if !guard.CanDeleteFile(d) {
		t.Fatal("d can delete file")
	}
}

// func NewHandler(next *Handler) *Handler {
// 	return &Handler{
// 		next: next,
// 	}
// }

// type Handler struct {
// 	next *Handler
// }

// func (h *Handler) Handle(param interface{}) interface{} {
// 	var result interface{}
// 	if result {
// 		return h.next.Handle(param)
// 	}
// 	return result
// }
