package chap4

import "fmt"

func CreateGraphFromDep(arr []int) []int {
	g := make([][]Node, 7)
	for i := 0; i < 7; i++ {
		g[i] = make([]Node, 7)
	}

	for i := 0; i < len(arr)-1; i = i + 2 {
		g[arr[i]][arr[i+1]].data = 1
		g[arr[i+1]][arr[i]].data = 1
	}

	q := Queue{}
	q.Enqueue(arr[len(arr)-1])

	var res []int
	for !q.IsEmpty() {
		x := q.Dqueue()
		fmt.Println(x)
		res = append(res, x)
		for j := 0; j < len(g[x]); j++ {
			if g[x][j].data != 0 && g[x][j].visited != true && g[j][x].visited != true {
				g[x][j].visited = true
				q.Enqueue(j)
			}
		}
	}
	return res
}
