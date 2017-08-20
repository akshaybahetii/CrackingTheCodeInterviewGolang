package chap2

func MoveAllNodeLessThanX(head *Node, x int) *Node {
	t := head
	for t.next != nil {
		if t.next.data < x {
			head = head.AddHead(t.next.data)
			t.DeleteNext()
		} else {
			t = t.next
		}
	}
	return head
}
