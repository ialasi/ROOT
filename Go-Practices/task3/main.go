package main

import "fmt"

func main() {
	a := 100
	b := 20

	sum := a + b
	product := a * b
	difference := a - b
	division := float64(a) / float64(b)

	fmt.Println("Sum:", sum)
	fmt.Println("Product:", product)
	fmt.Println("Difference:", difference)
	fmt.Println("Division:", division)
}
