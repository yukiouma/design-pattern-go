package biz

import "fmt"

type Apple struct {
	origin Sender
}

func NewApple(origin Sender) Sender {
	return &Apple{origin: origin}
}

func (t *Apple) Send() string {
	return fmt.Sprintf("APPLE: forwarding message: %s", t.origin.Send())
}
