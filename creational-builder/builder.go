package creationalbuilder

import "creational-builder/internal/biz"

func NewDirector() *biz.Director {
	return biz.NewDirector()
}

func NewChickenPizzaBuilder() biz.Builder {
	return biz.NewChickenPizzaBuilder()
}

func NewDurianPizzaBuilder() biz.Builder {
	return biz.NewDurianPizzaBuilder()
}
