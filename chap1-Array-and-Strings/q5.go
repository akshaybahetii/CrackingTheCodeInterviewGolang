package chap1

func checkEdits(s1 string, s2 string) bool {
	if len(s2) > (len(s1)+1) || len(s2) < (len(s1)-1) {
		return false
	}

	str1 := []rune(s1)
	str2 := []rune(s2)
	diff := 0

	minlen := len(str1)
	if len(str1) > len(str2) {
		minlen = len(str2)
		diff++
	}
	if len(str1) < len(str2) {
		diff++
	}
	for i := 0; i < minlen; i++ {
		if str1[i] != str2[i] {
			diff++
		}
	}

	if diff > 1 {
		return false
	}
	return true
}
