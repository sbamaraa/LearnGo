package divide_two_integers

import "testing"

func TestCase1(t *testing.T) {
	got := divide(10, 3)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase2(t *testing.T) {
	got := divide(7, -3)
	want := -2

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRegression1(t *testing.T) {
	got := divide(-2147483648, -1)
	want := 2147483647

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
