package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	greet("Tayo")
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}
