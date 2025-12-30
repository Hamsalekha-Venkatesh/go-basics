package main

import "math"
import "fmt"

type Geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type rectangle1 struct {
	width, height float64
}

type circle struct {
	radius float64
}

// *********** rectangle methods ***************
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//*********** rectangle methods ***************

// *********** Circle methods ***************
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// extra  methods are still possible!
func (c circle) dimeter() float64 {
	return 2 * c.radius
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

//*********** Circle methods ***************

func getArea(g Geometry) float64 {
	return g.area()
}

func (r rectangle1) area() float64 {
	return r.width * r.height
}

func main() {
	fmt.Println("Hello")
	r1 := rectangle{width: 10, height: 5}
	c1 := circle{radius: 5}
	//  r2 := rectangle1{width: 10, height: 5}

	fmt.Println("Area is r1 = ", getArea(r1))
	fmt.Println("Area is c1 = ", getArea(c1))

	myPrinter(1, "sdfsgfsfg", 23.4434, true)
}

func myPrinter(i ...interface{}) {
	for _, value := range i {
		fmt.Println(value)
	}
}
