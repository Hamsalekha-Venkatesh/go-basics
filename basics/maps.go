package main

import "fmt"

func main() {
	var mapName map[int]string
	mapName = make(map[int]string)
	mapName[1] = "a"
	mapName[2] = "b"
	mapName[3] = "c"

	fmt.Println(mapName)
	var freqMap = make(map[string]int)
	var words = make([]string, 0)

	words = append(words, "cat", "cat", "dog", "dog", "cat", "mouse", "mouse", "cat")
	for _, word := range words {
		freqMap[word]++
	}

	delete(freqMap, "cat")

	fmt.Println("Count of each words = ", freqMap)
	fmt.Println("freq of cow is  ", freqMap["cow"])
	fmt.Println("freq of cow is  ", freqMap["cat"])

	fmt.Println("Actual Word list", words)

}
