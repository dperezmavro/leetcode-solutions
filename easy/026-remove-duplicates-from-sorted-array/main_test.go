package lib

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 1, 2},
			output: 2,
		},
		{
			input:  []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			output: 5,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v", tt.input), func(*testing.T) {
			output := removeDuplicates(tt.input)
			if tt.output != output {
				t.Errorf("got %d instead of %d for %+v", output, tt.output, tt.input)
			}
		})
	}
}
