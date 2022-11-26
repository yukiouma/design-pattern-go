package biz

type ITDirector struct {
	name   string
	salary float64
}

var _ Director = (*ITDirector)(nil)

func (d *ITDirector) AdjustSalary(member Genenal, salary float64) error {
	if _, ok := member.(*ITGeneral); !ok {
		return ErrSetSalary
	}
	member.SetSalary(salary)
	return nil
}

func (d *ITDirector) Name() string {
	return d.name
}

func (d *ITDirector) Salary() float64 {
	return d.salary
}

func (d *ITDirector) SetSalary(s float64) {
	d.salary = s
}
