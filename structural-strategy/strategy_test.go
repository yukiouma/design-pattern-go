package structuralstrategy

import (
	"structural-strategy/internal/biz"
	"testing"
)

func TestStrategy(t *testing.T) {
	var (
		x, y, expect, result int
	)
	handler := biz.NewMathHandler()

	x, y, expect = 10, 20, 30
	handler.SetStrategy(biz.AddStrategy)
	result = handler.Execute(x, y)
	if expect != result {
		t.Fatalf("error: expect %v, got %v", expect, result)
	}

	x, y, expect = 100, 20, 80
	handler.SetStrategy(biz.SubstractStrategy)
	result = handler.Execute(x, y)
	if expect != result {
		t.Fatalf("error: expect %v, got %v", expect, result)
	}

	x, y, expect = 5, 20, 100
	handler.SetStrategy(biz.MultipleStrategy)
	result = handler.Execute(x, y)
	if expect != result {
		t.Fatalf("error: expect %v, got %v", expect, result)
	}

	x, y, expect = 10, 2, 5
	handler.SetStrategy(biz.DivideStrategy)
	result = handler.Execute(x, y)
	if expect != result {
		t.Fatalf("error: expect %v, got %v", expect, result)
	}
}
