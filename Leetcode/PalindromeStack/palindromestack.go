package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	stack := []int{}
	original := x
	for x > 0 {
		digit := x % 10
		stack = append(stack, digit)
		x /= 10
	}
	reversed := 0
	for _, digit := range stack {
		reversed = reversed*10 + digit
	}
	return original == reversed
}
func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(1221))
	fmt.Println(isPalindrome(-121))
}
