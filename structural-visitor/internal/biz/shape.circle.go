package biz

type circle struct {
	radium float64
}

func NewCircle(radium float64) Shape {
	return &circle{
		radium: radium,
	}
}

func (s *circle) Accept(v Visitor) float64 {
	return v.CaculateCircleArea(s)
}

func (s *circle) Radium() float64 {
	return s.radium
}
