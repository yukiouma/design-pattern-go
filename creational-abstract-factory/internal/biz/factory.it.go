package biz

type ITFactory struct{}

var _ Factory = (*ITFactory)(nil)

func (f *ITFactory) CreateGeneral(name string) Genenal {
	return &ITGeneral{
		name: name,
	}
}

func (f *ITFactory) CreateDirector(member Genenal) (Director, error) {
	m, ok := member.(*ITGeneral)
	if !ok {
		return nil, ErrNewDirector
	}
	return &ITDirector{
		name:   m.name,
		salary: m.salary,
	}, nil
}
