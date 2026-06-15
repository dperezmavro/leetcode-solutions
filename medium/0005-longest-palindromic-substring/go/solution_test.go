package logest_palindromic_substring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "aba",
			want:  "aba",
		},
		{
			input: "babad",
			want:  "bab",
		},
		{
			input: "ccc",
			want:  "ccc",
		},
		{
			input: "cbbd",
			want:  "bb",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			res := longestPalindrome(tt.input)
			if res != tt.want {
				t.Errorf("want %s got %s", tt.want, res)
			}
		})
	}
}
