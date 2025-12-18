package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
	salary    float64
}

func main() {

	person := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       42,
		salary:    1000,
	}

	fmt.Println(person.firstName, person.lastName, person.age, person.salary)

	user := struct {
		username string
		email    string
	}{
		username: "johndoe",
		email:    "userme@gmail.com",
	}

	fmt.Println(user.username, user.email)

}
