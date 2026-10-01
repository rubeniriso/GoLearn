package main

import (
	"fmt"
	"slices"
)

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func zeroSumTriplets(numbers []int) [][3]int {
	slices.Sort(numbers)
	result := [][3]int{}

	for i := 0; i < len(numbers)-2; i++ {
		// Skip duplicate values for the fixed pointer
		if i > 0 && numbers[i] == numbers[i-1] {
			continue
		}
		// Early exit: smallest possible sum is already positive
		if numbers[i] > 0 {
			break
		}

		left, right := i+1, len(numbers)-1

		for left < right {
			sum := numbers[i] + numbers[left] + numbers[right]
			switch {
			case sum == 0:
				result = append(result, [3]int{numbers[i], numbers[left], numbers[right]})
				// Skip duplicates on both sides
				for left < right && numbers[left] == numbers[left+1] {
					left++
				}
				for left < right && numbers[right] == numbers[right-1] {
					right--
				}
				left++
				right--
			case sum < 0:
				left++
			default:
				right--
			}
		}
	}

	return result
}

func main() {
	cases := [][]int{
		{-3, -2, -1, 0, 1, 1, 2},
		{},                    // empty
		{0, 0, 0},             // all zeros
		{1, 2, 3},             // no negatives
		{-1, -1, -1},          // no solution
		{-4, -1, -1, 0, 1, 2}, // multiple triplets
	}
	for _, c := range cases {
		fmt.Printf("%v → %v\n", c, zeroSumTriplets(c))
	}

}
