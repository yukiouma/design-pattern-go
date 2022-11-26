package biz

type cleanButoon struct {
	mediator Mediator
}

func NewClearButton(mediator Mediator) *cleanButoon {
	return &cleanButoon{
		mediator: mediator,
	}
}

func (c *cleanButoon) Clean() error {
	return c.mediator.Clean()
}
