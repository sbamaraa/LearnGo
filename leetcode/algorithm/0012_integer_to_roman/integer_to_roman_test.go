package integer_to_roman

import "testing"

func TestCase1(t *testing.T) {
	got := intToRoman(3749)
	wanted := "MMMDCCXLIX"

	if got != wanted {
		t.Errorf("got: %s, wanted: %s", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := intToRoman(58)
	wanted := "LVIII"

	if got != wanted {
		t.Errorf("got: %s, wanted: %s", got, wanted)
	}
}

func TestCase3(t *testing.T) {
	got := intToRoman(1994)
	wanted := "MCMXCIV"

	if got != wanted {
		t.Errorf("got: %s, wanted: %s", got, wanted)
	}
}
