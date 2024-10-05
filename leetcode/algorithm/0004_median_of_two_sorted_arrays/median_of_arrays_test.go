package median_of_arrays

import "testing"

func TestCase1(t *testing.T) {
	got := findMedianSortedArays([]int{1, 3}, []int{2})
	wanted := 2.0

	if got != wanted {
		t.Errorf("got %f, wanted %f", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := findMedianSortedArays([]int{1, 2}, []int{3, 4})
	wanted := 2.5

	if got != wanted {
		t.Errorf("got %f, wanted %f", got, wanted)
	}
}
