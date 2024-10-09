package zigzag_conversion

import "testing"

func TestCase1(t *testing.T) {
	got := convert("PAYPALISHIRING", 3)
	wanted := "PAHNAPLSIIGYIR"

	if got != wanted {
		t.Errorf("got: %s, wanted: %s", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := convert("PAYPALISHIRING", 4)
	wanted := "PINALSIGYAHRPI"

	if got != wanted {
		t.Errorf("got: %s, wanted: %s", got, wanted)
	}
}

func TestCase3(t *testing.T) {
	got := convert("A", 1)
	wanted := "A"

	if got != wanted {
		t.Errorf("got: %s, wanted: %s", got, wanted)
	}
}
