package chap1

import (
	"testing"
)

func TestStringEdits(t *testing.T) {
	type testSet struct {
		s1            string
		s2            string
		expectSuccess bool
	}

	testcases := []testSet{
		{"a", "a", true},
		{"a", "ab", true},
		{"aaa", "abcda", false},
		{"abc", "abc", true},
		{"abc", "ab", true},
		{"abc", "abcda", false},
		{"aaabbbccc", "abcabcabc", false},
		{"abcz", "abczabcz", false},
	}

	for _, s := range testcases {
		if s.expectSuccess != checkEdits(s.s1, s.s2) {
			t.Error("Fail")
		}
	}
}
