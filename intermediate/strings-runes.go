package main

import "fmt"

func main() {
	message := "Hello world!"
	backtick_strings := `raw\tstring literal` // does not formar escape sequence

	fmt.Println(backtick_strings, " ", message)
	fmt.Println(len(message))

	// string comparisons
	message1 := "apple"
	message2 := "banana"
	message3 := "app"
	fmt.Println(message1 < message2) // true - ASCII val comparison
	// | lexicographical comparisons..
	fmt.Println(message1 < message3) // false - lexicographical comparisons..

	/**
	%x - hex value
	%c - char value
	%d - int value
	%f - float value
	%v - actual value of rune
	%T - data type of variable
	*/
	for _, char := range message1 {
		fmt.Printf("%c\n", char) // %c - formatting verbs
	}

	// Runes
	var ch rune = 'a'
	fmt.Println(string(ch))
	fmt.Printf("Type of rune %T\n", message) //string
	fmt.Printf("Type of rune %T\n", ch)      //int32
	fmt.Printf("Type of rune %T\n", 23)      //int
}
