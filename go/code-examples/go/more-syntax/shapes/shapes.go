package shapes

import (
	"fmt"
	"math"
)

type Shape interface {
	CalculateArea() float64
	CalculatePerimeter() float64
	PurelySideEffects()
}

type Square struct {
	SideLength int
}

type Circle struct {
	Radius int
}

// "If it walks like a duck and it quacks like a duck, then it must be a duck"
// A rectangle is not a duck.
type Rectangle struct {
	Width  int
	Height int
}

// Square

func (s Square) CalculateArea() float64 {
	return float64(s.SideLength * s.SideLength)
}

func (s Square) CalculatePerimeter() float64 {
	return float64(s.SideLength) * 4
}

func (s Square) PurelySideEffects() {
	fmt.Println("I don't return anything, square.")
}

// Circle

func (c Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

func (c Circle) CalculatePerimeter() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

func (c Circle) PurelySideEffects() {
	fmt.Println("I don't return anything, circle.")
}
