package main

import "fmt"

func main() {
	arr1 := []int{4, 2, 6, 1, 0, 8}
	QucikSort(&arr1)
	fmt.Println(arr1)
}

func QucikSort(arr *[]int) {
	if len(*arr) <= 1 {
		return
	}
	lenA := len(*arr)
	mid := lenA / 2
	arr1 := (*arr)[:mid]
	arr2 := (*arr)[mid+1:]

	QucikSort(&arr1)
	QucikSort(&arr2)
	var i int
	for i = 0; i < mid-1; i++ {
		if arr1[i] < arr2[i] {
			(*arr)[(2 * i)] = arr1[i]
			(*arr)[2*i+1] = arr2[i]
		} else {
			(*arr)[2*i] = arr2[i]
			(*arr)[2*i+1] = arr1[i]
		}
	}
}
