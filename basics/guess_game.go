package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var low, high int = 0, 100
	var originalNum = rand.Intn(high-low+1) + low

	var totalGuessCount = 0
	fmt.Println("Enter the guess count:")
	fmt.Scanln(&totalGuessCount)

	for low <= high {
		var mid int = low + (high-low)/2
		totalGuessCount++

		if mid == originalNum {
			fmt.Print("Guessing stops for the number {", originalNum, "}.. you found the answer in ", totalGuessCount, " moves")
			break
		} else if mid > originalNum {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}
