package cor

type UpdaterHandler struct {
	app  *App
	next Handler
}

func NewUpdaterHandler(app *App) Handler {
	return &UpdaterHandler{
		app: app,
	}
}

func (h *UpdaterHandler) Handle(name string) bool {
	if h.app.IsUpdater(name) {
		return true
	}
	if h.next == nil {
		return false
	}
	return h.next.Handle(name)
}

func (h *UpdaterHandler) SetNext(handler Handler) {
	h.next = handler
}
