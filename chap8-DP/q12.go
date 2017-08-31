package main

/*
Eight Queens:Write an algorithm to print all ways of arranging eight queens on an 8x8 chess board
so that none of them share the same row, column, or diagonal. In this case, "diagonal" means all
diagonals, not just the two that bisect the board.

N-Queen Problem, Backtracking.
*/
import "fmt"

func CheckValid(board [8][8]int) bool {
	for i, _ := range board {
		for j, _ := range board[i] {
			if board[i][j] == 1 {
				//Check if current col is empty
				for k := 0; k < 8; k++ {
					if i != k && board[k][j] != 0 {
						return false
					}
				}
				//Check if current row is empty
				for k := 0; k < 8; k++ {
					if j != k && board[i][k] != 0 {
						return false
					}
				}
				//Check if diagonal is empty
				for k := 1; k < 8; k++ {
					if i-k < 8 && i-k > 0 && j-k > 0 && j-k < 8 && board[i-k][j-k] != 0 {
						return false
					}
					if i+k < 8 && j+k < 8 && board[i+k][j+k] != 0 {
						return false
					}
				}
			}
		}
	}
	return true
}
func BackTrack(board [8][8]int, start int, sol *[8][8]int) bool {
	if start > 7 {
		//Found solution
		return true
	}

	for j := 0; j < 8; j++ {
		board[start][j] = 1
		if CheckValid(board) && BackTrack(board, start+1, sol) {
			fmt.Println(start, j)
			(*sol)[start][j] = 1
			return true
		} else {
			board[start][j] = 0
		}
	}
	return false
}

func main() {
	var sol [8][8]int
	var board [8][8]int
	BackTrack(board, 0, &sol)
	fmt.Println(sol[0])
	fmt.Println(sol[1])
	fmt.Println(sol[2])
	fmt.Println(sol[3])
	fmt.Println(sol[4])
	fmt.Println(sol[5])
	fmt.Println(sol[6])
	fmt.Println(sol[7])

}
