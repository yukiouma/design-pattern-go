package biz

type square struct {
	side float64
}

func NewSquare(side float64) Shape {
	return &square{
		side: side,
	}
}

func (s *square) Accept(v Visitor) float64 {
	return v.CaculateSquareArea(s)
}

func (s *square) Side() float64 {
	return s.side
}
