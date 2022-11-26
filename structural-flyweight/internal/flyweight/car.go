package flyweight

type Car struct {
	x, y, speed float64
	base        *CarBase
}

func (c Car) Location() (float64, float64) {
	return c.x, c.y
}

func (c Car) Speed() float64 {
	return c.speed
}

func (c Car) Color() string {
	return c.base.Color
}

func (c Car) Brand() string {
	return c.base.Brand
}

func (c Car) SameBase(other Car) bool {
	return c.base == other.base
}
