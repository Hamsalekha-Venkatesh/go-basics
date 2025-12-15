package main

import "fmt"

func main() {
	fmt.Println("ADD oper() calling ", applyOperation(2, 3, add))
	fmt.Println("SUB oper() calling ", applyOperation(5, 3, sub))
	fmt.Println("MUL oper() calling ", applyOperation(2, 3, mul))

	fmt.Println("Demonstrating returning a func...")
	var multiplierByTwo = createMultipler(10)

	fmt.Println("Calling multiplierByTwo(...)...")
	fmt.Println(multiplierByTwo(5))
}

// func that takes another func as argument....
func applyOperation(a int, b int, oper func(int, int) int) int {
	return oper(a, b)
}

func createMultipler(factor int) func(int) int {
	return func(input int) int {
		return input * factor
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}
