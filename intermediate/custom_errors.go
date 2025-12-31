package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	message   string
	err       error
	errorCode int
}

// Error returns error message - implementing Error() method of error interface
func (e *CustomError) Error() string {
	return fmt.Sprintf("\n ERROR!!! Message - %s \n errorCode 0 - %d \n Wrapped Err - %s",
		e.message,
		e.errorCode,
		e.err,
	)
}

func divide(a, b int) (int, CustomError) {
	if b == 0 {
		var error = CustomError{message: "divide by zero", errorCode: -1, err: errors.New("Wrapped divide by zero")}

		return -1, error
	}

	return a / b, CustomError{message: "", errorCode: 0, err: nil}
}

func main() {
	var result, err = divide(10, 0)
	if err.errorCode != 0 {
		fmt.Println("Oops division failed", err.Error())
		// panic(err)
		return
	}

	fmt.Println(result)

}
