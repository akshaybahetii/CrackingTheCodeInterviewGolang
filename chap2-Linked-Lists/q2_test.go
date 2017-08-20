package chap2

import "testing"

func TestFindKthTerm(tt *testing.T) {
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
	k := FindKthTerm(head, 7)
	if k.data != 3 {
		tt.Fatalf("Expected 3 got %d", k.data)
	}
}
