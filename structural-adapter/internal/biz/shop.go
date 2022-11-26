package biz

import "errors"

var (
	ErrInvalidGood = errors.New("error: invalid good")
	ErrNotAfford   = errors.New("error: can not afford")
)

type Shop struct {
	goods []Good
}

func NewShop(goods ...Good) Shop {
	return Shop{goods: goods}
}

func (s Shop) Buy(id int, pay Money) error {
	if id >= len(s.goods) {
		return ErrInvalidGood
	}
	g := s.goods[id]
	if g.Price() > pay.Value() {
		return ErrNotAfford
	}
	return nil
}
