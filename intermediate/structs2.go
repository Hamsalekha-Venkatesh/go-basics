package main

import "fmt"

type Employee struct {
	id           int
	name         string
	age          int
	salary       float64
	home_address Address
	PhoneCell
}

type Address struct {
	address1 string
	address2 string
	zipcode  int
	city     string
}

type PhoneCell struct {
	mobile string
	home   string
}

// struct methods...
func (emp Employee) FullName() string {
	var fullName = emp.name + string(emp.id)
	fmt.Println(emp.home) // anonymous ce
	fmt.Println(emp.home_address.address1)
	return fullName
}

func (emp *Employee) incrementSalary(delta int) float64 {
	emp.salary += float64(delta)
	return emp.salary
}

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
	fmt.Println("New incremented salary: ", employee1.incrementSalary(1000))
	fmt.Println("FullName(): = ", employee1.FullName())
}
