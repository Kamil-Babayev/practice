package slicex

import (
	"errors"
	"slices"
)

var ErrInvalidInput = errors.New("invalid input value")

// Unique returns a new slice of unique elements of input slice keeping the order
func Unique(nums []int) []int {
	seen := make(map[int]bool)
	var res []int

	for i := range nums {
		if !seen[nums[i]] {
			res = append(res, nums[i])
			seen[nums[i]] = true
		}
	}

	return res
}

// Reverse reverces the slice in place
func Reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Chunk splits input slice into chunks of size parameter
// if length of input slice or size param is less than or equal to zero returns an error
// does not share the memory of the input slice, creates a new one
func Chunk(nums []int, size int) ([][]int, error) {
	if len(nums) <= 0 || size <= 0 {
		return nil, ErrInvalidInput
	}

	var res [][]int

	for i := 0; i < len(nums); i += size {
		res = append(res, slices.Clone(nums[i:min(i+size, len(nums))]))
	}

	return res, nil
}
