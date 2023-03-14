package main

import (
	"fmt"
	"math"
)

type Figure interface {
	Name() string
	Area() float64
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (rectangle Rectangle) Name() string {
	return "Rectangle"
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Base * rectangle.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Name() string {
	return "Circle"
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func main() {
	var rectangle Rectangle
	rectangle.Base = 2
	rectangle.Height = 8
	ShowArea(rectangle)

	var circle Circle
	circle.Radius = 5
	ShowArea(circle)
}

func ShowArea(figure Figure) {
	fmt.Println(fmt.Sprintf("Area of %s: %.2f", figure.Name(), figure.Area()))
}
