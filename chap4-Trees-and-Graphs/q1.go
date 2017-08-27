package chap4

type Node struct {
	data    int
	visited bool
}

type QNode struct {
	data int
	next *QNode
}
type Queue struct {
	head *QNode
	last *QNode
}

func (q *Queue) IsEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
func (q *Queue) Enqueue(d int) {
	n := &QNode{d, nil}

	if q.last == nil && q.head == nil {
		q.last = n
		q.head = n
		return
	}
	q.last.next = n
	q.last = n
}

func (q *Queue) Dqueue() int {
	if q.head == nil {
		return -1
	}
	t := q.head
	if q.head == q.last {
		q.last = nil
	}
	q.head = q.head.next

	return t.data
}

func FindRoute(g [][]Node, start int, end int) bool {
	q := Queue{}
	q.Enqueue(start)

	for !q.IsEmpty() {
		x := q.Dqueue()
		for j := 0; j < len(g[x]); j++ {
			if g[x][j].data != 0 && g[x][j].visited != true {
				if j == end {
					return true
				} else {
					g[x][j].visited = true
					q.Enqueue(j)
				}
			}
		}
	}
	return false
}
