package chap4

import "fmt"

type TreeNode struct {
	data        int
	right, left *TreeNode
	parent      *TreeNode
	visited     bool
}

func PrintTree(t *TreeNode) {
	if t == nil {
		return
	}
	if t.left == nil && t.right == nil {
		fmt.Printf("%d", t.data)
		return
	}
	PrintTree(t.left)
	fmt.Printf("%d", t.data)
	PrintTree(t.right)
}
func CreateBST(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	t := TreeNode{}
	if len(arr) == 1 {
		t.data = arr[0]
		return &t
	}
	if len(arr) == 2 {
		tl := TreeNode{}
		tl.data = arr[0]
		t.left = &tl
		tl.parent = &t
		t.data = arr[1]
		return &t
	}

	mid := len(arr) / 2

	t.left = CreateBST(arr[:mid])
	t.data = arr[mid]
	t.right = CreateBST(arr[mid+1:])
	t.left.parent = &t
	t.right.parent = &t

	return &t

}
