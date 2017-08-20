package chap2

import "testing"

func TestRemoveDuplicates(tt *testing.T) {
	head := &Node{1, nil}
	head.Add(2)
	head.Add(3)
	head.Add(4)
	head.Add(4)
	head.Add(4)
	head.Add(3)
	head.Add(3)
	head.Add(3)
	head.Add(3)
	RemoveDuplicates(head)
	head.Print()
}
