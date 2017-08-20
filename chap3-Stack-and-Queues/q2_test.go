package chap3

import "testing"

func TestStackMin(t *testing.T) {
	var s Stack
	s.Push(3)
	s.Push(6)
	s.Push(2)
	s.Push(7)
	s.Push(1)
	s.Push(9)
	if s.Min() != 1 {
		t.Fatalf("Min expected 1 got %d", s.Min())
	}
}
