package chap4

import (
	"fmt"
	"testing"
)

func TestCreateGraphFromDep(t *testing.T) {
	arr := []int{0, 3, 5, 1, 1, 3, 5, 0, 3, 2}
	fmt.Println(arr)
	fmt.Println(CreateGraphFromDep(arr))

}
