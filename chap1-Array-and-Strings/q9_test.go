package chap1

import (
	"testing"
)

type testcase struct {
	s, r          string
	expectSuccess bool
}

func TestStringRotation(tt *testing.T) {
	testcases := []testcase{
		{"waterbottle", "erbottlewat", true},
		{"watwaterbottle", "erbottlewatwat", true},
		{"waterbottle", "erbottlewwat", false},
		{"waterbottle", "erbottlexat", false},
	}
	for _, t := range testcases {
		if t.expectSuccess != CheckStringRotation(t.s, t.r) {
			tt.Error("Fail")
		}
	}
}
