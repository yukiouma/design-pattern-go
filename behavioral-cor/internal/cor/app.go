package cor

type App struct {
	state *State
}

func NewApp() *App {
	return &App{
		state: NewState(),
	}
}

func (a *App) Register(username string, role Role) {
	a.state.Register(username, role)
}

func (a *App) IsReader(username string) bool {
	return a.state.Findout(username, Reader)
}

func (a *App) IsAppender(username string) bool {
	return a.state.Findout(username, Appender)
}

func (a *App) IsUpdater(username string) bool {
	return a.state.Findout(username, Updater)
}

func (a *App) IsDeleter(username string) bool {
	return a.state.Findout(username, Deleter)
}
