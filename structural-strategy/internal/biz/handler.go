package biz

type handler struct {
	strategy Strategy
}

func NewMathHandler() handler {
	return handler{}
}

func (h *handler) SetStrategy(strategy Strategy) {
	h.strategy = strategy
}

func (h *handler) Execute(x, y int) int {
	return h.strategy.Execute(x, y)
}
