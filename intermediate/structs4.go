package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64

	area float64
}

// Pointer reciever
func (r *Rectangle) Area() float64 {
	r.area = r.length * r.width

	return r.area
}

// no need to pass the instance--- we
func (Rectangle) scale(factor int) {
	fmt.Println(factor)
}

func main() {
	rect1 := Rectangle{length: 10, width: 20}

	rect1.Area()
	fmt.Println("Area = ", rect1.width, rect1.length)
	fmt.Println("Area = ", rect1.area)
	rect1.scale(3)
	fmt.Println("Area = ", rect1.width, rect1.length)

}
