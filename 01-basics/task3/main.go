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
	for _, i := range nums {
		sum += i
	}

	//if sum == 0 then either len(nums) == 0 or sum of given nums is 0
	// both cases we should return 0 as avg
	if sum != 0 {
		avg = float64(sum) / float64(len(nums))
	}

	return
}
