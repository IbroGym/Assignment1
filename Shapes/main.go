package main

import (
	"fmt"
	"math"
)

// Shape interface defines methods for calculating area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle struct represents a rectangle with length and width
type Rectangle struct {
	Length float64
	Width  float64
}

// Area method calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Perimeter method calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// Circle struct represents a circle with radius
type Circle struct {
	Radius float64
}

// Area method calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method calculates the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Square struct represents a square with side length
type Square struct {
	Side float64
}

// Area method calculates the area of a square
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter method calculates the perimeter of a square
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

// Triangle struct represents a triangle with side lengths
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

// Area method calculates the area of a triangle using Heron's formula
func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

// Perimeter method calculates the perimeter of a triangle
func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

// PrintShapeDetails function takes a Shape interface and prints its area and perimeter
func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func main() {
	// Create a slice of shapes
	shapes := []Shape{
		Rectangle{Length: 5, Width: 3},
		Circle{Radius: 2},
		Square{Side: 4},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	// Iterate over the slice and print details for each shape
	for _, shape := range shapes {
		fmt.Println("** Shape Details **")
		PrintShapeDetails(shape)
		fmt.Println("------------------")
	}
}
