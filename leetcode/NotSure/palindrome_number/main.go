package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	div := 1
	for x >= 10*div {
		div *= 10
	}

	for x != 0 {
		if x/div != x%10 {
			return false
		}
		x = x % div / 10
		div /= 100
	}
	return true
}

func main() {
	palindrome := 12321
	notPalindrome := 12345

	fmt.Println(isPalindrome(palindrome))    // true
	fmt.Println(isPalindrome(notPalindrome)) // false
}
