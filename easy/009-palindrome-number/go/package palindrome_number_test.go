package palindrome_number

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		res   bool
		input int
	}{
		{
			input: -10,
		},
		{
			input: -11,
		},
		{
			input: 121,
			res:   true,
		},
		{
			input: 1221,
			res:   true,
		},
		{
			input: 1231,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.input), func(t *testing.T) {
			res := isPalindrome(tt.input)
			if res != tt.res {
				t.Errorf("failed test for input %d: wanted %t, got %t", tt.input, tt.res, res)
			}
		})
	}
}
