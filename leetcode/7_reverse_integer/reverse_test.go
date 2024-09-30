package reverse_integer

import "testing"

func TestCase1(t *testing.T) {
	got := reverse(123)
	want := 321

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase2(t *testing.T) {
	got := reverse(-123)
	want := -321

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase3(t *testing.T) {
	got := reverse(120)
	want := 21

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestOverflow(t *testing.T) {
	got := reverse(1534236469)
	want := 0

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
