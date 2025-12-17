package main

import "fmt"

func main() {
	var a = 10
	var ptr = &a // integer pointers

	fmt.Println("ptr = ", ptr)
	fmt.Println("ptr = ", *ptr)
}

func factorial(n int) int {
	if n == 2 {
		return 2
	}
	return n * factorial(n-1)
}

func digitsSum(n int) int64 {
	var sum int64 = 0
	for n > 0 {
		digit := n % 10
		sum += int64(digit)
		n = n / 10
	}

	return sum
}
