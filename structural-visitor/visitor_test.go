package structuralvisitor

import (
	"math"
	"structural-visitor/internal/biz"
	"testing"
)

func TestVisitor(t *testing.T) {
	visitor := biz.NewAreaVistor()
	circle := biz.NewCircle(10)
	rectangle := biz.NewRectangle(10, 20)
	square := biz.NewSquare(10)

	expect := math.Pow(10, 2) * math.Pi
	result := circle.Accept(visitor)
	if result != expect {
		t.Fatalf("expect %v, got %v", expect, result)
	}

	expect = 10 * 20
	result = rectangle.Accept(visitor)
	if result != expect {
		t.Fatalf("expect %v, got %v", expect, result)
	}

	expect = 10 * 10
	result = square.Accept(visitor)
	if result != expect {
		t.Fatalf("expect %v, got %v", expect, result)
	}
}
