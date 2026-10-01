package main

import (
	"fmt"
	"unicode"
)

func valid_palindrome(sentence string) bool {
	i, j := 0, len(sentence)-1
	for i < j {
		if unicode.ToLower(rune(sentence[i])) == unicode.ToLower(rune(sentence[j])) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(valid_palindrome("holaaloH"))
}
