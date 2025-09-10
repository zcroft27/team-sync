package main

import "fmt"

func main() {
  // var x int
  // x = 3 (commented out because you cannot declare and not use a variable, won't compile)
  y := 8
  // y:= 4 (errors, no new variables on the left side of :=)
  // y = 4.5 (errors, float != int)
  // y = "hello" (errors, string != int)
  fmt.Printf("Type: %T\n", y)
}
