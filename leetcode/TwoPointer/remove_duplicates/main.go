package main

import (
	"fmt"
	"slices"
)

func removeDuplicates(nums []int) int {
	l := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}

func test() bool {
	//test 1
	nums := []int{1, 2, 2, 2, 3, 4, 5}
	if k := removeDuplicates(nums); k != 5 {
		return false
	}
	expected := []int{1, 2, 3, 4, 5, 4, 5}
	if !slices.Equal(nums, expected) {
		return false
	}
	//test 2
	nums = []int{1, 2, 2, 3, 4, 4, 4, 4, 5, 6, 6, 7}
	if k := removeDuplicates(nums); k != 7 {
		return false
	}
	expected = []int{1, 2, 3, 4, 5, 6, 7, 4, 5, 6, 6, 7}
	if !slices.Equal(nums, expected) {
		return false
	}
	return true
}

func main() {
	fmt.Println(test())
}
