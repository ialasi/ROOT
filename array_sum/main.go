package main

import "fmt"

func main() {
	// Declare an array of integers
	numbers := [5]int{1, 2, 3, 4, 5}

	sum := 0 // variable to store sum

	// Loop through the array
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i] // add each element to sum
	}

	// Print the total sum
	fmt.Println("Sum of array:", sum)
}
