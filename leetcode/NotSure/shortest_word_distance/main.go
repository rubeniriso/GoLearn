package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func shortestWordDistance(words []string, first string, second string) int {
	firstPosition, secondPosition, distance := -1, -1, len(words)
	for pos, x := range words {
		switch x {
		case first:
			firstPosition = pos
		case second:
			secondPosition = pos
		}
		if firstPosition != -1 && secondPosition != -1 {
			distance = min(abs(firstPosition-secondPosition), distance)
		}
	}
	return distance
}

func shortestWordDistance2(words []string, first string, second string) int {
	firstPosition, secondPosition, distance := -1, -1, len(words)
	for pos, x := range words {
		if x == first {
			firstPosition = pos
		}
		if x == second {
			secondPosition = pos
		}
		if firstPosition != -1 && secondPosition != -1 {
			distance = min(abs(firstPosition-secondPosition), distance)
		}
	}
	return distance
}

func main() {
	words := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	word1 := "fox"
	word2 := "dog"
	fmt.Println(shortestWordDistance2(words, word1, word2))
	words = []string{"a", "c", "d", "b", "a"}
	word1 = "a"
	word2 = "b"
	fmt.Println(shortestWordDistance2(words, word1, word2))
	words = []string{"a", "c", "d", "b", "a"}
	word1 = "a"
	word2 = "a"
	fmt.Println(shortestWordDistance2(words, word1, word2))
}
