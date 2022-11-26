package biz

type rectangle struct {
	length, width float64
}

func NewRectangle(length, width float64) Shape {
	return &rectangle{
		length: length,
		width:  width,
	}
}

func (s *rectangle) Accept(v Visitor) float64 {
	return v.CaculateRectangleArea(s)
}

func (s *rectangle) Width() float64 {
	return s.width
}

func (s *rectangle) Length() float64 {
	return s.length
}
