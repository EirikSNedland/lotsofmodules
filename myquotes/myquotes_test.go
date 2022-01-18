package myquotes

import "testing"

func TestDisplay(t *testing.T) {
	wanted := "If a program is too slow, it must have a loop."
	state := Display()
	if state != wanted {
		t.Errorf("Wanted %v, got %v", wanted, state)
	}
}
