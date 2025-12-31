package main

import (
	"fmt"
	"strings"
)

func main() {
	var input strings.Builder

	input.WriteString("Hamsa")
	input.WriteString("lekha")

	fmt.Println(input.String())

	var input2 strings.Builder

	input2.WriteString("Go is awesome")
	input2.WriteString("lekha")
	input2.WriteRune('\n')
	input2.WriteString("Go is awesome")
	fmt.Println(input2.String())
}
