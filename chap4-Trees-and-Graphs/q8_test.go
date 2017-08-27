package chap4

import (
	"fmt"
	"testing"
)

func TestFirstComAncestor(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tt := CreateBST(arr)
	fmt.Println(FirstComAncestor(tt, 1, 7))
}
