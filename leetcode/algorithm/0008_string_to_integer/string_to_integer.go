package string_to_integer

import "unicode"
import "math"

func myAtoi(s string) int {
	expectWhiteSpaces := true
	expectSigns := true

	res := 0
	sign := '+'

	for _, theByte := range s {
		if theByte == ' ' {
			if ! expectWhiteSpaces {
				break
			}

			continue
		}

		if theByte == '-' || theByte == '+' {
			if ! expectSigns {
				break
			}

			sign = theByte
			expectWhiteSpaces = false
			expectSigns = false
			continue
		}

		if unicode.IsDigit(theByte) {
			res = res * 10 + int(theByte - '0')
			if sign == '+' && res > math.MaxInt32 {
				return math.MaxInt32
			} else if sign == '-' && -res < math.MinInt32 {
				return math.MinInt32
			}

			expectWhiteSpaces = false
			expectSigns = false
			continue
		}

		break
	}

	if sign == '-' {
		return -res
	}

	return res
}
