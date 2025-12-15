package main

import "fmt"

func main() {
	var message = "Hello World, welcome to my world!"

	for i, c := range message {
		fmt.Printf("%d: %c\n", i, c) // c is called as Rune in go.
		fmt.Println(i, c)
		fmt.Println("____________________________________")
	}

	var myMap = make(map[string]int)
	myMap["a"] = 10
	myMap["b"] = 20
	myMap["c"] = 30

	totalLength, lastValue := CalculateLengthOfMap(myMap)
	fmt.Println("Total count of key = ", totalLength, " last value = ", lastValue)
}

func CalculateLengthOfMap(namesMap map[string]int) (int, string) {
	if namesMap == nil {
		return 0, ""
	}

	totalLength := 0
	lastValue := ""
	for key, value := range namesMap {
		totalLength += value
		lastValue = key
	}
	return totalLength, lastValue
}
