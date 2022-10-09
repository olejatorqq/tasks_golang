package main

import "fmt"

// my solution
func twoSum(nums []int, target int) []int {

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	arr := []int{3, 2, 4}
	fmt.Println(twoSum(arr, 6))
}
