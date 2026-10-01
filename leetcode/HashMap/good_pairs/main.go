package main

import "fmt"

func goodPairs(numbers []int) [][]int {
	appearances := make(map[int][]int) // Using a map to mimic Java's HashSet
	var pairs [][]int

	for i, x := range numbers {
		for _, y := range appearances[x] {
			pairs = append(pairs, []int{y, i})
		}
		appearances[x] = append(appearances[x], i)
	}

	return pairs
}

func main() {
	numbers := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(goodPairs(numbers))
}
