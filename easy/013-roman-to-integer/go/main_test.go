package romantoint

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "III",
			output: 3,
		},
		{
			input:  "LVIII",
			output: 58,
		},
		{
			input:  "XIV",
			output: 14,
		},
		{
			input:  "MCMXCIV",
			output: 1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(*testing.T) {
			output := romanToInt(tt.input)
			if tt.output != output {
				t.Errorf("got %d instead of %d for %s", output, tt.output, tt.input)
			}
		})
	}
}
