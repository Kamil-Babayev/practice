package slicex

import (
	"errors"
	"slices"
	"testing"
)

func TestUnique(t *testing.T) {
	cases := map[string]struct {
		nums []int
		want []int
	}{
		"empty":          {nums: []int{}, want: []int{}},
		"nil":            {nums: nil, want: []int{}},
		"no duplicates":  {nums: []int{1, 2, 3, 4, 5, 6}, want: []int{1, 2, 3, 4, 5, 6}},
		"all duplicates": {nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, want: []int{1}},
		"mixed":          {nums: []int{2, 4, 4, 67, 69, 88, 67, 99, 2}, want: []int{2, 4, 67, 69, 88, 99}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := Unique(tc.nums)
			if slices.Compare(got, tc.want) != 0 {
				t.Errorf("Unique(%v) = %v; want %v", tc.nums, got, tc.want)
			}
		})
	}
}

type reverseCase[T any] struct {
	s    []T
	want []T
}

func runReverseTests[T comparable](t *testing.T, cases map[string]reverseCase[T]) {
	t.Helper()
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := slices.Clone(tc.s)
			Reverse(got)

			if !slices.Equal(got, tc.want) {
				t.Errorf("Reverse(%v) resulted in %v; want %v", tc.s, got, tc.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		runReverseTests(t, map[string]reverseCase[int]{
			"nil":      {s: nil, want: nil},
			"empty":    {s: []int{}, want: []int{}},
			"single":   {s: []int{1}, want: []int{1}},
			"even len": {s: []int{1, 2, 3, 4}, want: []int{4, 3, 2, 1}},
			"odd len":  {s: []int{1, 2, 3}, want: []int{3, 2, 1}},
		})
	})

	t.Run("strings", func(t *testing.T) {
		runReverseTests(t, map[string]reverseCase[string]{
			"words": {s: []string{"go", "is", "fun"}, want: []string{"fun", "is", "go"}},
		})
	})
}

func assert2DSliceEqual(t *testing.T, got, want [][]int) {
	t.Helper()
	if !slices.EqualFunc(got, want, slices.Equal) {
		t.Errorf("got %v; want %v", got, want)
	}
}

func TestChunk(t *testing.T) {
	cases := map[string]struct {
		nums      []int
		size      int
		wantSlice [][]int
		wantErr   error
	}{
		"zero size": {
			nums:      []int{1, 2},
			size:      0,
			wantSlice: nil,
			wantErr:   ErrInvalidInput,
		},
		"size bigger than slice len": {
			nums:      []int{1, 2},
			size:      3,
			wantSlice: [][]int{{1, 2}},
			wantErr:   nil,
		},
		"normal": {
			nums:      []int{1, 2, 3, 4, 5, 6, 7, 8},
			size:      2,
			wantSlice: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			wantErr:   nil,
		},
		"nil slice": {
			nums:      nil,
			size:      5,
			wantSlice: nil,
			wantErr:   ErrInvalidInput,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			gotSlice, gotErr := Chunk(tc.nums, tc.size)

			if tc.wantErr != nil {
				if gotErr == nil {
					t.Fatalf("got nil error; want %q", tc.wantErr)
				}
				if !errors.Is(gotErr, tc.wantErr) {
					t.Errorf("got error %q; want %q", gotErr.Error(), tc.wantErr)
				}
			} else if gotErr != nil {
				t.Fatalf("unexpected error: %v", gotErr)
			}

			assert2DSliceEqual(t, gotSlice, tc.wantSlice)
		})
	}
}
