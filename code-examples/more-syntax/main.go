package main

import (
	"fmt"
	"team-sync/code-examples/more-syntax/shapes"
)

func main() {
	sq := shapes.Square{SideLength: 3}
	cir := shapes.Circle{Radius: 2}
	printShapeFacts(sq)
	printShapeFacts(cir)
}

// func (shp Shape) printShapeFacts() { Invalid code, cannot define a receiver function for an interface type.

// }

func printShapeFacts(shp shapes.Shape) {
	fmt.Printf("Area: %d\n", shp.CalculateArea())
	fmt.Printf("Perimter: %d\n", shp.CalculatePerimeter())
}
