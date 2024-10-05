package longest_substring

func lengthOfLongestSubstring(s string) int {
	lastOccurence := make(map[byte]int)

	best := 0
	for l, r := 0, 0; r < len(s); r++ {
		if ind, exists := lastOccurence[s[r]]; exists {
			for ; l <= ind; l++ {
				delete(lastOccurence, s[l])
			}
		}

		if best < r - l + 1 {
			best = r - l + 1
		}

		lastOccurence[s[r]] = r
	}

	return best
}
