package structuraladapter

import "structural-adapter/internal/biz"

type RmbAdpter struct {
	rmb Rmb
}

func NewRmbAdpter(r Rmb) biz.Money {
	return &RmbAdpter{rmb: r}
}

var _ biz.Money = (*RmbAdpter)(nil)

func (r *RmbAdpter) Value() float64 {
	return r.rmb.Denomination()
}
