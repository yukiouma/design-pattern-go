package biz

type Pizza interface {
	ListFillings() []string
	WithCheess() bool
	WithTomatoSauce() bool
	IsThinCrust() bool
}
