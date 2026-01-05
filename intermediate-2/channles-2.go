package main

import "fmt"

func main() {
	// step 1: create a new channel that accepts a string
	var myChannel = make(chan string)

	// step 2: send values to this channel
	go func() {
		myChannel <- "Let's send some value"
		myChannel <- "What is your name?"
		myChannel <- " I'm slaza Slytherine..."

	}()

	// Step 3 : read the values from the channel...
	for _ = range 2 {
		myChannelRx := <-myChannel
		fmt.Println("myChannelRx == ", myChannelRx)
	}

	fmt.Println("I'm done....")

}
