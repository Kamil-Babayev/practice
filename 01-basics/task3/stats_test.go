package stats

import (
	"math"
	"testing"
)

func TestStats(t *testing.T) {
	cases := map[string]struct {
		nums    []int
		wantSum int
		wantAvg float64
	}{
		"positive numbers": {
			nums:    []int{1, 2, 3},
			wantSum: 6,
			wantAvg: 2,
		},
		"negative numbers": {
			nums:    []int{-10, -20},
			wantSum: -30,
			wantAvg: -15.0,
		},
		"empty slice": {
			nums:    []int{},
			wantSum: 0,
			wantAvg: 0.0,
		},
		"nil slice": {
			nums:    nil,
			wantSum: 0,
			wantAvg: 0.0,
		},
		"single element": {
			nums:    []int{5},
			wantSum: 5,
			wantAvg: 5.0,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			gotSum, gotAvg := stats(tc.nums)
			if gotSum != tc.wantSum {
				t.Errorf("stats(%v) sum = %d; want %d", tc.nums, gotSum, tc.wantSum)
			}
			if math.Abs(gotAvg-tc.wantAvg) > 1e-9 {
				t.Errorf("stats(%v) avg = %v; want %v", tc.nums, gotAvg, tc.wantAvg)
			}
		})
	}
}
