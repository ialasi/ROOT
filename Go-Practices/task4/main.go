package main

import "fmt"

func main() {
	number := 0

	if number > 0 {
		fmt.Println("Positive")
	} else if number < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}
