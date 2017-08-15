package chap1

func URLify(str1 string) string {
	strArray := []rune(str1)
	var res []rune

	for _, a := range strArray {
		if a == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, a)
		}
	}

	return string(res)
}
