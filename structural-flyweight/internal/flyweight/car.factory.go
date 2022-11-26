package flyweight

type CarFactory struct {
	carbase map[[2]string]*CarBase
}

func NewCarFacory() *CarFactory {
	return &CarFactory{
		carbase: make(map[[2]string]*CarBase),
	}
}

func (f *CarFactory) getCarBase(color, brand string) *CarBase {
	carbase := f.carbase
	key := [2]string{color, brand}
	base, ok := carbase[key]
	if !ok {
		base = &CarBase{Color: color, Brand: brand}
		carbase[key] = base
	}
	return base
}

func (f *CarFactory) NewCar(base CarBase, x, y, speed float64) *Car {
	carbase := f.getCarBase(base.Color, base.Brand)
	return &Car{
		base:  carbase,
		x:     x,
		y:     y,
		speed: speed,
	}
}

func (f *CarFactory) BaseSize() int {
	return len(f.carbase)
}
