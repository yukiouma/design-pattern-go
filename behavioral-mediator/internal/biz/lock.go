package biz

type lockButoon struct {
	lock     bool
	mediator Mediator
}

func NewLockButton(mediator Mediator) *lockButoon {
	return &lockButoon{
		mediator: mediator,
	}
}

func (l *lockButoon) IsLock() bool {
	return l.lock
}

func (l *lockButoon) SwitchStatus() {
	l.lock = !l.lock
}
