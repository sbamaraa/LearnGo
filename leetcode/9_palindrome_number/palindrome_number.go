package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed = 0
	var temp_x = x;
	for temp_x != 0 {
		reversed = reversed * 10 + temp_x % 10
		temp_x /= 10
	}

	return reversed == x
}
