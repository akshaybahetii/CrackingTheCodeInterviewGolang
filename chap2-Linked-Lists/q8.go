package chap2

func FindCircularHead(head1 *Node) *Node {

	s := head1.next
	t := head1.next.next
	for t != nil || s != nil {
		if s == t {
			x := head1
			for x != s {
				x = x.next
				x = s.next
			}
			return x
		}
		s = s.next
		t = t.next.next
	}
	return nil
}
