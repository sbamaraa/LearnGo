package longest_palindrome

import "testing"

func TestCase1(t *testing.T) {
	got := longestPalindrome("babad")
	wanted := "aba"

	if got != wanted {
		t.Errorf("got %s, wanted %s", got, wanted)
	}
}

func TestCase2(t *testing.T) {
	got := longestPalindrome("cbbd")
	wanted := "bb"

	if got != wanted {
		t.Errorf("got %s, wanted %s", got, wanted)
	}
}

func TestCaseRegression(t *testing.T) {
	got := longestPalindrome("aacabdkacaa")
	wanted := "aca"

	if got != wanted {
		t.Errorf("got %s, wanted %s", got, wanted)
	}
}
