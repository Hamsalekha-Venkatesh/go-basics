package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // step 1: create a waitgroups
	numWorkers := 3
	wg.Add(numWorkers) // step 2: Add all the workers count to this waitgroups...

	for i := range numWorkers {
		workers(i, &wg)
	}

	go func() {
		wg.Wait() // wait until all goroutines are returned ...
		fmt.Println("Done")
	}()

}

func workers(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Step 4: call done after execution of goroutine.
	// postpone this action till the end of the func call...

	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}
