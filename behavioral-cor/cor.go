package behavioralcor

import "behavioral-cor/internal/cor"

type guard struct {
	app *cor.App
}

func NewGuard(app *cor.App) *guard {
	return &guard{
		app: app,
	}
}

func (g *guard) CanReadFile(username string) bool {
	readHandler := cor.NewReaderHandler(g.app)
	appebdHandler := cor.NewAppenderHandler(g.app)
	updateHandler := cor.NewUpdaterHandler(g.app)
	deleteHandler := cor.NewDeleterHandler(g.app)

	appebdHandler.SetNext(readHandler)
	updateHandler.SetNext(appebdHandler)
	deleteHandler.SetNext(updateHandler)

	return deleteHandler.Handle(username)
}

func (g *guard) CanAppendFile(username string) bool {
	appebdHandler := cor.NewAppenderHandler(g.app)
	updateHandler := cor.NewUpdaterHandler(g.app)
	deleteHandler := cor.NewDeleterHandler(g.app)

	updateHandler.SetNext(appebdHandler)
	deleteHandler.SetNext(updateHandler)

	return deleteHandler.Handle(username)
}

func (g *guard) CanUpdateFile(username string) bool {
	updateHandler := cor.NewUpdaterHandler(g.app)
	deleteHandler := cor.NewDeleterHandler(g.app)

	deleteHandler.SetNext(updateHandler)

	return deleteHandler.Handle(username)
}

func (g *guard) CanDeleteFile(username string) bool {
	deleteHandler := cor.NewDeleterHandler(g.app)

	return deleteHandler.Handle(username)
}
