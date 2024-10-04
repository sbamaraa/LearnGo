package remove_element

import (
	"testing"
	"sort"
)

func TestCase1(t *testing.T) {
	toRemove := 3
	nums := []int{3, 2, 2, 3}
	
	trimmedLength := removeElement(nums, toRemove)
	expectedNums := []int{2, 2}

	if trimmedLength == len(expectedNums) {
		if len(nums) < trimmedLength {
			t.Error("Something off with the nums array")
		} else {
			for i := 0; i < len(expectedNums); i++ {
				if expectedNums[i] != nums[i] {
					t.Errorf("At the index %d, expectation is %d, but it was %d", i, expectedNums[i], nums[i])
				}
			}
		}
	} else {
		t.Errorf("Length we wanted [%d], but got [%d]", len(expectedNums), trimmedLength)
	}

}

func TestCase2(t *testing.T) {
	toRemove := 2
	nums := []int{0,1,2,2,3,0,4,2}
	
	trimmedLength := removeElement(nums, toRemove)
	expectedNums := []int{0,0,1,3,4}

	if trimmedLength == len(expectedNums) {
		if len(nums) < trimmedLength {
			t.Error("Something off with the nums array")
		} else {
			sort.Ints(nums[:trimmedLength])
			for i := 0; i < len(expectedNums); i++ {
				if expectedNums[i] != nums[i] {
					t.Errorf("At the index %d, expectation is %d, but it was %d", i, expectedNums[i], nums[i])
				}
			}
		}
	} else {
		t.Errorf("Length we wanted [%d], but got [%d]", len(expectedNums), trimmedLength)
	}

}
