package biz

type Builder interface {
	useThinCrust(bool)
	setBakeTime()
	addFilling()
	bake(float64)
	addCheess()
	addTomatoSauce()
	GetPizza() (Pizza, error)
}
