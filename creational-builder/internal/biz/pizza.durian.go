package biz

type DurianPizza struct {
	filling     []string
	cheess      bool
	tomatoSauce bool
	thinCrust   bool
}

var _ Pizza = (*DurianPizza)(nil)

func (e *DurianPizza) ListFillings() []string {
	return e.filling
}

func (e *DurianPizza) WithCheess() bool {
	return e.cheess
}

func (e *DurianPizza) WithTomatoSauce() bool {
	return e.tomatoSauce
}

func (e *DurianPizza) IsThinCrust() bool {
	return e.thinCrust
}
