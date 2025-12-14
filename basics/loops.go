package main

import "fmt"

func main() {
	k := 12
	var j int = 12
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
	}

	fmt.Println("\n k = ", k, " j = ", j)

	numbers := []int{12, 2, 54, 45, 33, 44}
	for index, number := range numbers {
		fmt.Println("index = ", index, "number = ", number)
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println("i = ", i)
	}

	for num := range 10 {
		fmt.Println("num = ", num)
	}

	i := 2
	for i < 10 {
		fmt.Println("pseudo while{} i =", i)
		i++
	}

}
