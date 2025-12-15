package main

import "fmt"

func main() {
	var inputWords = make([]string, 0)
	inputWords = append(inputWords, "word", "text", "book", "word", "cat", "horse")
	var freqMap = make(map[string]int)

	for _, word := range inputWords {
		freqMap[word]++

		if freqMap[word]%2 == 0 {
			delete(freqMap, word)
		}
	}

	for word, count := range freqMap {
		fmt.Println(word, count)
	}

	freqMap["Hamsa"] = 0

	value, isPresent := freqMap["Hamsa"]
	fmt.Println(value, isPresent)

	value1, isPresent1 := freqMap["Nanda"]
	fmt.Println(value1, isPresent1)

	var map4 = make(map[string]int)
	fmt.Println(len(map4))
	if map4 == nil {
		fmt.Println("map4 is nil")
	}

}
