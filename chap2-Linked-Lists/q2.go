package chap2

func FindKthTerm(head *Node, k int) *Node {
	i := head
	count := 0

	for i != nil {
		count++
		i = i.next
	}
	i = head
	count = count - k
	for count != 0 {
		count--
		i = i.next
	}
	return i
}
