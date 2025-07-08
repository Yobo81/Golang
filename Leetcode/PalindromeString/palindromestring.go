package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(-121))
}
