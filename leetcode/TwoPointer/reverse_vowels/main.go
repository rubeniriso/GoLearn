package main

import (
	"fmt"
	"unicode"
)

func isVowel(b byte) bool {
	vowels := map[rune]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	}
	_, exists := vowels[unicode.ToLower(rune(b))]
	return exists
}

func reverse_vowels(sentence string) string {
	s := []byte(sentence)
	i, j := 0, len(s)-1

	for i < j {
		if !isVowel(s[i]) {
			i++
		} else if !isVowel(s[j]) {
			j--
		} else {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	return string(s)
}

func main() {
	sentence := "meLlAmOmanOlo" // moLlOmamOnAle
	fmt.Println(reverse_vowels(sentence))
}
