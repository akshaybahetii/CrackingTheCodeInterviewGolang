package chap2

func DeleteCurrentNode(n *Node) {
	n.data = n.next.data
	n.next = n.next.next
}
