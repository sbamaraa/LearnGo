package palindrome_number

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var reversed = 0
	var tempX = x;
	for tempX != 0 {
		reversed = reversed * 10 + tempX % 10
		tempX /= 10
	}

	return reversed == x
}
