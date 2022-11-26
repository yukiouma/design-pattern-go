package biz

type ITGeneral struct {
	name   string
	salary float64
}

var _ Genenal = (*ITGeneral)(nil)

func (g *ITGeneral) Name() string {
	return g.name
}

func (g *ITGeneral) Salary() float64 {
	return g.salary
}

func (g *ITGeneral) SetSalary(s float64) {
	g.salary = s
}
