package biz

type wrapLinux struct {
	os Linux
}

var _ WrapOSAPI = (*wrapLinux)(nil)

func WrapLinux(app Linux) WrapOSAPI {
	return &wrapLinux{os: app}
}

func (a *wrapLinux) Read(path string) []byte {
	return a.os.Read(path)
}

func (a *wrapLinux) Write(path string, data []byte) {
	a.os.Write(path, data)
}
