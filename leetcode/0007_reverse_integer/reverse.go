package reverse_integer

import "math"

func reverse(x int) int {
	var reversed = 0;
	for x != 0 {
		reversed = reversed * 10 + x % 10
		x /= 10
	}

	if reversed < math.MinInt32 || math.MaxInt32 < reversed {
		return 0;
	}

	return reversed;
}
