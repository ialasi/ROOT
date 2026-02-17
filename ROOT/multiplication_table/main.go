package main

import "fmt"

func main() {
	// Outer loop for numbers 1 to 5

	for i := 1; i <= 5; i++ {
		// Inner loop for multiplying by 1 to 5
		for j := 1; j <= 5; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println() // Blank line after each table row
	}
}
