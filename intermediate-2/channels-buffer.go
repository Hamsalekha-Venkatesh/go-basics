package main

import (
	"fmt"
	"time"
)

func main() {
	channel2 := make(chan int, 2)
	channel2 <- 10
	channel2 <- 20
	go func() {
		time.Sleep(2 * time.Second)
	}()

	//fmt.Println(<-channel2)
	channel2 <- 30
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)
	fmt.Println(<-channel2)

}

func main1() {

	// make(chan Type, capacity)
	bufferChanel := make(chan int, 3)
	bufferChanel <- 10
	bufferChanel <- 20
	bufferChanel <- 30
	fmt.Println("I'm done....")
	reciver := <-bufferChanel
	fmt.Println(reciver)
	reciver = <-bufferChanel
	fmt.Println(reciver)
	bufferChanel <- 40
	bufferChanel <- 50
	reciver = <-bufferChanel
	fmt.Println(reciver)
	reciver = <-bufferChanel
	fmt.Println(reciver)

}
