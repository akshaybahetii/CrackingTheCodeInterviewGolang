package chap2

import "testing"

func TestMoveAllNodeLessThanX(t *testing.T) {
	head := &Node{1, nil}
	head = head.AddHead(2)
	head = head.AddHead(10)
	head = head.AddHead(5)
	head = head.AddHead(8)
	head = head.AddHead(5)
	head = head.AddHead(3)
	head.Print()
	head = MoveAllNodeLessThanX(head, 5)
	head.Print()
}
