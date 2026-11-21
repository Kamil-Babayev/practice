package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "World"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Hello, %s!\n", s)
}
