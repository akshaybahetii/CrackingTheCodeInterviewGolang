package chap2

func AddTwoNum(head1 *Node, head2 *Node) *Node {
	carry := 0
	var head3 *Node
	for head1 != nil && head2 != nil {
		x := head1.data + head2.data + carry
		carry = x / 10
		head3 = head3.AddLast(x % 10)
		head1 = head1.next
		head2 = head2.next
	}
	for head1 != nil {
		head3.AddLast(head1.data)
		head1 = head1.next
	}
	for head2 != nil {
		head3.AddLast(head2.data)
		head2 = head2.next
	}

	return head3
}
