package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int = 12, 11

	var result int = a + b
	fmt.Println("a + b = ", result)

	var res float32 = float32(a / b)
	fmt.Println("\n a / b = ", res)

	var largeValue int32 = math.MaxInt32
	var smallValue int32 = math.MinInt32

	fmt.Println("large value + 1 = ", largeValue+1)
	fmt.Println("small value - 1 = ", smallValue-1)

}
