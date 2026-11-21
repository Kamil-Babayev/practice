package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	arg := strings.Join(os.Args[1:], " ")

	fmt.Println(wordCount(arg))
}

func wordCount(s string) map[string]int {
	res := make(map[string]int)

	isSeparator := func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}

	for _, w := range strings.FieldsFunc(strings.ToLower(s), isSeparator) {
		res[w]++
	}

	return res
}
