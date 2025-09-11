package main

import (
	"fmt"
	"team-sync/more-syntax/shapes"
)

// func (shp Shape) printShapeFacts() { Invalid code, cannot define a receiver function for an interface type.

// }

func printShapeFacts(shp shapes.Shape) {
	fmt.Printf("Area: %f\n", shp.CalculateArea())
	fmt.Printf("Perimter: %f\n", shp.CalculatePerimeter())
}

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

	// Arrays and slices.

	var shapesArray [2]shapes.Shape
	var shapesSlice []shapes.Shape

	shapesArray[0] = sq
	shapesArray[1] = cir

	// This would panic, uninitialized slice of length 0!
	// shapesSlice[0] = sq
	// shapesSlice[1] = cir

	// make() is used for: slice, map, or channel.
	shapesSlice = make([]shapes.Shape, 2)
	shapesSlice[0] = sq
	shapesSlice[1] = cir
}
