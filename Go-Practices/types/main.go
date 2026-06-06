package main

import "fmt"

func main() {
	// Declare variables
	var x int = 10
	var y int = 3
	var name string = "Tayo"
	z := 4.5 // floated 64 inferred

	// Print variables
	fmt.Println("Name", name)
	fmt.Println("x + y =", x+y)
	fmt.Println("x - y =", x-y)
	fmt.Println("x * y =", x*y)
	fmt.Println("x / y =", x/y) // integer division
	fmt.Println("z =z", z)
}
