package main

import "fmt"

func main() {
	a := make([]int, 3, 5)
	b := append(a, 4)
	c := append(a, 5)
	fmt.Println(a, b, c)

	d := []int{1, 2, 3}
	e := d[:2]
	e = append(e, 99)
	fmt.Println(d, e)

	m := map[string][]int{}
	m["x"] = append(m["x"], 1)
	s := m["x"]
	s[0] = 42
	fmt.Println(m)
}
