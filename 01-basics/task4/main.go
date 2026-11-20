package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var nonAlphanumericRegex = regexp.MustCompile(`[^\p{L}\p{N} ]+`)

func main() {
	arg := strings.Join(os.Args[1:], " ")

	fmt.Println(wordCount(arg))
}

func wordCount(s string) map[string]int {
	res := make(map[string]int)
	for w := range strings.FieldsSeq(s) {
		res[strings.ToLower(nonAlphanumericRegex.ReplaceAllString(w, ""))]++
	}

	return res
}
