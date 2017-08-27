package chap4

import (
	"fmt"
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	tt := CreateBST(arr)
	PrettyPrintTree(tt)
}

func PrettyPrintTree(t *TreeNode) {
	var ll []*LLNode
	p := &ll
	CreateLinkedList(t, 0, &p)
	lp := *p
	for i := 0; i < len(lp); i++ {
		l := lp[i].next
		for l != nil {
			fmt.Printf("%d-", l.data)
			l = l.next
		}
		fmt.Println("        :level", i, ":")
	}

}
