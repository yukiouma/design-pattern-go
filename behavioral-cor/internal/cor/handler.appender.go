package cor

type AppenderHandler struct {
	app  *App
	next Handler
}

func NewAppenderHandler(app *App) Handler {
	return &AppenderHandler{
		app: app,
	}
}

func (h *AppenderHandler) Handle(name string) bool {
	if h.app.IsAppender(name) {
		return true
	}
	if h.next == nil {
		return false
	}
	return h.next.Handle(name)
}

func (h *AppenderHandler) SetNext(handler Handler) {
	h.next = handler
}
