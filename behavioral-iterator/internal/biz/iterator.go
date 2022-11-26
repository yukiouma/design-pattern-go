package biz

type Iterator interface {
	Next() int
	HasNext() bool
}
