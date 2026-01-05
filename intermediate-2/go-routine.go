package main

import (
	"errors"
	"fmt"
	"time"
)

var i int = 0

func main() {
	fmt.Println(time.Now())
	printNumbersUpto(10)
	go func() {
		err := doWork()
		if err != nil {
			fmt.Println("I'm done....")
		}
	}()
	//go increment()
	//go doubleIncrement()
	//go increment()
	//go doubleIncrement()
	go time.Sleep(10 * time.Second)
	//fmt.Println(time.Now())

}

func doWork() error {
	fmt.Println("do work...but Im returning error hahaha")
	increment()
	doubleIncrement()
	return errors.New("sdfsdfdsfsdfdsf ")

}

func increment() {
	i++
	fmt.Println("increment()", i)
}

func doubleIncrement() {
	i++
	fmt.Println("1/2 : doubleIncrement()", i)
	i++
	fmt.Println("2/2 : doubleIncrement()", i)

}

func printNumbersUpto(limit int) {
	for i := 0; i < limit; i++ {
		fmt.Println("i = ", i)
	}
}
