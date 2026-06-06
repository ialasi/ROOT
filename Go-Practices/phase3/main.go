package main

import "fmt"

func main() {
	text := "Palindrome"

	runes := []rune(text)

	isPalindrome := true

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrom")
	}
}
