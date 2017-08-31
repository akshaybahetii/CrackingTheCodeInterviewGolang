package main

import (
	"fmt"
	"strings"
)

func main() {
	t1 := []string{"apple", "banana", "orange"}
	t2 := []string{"apple", "lemon", "orange", "graples"}
	t3 := []string{"apple", "orange"}

	c := make(map[string]int)

	t1 = GetConb(t1)
	t2 = GetConb(t2)
	t3 = GetConb(t3)

	var res []string
	res = append(res, t1...)
	res = append(res, t2...)
	res = append(res, t3...)

	for _, x := range res {
		if !strings.Contains(x, ":") {
			continue
		}
		c[x]++
	}
	fmt.Println(c)
}

func GetConb(t []string) []string {
	if len(t) <= 0 {
		return nil
	}
	if len(t) == 1 {
		return t
	}
	s := GetConb(t[1:])
	z := t[0]
	var res []string
	for _, x := range s {
		res = append(res, strings.Join([]string{z, x}, ":"))
	}
	res = append(t, res...)
	return res
}
