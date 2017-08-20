package chap2

import (
	"fmt"
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	head1 := &Node{1, nil}
	head1 = head1.AddHead(2)
	head1 = head1.AddHead(3)
	head1 = head1.AddHead(3)
	head1 = head1.AddHead(8)
	head1 = head1.AddHead(2)
	head1 = head1.AddHead(1)
	fmt.Println(CheckPalindrome(head1))
}
