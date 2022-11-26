package biz

type TurnLeftCommand struct {
	backup   Direction
	receiver *Car
	app      *App
}

func NewTurnLeftCommand(receiver *Car, app *App) Command {
	command := &TurnLeftCommand{
		receiver: receiver,
		app:      app,
	}
	return command
}

func (c *TurnLeftCommand) Execute() {
	lastdir := c.receiver.Direction()
	newdir := lastdir + 1
	newdir = newdir % 4
	c.backup = lastdir
	c.receiver.ChangeDirection(newdir)
	c.app.Push(c)
}

func (c *TurnLeftCommand) Undo() {
	c.receiver.ChangeDirection(c.backup)
}
