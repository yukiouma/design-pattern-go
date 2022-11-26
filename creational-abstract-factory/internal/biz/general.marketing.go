package biz

type MarketingGeneral struct {
	name   string
	salary float64
}

var _ Genenal = (*MarketingGeneral)(nil)

func (g *MarketingGeneral) Name() string {
	return g.name
}

func (g *MarketingGeneral) Salary() float64 {
	return g.salary
}

func (g *MarketingGeneral) SetSalary(s float64) {
	g.salary = s
}
