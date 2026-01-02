package main

import (
	"fmt"
	"time"
)

type Student struct {
	Class  int
	Name   string
	Course string
}

func main() {
	var greeting = make(chan Student)

	go func() {
		greeting <- Student{1, "Gopher", "Golang"}
		greeting <- Student{2, "Gopher2", "Golang2"}
	}()

	go func() {
		reciever := <-greeting // blocking because its continously recieving values...
		fmt.Println(reciever)
		reciever = <-greeting // blocking because its continously recieving values...
		fmt.Println(reciever)
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("done....!")
}
