package main

import "fmt"

func main() {
	maze := make([][]int, 4)
	sol := make([][]int, 4)
	for i, _ := range sol {
		sol[i] = make([]int, 4)
	}
	maze[0] = []int{1, 0, 0, 0}
	maze[1] = []int{1, 1, 0, 1}
	maze[2] = []int{0, 1, 0, 0}
	maze[3] = []int{1, 1, 1, 1}

	FindSol(maze, 0, 0, &sol)
	fmt.Println(sol)
}

func FindSol(maze [][]int, i int, j int, sol *[][]int) bool {
	if (i == 3) && (j == 3) {
		return true
	}
	//Go right
	if i+1 < 4 && j < 4 && maze[i+1][j] != 0 && FindSol(maze, i+1, j, sol) {
		(*sol)[i+1][j] = 1
		return true
	}
	if j+1 < 4 && i < 4 && maze[i][j+1] != 0 && FindSol(maze, i, j+1, sol) {
		(*sol)[i][j+1] = 1
		return true
	}
	return false
}
