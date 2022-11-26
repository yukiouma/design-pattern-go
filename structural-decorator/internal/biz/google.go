package biz

import "fmt"

type Google struct {
	origin Sender
}

func NewGoogle(origin Sender) Sender {
	return &Google{origin: origin}
}

func (t *Google) Send() string {
	return fmt.Sprintf("GOOGLE: forwarding message: %s", t.origin.Send())
}
