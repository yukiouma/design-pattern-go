package biz

type App struct {
	commandHistory []Command
}

func NewApp() *App {
	return &App{
		commandHistory: make([]Command, 0, 10),
	}
}

func (a *App) Push(c Command) {
	h := a.commandHistory
	c = copyCommand(c)
	if c == nil {
		return
	}
	h = append(h, c)
	a.commandHistory = h
}

func (a *App) Pop() Command {
	h := a.commandHistory
	size := len(h)
	if size == 0 {
		return nil
	}
	c := h[size-1]
	h = h[:size-1]
	a.commandHistory = h
	return c
}

func copyCommand(c Command) Command {
	switch t := c.(type) {
	case *TurnLeftCommand:
		newT := *t
		return &newT
	case *TurnRightCommand:
		newT := *t
		return &newT
	}
	return nil
}
