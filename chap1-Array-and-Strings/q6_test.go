package chap1

import (
	"testing"
)

func TestTruncateString(t *testing.T) {
	type testSet struct {
		s1            string
		s2            string
		expectSuccess bool
	}

	testcases := []testSet{
		{"a", "a", true},
		{"aaaab", "a4b1", true},
		{"abbbb", "a1b4", true},
		{"aaaabbbb", "a4b4", true},
	}

	for _, s := range testcases {
		if s.expectSuccess != (s.s2 == truncateString(s.s1)) {
			t.Error("Fail %s:%s", s.s1, truncateString(s.s1))
		}
	}
}
