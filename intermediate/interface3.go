package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}

type Dog1 struct{}

func (d Dog1) Speak() string {
	return "Woof!"
}

func main() {
	var s Speaker
	d := Dog1{}
	fmt.Println(s == d)
}
