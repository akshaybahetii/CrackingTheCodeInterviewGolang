package chap1

import "testing"

func TestURLify(t *testing.T) {
	type testSet struct {
		s1 string
		s2 string
	}

	testcases := []testSet{
		{"aa a", "aa%20a"},
		{" ", "%20"},
	}

	for _, s := range testcases {
		if s.s2 != URLify(s.s1) {
			t.Error("Fail")
		}
	}
}
