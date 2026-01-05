package main

import (
	"fmt"
	"time"
)

func assignWorkerPool(workerId int, inputChan <-chan int, resultsChan chan<- int) {
	// this `workerId` processes the data from `inputChan` and writing the data to `resultsChan`

	for newJob := range inputChan {
		fmt.Println("Assigning worker", workerId, "to Job: ", newJob)
		time.Sleep(2 * time.Second)
		resultsChan <- newJob
		fmt.Println("Done processing job", newJob)
	}
}

func main() {
	numWorkers := 3
	numTasks := 15

	inputChan := make(chan int, numTasks)
	resultsChan := make(chan int, numTasks)

	// create multiple worker Pools
	for workerId := range numWorkers {
		go assignWorkerPool(workerId, inputChan, resultsChan) // concurrent routine
	}

	// start pushing tasks into the input channel....
	for taskId := range numTasks {
		inputChan <- taskId
	}

	close(inputChan)

	// Synchronize and print the results ...
	for completedTaskId := range resultsChan {
		fmt.Println("ResultChan: processed ", completedTaskId)
	}

	close(resultsChan)
}
