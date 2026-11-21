package main

import "fmt"

func main() {
	var data []int
	for range 30 {
		data = append(data, 1)
	}

	fmt.Println(stats(data))
}

func stats(nums []int) (sum int, avg float64) {
	if len(nums) == 0 {
		return 0, 0
	}

	for _, i := range nums {
		sum += i
	}

	return
}
