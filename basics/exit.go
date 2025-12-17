package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Exiting...")
		fmt.Println("more exits...trying some cleanups...")
	}()

	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}

	panic("OMG! I'm dying, help me asap")
	// os.Exit(1)
	//fmt.Println("Hello World")
}
