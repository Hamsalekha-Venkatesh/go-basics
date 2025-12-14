package main

import (
	"fmt"
)

func main() {
	var age = 51
	fmt.Println("age = ", age)

	switch {
	case age < 10:
		fmt.Print("age is less than 10")
	case age < 50:
		fmt.Print("age is less than 50")
	case age < 75:
		fmt.Print("age is less than 75")
	default:
		fmt.Println("age is greater than 10")
	}

	var fruit = "pineapple"
	switch fruit {
	case "apple":
		fmt.Println("apple")
	case "banana":
		fmt.Println("banana")
	case "cherry":
		fmt.Println("cherry")
	case "grape":
		fmt.Println("grape")
	default:
		fmt.Println(fruit, " is not a fruit!")
	}

	checkType(10)
	checkType(10.3434)
	checkType("Hamsa")
	checkType(true)

}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case float64:
		fmt.Println("float64")
	default:
		fmt.Println("I don't know how to handle")
	}
}
