package main

import (
	"fmt"
)

func main() {

	palindromestring := "lepel"
	result := isPalindrome(palindromestring)
	fmt.Println(result)
	if !result {
		fmt.Println(toMakePalindrome(palindromestring))
	}
}

func isPalindrome(palindromestring string) bool {

	reversedstring := Reverse(palindromestring)

	if palindromestring != reversedstring {
		return false
	}
	return true
}

func Reverse(palindrome string) string {

	reversedstring := ""
	for i := len(palindrome) - 1; i >= 0; i-- {
		reversedstring += string(palindrome[i])
	}
	fmt.Println(reversedstring)
	return string(reversedstring)
}

func toMakePalindrome(palindromestring string) int {

	var counter int = 0
	reversedstring := Reverse(palindromestring)
	for i := len(palindromestring) - 1; i >= 0; i-- {
		if reversedstring[i] != palindromestring[i] {
			counter++
		}

	}
	return counter
}
