package remove_element

import "sort"

func removeElement(nums []int, val int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] == val
	})

	var valCount = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			valCount++
		}
	}

	return len(nums) - valCount
}
