package chap2

import "testing"

func TestDeleteCurrentNode(tt *testing.T) {
	head := &Node{10, nil}
	head.Add(9)
	head.Add(8)
	head.Add(7)
	head.Add(6)
	head.Add(5)
	head.Add(4)
	head.Add(3)
	head.Add(2)
	head.Add(1)
	DeleteCurrentNode(head.next.next)
	head.Print()
}
