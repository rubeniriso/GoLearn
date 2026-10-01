package main

import "fmt"

func isAnagram(first string, second string) bool {
	if len(first) != len(second) {
		return false
	}
	anagram_map := make(map[rune]int)
	for i := 0; i < len(first); i++ {
		anagram_map[rune(first[i])]++
		anagram_map[rune(second[i])]--
		if anagram_map[rune(first[i])] == 0 {
			delete(anagram_map, rune(first[i]))
		}
		if anagram_map[rune(second[i])] == 0 {
			delete(anagram_map, rune(second[i]))
		}
	}
	return len(anagram_map) == 0
}

func main() {
	fmt.Println(isAnagram("silent", "listen")) //true
	fmt.Println(isAnagram("hola", "holo"))     //false
	fmt.Println(isAnagram("hola", "chola"))    //false
}
