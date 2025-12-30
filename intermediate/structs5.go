package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

type Employee1 struct {
	Person1
	company string
	id      int
}

func (p Person1) introduce() {
	fmt.Println("I am Person1")
}

func (e Employee1) introduce() {
	fmt.Println("overridden: I am Employee1")
}

func main() {
	var e1 = Employee1{
		Person1: Person1{name: "Hamsa", age: 30},
		company: "Netflix",
		id:      101,
	}
	e1.introduce()
	e1.Person1.introduce()
}
