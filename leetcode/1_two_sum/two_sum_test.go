package two_sum

import "testing"
import "reflect"

func TestCase1(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}

	if ! reflect.DeepEqual(got, want) {
		t.Error("test case 1 failed");
	}
}

func TestCase2(t *testing.T) {
	got := twoSum([]int{3, 2, 4}, 6)
	want := []int{1, 2}

	if ! reflect.DeepEqual(got, want) {
		t.Errorf("test case 2 failed")
	}
}

func TestCase3(t *testing.T) {
	got := twoSum([]int{3, 3}, 6)
	want := []int{0, 1}

	if ! reflect.DeepEqual(got, want) {
		t.Errorf("test case 3 failed")
	}
}
