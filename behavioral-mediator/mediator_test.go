package behavioralmediator

import (
	"behavioral-mediator/internal/biz"
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := biz.NewMediator()
	textArea := biz.NewTextArea(mediator)
	lockButoon := biz.NewLockButton(mediator)
	cleanButoon := biz.NewClearButton(mediator)
	mediator.BindTextArea(textArea)
	mediator.BindLockButton(lockButoon)
	mediator.BindCleanButton(cleanButoon)
	if err := textArea.Save("I have a pen"); err != nil {
		t.Fatalf("unexpected error")
	}
	if err := cleanButoon.Clean(); err != nil {
		t.Fatalf("unexpected error")
	}
	if len(textArea.Content()) > 0 {
		t.Fatalf("text area should be empty")
	}
	if err := textArea.Save("I have an apple"); err != nil {
		t.Fatalf("unexpected error")
	}
	lockButoon.SwitchStatus()
	if err := textArea.Save("ppap"); err == nil {
		t.Fatalf("expected error")
	}
	if err := cleanButoon.Clean(); err == nil {
		t.Fatalf("expected error")
	}
}
