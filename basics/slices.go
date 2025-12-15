package main

import (
	"fmt"
	"slices"
)

func main() {
	var numbers = []int{34, 55, 343, 4445, 56445, 23234, 3343}
	var numSlices = numbers[2:5]
	fmt.Println(numSlices)
	numSlices = append(numSlices, 12, 13, 14)
	fmt.Println(numSlices)

	newSlices := make([]int, 2)
	newSlices = append(numbers, 12, 21, 22, 444, 234534434)
	fmt.Println(newSlices)
	if slices.Equal(numbers, numSlices) {
		fmt.Println("slices equal")
	} else {
		fmt.Println("slices not equal")
	}

	// 2 D slices
	var matrixSlice = make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		matrixSlice[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			matrixSlice[i][j] = i + j
		}
	}

	fmt.Println(matrixSlice)

}
