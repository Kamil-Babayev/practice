package stats

func stats(nums []int) (sum int, avg float64) {
	if len(nums) == 0 {
		return 0, 0
	}

	for _, i := range nums {
		sum += i
	}

	avg = float64(sum) / float64(len(nums))

	return
}
