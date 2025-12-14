package main

import "fmt"

var globalName string

func main() {
	var age int
	var name string = "Hamsa"
	var name1 string = "lekha"
	var isFemale bool

	age = 32
	isFemale = true
	globalName = "UNNOTICEABLY"

	fmt.Println("name = ", name+name1)
	fmt.Println("age = ", age)
	fmt.Println("Is female = ", isFemale)
	fmt.Println("MIddle Name", globalName)
	getName()
	fmt.Println("MIddle Name", globalName)
}

func getName() {
	var firstName string = "Hamsalekha"
	fmt.Println(firstName, globalName)
	globalName = "Nandakumar"
}
