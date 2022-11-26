package structuraladapter

import (
	"structural-adapter/internal/biz"
	"testing"
)

func TestAdapter(t *testing.T) {
	goods := make([]biz.Good, 5)
	for k := range goods {
		goods[k] = biz.NewGood(string(rune(k+97)), float64((k+1)*10))
	}
	shop := biz.NewShop(goods...)
	rmb10, rmb20 := NewRmb(10), NewRmb(20)
	if err := shop.Buy(1, NewRmbAdpter(rmb10)); err == nil {
		t.Fatalf(err.Error())
	}
	if err := shop.Buy(1, NewRmbAdpter(rmb20)); err != nil {
		t.Fatalf("do not pass")
	}
}
