package wordcount

import (
	"maps"
	"testing"
)

func TestWordCount(t *testing.T) {
	cases := map[string]struct {
		a    string
		want map[string]int
	}{
		"empty string":  {a: "", want: map[string]int{}},
		"single word":   {a: "hello, Hello, hello!", want: map[string]int{"hello": 3}},
		"cirylic words": {a: "пиво? пиво!", want: map[string]int{"пиво": 2}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := wordCount(tc.a)
			if !maps.Equal(got, tc.want) {
				t.Errorf("wordCount(%v) = %v; want %v", tc.a, got, tc.want)
			}
		})
	}
}
