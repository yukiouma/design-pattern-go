package biz

type MarketingDirector struct {
	name   string
	salary float64
}

var _ Director = (*MarketingDirector)(nil)

func (d *MarketingDirector) AdjustSalary(member Genenal, salary float64) error {
	if _, ok := member.(*MarketingGeneral); !ok {
		return ErrSetSalary
	}
	member.SetSalary(salary)
	return nil
}

func (d *MarketingDirector) Name() string {
	return d.name
}

func (d *MarketingDirector) Salary() float64 {
	return d.salary
}

func (d *MarketingDirector) SetSalary(s float64) {
	d.salary = s
}
