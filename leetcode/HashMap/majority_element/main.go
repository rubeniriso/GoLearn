package main

import "fmt"

func majorityElement(nums []int) int {
	hashMap := make(map[int]int)

	for _, x := range nums {
		hashMap[x]++
		if hashMap[x] > len(nums)/2 {
			return x
		}
	}
	return -1
}

func main() {
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 2, 2}))
	fmt.Println(majorityElement([]int{1, 1, 1, 1, 2, 2, 2, 2}))
}
