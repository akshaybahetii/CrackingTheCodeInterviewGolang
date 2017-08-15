package chap1

func checkPer(str1 string, str2 string) bool {
	var counter [26]int

	for _, a := range str1 {
		counter[a-97]++
	}

	for _, a := range str2 {
		counter[a-97]--
	}

	for _, x := range counter {
		if x != 0 {
			return false
		}
	}
	return true
}
