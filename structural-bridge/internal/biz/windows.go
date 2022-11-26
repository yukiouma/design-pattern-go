package biz

type Windows struct {
}

func NewWindows() Windows {
	return Windows{}
}

func (w Windows) Open(path string) []byte {
	return make([]byte, 0)
}

func (w Windows) Save(path string, data []byte) {}
