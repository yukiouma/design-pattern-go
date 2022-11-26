package biz

type App struct {
	osapi   WrapOSAPI
	picture []byte
}

func NewApp(osapi WrapOSAPI) App {
	return App{osapi: osapi}
}

func (a *App) LoadPicture(path string) {
	a.picture = a.osapi.Read(path)
}

func (a *App) EditPicture() {}

func (a *App) SavePicture(path string) {
	a.osapi.Write(path, a.picture)
}
