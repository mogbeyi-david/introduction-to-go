package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	length float64
	width  float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

func (square Square) area() float64 {
	return square.length * square.width
}

func getArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	newCircle := Circle{7}
	newSquare := Square{10, 10}
	fmt.Println(getArea(newCircle), getArea(newSquare))
}
