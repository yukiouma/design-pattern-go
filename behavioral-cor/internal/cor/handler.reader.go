package cor

type ReaderHandler struct {
	app  *App
	next Handler
}

func NewReaderHandler(app *App) Handler {
	return &ReaderHandler{
		app: app,
	}
}

func (h *ReaderHandler) Handle(name string) bool {
	if h.app.IsReader(name) {
		return true
	}
	if h.next == nil {
		return false
	}
	return h.next.Handle(name)
}

func (h *ReaderHandler) SetNext(handler Handler) {
	h.next = handler
}
