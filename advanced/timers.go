package main

import (
	"fmt"
	"time"
)

func someLongRunnningOper() {
	for i := 0; i < 20; i++ {
		fmt.Println("i = ", i)
		time.Sleep(time.Second)
	}
}

func main() {
	status := make(chan bool)
	timeout := time.After(3 * time.Second)

	go func() {
		someLongRunnningOper()
		status <- true
	}()

	select {
	case <-timeout:
		fmt.Println("timeout operation")
	case <-status:
		fmt.Println("ok...done!")
	}

	// to delay the operation executed...
	newTimer := time.NewTimer(4 * time.Second)
	newChannel := make(chan string)

	newChannel <- "hello"
	newChannel <- "world"

	<-newTimer.C // let the timer expire...
	select {
	case <-newTimer.C:
		fmt.Println("timeout operation")
	case <-newChannel:
		for str := range newChannel {
			fmt.Println(str)
		}
	}
}
