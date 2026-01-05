package main

import (
	"context"
	"fmt"
	"time"
)

func checkEven(ctx context.Context, number int) (string, error) {
	select {
	case <-ctx.Done():
		return "Operation cancelled...:(", ctx.Err()
	default:
		if number%2 == 0 {
			fmt.Println("even number")
		} else {
			fmt.Println("odd number")
		}
	}
	return "", nil
}

func doWork(ctx context.Context) (string, error) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("work cancelled...")
			return "Operation cancelled...:(", ctx.Err()
		default:
			fmt.Println("do work... I'm working!", ctx.Value("requestId"))
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	rootContext := context.Background()
	rootContext, error := context.WithTimeout(rootContext, 2*time.Second) // timer starts here
	defer error()                                                         // defer to the end of main() execution....

	rootContext = context.WithValue(rootContext, "requestId", "1234")
	result, err := doWork(rootContext)
	fmt.Println("requestId = ", rootContext.Value("requestId"))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
