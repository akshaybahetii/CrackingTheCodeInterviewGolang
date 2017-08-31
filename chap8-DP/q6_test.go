package chap8

import (
	"fmt"
	"testing"
)

func TestTOI(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{}
	c := []int{}

	fmt.Println(a, b, c)
	TOI(5, &a, &b, &c)
	fmt.Println(a, b, c)

}
