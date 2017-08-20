package chap2

func CheckPalindrome(head1 *Node) bool {
	t := head1
	d := head1
	var head2 *Node
	for d != nil {
		if d.next == nil {
			d = nil
			t = t.next
			continue
		}
		d = d.next.next
		t = t.next
		head2 = head2.AddHead(t.data)
	}
	for head2 != nil {
		if head2.data != t.data {
			return false
		}
		head2 = head2.next
		t = t.next
	}
	//t is at halfway
	//d is at end
	return true
}
