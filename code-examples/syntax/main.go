package main

import "fmt"

type Person struct {
	Name              string
	Age               int    // exported, anyone can access the Age field on a Person struct.
	mothersMaidenName string // unexported, classes outside of the main package cannot access this.
}

func main() {
	// var x int
	// x = 3 (commented out because you cannot declare and not use a variable, won't compile)
	y := 8
	// y:= 4 (errors, no new variables on the left side of :=)
	// y = 4.5 (errors, float != int)
	// y = "hello" (errors, string != int)

	var z bool
	// z := true (errors, no new variables on the left side of :=)
	z = true

	fmt.Printf("Type: %T\n", y)

	foo()

	if z {
		demonstrateNameChange()
	}

	anonymousFunction := func() {
		fmt.Println("hello")
	}

	anonymousFunction()

	func() {
		fmt.Println("world")
	}()
}

// Private function scoped to the package.
func foo() {
	var bar string = "fizzbuzz"
	fmt.Printf("bar: %s\n", bar)
}

// Public receiver function (makes a copy of the struct).
func (p Person) AlterName(newName string) {
	p.Name = newName
	_ = p.Name // to silence static analysis complaints.
}

// Public pointer receiver function.
func (p *Person) AlterNamePointer(newName string) {
	p.Name = newName
}

func demonstrateNameChange() {
	george := Person{Name: "George", mothersMaidenName: "Watson"}

	fmt.Printf("George's Age (default value): %d\n", george.Age)

	fmt.Printf("George's name before: %s\n", george.Name)
	george.AlterName("michael")
	fmt.Printf("George's name after value receiver: %s\n", george.Name)

	// Don't need to explicitly reference george.
	george.AlterNamePointer("michael")
	fmt.Printf("George's name after first pointer receiver %s\n", george.Name)

	(&george).AlterNamePointer("david")
	fmt.Printf("George's name after second pointer receiver %s\n", george.Name)
}
