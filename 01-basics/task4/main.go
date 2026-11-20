package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := strings.Join(os.Args[1:], " ")

	fmt.Println(wordCount(arg))
}

func wordCount(s string) map[string]int {
	res := make(map[string]int)
	for _, w := range strings.Fields(s) {
		res[w]++
	}

	return res
}
