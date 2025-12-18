package main

import "fmt"

func main() {
	employee1 := Employee{
		id:     100,
		name:   "hamsalekha",
		age:    32,
		salary: 25000,
		home_address: Address{
			address1: "1234 west st",
			address2: "unit 5g",
			zipcode:  10001,
			city:     "nyc",
		},
	}
	fmt.Println(employee1.home_address)
}
