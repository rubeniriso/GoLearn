package main

import "fmt"

func containsDuplicate(nums []int) bool {
	unique_set := make(map[int]struct{}) // Using a map to mimic Java's HashSet
	for _, x := range nums {

		if _, exists := unique_set[x]; exists {
			// If the element is already in the set, return true
			return true
		}
		unique_set[x] = struct{}{} // Add the element to the set
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
