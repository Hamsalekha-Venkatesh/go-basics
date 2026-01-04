package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int, 100)
	for i := range 10 {
		channel <- i
		time.Sleep(100000 * time.Nanosecond)
		fmt.Println("sent", i, " to the channel")
	}

	for value := range channel {
		fmt.Println("received: ", value)
	}

	close(channel)

}
