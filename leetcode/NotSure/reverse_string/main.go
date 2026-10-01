package main

import "fmt"

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		aux := s[i]
		s[i] = s[j]
		s[j] = aux
		j--
		i++
	}
	fmt.Println(string(s))
}

func main() {
	bs := []byte("hola")
	reverseString(bs)
}
