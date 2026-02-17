package main

import "fmt"

func main() {
	var number int

	//Ask user for input
	fmt.Println("Enter a munber: ")
	fmt.Scan(&number)

	// Check if even or odd
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
