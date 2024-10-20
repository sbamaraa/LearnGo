package regular_expression_matching

import "testing"

func TestCase1(t *testing.T) {
	got := isMatch("aa", "a")
	wanted := false

	if got != wanted {
		t.Errorf("got: %t, wanted: %t", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := isMatch("aa", "a*")
	wanted := true

	if got != wanted {
		t.Errorf("got: %t, wanted: %t", got, wanted)
	}
}

func TestCase3(t *testing.T) {
	got := isMatch("ab", ".*")
	wanted := true

	if got != wanted {
		t.Errorf("got: %t, wanted: %t", got, wanted)
	}
}

func TestCase4(t *testing.T) {
	got := isMatch("aab", "c*a*b*")
	wanted := true

	if got != wanted {
		t.Errorf("got: %t, wanted: %t", got, wanted)
	}
}
