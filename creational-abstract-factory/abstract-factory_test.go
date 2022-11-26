package creationalabstractfactory

import "testing"

func TestAdjustSalary(t *testing.T) {
	itFactory, marketingFactory := NewFactory(IT),
		NewFactory(Marketing)
	itGeneral, marketingGeneral := itFactory.CreateGeneral("a"),
		marketingFactory.CreateGeneral("b")
	itDirector, err := itFactory.CreateDirector(itFactory.CreateGeneral("A"))
	if err != nil {
		t.Error(err.Error())
		return
	}
	if err := itDirector.AdjustSalary(itGeneral, 5000); err != nil {
		t.Error(err.Error())
		return
	}
	if err := itDirector.AdjustSalary(marketingGeneral, 5000); err == nil {
		t.Error("it director can not adjust salary of member in other department")
	}
}
