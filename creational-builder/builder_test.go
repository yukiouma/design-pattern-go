package creationalbuilder

import (
	"fmt"
	"testing"
)

func TestBuiler(t *testing.T) {
	durianBuiler := NewDurianPizzaBuilder()
	chickenBuilder := NewChickenPizzaBuilder()
	director := NewDirector()

	director.SetBuilder(chickenBuilder)
	director.MakeChickenPizza()
	pizza, err := chickenBuilder.GetPizza()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", pizza)

	director.ResetBuilder()

	director.SetBuilder(durianBuiler)
	director.MakeDurianPizza()
	pizza, err = durianBuiler.GetPizza()
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Printf("%#v\n", pizza)
}
