package structuraltemplatemethod

import (
	"structural-template-method/internal/biz"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	warrior := biz.NewWarrior().SetHp(100).SetPower(10)
	knight := biz.NewKnight().SetHp(100).SetPower(10).SetDefense(8)
	magician := biz.NewMagician().SetHp(100).SetPower(10).SetMana(10)

	warrior.Attack(knight, biz.Defend)
	if warrior.Hp() != 100 {
		t.Fatalf("warrior's hp expect 100, got %v", warrior.Hp())
	}
	if knight.Hp() != 98 {
		t.Fatalf("knight's hp expect 98, got %v", knight.Hp())
	}

	magician.Attack(knight, biz.Defend)
	if magician.Hp() != 100 {
		t.Fatalf("magician's hp expect 100, got %v", magician.Hp())
	}
	if knight.Hp() != 86 {
		t.Fatalf("knight's hp expect 86, got %v", knight.Hp())
	}

	knight.Attack(warrior, biz.Attack)
	if knight.Hp() != 76 {
		t.Fatalf("knight's hp expect 84, got %v", knight.Hp())
	}
	if warrior.Hp() != 90 {
		t.Fatalf("warrior's hp expect 90, got %v", warrior.Hp())
	}
}
