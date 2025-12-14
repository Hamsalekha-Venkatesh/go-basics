package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age                 int
}

func main() {
	const (
		MAX_RETRIES       = 10
		TOTAL_RETRIES     = 15
		HARD_RETRY    int = 20
		DEFAULT_NAME      = "Hamsalekha"
	)

	var id int = 2
	for id < MAX_RETRIES {
		fmt.Println("id = ", id)
		id++
	}
}
