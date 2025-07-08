package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				fmt.Printf("Output: [%d, %d]\n", i, j)
				return []int{i, j} // Return the result
			}
		}
	}
	return nil // If no result is found
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println("Result:", result)
}
