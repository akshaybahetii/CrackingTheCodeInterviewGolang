package chap1

func CheckPerPal(str1 string, str2 string) bool {
	strArray := []rune(str1)

	slen := len(str1)
	for i := 0; i <= slen/2; i++ {
		if strArray[i] != strArray[slen-i-1] {
			return false
		}
	}

	return checkPer(str1, str2)
}
