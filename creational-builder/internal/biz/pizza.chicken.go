package biz

type ChickenPizza struct {
	filling     []string
	cheess      bool
	tomatoSauce bool
	thinCrust   bool
}

var _ Pizza = (*ChickenPizza)(nil)

func (e *ChickenPizza) ListFillings() []string {
	return e.filling
}

func (e *ChickenPizza) WithCheess() bool {
	return e.cheess
}

func (e *ChickenPizza) WithTomatoSauce() bool {
	return e.tomatoSauce
}

func (e *ChickenPizza) IsThinCrust() bool {
	return e.thinCrust
}
