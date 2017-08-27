package chap4

import "testing"

func TestCreateBST(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tt := CreateBST(arr)
	PrintTree(tt)
}
