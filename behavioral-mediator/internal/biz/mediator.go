package biz

type Mediator interface {
	BindTextArea(*textArea)
	BindLockButton(*lockButoon)
	BindCleanButton(*cleanButoon)
	IsLock() bool
	Clean() error
}

type mediator struct {
	text  *textArea
	lock  *lockButoon
	clean *cleanButoon
}

func NewMediator() Mediator {
	return &mediator{}
}

func (m *mediator) BindTextArea(text *textArea) {
	m.text = text
}

func (m *mediator) BindLockButton(lock *lockButoon) {
	m.lock = lock
}

func (m *mediator) BindCleanButton(clean *cleanButoon) {
	m.clean = clean
}

func (m *mediator) IsLock() bool {
	return m.lock.IsLock()
}

func (m *mediator) Clean() error {
	return m.text.Save("")
}
