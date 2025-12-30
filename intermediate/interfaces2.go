package main

import "fmt"

// Animal Step 1: Define common interface with generic methods...
type Animal interface {
	getAge() uint
	canWalk() bool
	getName() string
}

// Step 2: Define generic methods calling the respective interface methods//
func getAge(a Animal) uint {
	return a.getAge()
}

func canWalk(a Animal) bool {
	return a.canWalk()
}

func getName(a Animal) string {
	return a.getName()
}

// Human Step 3: Define struct objects specific to interface type...
type Human struct {
	name     string
	walkMode bool
	age      uint
}

// Dog Step 3: Define struct objects specific to interface type...
type Dog struct {
	name     string
	age      uint
	walkMode bool
}

// Step 4: ***************** Human Implementation ************************
// this is how Go understands Human is of type Animal since we are implementing all its methods.
func (h Human) getAge() uint {
	fmt.Println("Human's lifespan is max 100 years")
	return h.age
}

func (h Human) getName() string {
	return h.name
}

func (h Human) canWalk() bool {
	fmt.Println("Human's can only use 2 legs to walk")
	return h.walkMode
}

// ***************** Dog Implementation ************************
// this is how Go understands Human is of type Animal since we are implementing all its methods.
func (d Dog) getAge() uint {
	fmt.Println("Dog's lifespan is max 15 years")
	fmt.Println("Dog's human life age is ", d.age*5)
	return d.age
}

func (d Dog) getName() string {
	return d.name
}

func (d Dog) canWalk() bool {
	fmt.Println("Dog's can only use 4 legs to walk")
	return d.walkMode
}

func main() {
	// Declare Human and Dog Objects
	h1 := Human{
		name:     "Hamsa",
		walkMode: true,
		age:      32,
	}

	d1 := Dog{
		name:     "Snoopy",
		age:      8,
		walkMode: false,
	}

	d2 := Dog{"Alfie", 10, true}

	// Now Go can understand which type of Animal object is passed and will invoke resp methods.
	fmt.Println("My Dogs data: ", getName(d1), getAge(d1), canWalk(d1))
	fmt.Println("My Human data: ", getName(h1), getAge(h1), canWalk(h1))
	fmt.Println(getAge(d2))

}
