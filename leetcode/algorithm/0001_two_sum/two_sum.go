package two_sum

func twoSum(nums []int, target int) []int {
	occurences := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if ind, exists := occurences[complement]; exists {
			return []int{ind, i}
		}
		occurences[nums[i]] = i
	}

	return []int{} // This should never happen thoug, because the problem guarantees solution.
}
