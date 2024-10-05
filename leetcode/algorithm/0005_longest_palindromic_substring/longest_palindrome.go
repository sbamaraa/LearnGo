package longest_palindrome

func longestPalindrome(theString string) string {
	n := len(theString)

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	s := 0
	f := 0
	for j := 1; j < n; j++ {
		for i := 0; i + j < n; i++ {
			if theString[i] == theString[i + j] {
				if dp[i + 1][i + j - 1] || j == 1{
					dp[i][i + j] = true
					s = i
					f = i + j
				}
			} else {
				dp[i][i + j] = false
			}
		}
	}


	return theString[s:f+1]
}
