package chap1

import (
	"testing"
)

func TestStringPer(t *testing.T) {
	type testSet struct {
		s1            string
		s2            string
		expectSuccess bool
	}

	testcases := []testSet{
		{"a", "a", true},
		{"aaa", "abcda", false},
		{"abc", "abc", true},
		{"abc", "abcda", false},
		{"aaabbbccc", "abcabcabc", true},
		{"abcz", "abczabcz", false},
	}

	for _, s := range testcases {
		if s.expectSuccess != checkPer(s.s1, s.s2) {
			t.Error("Fail")
		}
	}
}
