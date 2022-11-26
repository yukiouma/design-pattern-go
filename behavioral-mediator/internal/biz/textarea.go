package biz

import "errors"

var (
	errLockStatus = errors.New("error: can not edit text area in lock status")
)

type textArea struct {
	mediator Mediator
	text     string
}

func NewTextArea(mediator Mediator) *textArea {
	return &textArea{mediator: mediator}
}

func (t *textArea) Save(text string) error {
	if t.mediator.IsLock() {
		return errLockStatus
	}
	t.text = text
	return nil
}

func (t *textArea) Content() string {
	return t.text
}
