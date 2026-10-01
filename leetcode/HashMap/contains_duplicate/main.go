package main

import "fmt"

func containsDuplicate(nums []int) bool {
	unique_set := make(map[int]struct{})
	for _, x := range nums {

		if _, exists := unique_set[x]; exists {
			return true
		}
		unique_set[x] = struct{}{}
	}
	return false
}

func main() {
	var nums []int

	nums = []int{1, 2, 3}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1, 2, 2, 4}
	fmt.Println(containsDuplicate(nums))
}
