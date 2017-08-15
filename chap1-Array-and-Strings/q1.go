package chap1

func checkAllUnq(s string) bool {
	var found [256]int
	for _, c := range s {
		if found[c] > 0 {
			return false
		}
		found[c]++
	}
	return true
}
