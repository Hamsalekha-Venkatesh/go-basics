package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	sequence1 := adder()
	fmt.Println(sequence1()) // 1
	fmt.Println(sequence1()) // 2
	fmt.Println(sequence1()) // 3
	fmt.Println(sequence1()) // 4

	sequence2 := adder()
	fmt.Println(sequence2()) // 1
	fmt.Println(sequence2()) // 2
	fmt.Println(sequence2()) // 3
	fmt.Println(sequence2()) // 4

	fmt.Println(sequence1()) // 5

}

func adder() func() int {
	i := 0
	fmt.Println("i = ", i)
	return func() int {
		i++
		return i
	}
}
