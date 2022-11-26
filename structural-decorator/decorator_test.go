package structuraldecorator

import (
	"strings"
	"structural-decorator/internal/biz"
	"testing"
)

const (
	applePrefix   = "APPLE"
	googlePrefix  = "GOOGLE"
	tmobilePrefix = "TMOBILE"
)

func TestDecorator(t *testing.T) {
	tmobile := biz.NewTMobile()
	google := biz.NewGoogle(tmobile)
	apple := biz.NewApple(tmobile)

	msg1 := tmobile.Send()
	msg2 := google.Send()
	msg3 := apple.Send()

	if !strings.HasPrefix(msg1, tmobilePrefix) {
		t.Fatal("tmobile message error")
	}

	if !strings.HasPrefix(msg2, googlePrefix) {
		t.Fatal("google message error")
	}

	if !strings.HasPrefix(msg3, applePrefix) {
		t.Fatal("apple message error")
	}
}
