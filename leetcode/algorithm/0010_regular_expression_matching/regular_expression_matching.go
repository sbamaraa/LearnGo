package regular_expression_matching

import "regexp"

func isMatch(theString string, pattern string) bool {
	re := regexp.MustCompile("^" + pattern + "$")
	return re.MatchString(theString);
}
