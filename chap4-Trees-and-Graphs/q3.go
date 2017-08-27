package chap4

type LLNode struct {
	data int
	next *LLNode
}

func CreateLinkedList(t *TreeNode, level int, ll **[]*LLNode) {
	if t == nil {
		return
	}
	if len(**ll) < level+1 {
		l := append(**ll, &LLNode{})
		*ll = &l
	}
	CreateLinkedList(t.left, level+1, ll)
	newL := &LLNode{t.data, nil}
	CreateLinkedList(t.right, level+1, ll)

	d := **ll
	x := d[level]
	for x.next != nil {
		x = x.next
	}
	x.next = newL
	return
}
