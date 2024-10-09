package string_to_integer

import "testing"

func TestCase1(t *testing.T) {
	got := myAtoi("42")
	wanted := 42

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}
func TestCase2(t *testing.T) {
	got := myAtoi("-042")
	wanted := -42

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}
func TestCase3(t *testing.T) {
	got := myAtoi("1337c0d3")
	wanted := 1337

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}
func TestCase5(t *testing.T) {
	got := myAtoi("0-1")
	wanted := 0

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}
