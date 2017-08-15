package chap1

import "testing"

func TestPalPtr(t *testing.T) {
	type testSet struct {
		s1            string
		s2            string
		expectSuccess bool
	}

	testcases := []testSet{
		{"abbbba", "babbab", true},
	}

	for _, s := range testcases {
		if s.expectSuccess != CheckPerPal(s.s1, s.s2) {
			t.Error("Fail")
		}
	}
}
