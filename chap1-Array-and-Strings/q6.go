package chap1

import "fmt"

func truncateString(s string) string {
	str1 := []rune(s)
	var res []rune
	c := str1[0]
	cnum := 0
	for _, a := range str1 {
		if c != a {
			adds := fmt.Sprintf("%s%d", string(c), cnum)
			res = append(res, []rune(adds)...)
			cnum = 1
			c = a
		} else {
			cnum++
		}
	}
	adds := fmt.Sprintf("%s%d", string(c), cnum)
	res = append(res, []rune(adds)...)

	if len(res) > len(s) {
		return s
	}
	return string(res)
}
