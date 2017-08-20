package chap3

import "testing"

func TestTowerofAnnoi(t *testing.T) {
	s := &Stack{}
	s.Push(5)
	s.Push(1)
	s.Push(7)
	s.Push(3)
	s.Push(10)
	TowerofAnnoi(s)
	s.Print()
}
