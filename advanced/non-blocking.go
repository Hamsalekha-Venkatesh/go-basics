package main

import (
	"fmt"
	"time"
)

func producer(incoming chan int) {
	for i := range 20 {
		incoming <- i
	}

	close(incoming)
}

func filterEvenValues(incoming chan int, even chan int) {
	for data := range incoming {
		if data%2 == 0 {
			even <- data
		}
	}

	close(even)
}

func main() {
	incoming := make(chan int)
	even := make(chan int)

	go producer(incoming)
	go filterEvenValues(incoming, even)

	for val := range even {
		fmt.Println("received", val)
	}

}

func main1() {
	dataChan := make(chan int)

	// now sending data to channel
	go func() {
		fmt.Println("before to close the channel")
		for i := range 10 {
			dataChan <- i
			time.Sleep(1 * time.Millisecond)
		}
		fmt.Println("Going to close the channel")
		close(dataChan)
		close(dataChan)

	}()

	fmt.Println("Waiting for data...after closed the channels")
	for val := range dataChan {
		fmt.Println("data from data chan is: ", val)
	}
}
