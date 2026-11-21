package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	test := "1А роза упала на лапу Азора 1"
	fmt.Println(isPalindrome(test))
}

func isPalindrome(s string) bool {
	runes := []rune(strings.ToLower(s))

	n := 0

	for _, r := range runes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			runes[n] = r
			n++
		}
	}

	runes = runes[:n]

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}
