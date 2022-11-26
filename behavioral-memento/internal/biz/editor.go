package biz

type editor struct {
	text           string
	x              int
	y              int
	selectionWidth int
}

func NewTextEditor() *editor {
	return &editor{}
}

func (e *editor) SetText(text string) {
	e.text = text
}

func (e *editor) SetCursor(x, y int) {
	e.x = x
	e.y = y
}

func (e *editor) SetSelectionWidth(width int) {
	e.selectionWidth = width
}

func (e *editor) Text() string {
	return e.text
}

func (e *editor) X() int {
	return e.x
}

func (e *editor) Y() int {
	return e.y
}

func (e *editor) SelectionWidth() int {
	return e.selectionWidth
}

func (e *editor) CreateMemento() Memento {
	return NewMemento(e)
}
