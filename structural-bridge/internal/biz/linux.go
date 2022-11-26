package biz

type Linux struct {
}

func NewLinux() Linux {
	return Linux{}
}

func (w Linux) Read(path string) []byte {
	return make([]byte, 0)
}

func (w Linux) Write(path string, data []byte) {}
