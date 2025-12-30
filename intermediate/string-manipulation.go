package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var name1 = "Hamsa"
	var name2 = "lekha"

	fmt.Println(len(name1))
	fmt.Println(name1 + name2)
	fmt.Println(name1[1:4])

	// String conversions
	num := 10
	str3 := strconv.Itoa(num)
	fmt.Println(str3)
	fruits := "Apple,Orange,Banana,Grapes"
	var fruitsSeparated = strings.Split(fruits, ",")
	fmt.Println(fruitsSeparated)

	countries := []string{"Germany", "India", "Iceland"}
	fmt.Println(strings.Join(countries, ", "))

	str5 := "hello 123 go! vallhalla"
	re, _ := regexp.Compile(`\d`)
	fmt.Println("MATCHED STRINGS", re.FindAllString(str5, -1))

}
