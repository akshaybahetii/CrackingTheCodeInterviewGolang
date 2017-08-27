package chap5

import "fmt"

func Insertion(n uint, m uint, i uint, j uint) uint {

	x := ^uint(0)
	x2 := x << i
	x1 := uint((1 << j) - 1)
	x = x1 | x2
	m = m << j
	x = m | x
	n = n & x
	n = n ^ m
	fmt.Printf("%b %b  %b", x, m, n)

	return 0
}
