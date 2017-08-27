package chap4

import (
	"fmt"
	"testing"
)

func TestCheckBalanced(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tt := CreateBST(arr)
	fmt.Println(CheckBalanced(tt))
	tt.right = nil
	fmt.Println(CheckBalanced(tt))
}
