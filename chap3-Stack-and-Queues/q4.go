package chap3

type Queue struct {
	s1 *Stack
	s2 *Stack
}

func (q *Queue) Enqueue(d int) {
	q.s1.Push(d)
}

func (q *Queue) Dqueue() int {
	for q.s1.len != 0 {
		q.s2.Push(q.s1.Pop())
	}
	t := q.s2.Pop()
	for q.s2.len != 0 {
		q.s1.Push(q.s2.Pop())
	}
	return t
}
