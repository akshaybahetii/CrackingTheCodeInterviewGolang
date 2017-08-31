package main

import "fmt"

/*
This program implements a LRU cache using a doubly queue and hash map.
*/
type DQNode struct {
	data       *CustomURL
	next, prev *DQNode
}

type Queue struct {
	head *DQNode
	tail *DQNode
}
type CustomURL struct {
	Site string
	URL  string
}

type MapData struct {
	customData *CustomURL
	qNode      *DQNode
}

type LRU struct {
	cap int
	q   Queue
	HM  map[string]MapData
}

func (q *Queue) Enqueue(qdata *CustomURL) *DQNode {
	x := DQNode{data: qdata}

	if q.head == nil && q.tail == nil {
		//First Node in Queue
		q.tail = &x
	} else {
		// I am not onl y one
		x.next = q.head
		q.head.prev = &x
	}
	x.prev = nil
	q.head = &x
	return &x
}

func (q *Queue) Dequque() *CustomURL {
	if q.tail == nil {
		return nil
	}
	//last will be removed.
	last := q.tail
	if q.tail.prev != nil {
		//Not the last in queue
		q.tail = q.tail.prev
		q.tail.next = nil
	} else {
		//Deleting last node.
		q.tail = nil
		q.head = nil
	}
	return last.data
}

func (q *Queue) Print() {
	t := q.head
	for t != nil {
		fmt.Println(t.data)
		t = t.next
	}

}
func (q *Queue) AnyDequque(c *DQNode) {
	//Consider if this is last element LASTERTODO.
	if c.prev == nil && c.next == nil {
		q.head = nil
		q.tail = nil
	}
	prev := c.prev //(facebook)
	next := c.next
	if prev != nil {
		//Update next for the one before me
		prev.next = next
	} else {
		//We are deleting the head.
		q.head = next
	}
	if next != nil {
		// Update prev for the one after me
		next.prev = prev
	} else {
		//We are deleting the tail
		q.tail = prev
	}

}

func (lru *LRU) LRUAdd(c *CustomURL) {
	if len(lru.HM) == lru.cap {
		//Dequeue last.
		last := lru.q.Dequque()
		//Delete from HasmMAp
		delete(lru.HM, last.Site)
	}
	added := lru.q.Enqueue(c)

	lru.HM[c.Site] = MapData{customData: c, qNode: added}
}

func (lru *LRU) LRUDelete(site string) {
	n, ok := lru.HM[site]
	if !ok {
		return
	}
	lru.q.AnyDequque(n.qNode)
	delete(lru.HM, site)
}

func (lru *LRU) LRULookUp(site string) *CustomURL {
	gotit := lru.HM[site]
	lru.q.AnyDequque(gotit.qNode)
	gotit.qNode = lru.q.Enqueue(gotit.qNode.data)
	lru.HM[site] = gotit
	return gotit.customData
}

func main() {
	lru := LRU{}
	lru.cap = 5
	lru.HM = make(map[string]MapData)
	lru.LRUAdd(&CustomURL{"Google", "www.google.com"})
	lru.LRUAdd(&CustomURL{"FaceBook", "www.facebook.com"})
	lru.LRUAdd(&CustomURL{"Twitter", "www.twitter.com"})
	lru.q.Print()
	fmt.Println("Looked up google the new queue is ")
	lru.LRULookUp("Google")
	lru.q.Print()
	fmt.Println("Looked up fb the new queue is ")
	lru.LRULookUp("FaceBook")
	lru.q.Print()
	lru.LRUDelete("Google")
	fmt.Println(lru.HM)
	fmt.Println("Delete Google the new queue is ")
	lru.q.Print()
	fmt.Println("Looked up twitter the new queue is ")
	lru.LRULookUp("Twitter")
	lru.q.Print()

}
