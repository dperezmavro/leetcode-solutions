package palindrome_number

import (
	"fmt"
)

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%d", x)
	var i int = 0
	for i = 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPalindromeNoConv(x int) bool {
	s := fmt.Sprintf("%d", x)
	var i int = 0
	for i = 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
