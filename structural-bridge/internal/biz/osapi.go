package biz

type WrapOSAPI interface {
	Read(path string) []byte
	Write(path string, data []byte)
}
