package biz

type Good struct {
	name  string
	price float64
}

func NewGood(name string, price float64) Good {
	return Good{name: name, price: price}
}

func (g Good) Name() string {
	return g.name
}

func (g Good) Price() float64 {
	return g.price
}
