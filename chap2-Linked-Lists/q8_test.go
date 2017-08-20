package chap2

import (
	"fmt"
	"testing"
)

func TestFindCircularHead(t *testing.T) {
	head1 := &Node{1, nil}
	head1 = head1.AddHead(2)
	head1 = head1.AddHead(4)
	x := head1
	head1 = head1.AddHead(8)
	head1 = head1.AddHead(9)
	head1 = head1.AddHead(111)
	x.next = head1
	fmt.Println(FindCircularHead(head1).data)
}
