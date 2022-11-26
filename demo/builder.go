package demo

import (
	creationalbuilder "creational-builder"
	"fmt"
)

func Builder() {
	fmt.Println("Builder:")
	durianBuiler := creationalbuilder.NewDurianPizzaBuilder()
	chickenBuilder := creationalbuilder.NewChickenPizzaBuilder()
	director := creationalbuilder.NewDirector()

	director.SetBuilder(chickenBuilder)
	director.MakeChickenPizza()
	pizza, err := chickenBuilder.GetPizza()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\t%#v\n", pizza)

	director.ResetBuilder()

	director.SetBuilder(durianBuiler)
	director.MakeDurianPizza()
	pizza, err = durianBuiler.GetPizza()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\t%#v\n", pizza)
}
