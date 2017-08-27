package chap4

import "fmt"

type SNode struct {
	ptr  interface{}
	next *SNode
}

type Stack struct {
	head *SNode
	len  int
}

func (s *Stack) IsEmpty() bool {
	if s.len == 0 {
		return true
	}
	return false
}
func (s *Stack) Push(d interface{}) {
	new := SNode{d, s.head}
	s.len++
	s.head = &new
}

func (s *Stack) Peep() interface{} {
	if s.head == nil {
		s.len = 0
		return nil
	}
	return s.head.ptr
}

func (s *Stack) Pop() interface{} {
	if s.head == nil {
		s.len = 0
		return nil
	}
	n := s.head
	s.head = s.head.next
	s.len--
	return n.ptr
}

func FirstComAncestor(t *TreeNode, x int, y int) int {
	s := Stack{}

	s.Push(t)

	for !s.IsEmpty() {
		pp := s.Peep()
		p := pp.(*TreeNode)
		if p.data == x {
			fmt.Println("Found X", x)
			//Found X
		}
		if p.data == y {
			//Found Y
			gg := s.Peep()
			g := gg.(*TreeNode)
			fmt.Println("Found Y", y, g.data)
		}
		p.visited = true

		if p.left != nil && p.left.visited != true {
			s.Push(p.left)
			continue
		}
		_ = s.Pop()
		if p.right != nil && p.right.visited != true {
			s.Push(p.right)
			continue
		}

	}
	return -1
}
