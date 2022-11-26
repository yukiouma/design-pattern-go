package behavioralmemento

import (
	"behavioral-memento/internal/biz"
	"testing"
)

func TestMemento(t *testing.T) {
	editor := biz.NewTextEditor()
	command := biz.NewCommand(editor)

	// editing
	editor.SetText("I have a pen")
	editor.SetCursor(5, 1)
	editor.SetSelectionWidth(3)

	// save state
	command.Backup()

	// editing
	editor.SetText("I have a pen\nI have a apple")
	editor.SetCursor(3, 2)
	editor.SetSelectionWidth(5)

	if editor.Text() != "I have a pen\nI have a apple" {
		t.Fatalf("expect [I have a pen\nI have a apple], got [%s]", editor.Text())
	}

	if editor.X() != 3 {
		t.Fatalf("expect [3], got [%d]", editor.X())
	}

	if editor.Y() != 2 {
		t.Fatalf("expect [2], got [%d]", editor.Y())
	}

	if editor.SelectionWidth() != 5 {
		t.Fatalf("expect [5], got [%d]", editor.SelectionWidth())
	}

	if !command.Restore() {
		t.Fatalf("restore should succeed")
	}

	if command.Restore() {
		t.Fatalf("restore should failed")
	}

	if editor.Text() != "I have a pen" {
		t.Fatalf("expect [I have a pen], got [%s]", editor.Text())
	}

	if editor.X() != 5 {
		t.Fatalf("expect [5], got [%d]", editor.X())
	}

	if editor.Y() != 1 {
		t.Fatalf("expect [1], got [%d]", editor.Y())
	}

	if editor.SelectionWidth() != 3 {
		t.Fatalf("expect [3], got [%d]", editor.SelectionWidth())
	}
}
