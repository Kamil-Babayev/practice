package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := map[string]struct {
		a    string
		want bool
	}{
		"numbers": {a: "12321", want: true},
		"string":  {a: "qwertyuiop", want: false},
		"cirylic": {a: "йййцукенгшщз", want: false},
		"empty":   {a: "", want: true},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := isPalindrome(tc.a)
			if got != tc.want {
				t.Errorf("isPalindrome(%v) = %v; want %v", tc.a, got, tc.want)
			}
		})
	}
}
