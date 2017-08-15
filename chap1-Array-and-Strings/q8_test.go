package chap1

import (
	"fmt"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	arr := [][]int{
		{1, 2, 3},
		{4, 0, 6},
		{7, 8, 9},
	}
	resP := ZeroMatrix(&arr)
	res := *resP
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println("-")
	}
}
