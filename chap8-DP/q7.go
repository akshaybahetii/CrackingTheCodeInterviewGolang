package chap8

import (
	"strings"
)

func AddC(x rune, s string) []string {
	var res []string

	str := []rune(s)
	for i, _ := range str {
		ns := strings.Join([]string{string(str[:i]), string(x)}, "")
		ns = strings.Join([]string{ns, string(str[i:])}, "")
		res = append(res, ns)
	}
	ns := strings.Join([]string{s, string(x)}, "")

	res = append(res, ns)
	return res
}

func GetPerm(s string) []string {
	str := []rune(s)

	if len(str) == 0 {
		return nil
	}
	if len(str) == 1 {
		return []string{s}
	}
	var res2 []string
	res := GetPerm(string(str[1:]))
	for _, x := range res {
		f := AddC(str[0], x)
		res2 = append(res2, f...)
	}
	return res2
}
