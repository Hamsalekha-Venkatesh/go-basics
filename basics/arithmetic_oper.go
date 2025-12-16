package main

import "fmt"

func main() {
	var nums = make([]int, 10)
	nums = append(nums, 0, 1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println(rangeSum(nums...))
}

// we are saying there can be unlimited params of type int.
func rangeSum(nums ...int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}

	return
}
