package main

import (
	"errors"
	"fmt"
	"math"
)

func processData(data []byte) error {
	if len(data) == 0 || data == nil {
		return errors.New("invalid data")
	}

	return nil
}

func sqrt(input float64) (float64, error) {
	if input < 0 {
		fmt.Println("input is negative, so cant compute sqrt()")

		return -1, errors.New("input is negative")
	}

	return math.Sqrt(input), nil
}

func main() {
	var ans1, err1 = sqrt(4)
	if err1 != nil {
		fmt.Println(ans1)
	}

	var ans2, err2 = sqrt(-4)
	if err2 != nil {
		fmt.Println(ans2)
	}

	var data []byte
	if res := processData(data); res != nil {
		fmt.Println("Error processing data")
	} else {
		fmt.Println("Correct data is procesed")
	}

}
