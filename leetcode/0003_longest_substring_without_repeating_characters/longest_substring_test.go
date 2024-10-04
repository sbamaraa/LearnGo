package longest_substring

import "testing"

func TestCase1(t *testing.T) {
	got := lengthOfLongestSubstring("abcabcbb")
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase2(t *testing.T) {
	got := lengthOfLongestSubstring("bbbbb")
	want := 1

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase3(t *testing.T) {
	got := lengthOfLongestSubstring("pwwkew")
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
