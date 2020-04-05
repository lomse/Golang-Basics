package main

import (
	"fmt"
	"math"
)

// Shape is an interface
type Shape interface {
	area() float64
}

// Circle is a struct
type Circle struct {
	radius float64
}

// Rectangle is a struct
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{height: 45, width: 300}

	fmt.Printf("Area of circle is %f\n", getArea(circle))
	fmt.Printf("Area of rectangle is %f\n", getArea(rectangle))

}
