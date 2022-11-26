package biz

import "errors"

type stateType int

var (
	ErrUpdateInModeration = errors.New("error: news can not be modified in moderation state")
	ErrUpdateInPublished  = errors.New("error: news can not be modified in published state")
)

const (
	StateDraft stateType = iota
	StateModeration
	StatePublished
)

type State interface {
	Publish()
	Update(string) error
	Revert()
}
