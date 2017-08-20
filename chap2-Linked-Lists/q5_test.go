package chap2

import "testing"

func TestAddTwoNum(t *testing.T) {
	head1 := &Node{6, nil}
	head1 = head1.AddHead(1)
	head1 = head1.AddHead(7)
	head2 := &Node{2, nil}
	head2 = head2.AddHead(9)
	head2 = head2.AddHead(5)
	head3 := AddTwoNum(head1, head2)
	head3.Print()
}
