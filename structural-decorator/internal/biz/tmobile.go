package biz

type TMobile struct{}

func NewTMobile() Sender {
	return &TMobile{}
}

func (t *TMobile) Send() string {
	return "TMOBILE: hello"
}
