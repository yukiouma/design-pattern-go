package biz

type UndoCommand struct {
	app *App
}

func NewUndoCommand(receiver *Car, app *App) Command {
	command := &UndoCommand{
		app: app,
	}
	return command
}

func (c *UndoCommand) Execute() {
	command := c.app.Pop()
	if command == nil {
		return
	}
	command.Undo()
}

func (c *UndoCommand) Undo() {
	// Nothing need to do
}
