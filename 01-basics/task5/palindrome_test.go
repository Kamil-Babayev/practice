package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := map[string]struct {
		s    string
		want bool
	}{
		"numbers": {s: "12321", want: true},
		"string":  {s: "qwertyuiop", want: false},
		"cirylic": {s: "йййцукенгшщз", want: false},
		"empty":   {s: "", want: true},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := isPalindrome(tc.s)
			if got != tc.want {
				t.Errorf("isPalindrome(%v) = %v; want %v", tc.s, got, tc.want)
			}
		})
	}
}
