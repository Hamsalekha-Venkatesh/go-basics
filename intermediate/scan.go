package main

import "fmt"

func main() {
	var input, password string
	fmt.Printf("Enter name")
	fmt.Scanln(&input)
	fmt.Printf("Enter password")
	fmt.Scanln(&password)

	fmt.Println("Input = ", input)
	fmt.Println("Password = ", password)
}
