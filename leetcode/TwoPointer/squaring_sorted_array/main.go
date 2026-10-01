package main

import (
	"fmt"
)

func squaringSortedArray(numbers []int) []int {
	left, right, insertIndex := 0, len(numbers)-1, len(numbers)-1
	sortedArray := make([]int, len(numbers))

	for left <= right {
		switch {
		case numbers[left]*numbers[left] >= numbers[right]*numbers[right]:
			sortedArray[insertIndex] = numbers[left] * numbers[left]
			left++
		case numbers[left]*numbers[left] < numbers[right]*numbers[right]:
			sortedArray[insertIndex] = numbers[right] * numbers[right]
			right--
		}
		insertIndex--
	}
	return sortedArray
}

func main() {
	numbers := []int{-2, -1, 0, 2, 3}
	fmt.Println(squaringSortedArray(numbers))
	numbers = []int{-3, -1, 0, 1, 2}
	fmt.Println(squaringSortedArray(numbers))
}
