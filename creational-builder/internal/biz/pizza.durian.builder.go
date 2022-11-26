package biz

type DurianPizzaBuilder struct {
	bakeMin     float64
	pizza       *DurianPizza
	needBakeMin float64
}

var _ Builder = (*DurianPizzaBuilder)(nil)

func NewDurianPizzaBuilder() Builder {
	b := new(DurianPizzaBuilder)
	b.pizza = &DurianPizza{}
	return b
}

func (e *DurianPizzaBuilder) useThinCrust(thinCrust bool) {
	e.pizza.thinCrust = thinCrust
}

func (e *DurianPizzaBuilder) setBakeTime() {
	if e.pizza.thinCrust {
		e.needBakeMin = 7
	} else {
		e.needBakeMin = 5
	}
}

func (e *DurianPizzaBuilder) addFilling() {
	e.pizza.filling = []string{"durian", "pineapple"}
}

func (e *DurianPizzaBuilder) bake(min float64) {
	e.bakeMin += min
}

func (e *DurianPizzaBuilder) addCheess() {
	e.pizza.cheess = true
}

func (e *DurianPizzaBuilder) addTomatoSauce() {
	e.pizza.tomatoSauce = true
}

func (e *DurianPizzaBuilder) GetPizza() (Pizza, error) {
	if e.bakeMin < e.needBakeMin {
		return nil, ErrPizzaNotReady
	}
	return e.pizza, nil
}
