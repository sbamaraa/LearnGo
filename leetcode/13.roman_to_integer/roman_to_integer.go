package roman_to_number

func romanToInt(s string) int {
	romanChars := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	};

	result := 0
	highestSoFar := 0

	for i := len(s) - 1; i >= 0; i-- {
		cVal := romanChars[s[i]]
		if cVal < highestSoFar {
			result -= cVal;
		} else {
			result += cVal;
		}

		if highestSoFar < romanChars[s[i]] {
			highestSoFar = romanChars[s[i]]
		}
	}

	return result
}
