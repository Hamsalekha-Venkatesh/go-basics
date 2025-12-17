package main

import "fmt"

func main() {
	// func --> func(int) int
	subtractor := func() func(int) int {
		var countdown = 100
		return func(delta int) int {
			countdown -= delta
			//fmt.Println(countdown)

			return countdown
		}
	}()

	fmt.Println(subtractor(10))
	fmt.Println(subtractor(10))
	fmt.Println(subtractor(10))

}
