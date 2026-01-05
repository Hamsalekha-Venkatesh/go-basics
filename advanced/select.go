package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch2 <- 3
	ch2 <- 4

	for {
		select {
		case message, ok := <-ch1:
			if !ok {
				fmt.Println("channel 1 closed")
				return
			}
			fmt.Println("received from ch1", message)
		case message, ok := <-ch2:
			if !ok {
				fmt.Println("channel 2 closed")
				return
			}
			fmt.Println("received from ch2", message)
		default:
			fmt.Println("no message received")
		}
	}

}
