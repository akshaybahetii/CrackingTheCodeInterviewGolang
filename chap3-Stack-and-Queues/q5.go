package chap3

// s -> 5 1 3 10
// t ->
// temp = 11
func TowerofAnnoi(s *Stack) {
	t := &Stack{}
	t.Push(s.Pop())

	for s.len > 0 {
		if t.Peek() > s.Peek() {
			t.Push(s.Pop())
		} else {
			temp := s.Pop()
			for t.Peek() < temp && t.len > 0 {
				s.Push(t.Pop())
			}
			t.Push(temp)
		}
	}
	*s = *t
	return
}
