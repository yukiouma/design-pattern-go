package biz

type Memento interface {
	SetNext(Memento)
	Next() Memento
	Restore()
}

type memento struct {
	text           string
	x              int
	y              int
	selectionWidth int
	editor         *editor
	next           Memento
}

func NewMemento(e *editor) Memento {
	return &memento{
		text:           e.Text(),
		x:              e.X(),
		y:              e.Y(),
		selectionWidth: e.SelectionWidth(),
		editor:         e,
	}
}

func (m *memento) SetNext(mem Memento) {
	m.next = mem
}

func (m *memento) Next() Memento {
	return m.next
}

func (m *memento) Restore() {
	m.editor.SetText(m.text)
	m.editor.SetCursor(m.x, m.y)
	m.editor.SetSelectionWidth(m.selectionWidth)
}
