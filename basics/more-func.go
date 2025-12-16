package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("ADD oper() calling ", applyOperation(2, 3, add))
	fmt.Println("SUB oper() calling ", applyOperation(5, 3, sub))
	fmt.Println("MUL oper() calling ", applyOperation(2, 3, mul))

	fmt.Println("Demonstrating returning a func...")
	var multiplierByTwo = createMultipler(10)

	fmt.Println("Calling multiplierByTwo(...)...")
	fmt.Println(multiplierByTwo(5))

	//q1, r1, isOK1 := divide(12, 0)
	//q2, r2, isOK2 := divide(26, 12)

}

func divide(num1 int, num2 int) (quotient int, remainder int, error) {
	if num2 == 0 {
		return -1, -1, errors.New("Can't divide, you will get divide by zero...")
	}

	quotient = num1 / num2
	remainder = num1 % num2

	return quotient, remainder, nil

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
