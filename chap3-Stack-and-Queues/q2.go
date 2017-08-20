package chap3

import "fmt"

type Node struct {
	data int
	next *Node
}

type Stack struct {
	head *Node
	len  int
	min  int
}

func (s *Stack) Push(d int) {
	new := Node{d, s.head}
	s.len++
	if s.head == nil {
		s.min = d
	}
	if s.min > d {
		s.min = d
	}
	s.head = &new
}

func (s *Stack) Peek() int {
	if s.head == nil {
		return -1
	}
	return s.head.data

}
func (s *Stack) Pop() int {
	if s.head == nil {
		s.len = 0
		return -999
	}
	n := s.head
	s.head = s.head.next
	s.len--
	return n.data
}

func (s *Stack) Print() {
	t := s.head
	fmt.Println("List: ")
	for t != nil {
		fmt.Printf("%d-", t.data)
		t = t.next
	}
}

func (s *Stack) Min() int {
	return s.min
}
