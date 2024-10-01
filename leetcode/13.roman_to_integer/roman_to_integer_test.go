package roman_to_number

import "testing"

func TestCase1(t *testing.T) {
	got := romanToInt("III")
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase2(t *testing.T) {
	got := romanToInt("LVIII")
	want := 58

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase3(t *testing.T) {
	got := romanToInt("MCMXCIV")
	want := 1994

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
