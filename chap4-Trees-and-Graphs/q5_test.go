package chap4

import (
	"fmt"
	"testing"
)

func TestValidateBST(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	tt := CreateBST(arr)
	PrettyPrintTree(tt)
	fmt.Println(ValidateBST(tt))
	tt.right.data = -1
	PrettyPrintTree(tt)
	fmt.Println(ValidateBST(tt))

}
