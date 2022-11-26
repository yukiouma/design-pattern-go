package biz

type command struct {
	editor  *editor
	history Memento
}

func NewCommand(editor *editor) *command {
	return &command{
		editor: editor,
	}
}

func (c *command) Backup() {
	memento := c.editor.CreateMemento()
	memento.SetNext(c.history)
	c.history = memento
}

func (c *command) Restore() bool {
	memento := c.history
	if memento == nil {
		return false
	}
	c.history = memento.Next()
	memento.Restore()
	return true
}
