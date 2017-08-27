package chap4

import (
	"fmt"
	"testing"
)

func TestFindRoute(t *testing.T) {
	graph := make([][]Node, 5)
	for i := 0; i < 5; i++ {
		graph[i] = make([]Node, 5)
	}

	graph[0][1].data = 1
	graph[1][2].data = 1
	graph[2][3].data = 1
	graph[3][4].data = 1

	fmt.Println(FindRoute(graph, 0, 4))

}
