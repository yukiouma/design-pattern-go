package biz

var (
	AddStrategy, SubstractStrategy, MultipleStrategy, DivideStrategy Strategy
)

func init() {
	AddStrategy = newAddStrategy()
	SubstractStrategy = newSubstractStrategy()
	MultipleStrategy = newMultipleStrategy()
	DivideStrategy = newDivideStrategy()
}

type Strategy interface {
	Execute(x, y int) int
}

type addStrategy struct{}

func newAddStrategy() Strategy {
	return &addStrategy{}
}

func (s *addStrategy) Execute(x, y int) int {
	return x + y
}

type substractStrategy struct{}

func newSubstractStrategy() Strategy {
	return &substractStrategy{}
}

func (s *substractStrategy) Execute(x, y int) int {
	return x - y
}

type mutipleStrategy struct{}

func newMultipleStrategy() Strategy {
	return &mutipleStrategy{}
}

func (s *mutipleStrategy) Execute(x, y int) int {
	return x * y
}

type divideStrategy struct{}

func newDivideStrategy() Strategy {
	return &divideStrategy{}
}

func (s *divideStrategy) Execute(x, y int) int {
	return x / y
}
