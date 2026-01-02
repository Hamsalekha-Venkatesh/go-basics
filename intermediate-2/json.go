package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"first_name"`    // name is string-- so its in ""
	Age     int      `json:"age,omitempty"` //json naming convention to go naming convention
	Email   string   `json:"email,omitempty"`
	Address *Address `json:"home_address,omitempty"`
}

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func main() {
	person := Person{
		Name:  "John",
		Age:   32,
		Email: "hamsa@gmail.com",
		//Address: &Address{
		//	Street:  "123 Elm Street",
		//	City:    "Wisconsin",
		//	State:   "MN",
		//	ZipCode: "1234",
		//},
	}

	jsonData, err := json.MarshalIndent(person, " ", " ")
	if err != nil {
		return
	}
	fmt.Println(string(jsonData))

	var person2 Person
	err = json.Unmarshal(jsonData, &person2)
	if err != nil {
		return
	}
	fmt.Println(person2)
	// fmt.Println(person2.Address.State)
}
