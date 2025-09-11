package main

import (
	"fmt"
	"team-sync/more-syntax/shapes"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic. phew!")
		}
		fmt.Println("woohoo shapes!")
	}()

	sq := shapes.Square{SideLength: 3}
	cir := shapes.Circle{Radius: 4}
	// panic("shapes!")
	printShapeFacts(sq)
	printShapeFacts(cir)
}

// func (shp Shape) printShapeFacts() { Invalid code, cannot define a receiver function for an interface type.

// }

func printShapeFacts(shp shapes.Shape) {
	fmt.Printf("Area: %f\n", shp.CalculateArea())
	fmt.Printf("Perimter: %f\n", shp.CalculatePerimeter())
}
