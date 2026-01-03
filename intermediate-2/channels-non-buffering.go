package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 1
	go func() {
		ch <- 2
		reciver := <-ch //immediately they need to send the value...
		fmt.Println(reciver)
	}()
}
