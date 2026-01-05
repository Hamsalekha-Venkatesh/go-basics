package main

import (
	"fmt"
	"strconv"
	"time"
)

type Job struct {
	JobId  int
	Name   string
	Status string
}

func assignWorker(workerId int, tasks <-chan Job, results chan<- int) {
	for job := range tasks {
		time.Sleep(1 * time.Second)
		job.Status = "COMPLETED"
		fmt.Println("TASK CHANNEL: Worker ", workerId, "completed job", job.JobId)
		results <- 10 * job.JobId
	}
}

func main() {
	numWorkers := 4
	numJobs := 10

	tasks := make(chan Job, numJobs)   // new jobs
	results := make(chan int, numJobs) // completed jobs

	// start workers
	for i := 0; i < numWorkers; i++ {
		go assignWorker(i, tasks, results) //
	}

	// Create worker pools
	for i := range numJobs {
		nextJob := Job{
			JobId:  i,
			Name:   "Worker" + strconv.Itoa(i),
			Status: "NOT_STARTED",
		}
		tasks <- nextJob
	}

	close(tasks)

	for i := 0; i < numJobs; i++ {
		jobId := <-results
		fmt.Println("RESULT CHANNEL: ", jobId, " is done executing....")

	}

	fmt.Println("All jobs completed")

}
