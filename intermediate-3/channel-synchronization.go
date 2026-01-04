package main

import (
	"fmt"
	"time"
)

func main() {
	numGoroutines := 3
	doneChannel := make(chan int, numGoroutines)
	for i := range numGoroutines {
		go func(id int) {
			fmt.Println(id)
			time.Sleep(1 * time.Nanosecond)
			doneChannel <- id
		}(i)
	}

	for _ = range numGoroutines {
		<-doneChannel
		fmt.Println("wait for this go routine to be finished...")
	}

	fmt.Println(doneChannel)
}

func main1() {
	done := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		done <- struct{}{} // send channel
	}()

	<-done // recieve channel
	fmt.Println("done...!")
}
