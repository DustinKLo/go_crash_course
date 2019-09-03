package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x      float64
	y      float64
	radius float64
}

type Rectangle struct {
	width, height float64
}

// no need for *Circle because we're not changing the data in the struct (or object)
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

// INTERFACE
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("circle area %f\n", getArea(circle))
	fmt.Printf("rectangle area %f\n", getArea(rectangle))

	fmt.Printf("circle area %f\n", circle.area())
	fmt.Printf("rectangle area %f\n", rectangle.area())
}
