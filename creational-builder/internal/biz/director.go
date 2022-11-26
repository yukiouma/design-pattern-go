package biz

type Director struct {
	builder Builder
}

func NewDirector() *Director {
	return &Director{}
}

func (e *Director) SetBuilder(b Builder) {
	e.builder = b
}

func (e *Director) MakeChickenPizza() {
	if e.builder == nil {
		return
	}
	e.builder.useThinCrust(true)
	e.builder.setBakeTime()
	e.builder.bake(2)
	e.builder.addFilling()
	e.builder.addTomatoSauce()
	e.builder.bake(5)
}

func (e *Director) MakeDurianPizza() {
	if e.builder == nil {
		return
	}
	e.builder.useThinCrust(false)
	e.builder.setBakeTime()
	e.builder.addFilling()
	e.builder.addCheess()
	e.builder.bake(5)
}

func (e *Director) ResetBuilder() {
	e.builder = nil
}
