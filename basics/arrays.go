package main

import (
	"fmt"
)

func main() {
	var numbers [10]int
	i := 1

	for index, _ := range numbers {
		numbers[index] = i
		i++
		fmt.Println(index, numbers[index])
	}

	fmt.Print(numbers)
	fruits := [4]string{"Apple", "Orange", "Pear", "Grape"}

	for fruit := range fruits {
		fmt.Println("Fruit: ", fruit)
	}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Formatted str  %s", fruits[i])
	}

	q, rem := divide(100, 11)
	fmt.Println("Let's practice some division...", q, rem)

}

func divide(a int, b int) (quotient, remainder int) {
	quotient = a / b
	remainder = a % b
	return quotient, remainder
}
