package chap1

import (
	"fmt"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	arr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	res := RotateMatrix(arr)
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println("-")
	}
}
