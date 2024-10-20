package container

import "testing"

func TestCase1(t *testing.T) {
	got := maxArea([]int{1,8,6,2,5,4,8,3,7})
	wanted := 49

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := maxArea([]int{1, 1})
	wanted := 1

	if got != wanted {
		t.Errorf("got: %d, wanted: %d", got, wanted)
	}
}
