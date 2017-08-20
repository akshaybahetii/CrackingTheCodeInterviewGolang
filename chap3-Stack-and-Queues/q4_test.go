package chap3

import (
	"fmt"
	"testing"
)

func TestQueueWithStacks(t *testing.T) {
	var q Queue
	q.s1 = &Stack{}
	q.s2 = &Stack{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	fmt.Println(q.Dqueue())
	fmt.Println(q.Dqueue())
	fmt.Println(q.Dqueue())
	fmt.Println(q.Dqueue())
	fmt.Println(q.Dqueue())
}
