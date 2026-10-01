package main

import "fmt"

func pairWithTargetSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		switch {
		case sum == target:
			return []int{left, right}
		case sum < target:
			left++
		default:
			right--
		}
	}
	return []int{-1, -1}
}

func main() {
	numbers := []int{0, 1, 3, 5, 8, 12}
	target := 9
	fmt.Println(pairWithTargetSum(numbers, target))
}
