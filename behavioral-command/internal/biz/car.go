package biz

const defaultDirection = Up

// receiver
type Car struct {
	direction Direction
}

// create car receiver
func NewCar() *Car {
	return &Car{
		direction: defaultDirection,
	}
}

func (c *Car) ChangeDirection(d Direction) {
	c.direction = d
}

func (c *Car) Direction() Direction {
	return c.direction
}
