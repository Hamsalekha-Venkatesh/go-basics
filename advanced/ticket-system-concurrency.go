package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
)

/**
Ticket selling and buying
*/

var (
	TOTAL_AVAILABLE_TICKETS = 100
	CONSUMED_TICKETS        = 0
	totalRequest            = 15
	totalWorker             = 3
)

type TicketRequest struct {
	personId   int
	numTickets int
}

type Result struct {
	personId int
	cost     int
	err      error
}

func ticketProcessor(requestChan <-chan TicketRequest, resultsChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for request := range requestChan {
		if request.numTickets+CONSUMED_TICKETS > TOTAL_AVAILABLE_TICKETS {
			fmt.Println("Show is full, can't purchase. Aborting tx")
			resultsChan <- Result{personId: -1, cost: -1, err: errors.New("Show is full, can't purchase. Aborting tx")}
			continue
		}

		CONSUMED_TICKETS += request.numTickets

		finalCost := 20 * request.numTickets
		fmt.Printf("Person %d Booked %d tickets. Total cost: %d \n ", request.personId, request.numTickets, finalCost)
		resultsChan <- Result{personId: request.personId, cost: finalCost, err: nil}

	}
}

func main() {
	var waitGroups sync.WaitGroup

	requestChan := make(chan TicketRequest, totalRequest)
	resultsChan := make(chan Result, totalRequest)

	waitGroups.Add(totalWorker)

	for range totalWorker {
		go ticketProcessor(requestChan, resultsChan, &waitGroups)
	}

	for i := 1; i <= totalRequest; i++ {
		ticketRequest := TicketRequest{
			personId:   i,
			numTickets: rand.Intn(TOTAL_AVAILABLE_TICKETS / 5),
		}
		requestChan <- ticketRequest
	}

	close(requestChan)

	// non-blocking mechanisms... get values in real time...
	go func() {
		waitGroups.Wait()
		close(resultsChan)
	}()

	for result := range resultsChan {
		if result.err != nil {
			fmt.Println("person: ", result.personId, "cost: ", result.cost)
		}
	}
}
