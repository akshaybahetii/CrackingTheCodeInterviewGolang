package chap2

import (
	"fmt"
	"testing"
)

func TestFindIntersection(t *testing.T) {
	head1 := &Node{1, nil}
	head1 = head1.AddHead(2)
	head1 = head1.AddHead(3)
	head1 = head1.AddHead(3)
	head2 := &Node{3, head1}
	head2 = head2.AddHead(8)
	head2 = head2.AddHead(2)
	head2 = head2.AddHead(12)
	head1 = head1.AddHead(2)
	head1 = head1.AddHead(3)
	head1 = head1.AddHead(3)
	fmt.Println(FindIntersection(head1, head2))
}
