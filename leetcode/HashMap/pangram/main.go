package main

import (
	"fmt"
	"strings"
	"unicode"
)

func pangram(sentence string) bool {
	lowercase_string := strings.ToLower(sentence)

	seen_letter := make(map[rune]struct{}) // Using a map to mimic Java's HashSet

	for _, x := range lowercase_string {
		if unicode.IsLetter(x) {
			seen_letter[x] = struct{}{}
		}
	}
	return len(seen_letter) == 26
}

func main() {
	var nums string

	nums = "me llamo manolo"
	fmt.Println(pangram(nums))

	nums = "TheQuickBrownFoxJumpsOverTheLazyDog"
	fmt.Println(pangram(nums))
}
