package biz

import "math"

type Visitor interface {
	CaculateCircleArea(*circle) float64
	CaculateSquareArea(*square) float64
	CaculateRectangleArea(*rectangle) float64
}

type areaVisitor struct{}

func NewAreaVistor() Visitor {
	return &areaVisitor{}
}

func (v *areaVisitor) CaculateCircleArea(s *circle) float64 {
	return math.Pow(s.Radium(), 2) * math.Pi
}

func (v *areaVisitor) CaculateSquareArea(s *square) float64 {
	return math.Pow(s.Side(), 2)
}

func (v *areaVisitor) CaculateRectangleArea(s *rectangle) float64 {
	return s.Width() * s.Length()
}
