package main

import (
	"errors"
	"fmt"
)

func main() {
	if ans := processErrorData(-1); ans != nil {
		_ = fmt.Errorf("%w", ans)
	}
}

func processErrorData(d int) error {
	if d < 0 {
		return errors.New("cannot process negative numbers")
	}
	return nil
}
