package wordcount

import (
	"maps"
	"testing"
)

func TestWordCount(t *testing.T) {
	cases := map[string]struct {
		s    string
		want map[string]int
	}{
		"empty string":  {s: "", want: map[string]int{}},
		"single word":   {s: "hello, Hello, hello!", want: map[string]int{"hello": 3}},
		"cirylic words": {s: "пиво? пиво!", want: map[string]int{"пиво": 2}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := wordCount(tc.s)
			if !maps.Equal(got, tc.want) {
				t.Errorf("wordCount(%v) = %v; want %v", tc.s, got, tc.want)
			}
		})
	}
}
