package main

import "fmt"

func main() {
	// process(10)
	process(-10)
	fmt.Println("Normal stmts after panic")
	// process(20)
}

func process(i int) {
	defer fmt.Println("process end...a deferreds stmt 1", i)

	//defer func() {
	//	r := recover()
	//	if r != nil {
	//		fmt.Println(r)
	//	}
	//} ()
	//
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic:", r)
		}
	}()

	if i < 0 {
		panic("Can't process negative numbers")
	}

	i++

}
