package demo

import (
	creationalabstractfactory "creational-abstract-factory"
	"fmt"
)

func AbstractFactory() {
	marketingFactory := creationalabstractfactory.NewFactory(creationalabstractfactory.Marketing)
	general1 := marketingFactory.CreateGeneral("general a")
	director1, err := marketingFactory.CreateDirector(marketingFactory.CreateGeneral("director a"))
	if err != nil {
		panic(err)
	}
	director1.AdjustSalary(general1, 5000)
	fmt.Println("AbstractFactory:")
	fmt.Printf("\tsalary of [%s] is $%5.1f\n", general1.Name(), general1.Salary())
	fmt.Printf("\tsalary of [%s] is $%5.1f\n", director1.Name(), director1.Salary())
}
