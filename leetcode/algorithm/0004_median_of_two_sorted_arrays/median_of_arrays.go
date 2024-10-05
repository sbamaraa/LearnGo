package median_of_arrays

func findMedianSortedArays(nums1 []int, nums2 []int) float64 {
	m1 := float64(nthElement(nums1, nums2, len(nums1) + len(nums2) / 2))
	m2 := float64(nthElement(nums1, nums2, (len(nums1) + len(nums2) + 1) / 2))
	return (m1 + m2) / 2.0
}

func nthElement(nums1 []int, nums2 []int, n int) int {
	el, i1, i2 := 0, 0, 0
	for n > 0 {
		if i1 < len(nums1) && (nums1[i1] < nums2[i2] || i2 == len(nums2)) {
			el = nums1[i1]
			i1++
		} else {
			el = nums2[i2]
			i2++
		}
		n--
	}
	return el
}
