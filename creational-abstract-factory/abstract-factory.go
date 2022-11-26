package creationalabstractfactory

import "creational-abstract-factory/internal/biz"

type department int

const (
	IT department = iota
	Marketing
)

// Get factory, default is IT
func NewFactory(dept department) biz.Factory {
	switch dept {
	case Marketing:
		return &biz.MarketingFactory{}
	default:
		return &biz.ITFactory{}
	}
}
