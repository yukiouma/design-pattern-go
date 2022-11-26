package biz

type Shape interface {
	Accept(Visitor) float64
}
