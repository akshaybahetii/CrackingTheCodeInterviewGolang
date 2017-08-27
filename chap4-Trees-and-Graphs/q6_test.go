package chap4

import (
	"fmt"
	"testing"
)

func TestFindNodeInBST(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tt := CreateBST(arr)
	PrettyPrintTree(tt)
	fmt.Println(FindNodeInBST(tt, 5).data)
	fmt.Println(FindNodeInBST(tt, 10).data)
	fmt.Println(FindNodeInBST(tt, 9).data)
	fmt.Println(FindNodeInBST(tt, 1).data)
	fmt.Println(FindNodeInBST(tt, 6).data)
	fmt.Println(FindNodeInBST(tt, 7).data)
	//	tt.right.data = -1
	//	PrettyPrintTree(tt)
	//	fmt.Println(ValidateBST(tt))

}
