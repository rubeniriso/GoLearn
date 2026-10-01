package main

import "fmt"

func findNonDuplicates(nums []int) int {
	next := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[next] = nums[i]
			next++
		}
	}
	return next
}

func main() {
	fmt.Println(findNonDuplicates([]int{2, 3, 3, 3, 6, 9, 9}))
}
