package biz

type MarketingFactory struct{}

var _ Factory = (*MarketingFactory)(nil)

func (f *MarketingFactory) CreateGeneral(name string) Genenal {
	return &MarketingGeneral{
		name: name,
	}
}

func (f *MarketingFactory) CreateDirector(member Genenal) (Director, error) {
	m, ok := member.(*MarketingGeneral)
	if !ok {
		return nil, ErrNewDirector
	}
	return &MarketingDirector{
		name:   m.name,
		salary: m.salary,
	}, nil
}
