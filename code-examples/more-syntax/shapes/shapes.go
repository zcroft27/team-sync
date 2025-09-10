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

// CalculateArea implements Shape.
func (s Square) CalculateArea() float64 {
	return float64(s.SideLength * s.SideLength)
}

// CalculatePerimeter implements Shape.
func (s Square) CalculatePerimeter() float64 {
	return float64(s.SideLength) * 4
}

// PurelySideEffects implements Shape.
func (s Square) PurelySideEffects() {
	fmt.Println("I don't return anything, square.")
}

// CalculateArea implements Shape.
func (c Circle) CalculateArea() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

// CalculatePerimeter implements Shape.
func (c Circle) CalculatePerimeter() float64 {
	return 2 * math.Pi * float64(c.Radius)
}

// PurelySideEffects implements Shape.
func (c Circle) PurelySideEffects() {
	fmt.Println("I don't return anything, circle.")
}
