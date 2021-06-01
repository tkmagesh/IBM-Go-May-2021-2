package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2*r.Height + 2*r.Width
}

type ShapeWithArea interface {
	Area() float32
}

func printArea(shape ShapeWithArea) {
	fmt.Println("Area = ", shape.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func printPerimeter(shape ShapeWithPerimeter) {
	fmt.Println("Perimeter = ", shape.Perimeter())
}

/*
type ShapeWithDimensions interface {
	Area() float32
	Perimeter() float32
}
*/

type ShapeWithDimensions interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func printDimension(shape ShapeWithDimensions) {
	fmt.Println("Area = ", shape.Area())
	fmt.Println("Perimiter = ", shape.Perimeter())
}
func main() {
	circle := Circle{Radius: 10}
	/* printArea(circle)
	printPerimeter(circle) */
	printDimension(circle)
	rectangle := Rectangle{Height: 12, Width: 15}
	/* printArea(rectangle)
	printPerimeter(rectangle) */
	printDimension(rectangle)
}

/*
	Perimeter for Circle = 2 * Pi * r
	Perimeter for Rectangle = 2 * Height + 2 * Width
*/
