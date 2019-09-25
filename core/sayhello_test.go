package core

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello World !"
	got := SayHello("World")
	if got != want {
		t.Errorf("Incorrect message. Want %s, got: %s", want, got)
	}
}
