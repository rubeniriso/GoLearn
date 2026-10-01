package main

import "fmt"

func sqrt(number int, min int, max int) int {
	current := (min + max) / 2
	if current*current <= number && (current+1)*(current+1) > number {
		return current
	} else if current*current > number {
		return sqrt(number, min, current)
	} else {
		return sqrt(number, current, max)
	}
}

func sqrt_iterative(number int) int {
	min, max := 0, number
	for min < max {
		mid := (min + max + 1) / 2
		if mid*mid <= number {
			min = mid
		} else {
			max = mid - 1
		}
	}
	return min
}

func main() {
	var number int
	number = 100
	fmt.Println(sqrt(number, 1, number))
}
