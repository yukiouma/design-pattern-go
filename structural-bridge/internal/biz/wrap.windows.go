package biz

type wrapWindows struct {
	os Windows
}

var _ WrapOSAPI = (*wrapWindows)(nil)

func WrapWindows(app Windows) WrapOSAPI {
	return &wrapWindows{os: app}
}

func (a *wrapWindows) Read(path string) []byte {
	return a.os.Open(path)
}

func (a *wrapWindows) Write(path string, data []byte) {
	a.os.Save(path, data)
}
