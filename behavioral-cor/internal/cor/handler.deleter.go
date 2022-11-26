package cor

type DeleterHandler struct {
	app  *App
	next Handler
}

func NewDeleterHandler(app *App) Handler {
	return &DeleterHandler{
		app: app,
	}
}

func (h *DeleterHandler) Handle(name string) bool {
	if h.app.IsDeleter(name) {
		return true
	}
	if h.next == nil {
		return false
	}
	return h.next.Handle(name)
}

func (h *DeleterHandler) SetNext(handler Handler) {
	h.next = handler
}
