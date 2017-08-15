package chap1

import (
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	testcases := []string{"a", "z", "aaa", "abcda", "abd", "abcabc", "kdsjflasfa", "qwert"}

	for _, s := range testcases {

		if !checkAllUnq(s) {
			t.Error("\nFail for [%s]", s)
		}
	}
}
