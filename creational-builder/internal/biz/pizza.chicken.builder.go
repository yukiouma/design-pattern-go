package biz

type ChickenPizzaBuilder struct {
	bakeMin     float64
	pizza       *ChickenPizza
	needBakeMin float64
}

var _ Builder = (*ChickenPizzaBuilder)(nil)

func NewChickenPizzaBuilder() Builder {
	b := new(ChickenPizzaBuilder)
	b.pizza = &ChickenPizza{}
	return b
}

func (e *ChickenPizzaBuilder) useThinCrust(thinCrust bool) {
	e.pizza.thinCrust = thinCrust
}

func (e *ChickenPizzaBuilder) setBakeTime() {
	if e.pizza.thinCrust {
		e.needBakeMin = 7
	} else {
		e.needBakeMin = 5
	}
}

func (e *ChickenPizzaBuilder) addFilling() {
	e.pizza.filling = []string{"chicken", "potato"}
}

func (e *ChickenPizzaBuilder) bake(min float64) {
	e.bakeMin += min
}

func (e *ChickenPizzaBuilder) addCheess() {
	e.pizza.cheess = true
}

func (e *ChickenPizzaBuilder) addTomatoSauce() {
	e.pizza.tomatoSauce = true
}

func (e *ChickenPizzaBuilder) GetPizza() (Pizza, error) {
	if e.bakeMin < e.needBakeMin {
		return nil, ErrPizzaNotReady
	}
	return e.pizza, nil
}
