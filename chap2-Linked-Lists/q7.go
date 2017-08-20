package chap2

func FindIntersection(head1 *Node, head2 *Node) bool {
	t := head1
	s := head2
	for t != nil {
		s = head2
		for s != nil {
			if s == t {
				return true
			}
			s = s.next
		}
		t = t.next
	}
	return false
}
