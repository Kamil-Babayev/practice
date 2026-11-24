package main

import "fmt"

func main() {
	a := make([]int, 3, 5) //{0, 0, 0} capacity - 5
	b := append(a, 4)      // same pointer, {0, 0, 0, 4} capacity - 5
	c := append(a, 5)      // same pointer, {0, 0, 0, 4, 5} capacity - 5
	fmt.Println(a, b, c)   // {0, 0, 0} {0, 0, 0, 4, 5} {0, 0, 0, 4, 5}

	d := []int{1, 2, 3} // {1, 2, 3}
	e := d[:2]          // {1, 2}
	e = append(e, 99)   // {1, 2, 99}
	fmt.Println(d, e)   // {1, 2, 99} {1, 2, 99}

	m := map[string][]int{}
	m["x"] = append(m["x"], 1)
	s := m["x"]
	s[0] = 42
	fmt.Println(m) // {x: {42}}
}
