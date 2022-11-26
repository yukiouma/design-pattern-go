package structuralflyweight

import (
	"structural-flyweight/internal/flyweight"
	"testing"
)

func TestFlyweight(t *testing.T) {
	factory := flyweight.NewCarFacory()
	carData := []flyweight.CarBase{
		{Color: "red", Brand: "toyota"},
		{Color: "black", Brand: "honda"},
		{Color: "white", Brand: "tesla"},
	}
	cars := make([]*flyweight.Car, 0, 50)
	for i := 0; i < 50; i++ {
		base := carData[i%3]
		car := factory.NewCar(flyweight.CarBase{Color: base.Color, Brand: base.Brand}, 0, 0, 0)
		cars = append(cars, car)
	}
	if len(cars) != 50 {
		t.Fatalf("wrong car number")
	}
	if factory.BaseSize() != 3 {
		t.Fatalf("wrong car base number")
	}
	if !cars[0].SameBase(*cars[3]) {
		t.Fatalf("wrong base")
	}
}
