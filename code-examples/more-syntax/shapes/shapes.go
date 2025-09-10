package shapes

type Shape interface {
	CalculateArea() int
	CalculatePerimeter() int
	PurelySideEffects()
}

type Square struct {
	SideLength int
}

type Circle struct {
	Radius int
}

// CalculateArea implements Shape.
func (s Square) CalculateArea() int {
	panic("unimplemented")
}

// CalculatePerimeter implements Shape.
func (s Square) CalculatePerimeter() int {
	panic("unimplemented")
}

// PurelySideEffects implements Shape.
func (s Square) PurelySideEffects() {
	panic("unimplemented")
}

// CalculateArea implements Shape.
func (c Circle) CalculateArea() int {
	panic("unimplemented")
}

// CalculatePerimeter implements Shape.
func (c Circle) CalculatePerimeter() int {
	panic("unimplemented")
}

// PurelySideEffects implements Shape.
func (c Circle) PurelySideEffects() {
	panic("unimplemented")
}
