package chap2

import "fmt"

type Node struct {
	data int
	next *Node
}

//Add node to linkedlist
func (n *Node) Add(d int) {
	new := Node{d, n.next}
	n.next = &new
}

//Add head node to linkedlist
func (n *Node) AddHead(d int) *Node {
	new := Node{d, n}
	return &new
}

//Add last node to linkedlist
func (n *Node) AddLast(d int) *Node {
	if n == nil {
		new := Node{d, nil}
		return &new
	} else {
		t := n
		for t.next != nil {
			t = t.next
		}
		new := Node{d, nil}
		t.next = &new
		return n
	}
}

//Delete next node
func (n *Node) DeleteNext() {
	if n.next == nil {
		fmt.Println("Fail cannot delete this is end of linkedlist")
		return
	}
	n.next = n.next.next
}

func (head *Node) Print() {
	var t *Node
	t = head
	fmt.Println("List: ")
	for t != nil {
		fmt.Printf("%d-", t.data)
		t = t.next
	}
}

func RemoveDuplicates(head *Node) {
	i := head
	for i != nil {
		j := i
		for j.next != nil {
			if j.next.data == i.data {
				j.DeleteNext()
			} else {
				j = j.next
			}
		}
		i = i.next
	}
}
