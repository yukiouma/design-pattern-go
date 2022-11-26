package biz

type TurnRightCommand struct {
	backup   Direction
	receiver *Car
	app      *App
}

func NewTurnRightCommand(receiver *Car, app *App) Command {
	command := &TurnRightCommand{
		receiver: receiver,
		app:      app,
	}
	return command
}

func (c *TurnRightCommand) Execute() {
	lastdir := c.receiver.Direction()
	newdir := lastdir
	newdir -= 1
	if newdir < 0 {
		newdir += 4
	}
	c.backup = lastdir
	c.receiver.ChangeDirection(newdir)
	c.app.Push(c)
}

func (c *TurnRightCommand) Undo() {
	c.receiver.ChangeDirection(c.backup)
}
