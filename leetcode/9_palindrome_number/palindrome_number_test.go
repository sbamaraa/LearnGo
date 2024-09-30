package palindrome_number

import "testing"
import "strconv"

func TestCase1(t *testing.T) {
	got := isPalindrome(121)
	want := true

	if got != want {
		t.Errorf("got %s, wanted %s", strconv.FormatBool(got), strconv.FormatBool(want));
	}
}

func TestCase2(t *testing.T) {
	got := isPalindrome(-121)
	want := false

	if got != want {
		t.Errorf("got %s, wanted %s", strconv.FormatBool(got), strconv.FormatBool(want));
	}
}


func TestCase3(t *testing.T) {
	got := isPalindrome(10)
	want := false

	if got != want {
		t.Errorf("got %s, wanted %s", strconv.FormatBool(got), strconv.FormatBool(want));
	}
}
