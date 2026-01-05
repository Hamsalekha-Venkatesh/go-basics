package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	constructionWorkerGroup sync.WaitGroup
)

type ConstructionWorker struct {
	workerId int
	taskName string
}

// Simulate worker pointer receiver....
func (worker *ConstructionWorker) performTask(constructionWorkerGroup *sync.WaitGroup) {
	defer constructionWorkerGroup.Done()
	fmt.Printf("Worker %d started task %s\n", worker.workerId, worker.taskName)
	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func main() {
	// defines tasks by the worker...
	taskNames := []string{
		"brushing teeth",
		"going to gym",
		"bathing",
		"Running to work",
		"Lunching",
		"Sending family time",
	}

	for index, taskName := range taskNames {
		myTask := ConstructionWorker{workerId: index, taskName: taskName}

		// Add each worker one by one...
		constructionWorkerGroup.Add(1)

		// each task is a concurrent go routine...
		go myTask.performTask(&constructionWorkerGroup)
	}

	constructionWorkerGroup.Wait()
	fmt.Println("Done with ALL TASKS")
}
